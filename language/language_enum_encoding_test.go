// Code generated by go-enum-encoding; DO NOT EDIT.

package language

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"slices"
	"testing"
)

func ExampleLanguage_MarshalText() {
	for _, v := range []Language{Unknown, Abkhaz, Afar, Afrikaans, Akan, Albanian, Amharic, Arabic, Aragonese, Armenian, Assamese, Avaric, Avestan, Aymara, Azerbaijani, Bambara, Bashkir, Basque, Belarusian, Bengali, Bihari, Bislama, Bosnian, Breton, Bulgarian, Burmese, Catalan, Chamorro, Chechen, Chichewa, Chinese, ChineseTraditional, Chuvash, Cornish, Corsican, Cree, Croatian, Czech, Danish, Divehi, Dutch, Dzongkha, English, Esperanto, Estonian, Ewe, Faroese, Fijian, Finnish, French, Fula, Galician, Ganda, Georgian, German, Greek, Guaraní, Gujarati, Haitian, Hausa, Hebrew, Herero, Hindi, Hiri, Hungarian, Icelandic, Ido, Igbo, Indonesian, Interlingua, Interlingue, Inuktitut, Inupiaq, Irish, Italian, Japanese, Javanese, Kalaallisut, Kannada, Kanuri, Kashmiri, Kazakh, Khmer, Kikuyu, Kinyarwanda, Kirundi, Komi, Kongo, Korean, Kurdish, Kwanyama, Kyrgyz, Lao, Latin, Latvian, Limburgish, Lingala, Lithuanian, Luba, Luxembourgish, Macedonian, Malagasy, Malay, Malayalam, Maltese, Manx, Māori, Marathi, Marshallese, Mongolian, Nauru, Navajo, Ndonga, Nepali, Northern, NorthernNdebele, Norwegian, NorwegianBokmal, NorwegianNynorsk, Nuosu, Occitan, Ojibwe, Old, Oriya, Oromo, Ossetian, Pāli, Panjabi, Pashto, Persian, Polish, Portuguese, Quechua, Romanian, Romansh, Russian, Samoan, Sango, Sanskrit, Sardinian, Scottish, Serbian, Shona, Sindhi, Sinhala, Slovak, Slovene, Somali, SouthernNdebele, SouthernSotho, Spanish, Sundanese, Swahili, Swati, Swedish, Tagalog, Tahitian, Tajik, Tamil, Tatar, Telugu, Thai, Tibetan, Tigrinya, Tonga, Tsonga, Tswana, Turkish, Turkmen, Twi, Ukrainian, Urdu, Uyghur, Uzbek, Venda, Vietnamese, Volapük, Walloon, Welsh, Western, Wolof, Xhosa, Yiddish, Yoruba, Zhuang, Zulu} {
		b, _ := v.MarshalText()
		fmt.Printf("%s ", string(b))
	}
	// Output:  ab aa af ak sq am ar an hy as av ae ay az bm ba eu be bn bh bi bs br bg my ca ch ce ny zh zh_Hant cv kw co cr hr cs da dv nl dz en eo et ee fo fj fi fr ff gl lg ka de el gn gu ht ha he hz hi ho hu is io ig id ia ie iu ik ga it ja jv kl kn kr ks kk km ki rw rn kv kg ko ku kj ky lo la lv li ln lt lu lb mk mg ms ml mt gv mi mr mh mn na nv ng ne se nd no nb nn ii oc oj cu or om os pi pa ps fa pl pt qu ro rm ru sm sg sa sc gd sr sn sd si sk sl so nr st es su sw ss sv tl ty tg ta tt te th bo ti to ts tn tr tk tw uk ur ug uz ve vi vo wa cy fy wo xh yi yo za zu
}

