package easymodels

import (
	"github.com/mailru/easyjson/jlexer"
	"github.com/mailru/easyjson/jwriter"
)

type PetSliceAdapter []Pet

func (v PetSliceAdapter) MarshalEasyJSON(w *jwriter.Writer) {
	panic("implement me")
}

func (v *PetSliceAdapter) UnmarshalEasyJSON(l *jlexer.Lexer) {
	unmarshalPetSlice(l, v)
}

func unmarshalPetSlice(in *jlexer.Lexer, out *PetSliceAdapter) {
	in.Delim('[')
	if out == nil {
		if !in.IsDelim(']') {
			r := PetSliceAdapter(make([]Pet, 0, 8))
			out = &r
		} else {
			out = &PetSliceAdapter{}
		}
	} else {
		*out = (*out)[:0]
	}
	for !in.IsDelim(']') {
		getTypeLexer := *in
		switch d := fetchDiscriminator(&getTypeLexer); d {
		case "cat":
			var v1 *Cat
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(Cat)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		case "dog":
			var v1 *Dog
			if in.IsNull() {
				in.Skip()
				v1 = nil
			} else {
				if v1 == nil {
					v1 = new(Dog)
				}
				(*v1).UnmarshalEasyJSON(in)
			}
			*out = append(*out, v1)
			in.WantComma()
		default:
			in.AddError(&jlexer.LexerError{
				Reason: "unknown discriminator '" + d + "' for type Pet",
				Offset: in.GetPos(),
				Data:   string(in.Data[in.GetPos():]),
			})
		}
	}
	in.Delim(']')
}

func fetchDiscriminator(in *jlexer.Lexer) string {
	var discriminator string
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "kind":
			discriminator = in.String()
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	return discriminator
}
