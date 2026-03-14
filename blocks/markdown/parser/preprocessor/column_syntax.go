package preprocessor

import (
	"fmt"
	"math"
	"regexp"
	"strings"
)

const (
	columnListDelimiter = "::: columns"
	ratioTolerance      = 0.0001
	minimumColumns      = 2
)

var columnPattern = regexp.MustCompile(`:::\s+column(?:\s+([\d.]+))?`)

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

type ColumnSyntaxPreProcessor struct{}

func NewColumnSyntaxPreProcessor() *ColumnSyntaxPreProcessor {
	return &ColumnSyntaxPreProcessor{}
}

func (p *ColumnSyntaxPreProcessor) Process(markdownText string) (string, error) {
	if !strings.Contains(markdownText, columnListDelimiter) {
		return markdownText, nil
	}

	if err := p.validateAllColumnLists(markdownText); err != nil {
		return "", err
	}

	return markdownText, nil
}

func (p *ColumnSyntaxPreProcessor) validateAllColumnLists(markdownText string) error {
	blocks := p.extractColumnListBlocks(markdownText)
	for _, block := range blocks {
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
		if strings.TrimSpace(line) == columnListDelimiter {
			block := p.extractIndentedBlock(lines, i+1)
			blocks = append(blocks, block)
		}
	}

	return blocks
}

func (p *ColumnSyntaxPreProcessor) extractIndentedBlock(lines []string, startIndex int) string {
	if startIndex >= len(lines) {
		return ""
	}

	baseLevel := indentationLevel(lines[startIndex])
	baseSpaces := baseLevel * spacesPerNestingLevel
	var blockLines []string

	for _, line := range lines[startIndex:] {
		if strings.TrimSpace(line) == "" {
			blockLines = append(blockLines, line)
			continue
		}

		currentLevel := indentationLevel(line)
		if currentLevel < baseLevel {
			break
		}

		if len(line) >= baseSpaces {
			blockLines = append(blockLines, line[baseSpaces:])
		} else {
			blockLines = append(blockLines, line)
		}
	}

	return strings.Join(blockLines, "\n")
}

func (p *ColumnSyntaxPreProcessor) validateColumnListBlock(blockContent string) error {
	matches := columnPattern.FindAllStringSubmatch(blockContent, -1)
	columnCount := len(matches)

	if columnCount < minimumColumns {
		return &InsufficientColumnsError{Count: columnCount}
	}

	ratios := p.extractColumnRatios(matches)
	return p.validateRatioSum(ratios, columnCount)
}

func (p *ColumnSyntaxPreProcessor) extractColumnRatios(matches [][]string) []float64 {
	var ratios []float64
	for _, match := range matches {
		if len(match) > 1 && match[1] != "" && match[1] != "1" {
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

func indentationLevel(line string) int {
	leading := len(line) - len(strings.TrimLeft(line, " \t"))
	return leading / spacesPerNestingLevel
}