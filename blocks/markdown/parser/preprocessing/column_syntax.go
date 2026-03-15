package pre

import (
	"fmt"
	"math"
	"strings"

	"github.com/mathisbot/notionary-go/blocks/markdown/syntax"
)

const (
	ratioTolerance = 0.0001
	minimumColumns = 2
)

type InsufficientColumnsError struct {
	Count int
}

func (e *InsufficientColumnsError) Error() string {
	return fmt.Sprintf("column list must contain at least %d columns, found %d", minimumColumns, e.Count)
}

type InvalidColumnRatioSumError struct {
	Sum       float64
	Tolerance float64
}

func (e *InvalidColumnRatioSumError) Error() string {
	return fmt.Sprintf("column ratios must sum to 1.0 (±%.4f), but sum to %.4f", e.Tolerance, e.Sum)
}

func mustTagSyntax(key syntax.RegistryKey) syntax.TagSyntax {
	def, ok := syntax.Registry[key].(syntax.TagSyntax)
	if !ok {
		panic(fmt.Sprintf("syntax: %q is not a TagSyntax", key))
	}
	return def
}

type ColumnSyntaxPreProcessor struct {
	columnList syntax.TagSyntax
	column     syntax.TagSyntax
}

func NewColumnSyntaxPreProcessor() *ColumnSyntaxPreProcessor {
	return &ColumnSyntaxPreProcessor{
		columnList: mustTagSyntax(syntax.ColumnList),
		column:     mustTagSyntax(syntax.Column),
	}
}

func (p *ColumnSyntaxPreProcessor) Process(markdownText string) (string, error) {
	if !strings.Contains(markdownText, p.columnList.OpenTag) {
		return markdownText, nil
	}

	if err := p.validateAllColumnLists(markdownText); err != nil {
		return "", err
	}

	return markdownText, nil
}

func (p *ColumnSyntaxPreProcessor) validateAllColumnLists(markdownText string) error {
	for _, block := range p.extractColumnListBlocks(markdownText) {
		if err := p.validateColumnListBlock(block); err != nil {
			return err
		}
	}
	return nil
}

func (p *ColumnSyntaxPreProcessor) extractColumnListBlocks(markdownText string) []string {
	lines := strings.Split(markdownText, "\n")
	var blocks []string

	for i, line := range lines {
		if p.columnList.Pattern.MatchString(strings.TrimSpace(line)) {
			blocks = append(blocks, p.extractUntilClosingTag(lines, i+1))
		}
	}

	return blocks
}

func (p *ColumnSyntaxPreProcessor) extractUntilClosingTag(lines []string, startIndex int) string {
	var blockLines []string
	for _, line := range lines[startIndex:] {
		if p.columnList.EndPattern.MatchString(strings.TrimSpace(line)) {
			break
		}
		blockLines = append(blockLines, line)
	}
	return strings.Join(blockLines, "\n")
}

func (p *ColumnSyntaxPreProcessor) validateColumnListBlock(blockContent string) error {
	matches := p.column.Pattern.FindAllStringSubmatch(blockContent, -1)

	if len(matches) < minimumColumns {
		return &InsufficientColumnsError{Count: len(matches)}
	}

	return p.validateRatioSum(p.extractColumnRatios(matches), len(matches))
}

func (p *ColumnSyntaxPreProcessor) extractColumnRatios(matches [][]string) []float64 {
	var ratios []float64
	for _, match := range matches {
		if len(match) > 1 && match[1] != "" {
			var ratio float64
			fmt.Sscanf(match[1], "%f", &ratio)
			ratios = append(ratios, ratio)
		}
	}
	return ratios
}

func (p *ColumnSyntaxPreProcessor) validateRatioSum(ratios []float64, columnCount int) error {
	if len(ratios) == 0 || len(ratios) != columnCount {
		return nil
	}

	var total float64
	for _, r := range ratios {
		total += r
	}

	if math.Abs(total-1.0) > ratioTolerance {
		return &InvalidColumnRatioSumError{Sum: total, Tolerance: ratioTolerance}
	}

	return nil
}