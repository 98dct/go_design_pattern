package decorator

type TextProcessor interface {
	process(text string) string
}

type BaseProcessor struct {
}

func (p *BaseProcessor) process(text string) string {
	return text
}

//加密装饰器
type EncryptDecorator struct {
	processor TextProcessor
}

func (d *EncryptDecorator) process(text string) string {
	encryptedText := "encrypted(" + d.processor.process(text) + ")"
	return encryptedText
}

func NewEncryptDecorator(processor TextProcessor) TextProcessor {
	return &EncryptDecorator{processor: processor}
}

//格式化装饰器
type FormatDecorator struct {
	processor TextProcessor
}

func (d *FormatDecorator) process(text string) string {
	formatedText := "formatted(" + d.processor.process(text) + ")"
	return formatedText
}

func NewFormatDecorator(processor TextProcessor) TextProcessor {
	return &FormatDecorator{processor: processor}
}
