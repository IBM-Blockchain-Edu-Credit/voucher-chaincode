// records := []Record{
// 	Record{sid: "243506374", fullName: "Dong hoon", level: "A", starEarned: "4320",logins: "9",listen: "23",read: "24",worksheet: "0",quiz: "44",passedQuizCount: "25", practiceRecording: "0"},
// 	Record{sid: "243016624", fullName: "Kang yukyung", level: "A", starEarned: "5300",logins: "25",listen: "48",read: "44",worksheet: "0",quiz: "34",passedQuizCount: "22",practiceRecording: "7"},
// 	Record{sid: "243016627", fullName: "Kang Denis", level: "A", starEarned: "6580",logins: "14",listen: "37",read: "36",worksheet: "0",quiz: "82",passedQuizCount: "48",practiceRecording: "1"},
// 	Record{sid: "241649412", fullName: "Jaebin", level: "B", starEarned: "3470",logins: "10",listen: "24",read: "24",worksheet: "0",quiz: "28",passedQuizCount: "16",practiceRecording: "4"},
// 	Record{sid: "241392752", fullName: "Ye Seo", level: "C", starEarned: "13470",logins: "35",listen: "131",read: "162",worksheet: "0",quiz: "139",passedQuizCount: "61",practiceRecording: "16"},
// 	Record{sid: "241392846", fullName: "Erika", level: "A", starEarned: "4610",logins: "18",listen: "27",read: "38",worksheet: "0",quiz: "73",passedQuizCount: "32",practiceRecording: "0"},
// 	Record{sid: "241392874", fullName: "Yeon Woo", level: "A", starEarned: "2490",logins: "11",listen: "21",read: "25",worksheet: "0",quiz: "32",passedQuizCount: "18",practiceRecording: "0"},
// 	Record{sid: "241391885", fullName: "Ye Joon", level: "C", starEarned: "23200",logins: "49",listen: "162",read: "145",worksheet: "0",quiz: "271",passedQuizCount: "162",practiceRecording: "15"},
// 	Record{sid: "232199169", fullName: "Tan", level: "B", starEarned: "1680",logins: "9",listen: "8",read: "12",worksheet: "0",quiz: "26",passedQuizCount: "13",practiceRecording: "0"},
// 	Record{sid: "232199170", fullName: "jiwook kim", level: "A", starEarned: "2830",logins: "8",listen: "12",read: "15",worksheet: "0",quiz: "44",passedQuizCount: "27",practiceRecording: "3"},``
// 	Record{sid: "232199171", fullName: "junhyub park", level: "B", starEarned: "8540",logins: "15",listen: "41",read: "42",worksheet: "0",quiz: "96",passedQuizCount: "53",practiceRecording: "14"},
// 	Record{sid: "232199174", fullName: "hyejung park", level: "E", starEarned: "9380",logins: "41",listen: "44",read: "136",worksheet: "0",quiz: "344",passedQuizCount: "123",practiceRecording: "10"},
// 	Record{sid: "232199181", fullName: "Bladic", level: "C", starEarned: "13500",logins: "14",listen: "85",read: "82",worksheet: "0",quiz: "105",passedQuizCount: "66",practiceRecording: "6"},
// 	Record{sid: "232199182", fullName: "Sun Ha", level: "C", starEarned: "8280",logins: "20",listen: "60",read: "87",worksheet: "0",quiz: "181",passedQuizCount: "91",practiceRecording: "11"},
// 	Record{sid: "232199184", fullName: "sungsoo moon", level: "E", starEarned: "12960",logins: "36",listen: "98",read: "68",worksheet: "0",quiz: "208",passedQuizCount: "103",practiceRecording: "5"},
// 	Record{sid: "232173421", fullName: "Hana", level: "B", starEarned: "7030",logins: "18",listen: "35",read: "46",worksheet: "0",quiz: "50",passedQuizCount: "30",practiceRecording: "0"},
// 	Record{sid: "232170417", fullName: "kyuhyung lee", level: "B", starEarned: "950",logins: "11",listen: "23",read: "22",worksheet: "0",quiz: "9",passedQuizCount: "4",practiceRecording: "0"},
// 	Record{sid: "232170418", fullName: "nayeon park", level: "B", starEarned: "43270",logins: "52",listen: "237",read: "313",worksheet: "0",quiz: "488",passedQuizCount: "256",practiceRecording: "133"},
// 	Record{sid: "232170420", fullName: "minji kang", level: "B", starEarned: "9280",logins: "17",listen: "41",read: "46",worksheet: "0",quiz: "78",passedQuizCount: "44",practiceRecording: "1"},
// 	Record{sid: "232170422", fullName: "youngmi kim", level: "C", starEarned: "10140",logins: "33",listen: "47",read: "88",worksheet: "0",quiz: "215",passedQuizCount: "93",practiceRecording: "42"},
// 	Record{sid: "232170423", fullName: "jaesung kim", level: "B", starEarned: "6950",logins: "23",listen: "37",read: "40",worksheet: "0",quiz: "81",passedQuizCount: "44",practiceRecording: "2"},
// 	Record{sid: "232170426", fullName: "hyunji park", level: "B", starEarned: "9960",logins: "22",listen: "64",read: "51",worksheet: "0",quiz: "157",passedQuizCount: "80",practiceRecording: "0"},
// 	Record{sid: "232170427", fullName: "hyewon shin", level: "C", starEarned: "6990",logins: "20",listen: "35",read: "37",worksheet: "0",quiz: "103",passedQuizCount: "50",practiceRecording: "6"},
// 	Record{sid: "232170428", fullName: "hyeeun shin", level: "A", starEarned: "3500",logins: "9",listen: "4",read: "17",worksheet: "0",quiz: "34",passedQuizCount: "17",practiceRecording: "8"},
// 	Record{sid: "232170429", fullName: "hyejin shin", level: "A", starEarned: "390",logins: "3",listen: "9",read: "3",worksheet: "0",quiz: "7",passedQuizCount: "4",practiceRecording: "2"},
// 	Record{sid: "232170430", fullName: "hosung lee", level: "A", starEarned: "6410",logins: "30",listen: "55",read: "54",worksheet: "0",quiz: "102",passedQuizCount: "61",practiceRecording: "0"},
// 	Record{sid: "232170431", fullName: "hyokeun kim", level: "A", starEarned: "4340",logins: "26",listen: "55",read: "32",worksheet: "0",quiz: "65",passedQuizCount: "44",practiceRecording: "21"},
// }