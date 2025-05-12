package language

// Language is encoded into ISO 639 encoding
type Language uint8

//go:generate go-enum-encoding -type=Language -string
const (
	Unknown            Language = iota // json:""
	Abkhaz                             // json:"ab"
	Afar                               // json:"aa"
	Afrikaans                          // json:"af"
	Akan                               // json:"ak"
	Albanian                           // json:"sq"
	Amharic                            // json:"am"
	Arabic                             // json:"ar"
	Aragonese                          // json:"an"
	Armenian                           // json:"hy"
	Assamese                           // json:"as"
	Avaric                             // json:"av"
	Avestan                            // json:"ae"
	Aymara                             // json:"ay"
	Azerbaijani                        // json:"az"
	Bambara                            // json:"bm"
	Bashkir                            // json:"ba"
	Basque                             // json:"eu"
	Belarusian                         // json:"be"
	Bengali                            // json:"bn"
	Bihari                             // json:"bh"
	Bislama                            // json:"bi"
	Bosnian                            // json:"bs"
	Breton                             // json:"br"
	Bulgarian                          // json:"bg"
	Burmese                            // json:"my"
	Catalan                            // json:"ca"
	Chamorro                           // json:"ch"
	Chechen                            // json:"ce"
	Chichewa                           // json:"ny"
	Chinese                            // json:"zh"
	ChineseTraditional                 // json:"zh_Hant"
	Chuvash                            // json:"cv"
	Cornish                            // json:"kw"
	Corsican                           // json:"co"
	Cree                               // json:"cr"
	Croatian                           // json:"hr"
	Czech                              // json:"cs"
	Danish                             // json:"da"
	Divehi                             // json:"dv"
	Dutch                              // json:"nl"
	Dzongkha                           // json:"dz"
	English                            // json:"en"
	Esperanto                          // json:"eo"
	Estonian                           // json:"et"
	Ewe                                // json:"ee"
	Faroese                            // json:"fo"
	Fijian                             // json:"fj"
	Finnish                            // json:"fi"
	French                             // json:"fr"
	Fula                               // json:"ff"
	Galician                           // json:"gl"
	Ganda                              // json:"lg"
	Georgian                           // json:"ka"
	German                             // json:"de"
	Greek                              // json:"el"
	Guaraní                            // json:"gn"
	Gujarati                           // json:"gu"
	Haitian                            // json:"ht"
	Hausa                              // json:"ha"
	Hebrew                             // json:"he"
	Herero                             // json:"hz"
	Hindi                              // json:"hi"
	Hiri                               // json:"ho"
	Hungarian                          // json:"hu"
	Icelandic                          // json:"is"
	Ido                                // json:"io"
	Igbo                               // json:"ig"
	Indonesian                         // json:"id"
	Interlingua                        // json:"ia"
	Interlingue                        // json:"ie"
	Inuktitut                          // json:"iu"
	Inupiaq                            // json:"ik"
	Irish                              // json:"ga"
	Italian                            // json:"it"
	Japanese                           // json:"ja"
	Javanese                           // json:"jv"
	Kalaallisut                        // json:"kl"
	Kannada                            // json:"kn"
	Kanuri                             // json:"kr"
	Kashmiri                           // json:"ks"
	Kazakh                             // json:"kk"
	Khmer                              // json:"km"
	Kikuyu                             // json:"ki"
	Kinyarwanda                        // json:"rw"
	Kirundi                            // json:"rn"
	Komi                               // json:"kv"
	Kongo                              // json:"kg"
	Korean                             // json:"ko"
	Kurdish                            // json:"ku"
	Kwanyama                           // json:"kj"
	Kyrgyz                             // json:"ky"
	Lao                                // json:"lo"
	Latin                              // json:"la"
	Latvian                            // json:"lv"
	Limburgish                         // json:"li"
	Lingala                            // json:"ln"
	Lithuanian                         // json:"lt"
	Luba                               // json:"lu"
	Luxembourgish                      // json:"lb"
	Macedonian                         // json:"mk"
	Malagasy                           // json:"mg"
	Malay                              // json:"ms"
	Malayalam                          // json:"ml"
	Maltese                            // json:"mt"
	Manx                               // json:"gv"
	Māori                              // json:"mi"
	Marathi                            // json:"mr"
	Marshallese                        // json:"mh"
	Mongolian                          // json:"mn"
	Nauru                              // json:"na"
	Navajo                             // json:"nv"
	Ndonga                             // json:"ng"
	Nepali                             // json:"ne"
	Northern                           // json:"se"
	NorthernNdebele                    // json:"nd"
	Norwegian                          // json:"no"
	NorwegianBokmal                    // json:"nb"
	NorwegianNynorsk                   // json:"nn"
	Nuosu                              // json:"ii"
	Occitan                            // json:"oc"
	Ojibwe                             // json:"oj"
	Old                                // json:"cu"
	Oriya                              // json:"or"
	Oromo                              // json:"om"
	Ossetian                           // json:"os"
	Pāli                               // json:"pi"
	Panjabi                            // json:"pa"
	Pashto                             // json:"ps"
	Persian                            // json:"fa"
	Polish                             // json:"pl"
	Portuguese                         // json:"pt"
	Quechua                            // json:"qu"
	Romanian                           // json:"ro"
	Romansh                            // json:"rm"
	Russian                            // json:"ru"
	Samoan                             // json:"sm"
	Sango                              // json:"sg"
	Sanskrit                           // json:"sa"
	Sardinian                          // json:"sc"
	Scottish                           // json:"gd"
	Serbian                            // json:"sr"
	Shona                              // json:"sn"
	Sindhi                             // json:"sd"
	Sinhala                            // json:"si"
	Slovak                             // json:"sk"
	Slovene                            // json:"sl"
	Somali                             // json:"so"
	SouthernNdebele                    // json:"nr"
	SouthernSotho                      // json:"st"
	Spanish                            // json:"es"
	Sundanese                          // json:"su"
	Swahili                            // json:"sw"
	Swati                              // json:"ss"
	Swedish                            // json:"sv"
	Tagalog                            // json:"tl"
	Tahitian                           // json:"ty"
	Tajik                              // json:"tg"
	Tamil                              // json:"ta"
	Tatar                              // json:"tt"
	Telugu                             // json:"te"
	Thai                               // json:"th"
	Tibetan                            // json:"bo"
	Tigrinya                           // json:"ti"
	Tonga                              // json:"to"
	Tsonga                             // json:"ts"
	Tswana                             // json:"tn"
	Turkish                            // json:"tr"
	Turkmen                            // json:"tk"
	Twi                                // json:"tw"
	Ukrainian                          // json:"uk"
	Urdu                               // json:"ur"
	Uyghur                             // json:"ug"
	Uzbek                              // json:"uz"
	Venda                              // json:"ve"
	Vietnamese                         // json:"vi"
	Volapük                            // json:"vo"
	Walloon                            // json:"wa"
	Welsh                              // json:"cy"
	Western                            // json:"fy"
	Wolof                              // json:"wo"
	Xhosa                              // json:"xh"
	Yiddish                            // json:"yi"
	Yoruba                             // json:"yo"
	Zhuang                             // json:"za"
	Zulu                               // json:"zu"
)

