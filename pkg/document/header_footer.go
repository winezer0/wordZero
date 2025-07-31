// Package document 提供Word文档的页眉页脚操作功能
package document

import (
	"encoding/xml"
	"fmt"
	"strconv"
	"strings"
)

// HeaderFooterType 页眉页脚类型
type HeaderFooterType string

const (
	// HeaderFooterTypeDefault 默认页眉页脚
	HeaderFooterTypeDefault HeaderFooterType = "default"
	// HeaderFooterTypeFirst 首页页眉页脚
	HeaderFooterTypeFirst HeaderFooterType = "first"
	// HeaderFooterTypeEven 偶数页页眉页脚
	HeaderFooterTypeEven HeaderFooterType = "even"
)

// Header 页眉结构
type Header struct {
	XMLName     xml.Name     `xml:"w:hdr"`
	XmlnsWPC    string       `xml:"xmlns:wpc,attr"`
	XmlnsMC     string       `xml:"xmlns:mc,attr"`
	XmlnsO      string       `xml:"xmlns:o,attr"`
	XmlnsR      string       `xml:"xmlns:r,attr"`
	XmlnsM      string       `xml:"xmlns:m,attr"`
	XmlnsV      string       `xml:"xmlns:v,attr"`
	XmlnsWP14   string       `xml:"xmlns:wp14,attr"`
	XmlnsWP     string       `xml:"xmlns:wp,attr"`
	XmlnsW10    string       `xml:"xmlns:w10,attr"`
	XmlnsW      string       `xml:"xmlns:w,attr"`
	XmlnsW14    string       `xml:"xmlns:w14,attr"`
	XmlnsW15    string       `xml:"xmlns:w15,attr"`
	XmlnsWPG    string       `xml:"xmlns:wpg,attr"`
	XmlnsWPI    string       `xml:"xmlns:wpi,attr"`
	XmlnsWNE    string       `xml:"xmlns:wne,attr"`
	XmlnsWPS    string       `xml:"xmlns:wps,attr"`
	XmlnsWPSCD  string       `xml:"xmlns:wpsCustomData,attr"`
	MCIgnorable string       `xml:"mc:Ignorable,attr"`
	Paragraphs  []*Paragraph `xml:"w:p"`
}

// Footer 页脚结构
type Footer struct {
	XMLName     xml.Name     `xml:"w:ftr"`
	XmlnsWPC    string       `xml:"xmlns:wpc,attr"`
	XmlnsMC     string       `xml:"xmlns:mc,attr"`
	XmlnsO      string       `xml:"xmlns:o,attr"`
	XmlnsR      string       `xml:"xmlns:r,attr"`
	XmlnsM      string       `xml:"xmlns:m,attr"`
	XmlnsV      string       `xml:"xmlns:v,attr"`
	XmlnsWP14   string       `xml:"xmlns:wp14,attr"`
	XmlnsWP     string       `xml:"xmlns:wp,attr"`
	XmlnsW10    string       `xml:"xmlns:w10,attr"`
	XmlnsW      string       `xml:"xmlns:w,attr"`
	XmlnsW14    string       `xml:"xmlns:w14,attr"`
	XmlnsW15    string       `xml:"xmlns:w15,attr"`
	XmlnsWPG    string       `xml:"xmlns:wpg,attr"`
	XmlnsWPI    string       `xml:"xmlns:wpi,attr"`
	XmlnsWNE    string       `xml:"xmlns:wne,attr"`
	XmlnsWPS    string       `xml:"xmlns:wps,attr"`
	XmlnsWPSCD  string       `xml:"xmlns:wpsCustomData,attr"`
	MCIgnorable string       `xml:"mc:Ignorable,attr"`
	Paragraphs  []*Paragraph `xml:"w:p"`
}

// HeaderFooterReference 页眉页脚引用
type HeaderFooterReference struct {
	XMLName xml.Name `xml:"w:headerReference"`
	Type    string   `xml:"w:type,attr"`
	ID      string   `xml:"r:id,attr"`
}

// FooterReference 页脚引用
type FooterReference struct {
	XMLName xml.Name `xml:"w:footerReference"`
	Type    string   `xml:"w:type,attr"`
	ID      string   `xml:"r:id,attr"`
}

// TitlePage 首页不同设置
type TitlePage struct {
	XMLName xml.Name `xml:"w:titlePg"`
}

// PageNumber 页码字段
type PageNumber struct {
	XMLName xml.Name `xml:"w:fldSimple"`
	Instr   string   `xml:"w:instr,attr"`
	Text    *Text    `xml:"w:t,omitempty"`
}