func ExampleLanguage_UnmarshalText() {
	for _, s := range []string{"", "ab", "aa", "af", "ak", "sq", "am", "ar", "an", "hy", "as", "av", "ae", "ay", "az", "bm", "ba", "eu", "be", "bn", "bh", "bi", "bs", "br", "bg", "my", "ca", "ch", "ce", "ny", "zh", "zh_Hant", "cv", "kw", "co", "cr", "hr", "cs", "da", "dv", "nl", "dz", "en", "eo", "et", "ee", "fo", "fj", "fi", "fr", "ff", "gl", "lg", "ka", "de", "el", "gn", "gu", "ht", "ha", "he", "hz", "hi", "ho", "hu", "is", "io", "ig", "id", "ia", "ie", "iu", "ik", "ga", "it", "ja", "jv", "kl", "kn", "kr", "ks", "kk", "km", "ki", "rw", "rn", "kv", "kg", "ko", "ku", "kj", "ky", "lo", "la", "lv", "li", "ln", "lt", "lu", "lb", "mk", "mg", "ms", "ml", "mt", "gv", "mi", "mr", "mh", "mn", "na", "nv", "ng", "ne", "se", "nd", "no", "nb", "nn", "ii", "oc", "oj", "cu", "or", "om", "os", "pi", "pa", "ps", "fa", "pl", "pt", "qu", "ro", "rm", "ru", "sm", "sg", "sa", "sc", "gd", "sr", "sn", "sd", "si", "sk", "sl", "so", "nr", "st", "es", "su", "sw", "ss", "sv", "tl", "ty", "tg", "ta", "tt", "te", "th", "bo", "ti", "to", "ts", "tn", "tr", "tk", "tw", "uk", "ur", "ug", "uz", "ve", "vi", "vo", "wa", "cy", "fy", "wo", "xh", "yi", "yo", "za", "zu"} {
		var v Language
		if err := (&v).UnmarshalText([]byte(s)); err != nil {
			fmt.Println(err)
		}
	}
}

func TestLanguage_MarshalText_UnmarshalText(t *testing.T) {
	for _, v := range []Language{Unknown, Abkhaz, Afar, Afrikaans, Akan, Albanian, Amharic, Arabic, Aragonese, Armenian, Assamese, Avaric, Avestan, Aymara, Azerbaijani, Bambara, Bashkir, Basque, Belarusian, Bengali, Bihari, Bislama, Bosnian, Breton, Bulgarian, Burmese, Catalan, Chamorro, Chechen, Chichewa, Chinese, ChineseTraditional, Chuvash, Cornish, Corsican, Cree, Croatian, Czech, Danish, Divehi, Dutch, Dzongkha, English, Esperanto, Estonian, Ewe, Faroese, Fijian, Finnish, French, Fula, Galician, Ganda, Georgian, German, Greek, Guaraní, Gujarati, Haitian, Hausa, Hebrew, Herero, Hindi, Hiri, Hungarian, Icelandic, Ido, Igbo, Indonesian, Interlingua, Interlingue, Inuktitut, Inupiaq, Irish, Italian, Japanese, Javanese, Kalaallisut, Kannada, Kanuri, Kashmiri, Kazakh, Khmer, Kikuyu, Kinyarwanda, Kirundi, Komi, Kongo, Korean, Kurdish, Kwanyama, Kyrgyz, Lao, Latin, Latvian, Limburgish, Lingala, Lithuanian, Luba, Luxembourgish, Macedonian, Malagasy, Malay, Malayalam, Maltese, Manx, Māori, Marathi, Marshallese, Mongolian, Nauru, Navajo, Ndonga, Nepali, Northern, NorthernNdebele, Norwegian, NorwegianBokmal, NorwegianNynorsk, Nuosu, Occitan, Ojibwe, Old, Oriya, Oromo, Ossetian, Pāli, Panjabi, Pashto, Persian, Polish, Portuguese, Quechua, Romanian, Romansh, Russian, Samoan, Sango, Sanskrit, Sardinian, Scottish, Serbian, Shona, Sindhi, Sinhala, Slovak, Slovene, Somali, SouthernNdebele, SouthernSotho, Spanish, Sundanese, Swahili, Swati, Swedish, Tagalog, Tahitian, Tajik, Tamil, Tatar, Telugu, Thai, Tibetan, Tigrinya, Tonga, Tsonga, Tswana, Turkish, Turkmen, Twi, Ukrainian, Urdu, Uyghur, Uzbek, Venda, Vietnamese, Volapük, Walloon, Welsh, Western, Wolof, Xhosa, Yiddish, Yoruba, Zhuang, Zulu} {
		b, err := v.MarshalText()
		if err != nil {
			t.Errorf("cannot encode: %s", err)
		}

		var d Language
		if err := (&d).UnmarshalText(b); err != nil {
			t.Errorf("cannot decode: %s", err)
		}

		if d != v {
			t.Errorf("exp(%v) != got(%v)", v, d)
		}
	}

	t.Run("when unknown value, then error", func(t *testing.T) {
		s := `something`
		var v Language
		err := (&v).UnmarshalText([]byte(s))
		if err == nil {
			t.Error("must be error")
		}
		if !errors.Is(err, ErrUnknownLanguage) {
			t.Error("wrong error", err)
		}
	})
}

