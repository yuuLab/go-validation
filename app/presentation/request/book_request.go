package request

import validation "github.com/go-ozzo/ozzo-validation"

type Book struct {
	Title  string   `json:"title"`
	Author string   `json:"author"`
	Price  *float64 `json:"price"`
}

func (b Book) Validate() error {
	//NOTE: 日本語のエラー文が不要で、デフォルトの英語のエラー文で必要十分である場合、`.Error("xxx")`は不要でOK
	return validation.ValidateStruct(&b,
		validation.Field(
			&b.Title,
			validation.Required.Error("タイトルは必須項目です。"),
			validation.RuneLength(1, 50).Error("タイトルは 1文字 以上 50文字 以内です。"),
		),
		validation.Field(
			&b.Author,
			validation.Required.Error("著者名は必須項目です。"),
			validation.RuneLength(1, 50).Error("著者名は 1文字 以上 50文字 以内です。"),
		),
		validation.Field(
			&b.Price,
			validation.Required.Error("価格は必須項目です。"),
			validation.Max(1000000.0).Error("価格は 1,000,000円 以下で指定してください。"),
			validation.Min(1.0).Error("価格は 1円 以上で指定してください。"),
		),
	)
}