// createStandardHeader 创建标准页眉结构
func createStandardHeader() *Header {
	return &Header{
		XmlnsWPC:    "http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas",
		XmlnsMC:     "http://schemas.openxmlformats.org/markup-compatibility/2006",
		XmlnsO:      "urn:schemas-microsoft-com:office:office",
		XmlnsR:      "http://schemas.openxmlformats.org/officeDocument/2006/relationships",
		XmlnsM:      "http://schemas.openxmlformats.org/officeDocument/2006/math",
		XmlnsV:      "urn:schemas-microsoft-com:vml",
		XmlnsWP14:   "http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing",
		XmlnsWP:     "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing",
		XmlnsW10:    "urn:schemas-microsoft-com:office:word",
		XmlnsW:      "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
		XmlnsW14:    "http://schemas.microsoft.com/office/word/2010/wordml",
		XmlnsW15:    "http://schemas.microsoft.com/office/word/2012/wordml",
		XmlnsWPG:    "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup",
		XmlnsWPI:    "http://schemas.microsoft.com/office/word/2010/wordprocessingInk",
		XmlnsWNE:    "http://schemas.microsoft.com/office/word/2006/wordml",
		XmlnsWPS:    "http://schemas.microsoft.com/office/word/2010/wordprocessingShape",
		XmlnsWPSCD:  "http://www.wps.cn/officeDocument/2013/wpsCustomData",
		MCIgnorable: "w14 w15 wp14",
		Paragraphs:  make([]*Paragraph, 0),
	}
}

// createStandardFooter 创建标准页脚结构
func createStandardFooter() *Footer {
	return &Footer{
		XmlnsWPC:    "http://schemas.microsoft.com/office/word/2010/wordprocessingCanvas",
		XmlnsMC:     "http://schemas.openxmlformats.org/markup-compatibility/2006",
		XmlnsO:      "urn:schemas-microsoft-com:office:office",
		XmlnsR:      "http://schemas.openxmlformats.org/officeDocument/2006/relationships",
		XmlnsM:      "http://schemas.openxmlformats.org/officeDocument/2006/math",
		XmlnsV:      "urn:schemas-microsoft-com:vml",
		XmlnsWP14:   "http://schemas.microsoft.com/office/word/2010/wordprocessingDrawing",
		XmlnsWP:     "http://schemas.openxmlformats.org/drawingml/2006/wordprocessingDrawing",
		XmlnsW10:    "urn:schemas-microsoft-com:office:word",
		XmlnsW:      "http://schemas.openxmlformats.org/wordprocessingml/2006/main",
		XmlnsW14:    "http://schemas.microsoft.com/office/word/2010/wordml",
		XmlnsW15:    "http://schemas.microsoft.com/office/word/2012/wordml",
		XmlnsWPG:    "http://schemas.microsoft.com/office/word/2010/wordprocessingGroup",
		XmlnsWPI:    "http://schemas.microsoft.com/office/word/2010/wordprocessingInk",
		XmlnsWNE:    "http://schemas.microsoft.com/office/word/2006/wordml",
		XmlnsWPS:    "http://schemas.microsoft.com/office/word/2010/wordprocessingShape",
		XmlnsWPSCD:  "http://www.wps.cn/officeDocument/2013/wpsCustomData",
		MCIgnorable: "w14 w15 wp14",
		Paragraphs:  make([]*Paragraph, 0),
	}
}

// createPageNumberRuns 创建页码域代码的Run集合
func createPageNumberRuns() []Run {
	return []Run{
		{
			FieldChar: &FieldChar{
				FieldCharType: "begin",
			},
		},
		{
			InstrText: &InstrText{
				Space:   "preserve",
				Content: " PAGE  \\* MERGEFORMAT ",
			},
		},
		{
			FieldChar: &FieldChar{
				FieldCharType: "separate",
			},
		},
		{
			Text: Text{
				Content: "1",
			},
		},
		{
			FieldChar: &FieldChar{
				FieldCharType: "end",
			},
		},
	}
}

// getFileNameForType 获取页眉页脚文件名
func getFileNameForType(typePrefix string, headerType HeaderFooterType) string {
	switch headerType {
	case HeaderFooterTypeDefault:
		return fmt.Sprintf("%s1.xml", typePrefix)
	case HeaderFooterTypeFirst:
		return fmt.Sprintf("%sfirst.xml", typePrefix)
	case HeaderFooterTypeEven:
		return fmt.Sprintf("%seven.xml", typePrefix)
	default:
		return fmt.Sprintf("%s1.xml", typePrefix)
	}
}