func TestLanguage_JSON(t *testing.T) {
	type V struct {
		Values []Language `json:"values"`
	}

	values := []Language{Unknown, Abkhaz, Afar, Afrikaans, Akan, Albanian, Amharic, Arabic, Aragonese, Armenian, Assamese, Avaric, Avestan, Aymara, Azerbaijani, Bambara, Bashkir, Basque, Belarusian, Bengali, Bihari, Bislama, Bosnian, Breton, Bulgarian, Burmese, Catalan, Chamorro, Chechen, Chichewa, Chinese, ChineseTraditional, Chuvash, Cornish, Corsican, Cree, Croatian, Czech, Danish, Divehi, Dutch, Dzongkha, English, Esperanto, Estonian, Ewe, Faroese, Fijian, Finnish, French, Fula, Galician, Ganda, Georgian, German, Greek, Guaraní, Gujarati, Haitian, Hausa, Hebrew, Herero, Hindi, Hiri, Hungarian, Icelandic, Ido, Igbo, Indonesian, Interlingua, Interlingue, Inuktitut, Inupiaq, Irish, Italian, Japanese, Javanese, Kalaallisut, Kannada, Kanuri, Kashmiri, Kazakh, Khmer, Kikuyu, Kinyarwanda, Kirundi, Komi, Kongo, Korean, Kurdish, Kwanyama, Kyrgyz, Lao, Latin, Latvian, Limburgish, Lingala, Lithuanian, Luba, Luxembourgish, Macedonian, Malagasy, Malay, Malayalam, Maltese, Manx, Māori, Marathi, Marshallese, Mongolian, Nauru, Navajo, Ndonga, Nepali, Northern, NorthernNdebele, Norwegian, NorwegianBokmal, NorwegianNynorsk, Nuosu, Occitan, Ojibwe, Old, Oriya, Oromo, Ossetian, Pāli, Panjabi, Pashto, Persian, Polish, Portuguese, Quechua, Romanian, Romansh, Russian, Samoan, Sango, Sanskrit, Sardinian, Scottish, Serbian, Shona, Sindhi, Sinhala, Slovak, Slovene, Somali, SouthernNdebele, SouthernSotho, Spanish, Sundanese, Swahili, Swati, Swedish, Tagalog, Tahitian, Tajik, Tamil, Tatar, Telugu, Thai, Tibetan, Tigrinya, Tonga, Tsonga, Tswana, Turkish, Turkmen, Twi, Ukrainian, Urdu, Uyghur, Uzbek, Venda, Vietnamese, Volapük, Walloon, Welsh, Western, Wolof, Xhosa, Yiddish, Yoruba, Zhuang, Zulu}

	var v V
	s := `{"values":["","ab","aa","af","ak","sq","am","ar","an","hy","as","av","ae","ay","az","bm","ba","eu","be","bn","bh","bi","bs","br","bg","my","ca","ch","ce","ny","zh","zh_Hant","cv","kw","co","cr","hr","cs","da","dv","nl","dz","en","eo","et","ee","fo","fj","fi","fr","ff","gl","lg","ka","de","el","gn","gu","ht","ha","he","hz","hi","ho","hu","is","io","ig","id","ia","ie","iu","ik","ga","it","ja","jv","kl","kn","kr","ks","kk","km","ki","rw","rn","kv","kg","ko","ku","kj","ky","lo","la","lv","li","ln","lt","lu","lb","mk","mg","ms","ml","mt","gv","mi","mr","mh","mn","na","nv","ng","ne","se","nd","no","nb","nn","ii","oc","oj","cu","or","om","os","pi","pa","ps","fa","pl","pt","qu","ro","rm","ru","sm","sg","sa","sc","gd","sr","sn","sd","si","sk","sl","so","nr","st","es","su","sw","ss","sv","tl","ty","tg","ta","tt","te","th","bo","ti","to","ts","tn","tr","tk","tw","uk","ur","ug","uz","ve","vi","vo","wa","cy","fy","wo","xh","yi","yo","za","zu"]}`
	json.Unmarshal([]byte(s), &v)

	if len(v.Values) != len(values) {
		t.Errorf("cannot decode: %d", len(v.Values))
	}
	if !slices.Equal(v.Values, values) {
		t.Errorf("wrong decoded: %v", v.Values)
	}

	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("cannot encode: %s", err)
	}
	if string(b) != s {
		t.Errorf("wrong encoded: %s != %s", string(b), s)
	}

	t.Run("when unknown value, then error", func(t *testing.T) {
		s := `{"values":["something"]}`
		var v V
		err := json.Unmarshal([]byte(s), &v)
		if err == nil {
			t.Error("must be error")
		}
		if !errors.Is(err, ErrUnknownLanguage) {
			t.Error("wrong error", err)
		}
	})
}