var All = [...]Language{
	Abkhaz,
	Afar,
	Afrikaans,
	Akan,
	Albanian,
	Amharic,
	Arabic,
	Aragonese,
	Armenian,
	Assamese,
	Avaric,
	Avestan,
	Aymara,
	Azerbaijani,
	Bambara,
	Bashkir,
	Basque,
	Belarusian,
	Bengali,
	Bihari,
	Bislama,
	Bosnian,
	Breton,
	Bulgarian,
	Burmese,
	Catalan,
	Chamorro,
	Chechen,
	Chichewa,
	Chinese,
	ChineseTraditional,
	Chuvash,
	Cornish,
	Corsican,
	Cree,
	Croatian,
	Czech,
	Danish,
	Divehi,
	Dutch,
	Dzongkha,
	English,
	Esperanto,
	Estonian,
	Ewe,
	Faroese,
	Fijian,
	Finnish,
	French,
	Fula,
	Galician,
	Ganda,
	Georgian,
	German,
	Greek,
	Guaraní,
	Gujarati,
	Haitian,
	Hausa,
	Hebrew,
	Herero,
	Hindi,
	Hiri,
	Hungarian,
	Icelandic,
	Ido,
	Igbo,
	Indonesian,
	Interlingua,
	Interlingue,
	Inuktitut,
	Inupiaq,
	Irish,
	Italian,
	Japanese,
	Javanese,
	Kalaallisut,
	Kannada,
	Kanuri,
	Kashmiri,
	Kazakh,
	Khmer,
	Kikuyu,
	Kinyarwanda,
	Kirundi,
	Komi,
	Kongo,
	Korean,
	Kurdish,
	Kwanyama,
	Kyrgyz,
	Lao,
	Latin,
	Latvian,
	Limburgish,
	Lingala,
	Lithuanian,
	Luba,
	Luxembourgish,
	Macedonian,
	Malagasy,
	Malay,
	Malayalam,
	Maltese,
	Manx,
	Māori,
	Marathi,
	Marshallese,
	Mongolian,
	Nauru,
	Navajo,
	Ndonga,
	Nepali,
	Northern,
	NorthernNdebele,
	Norwegian,
	NorwegianBokmal,
	NorwegianNynorsk,
	Nuosu,
	Occitan,
	Ojibwe,
	Old,
	Oriya,
	Oromo,
	Ossetian,
	Pāli,
	Panjabi,
	Pashto,
	Persian,
	Polish,
	Portuguese,
	Quechua,
	Romanian,
	Romansh,
	Russian,
	Samoan,
	Sango,
	Sanskrit,
	Sardinian,
	Scottish,
	Serbian,
	Shona,
	Sindhi,
	Sinhala,
	Slovak,
	Slovene,
	Somali,
	SouthernNdebele,
	SouthernSotho,
	Spanish,
	Sundanese,
	Swahili,
	Swati,
	Swedish,
	Tagalog,
	Tahitian,
	Tajik,
	Tamil,
	Tatar,
	Telugu,
	Thai,
	Tibetan,
	Tigrinya,
	Tonga,
	Tsonga,
	Tswana,
	Turkish,
	Turkmen,
	Twi,
	Ukrainian,
	Urdu,
	Uyghur,
	Uzbek,
	Venda,
	Vietnamese,
	Volapük,
	Walloon,
	Welsh,
	Western,
	Wolof,
	Xhosa,
	Yiddish,
	Yoruba,
	Zhuang,
	Zulu,
}