// AddHeader 添加页眉
func (d *Document) AddHeader(headerType HeaderFooterType, text string) error {
	header := createStandardHeader()

	// 创建页眉段落
	paragraph := &Paragraph{}
	if text != "" {
		run := Run{
			Text: Text{
				Content: text,
				Space:   "preserve",
			},
		}
		paragraph.Runs = append(paragraph.Runs, run)
	}
	header.Paragraphs = append(header.Paragraphs, paragraph)

	// 生成关系ID
	headerID := fmt.Sprintf("rId%d", len(d.documentRelationships.Relationships)+2) // +2因为rId1保留给styles

	// 序列化页眉
	headerXML, err := xml.MarshalIndent(header, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化页眉失败: %v", err)
	}

	// 添加XML声明
	fullXML := append([]byte(xml.Header), headerXML...)

	// 获取文件名
	fileName := getFileNameForType("header", headerType)
	headerPartName := fmt.Sprintf("word/%s", fileName)

	// 存储页眉内容
	d.parts[headerPartName] = fullXML

	// 添加关系到文档关系
	relationship := Relationship{
		ID:     headerID,
		Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/header",
		Target: fileName,
	}
	d.documentRelationships.Relationships = append(d.documentRelationships.Relationships, relationship)

	// 添加内容类型
	d.addContentType(headerPartName, "application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml")

	// 更新节属性
	d.addHeaderReference(headerType, headerID)

	return nil
}

// AddFooter 添加页脚
func (d *Document) AddFooter(footerType HeaderFooterType, text string) error {
	footer := createStandardFooter()

	// 创建页脚段落
	paragraph := &Paragraph{}
	if text != "" {
		run := Run{
			Text: Text{
				Content: text,
				Space:   "preserve",
			},
		}
		paragraph.Runs = append(paragraph.Runs, run)
	}
	footer.Paragraphs = append(footer.Paragraphs, paragraph)

	// 生成关系ID
	footerID := fmt.Sprintf("rId%d", len(d.documentRelationships.Relationships)+2) // +2因为rId1保留给styles

	// 序列化页脚
	footerXML, err := xml.MarshalIndent(footer, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化页脚失败: %v", err)
	}

	// 添加XML声明
	fullXML := append([]byte(xml.Header), footerXML...)

	// 获取文件名
	fileName := getFileNameForType("footer", footerType)
	footerPartName := fmt.Sprintf("word/%s", fileName)

	// 存储页脚内容
	d.parts[footerPartName] = fullXML

	// 添加关系到文档关系
	relationship := Relationship{
		ID:     footerID,
		Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footer",
		Target: fileName,
	}
	d.documentRelationships.Relationships = append(d.documentRelationships.Relationships, relationship)

	// 添加内容类型
	d.addContentType(footerPartName, "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml")

	// 更新节属性
	d.addFooterReference(footerType, footerID)

	return nil
}

// AddHeaderWithPageNumber 添加带页码的页眉
func (d *Document) AddHeaderWithPageNumber(headerType HeaderFooterType, text string, showPageNum bool) error {
	header := createStandardHeader()

	// 创建页眉段落
	paragraph := &Paragraph{}

	if text != "" {
		run := Run{
			Text: Text{
				Content: text,
				Space:   "preserve",
			},
		}
		paragraph.Runs = append(paragraph.Runs, run)
	}

	if showPageNum {
		// 添加"第"字
		pageNumRun := Run{
			Text: Text{
				Content: " 第 ",
				Space:   "preserve",
			},
		}
		paragraph.Runs = append(paragraph.Runs, pageNumRun)

		// 添加页码域代码
		pageNumberRuns := createPageNumberRuns()
		paragraph.Runs = append(paragraph.Runs, pageNumberRuns...)

		// 添加"页"字
		pageNumRun2 := Run{
			Text: Text{
				Content: " 页",
				Space:   "preserve",
			},
		}
		paragraph.Runs = append(paragraph.Runs, pageNumRun2)
	}

	header.Paragraphs = append(header.Paragraphs, paragraph)

	// 生成关系ID
	headerID := fmt.Sprintf("rId%d", len(d.documentRelationships.Relationships)+2) // +2因为rId1保留给styles

	// 序列化页眉
	headerXML, err := xml.MarshalIndent(header, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化页眉失败: %v", err)
	}

	// 添加XML声明
	fullXML := append([]byte(xml.Header), headerXML...)

	// 获取文件名
	fileName := getFileNameForType("header", headerType)
	headerPartName := fmt.Sprintf("word/%s", fileName)

	// 存储页眉内容
	d.parts[headerPartName] = fullXML

	// 添加关系到文档关系
	relationship := Relationship{
		ID:     headerID,
		Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/header",
		Target: fileName,
	}
	d.documentRelationships.Relationships = append(d.documentRelationships.Relationships, relationship)

	// 添加内容类型
	d.addContentType(headerPartName, "application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml")

	// 更新节属性
	d.addHeaderReference(headerType, headerID)

	return nil
}