func BenchmarkLanguage_UnmarshalText(b *testing.B) {
	vb := seq_bytes_Language[rand.Intn(len(seq_bytes_Language))]

	var x Language

	for b.Loop() {
		_ = x.UnmarshalText(vb)
	}
}

func BenchmarkLanguage_AppendText(b *testing.B) {
	bb := make([]byte, 10, 1000)

	vs := []Language{Unknown, Abkhaz, Afar, Afrikaans, Akan, Albanian, Amharic, Arabic, Aragonese, Armenian, Assamese, Avaric, Avestan, Aymara, Azerbaijani, Bambara, Bashkir, Basque, Belarusian, Bengali, Bihari, Bislama, Bosnian, Breton, Bulgarian, Burmese, Catalan, Chamorro, Chechen, Chichewa, Chinese, ChineseTraditional, Chuvash, Cornish, Corsican, Cree, Croatian, Czech, Danish, Divehi, Dutch, Dzongkha, English, Esperanto, Estonian, Ewe, Faroese, Fijian, Finnish, French, Fula, Galician, Ganda, Georgian, German, Greek, Guaraní, Gujarati, Haitian, Hausa, Hebrew, Herero, Hindi, Hiri, Hungarian, Icelandic, Ido, Igbo, Indonesian, Interlingua, Interlingue, Inuktitut, Inupiaq, Irish, Italian, Japanese, Javanese, Kalaallisut, Kannada, Kanuri, Kashmiri, Kazakh, Khmer, Kikuyu, Kinyarwanda, Kirundi, Komi, Kongo, Korean, Kurdish, Kwanyama, Kyrgyz, Lao, Latin, Latvian, Limburgish, Lingala, Lithuanian, Luba, Luxembourgish, Macedonian, Malagasy, Malay, Malayalam, Maltese, Manx, Māori, Marathi, Marshallese, Mongolian, Nauru, Navajo, Ndonga, Nepali, Northern, NorthernNdebele, Norwegian, NorwegianBokmal, NorwegianNynorsk, Nuosu, Occitan, Ojibwe, Old, Oriya, Oromo, Ossetian, Pāli, Panjabi, Pashto, Persian, Polish, Portuguese, Quechua, Romanian, Romansh, Russian, Samoan, Sango, Sanskrit, Sardinian, Scottish, Serbian, Shona, Sindhi, Sinhala, Slovak, Slovene, Somali, SouthernNdebele, SouthernSotho, Spanish, Sundanese, Swahili, Swati, Swedish, Tagalog, Tahitian, Tajik, Tamil, Tatar, Telugu, Thai, Tibetan, Tigrinya, Tonga, Tsonga, Tswana, Turkish, Turkmen, Twi, Ukrainian, Urdu, Uyghur, Uzbek, Venda, Vietnamese, Volapük, Walloon, Welsh, Western, Wolof, Xhosa, Yiddish, Yoruba, Zhuang, Zulu}
	v := vs[rand.Intn(len(vs))]

	for b.Loop() {
		_, _ = v.AppendText(bb)
	}
}

func BenchmarkLanguage_MarshalText(b *testing.B) {
	vs := []Language{Unknown, Abkhaz, Afar, Afrikaans, Akan, Albanian, Amharic, Arabic, Aragonese, Armenian, Assamese, Avaric, Avestan, Aymara, Azerbaijani, Bambara, Bashkir, Basque, Belarusian, Bengali, Bihari, Bislama, Bosnian, Breton, Bulgarian, Burmese, Catalan, Chamorro, Chechen, Chichewa, Chinese, ChineseTraditional, Chuvash, Cornish, Corsican, Cree, Croatian, Czech, Danish, Divehi, Dutch, Dzongkha, English, Esperanto, Estonian, Ewe, Faroese, Fijian, Finnish, French, Fula, Galician, Ganda, Georgian, German, Greek, Guaraní, Gujarati, Haitian, Hausa, Hebrew, Herero, Hindi, Hiri, Hungarian, Icelandic, Ido, Igbo, Indonesian, Interlingua, Interlingue, Inuktitut, Inupiaq, Irish, Italian, Japanese, Javanese, Kalaallisut, Kannada, Kanuri, Kashmiri, Kazakh, Khmer, Kikuyu, Kinyarwanda, Kirundi, Komi, Kongo, Korean, Kurdish, Kwanyama, Kyrgyz, Lao, Latin, Latvian, Limburgish, Lingala, Lithuanian, Luba, Luxembourgish, Macedonian, Malagasy, Malay, Malayalam, Maltese, Manx, Māori, Marathi, Marshallese, Mongolian, Nauru, Navajo, Ndonga, Nepali, Northern, NorthernNdebele, Norwegian, NorwegianBokmal, NorwegianNynorsk, Nuosu, Occitan, Ojibwe, Old, Oriya, Oromo, Ossetian, Pāli, Panjabi, Pashto, Persian, Polish, Portuguese, Quechua, Romanian, Romansh, Russian, Samoan, Sango, Sanskrit, Sardinian, Scottish, Serbian, Shona, Sindhi, Sinhala, Slovak, Slovene, Somali, SouthernNdebele, SouthernSotho, Spanish, Sundanese, Swahili, Swati, Swedish, Tagalog, Tahitian, Tajik, Tamil, Tatar, Telugu, Thai, Tibetan, Tigrinya, Tonga, Tsonga, Tswana, Turkish, Turkmen, Twi, Ukrainian, Urdu, Uyghur, Uzbek, Venda, Vietnamese, Volapük, Walloon, Welsh, Western, Wolof, Xhosa, Yiddish, Yoruba, Zhuang, Zulu}
	v := vs[rand.Intn(len(vs))]

	for b.Loop() {
		_, _ = v.MarshalText()
	}
}

func TestLanguage_String(t *testing.T) {
	values := []Language{Unknown, Abkhaz, Afar, Afrikaans, Akan, Albanian, Amharic, Arabic, Aragonese, Armenian, Assamese, Avaric, Avestan, Aymara, Azerbaijani, Bambara, Bashkir, Basque, Belarusian, Bengali, Bihari, Bislama, Bosnian, Breton, Bulgarian, Burmese, Catalan, Chamorro, Chechen, Chichewa, Chinese, ChineseTraditional, Chuvash, Cornish, Corsican, Cree, Croatian, Czech, Danish, Divehi, Dutch, Dzongkha, English, Esperanto, Estonian, Ewe, Faroese, Fijian, Finnish, French, Fula, Galician, Ganda, Georgian, German, Greek, Guaraní, Gujarati, Haitian, Hausa, Hebrew, Herero, Hindi, Hiri, Hungarian, Icelandic, Ido, Igbo, Indonesian, Interlingua, Interlingue, Inuktitut, Inupiaq, Irish, Italian, Japanese, Javanese, Kalaallisut, Kannada, Kanuri, Kashmiri, Kazakh, Khmer, Kikuyu, Kinyarwanda, Kirundi, Komi, Kongo, Korean, Kurdish, Kwanyama, Kyrgyz, Lao, Latin, Latvian, Limburgish, Lingala, Lithuanian, Luba, Luxembourgish, Macedonian, Malagasy, Malay, Malayalam, Maltese, Manx, Māori, Marathi, Marshallese, Mongolian, Nauru, Navajo, Ndonga, Nepali, Northern, NorthernNdebele, Norwegian, NorwegianBokmal, NorwegianNynorsk, Nuosu, Occitan, Ojibwe, Old, Oriya, Oromo, Ossetian, Pāli, Panjabi, Pashto, Persian, Polish, Portuguese, Quechua, Romanian, Romansh, Russian, Samoan, Sango, Sanskrit, Sardinian, Scottish, Serbian, Shona, Sindhi, Sinhala, Slovak, Slovene, Somali, SouthernNdebele, SouthernSotho, Spanish, Sundanese, Swahili, Swati, Swedish, Tagalog, Tahitian, Tajik, Tamil, Tatar, Telugu, Thai, Tibetan, Tigrinya, Tonga, Tsonga, Tswana, Turkish, Turkmen, Twi, Ukrainian, Urdu, Uyghur, Uzbek, Venda, Vietnamese, Volapük, Walloon, Welsh, Western, Wolof, Xhosa, Yiddish, Yoruba, Zhuang, Zulu}
	tags := []string{"", "ab", "aa", "af", "ak", "sq", "am", "ar", "an", "hy", "as", "av", "ae", "ay", "az", "bm", "ba", "eu", "be", "bn", "bh", "bi", "bs", "br", "bg", "my", "ca", "ch", "ce", "ny", "zh", "zh_Hant", "cv", "kw", "co", "cr", "hr", "cs", "da", "dv", "nl", "dz", "en", "eo", "et", "ee", "fo", "fj", "fi", "fr", "ff", "gl", "lg", "ka", "de", "el", "gn", "gu", "ht", "ha", "he", "hz", "hi", "ho", "hu", "is", "io", "ig", "id", "ia", "ie", "iu", "ik", "ga", "it", "ja", "jv", "kl", "kn", "kr", "ks", "kk", "km", "ki", "rw", "rn", "kv", "kg", "ko", "ku", "kj", "ky", "lo", "la", "lv", "li", "ln", "lt", "lu", "lb", "mk", "mg", "ms", "ml", "mt", "gv", "mi", "mr", "mh", "mn", "na", "nv", "ng", "ne", "se", "nd", "no", "nb", "nn", "ii", "oc", "oj", "cu", "or", "om", "os", "pi", "pa", "ps", "fa", "pl", "pt", "qu", "ro", "rm", "ru", "sm", "sg", "sa", "sc", "gd", "sr", "sn", "sd", "si", "sk", "sl", "so", "nr", "st", "es", "su", "sw", "ss", "sv", "tl", "ty", "tg", "ta", "tt", "te", "th", "bo", "ti", "to", "ts", "tn", "tr", "tk", "tw", "uk", "ur", "ug", "uz", "ve", "vi", "vo", "wa", "cy", "fy", "wo", "xh", "yi", "yo", "za", "zu"}

	for i := range values {
		if s := values[i].String(); s != tags[i] {
			t.Error(s, tags[i])
		}
	}
}