// AddFooterWithPageNumber 添加带页码的页脚
func (d *Document) AddFooterWithPageNumber(footerType HeaderFooterType, text string, showPageNum bool) error {
	footer := createStandardFooter()

	// 创建页脚段落
	paragraph := &Paragraph{}
	paragraph.SetAlignment(AlignCenter)
	if text != "" {
		run := Run{
			Text: Text{
				Content: text,
				Space:   "preserve",
			},
		}
		paragraph.Runs = append(paragraph.Runs, run)
	}

	if showPageNum {
		// 添加"第"字
		pageNumRun := Run{
			Text: Text{
				Content: " 第 ",
				Space:   "preserve",
			},
		}
		paragraph.Runs = append(paragraph.Runs, pageNumRun)

		// 添加页码域代码
		pageNumberRuns := createPageNumberRuns()
		paragraph.Runs = append(paragraph.Runs, pageNumberRuns...)

		// 添加"页"字
		pageNumRun2 := Run{
			Text: Text{
				Content: " 页",
				Space:   "preserve",
			},
		}
		paragraph.Runs = append(paragraph.Runs, pageNumRun2)
	}

	footer.Paragraphs = append(footer.Paragraphs, paragraph)

	// 生成关系ID
	footerID := fmt.Sprintf("rId%d", len(d.documentRelationships.Relationships)+2) // +2因为rId1保留给styles

	// 序列化页脚
	footerXML, err := xml.MarshalIndent(footer, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化页脚失败: %v", err)
	}

	// 添加XML声明
	fullXML := append([]byte(xml.Header), footerXML...)

	// 获取文件名
	fileName := getFileNameForType("footer", footerType)
	footerPartName := fmt.Sprintf("word/%s", fileName)

	// 存储页脚内容
	d.parts[footerPartName] = fullXML

	// 添加关系到文档关系
	relationship := Relationship{
		ID:     footerID,
		Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footer",
		Target: fileName,
	}
	d.documentRelationships.Relationships = append(d.documentRelationships.Relationships, relationship)

	// 添加内容类型
	d.addContentType(footerPartName, "application/vnd.openxmlformats-officedocument.wordprocessingml.footer+xml")

	// 更新节属性
	d.addFooterReference(footerType, footerID)

	return nil
}

// SetDifferentFirstPage 设置首页不同
func (d *Document) SetDifferentFirstPage(different bool) {
	sectPr := d.getSectionPropertiesForHeaderFooter()
	if different {
		sectPr.TitlePage = &TitlePage{}
	} else {
		sectPr.TitlePage = nil
	}
}

// addHeaderReference 添加页眉引用到节属性
func (d *Document) addHeaderReference(headerType HeaderFooterType, headerID string) {
	sectPr := d.getSectionPropertiesForHeaderFooter()

	// 确保设置关系命名空间
	if sectPr.XmlnsR == "" {
		sectPr.XmlnsR = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"
	}

	headerRef := &HeaderFooterReference{
		Type: string(headerType),
		ID:   headerID,
	}

	sectPr.HeaderReferences = append(sectPr.HeaderReferences, headerRef)
}

// addFooterReference 添加页脚引用到节属性
func (d *Document) addFooterReference(footerType HeaderFooterType, footerID string) {
	sectPr := d.getSectionPropertiesForHeaderFooter()

	// 确保设置关系命名空间
	if sectPr.XmlnsR == "" {
		sectPr.XmlnsR = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"
	}

	footerRef := &FooterReference{
		Type: string(footerType),
		ID:   footerID,
	}

	sectPr.FooterReferences = append(sectPr.FooterReferences, footerRef)
}