func BenchmarkLanguage_String(b *testing.B) {
	vs := []Language{Unknown, Abkhaz, Afar, Afrikaans, Akan, Albanian, Amharic, Arabic, Aragonese, Armenian, Assamese, Avaric, Avestan, Aymara, Azerbaijani, Bambara, Bashkir, Basque, Belarusian, Bengali, Bihari, Bislama, Bosnian, Breton, Bulgarian, Burmese, Catalan, Chamorro, Chechen, Chichewa, Chinese, ChineseTraditional, Chuvash, Cornish, Corsican, Cree, Croatian, Czech, Danish, Divehi, Dutch, Dzongkha, English, Esperanto, Estonian, Ewe, Faroese, Fijian, Finnish, French, Fula, Galician, Ganda, Georgian, German, Greek, Guaraní, Gujarati, Haitian, Hausa, Hebrew, Herero, Hindi, Hiri, Hungarian, Icelandic, Ido, Igbo, Indonesian, Interlingua, Interlingue, Inuktitut, Inupiaq, Irish, Italian, Japanese, Javanese, Kalaallisut, Kannada, Kanuri, Kashmiri, Kazakh, Khmer, Kikuyu, Kinyarwanda, Kirundi, Komi, Kongo, Korean, Kurdish, Kwanyama, Kyrgyz, Lao, Latin, Latvian, Limburgish, Lingala, Lithuanian, Luba, Luxembourgish, Macedonian, Malagasy, Malay, Malayalam, Maltese, Manx, Māori, Marathi, Marshallese, Mongolian, Nauru, Navajo, Ndonga, Nepali, Northern, NorthernNdebele, Norwegian, NorwegianBokmal, NorwegianNynorsk, Nuosu, Occitan, Ojibwe, Old, Oriya, Oromo, Ossetian, Pāli, Panjabi, Pashto, Persian, Polish, Portuguese, Quechua, Romanian, Romansh, Russian, Samoan, Sango, Sanskrit, Sardinian, Scottish, Serbian, Shona, Sindhi, Sinhala, Slovak, Slovene, Somali, SouthernNdebele, SouthernSotho, Spanish, Sundanese, Swahili, Swati, Swedish, Tagalog, Tahitian, Tajik, Tamil, Tatar, Telugu, Thai, Tibetan, Tigrinya, Tonga, Tsonga, Tswana, Turkish, Turkmen, Twi, Ukrainian, Urdu, Uyghur, Uzbek, Venda, Vietnamese, Volapük, Walloon, Welsh, Western, Wolof, Xhosa, Yiddish, Yoruba, Zhuang, Zulu}
	v := vs[rand.Intn(len(vs))]

	for b.Loop() {
		_ = v.String()
	}
}