// getSectionPropertiesForHeaderFooter 获取或创建带页眉页脚支持的节属性
func (d *Document) getSectionPropertiesForHeaderFooter() *SectionProperties {
	// 查找文档中是否已存在节属性
	for _, element := range d.Body.Elements {
		if sectPr, ok := element.(*SectionProperties); ok {
			// 确保设置了关系命名空间
			if sectPr.XmlnsR == "" {
				sectPr.XmlnsR = "http://schemas.openxmlformats.org/officeDocument/2006/relationships"
			}
			return sectPr
		}
	}

	// 如果不存在，创建新的节属性
	sectPr := &SectionProperties{
		XMLName: xml.Name{Local: "w:sectPr"},
		XmlnsR:  "http://schemas.openxmlformats.org/officeDocument/2006/relationships",
		PageNumType: &PageNumType{
			Fmt: "decimal",
		},
		Columns: &Columns{
			Space: "720",
			Num:   "1",
		},
	}
	d.Body.Elements = append(d.Body.Elements, sectPr)
	return sectPr
}

// addContentType 添加内容类型
func (d *Document) addContentType(partName, contentType string) {
	// 检查是否已存在
	for _, override := range d.contentTypes.Overrides {
		if override.PartName == "/"+partName {
			return
		}
	}

	// 添加新的内容类型覆盖
	override := Override{
		PartName:    "/" + partName,
		ContentType: contentType,
	}
	d.contentTypes.Overrides = append(d.contentTypes.Overrides, override)
}

// AddStyleHeader 添加带有样式的页眉
func (d *Document) AddStyleHeader(headerType HeaderFooterType, text, redText string, format *TextFormat) error {
	header := createStandardHeader()
	// 创建运行属性
	runProps := setFormat(format)
	// 创建页眉段落
	paragraph := &Paragraph{}
	paragraph.SetAlignment(AlignCenter)
	if text != "" {
		if strings.Contains(text, "\n") {
			runs := strings.Split(text, "\n")
			for i, runText := range runs {
				run := Run{
					Properties: runProps,
					Text: Text{
						Content: runText,
						Space:   "preserve",
					},
				}
				paragraph.Runs = append(paragraph.Runs, run)
				if i == len(runs)-1 {
					break
				}
				paragraph.AddLineBreak("") // 换行
			}

			if redText != "" {
				// 设置成红色
				format.FontColor = "FF0000"
				paragraph.Runs = append(paragraph.Runs, Run{
					Properties: setFormat(format),
					Text: Text{
						Content: redText,
						Space:   "preserve",
					},
				})
				paragraph.AddLineBreak("") // 换行
			}
		} else {
			run := Run{
				Properties: runProps,
				Text: Text{
					Content: text,
					Space:   "preserve",
				},
			}
			paragraph.Runs = append(paragraph.Runs, run)
		}
	}
	header.Paragraphs = append(header.Paragraphs, paragraph)

	// 生成关系ID
	headerID := fmt.Sprintf("rId%d", len(d.documentRelationships.Relationships)+2) // +2因为rId1保留给styles

	// 序列化页眉
	headerXML, err := xml.MarshalIndent(header, "", "  ")
	if err != nil {
		return fmt.Errorf("序列化页眉失败: %v", err)
	}

	// 添加XML声明
	fullXML := append([]byte(xml.Header), headerXML...)

	// 获取文件名
	fileName := getFileNameForType("header", headerType)
	headerPartName := fmt.Sprintf("word/%s", fileName)

	// 存储页眉内容
	d.parts[headerPartName] = fullXML

	// 添加关系到文档关系
	relationship := Relationship{
		ID:     headerID,
		Type:   "http://schemas.openxmlformats.org/officeDocument/2006/relationships/header",
		Target: fileName,
	}
	d.documentRelationships.Relationships = append(d.documentRelationships.Relationships, relationship)

	// 添加内容类型
	d.addContentType(headerPartName, "application/vnd.openxmlformats-officedocument.wordprocessingml.header+xml")

	// 更新节属性
	d.addHeaderReference(headerType, headerID)

	return nil
}

func setFormat(format *TextFormat) *RunProperties {
	runProps := &RunProperties{}
	if format != nil {
		if format.FontFamily != "" {
			runProps.FontFamily = &FontFamily{ASCII: format.FontFamily}
		}

		if format.Bold {
			runProps.Bold = &Bold{}
		}

		if format.Italic {
			runProps.Italic = &Italic{}
		}

		if format.FontColor != "" {
			color := strings.TrimPrefix(format.FontColor, "#")
			runProps.Color = &Color{Val: color}
		}

		if format.FontSize > 0 {
			runProps.FontSize = &FontSize{Val: strconv.Itoa(format.FontSize * 2)}
		}
	}
	return runProps
}
