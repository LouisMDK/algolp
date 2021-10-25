package algolp

import (
	"testing"
)

func TestRapideVide(t *testing.T){
	var liste []int
	TriRapide(liste)
	if len(liste) != 0 || liste != nil{
		t.Fail()
	}
}

func TestRapideVide2(t *testing.T){
	var liste []int = nil
	TriRapide(liste)
	if len(liste) != 0 || liste != nil {
		t.Fail()
	}
}

func TestRapideDejaTrie(t *testing.T) {
	var liste []int = []int{1, 2, 36, 72, 80}
	TriRapide(liste)
	if !listeEgal(liste, []int{1, 2, 36, 72, 80}) {
		t.Fail()
	}
}

func TestRapideValeursNegatives(t *testing.T) {
	var liste []int = []int{-5, -50, -63, -05, -98}
	TriRapide(liste)
	if !listeEgal(liste, []int{-98, -63, -50, -5, -5}) {
		t.Fail()
	}
}


func TestRapideUneValeur(t *testing.T) {
	var liste []int = []int{1}
	TriRapide(liste)
	if !listeEgal(liste, []int{1}) {
		t.Fail()
	}
}

func TestRapidePlusieursValeurs(t *testing.T) {
	var liste []int = []int{9, 5, 69, 1, 0}
	TriRapide(liste)
	if !listeEgal(liste, []int{0, 1, 5, 9, 69}) {
		t.Fail()
	}
}

func TestRapideDoublons(t *testing.T) {
	var liste []int = []int{-5, 2, 30, 2, 99, 85, 99, 30, 85, 85, 10}
	TriRapide(liste)
	if !listeEgal(liste, []int{-5, 2, 2, 10, 30, 30, 85, 85, 85, 99, 99}) {
		t.Fail()
	}
}

func BenchmarkRapide(b *testing.B){
	for n:=0; n<b.N;n++{
		TriRapide([]int{106,
			223,
			-81,
			637,
			-628,
			69,
			673,
			-666,
			666,
			-869,
			889,
			25,
			179,
			289,
			-727,
			499,
			202,
			-262,
			-47,
			-441,
			726,
			533,
			816,
			-158,
			-705,
			-989,
			-614,
			213,
			-83,
			472,
			-982,
			-550,
			-31,
			905,
			-189,
			-428,
			-185,
			718,
			545,
			46,
			509,
			27,
			648,
			487,
			479,
			522,
			-259,
			569,
			189,
			-234,
			-664,
			-712,
			-825,
			-428,
			928,
			-874,
			-951,
			-615,
			-338,
			-935,
			656,
			526,
			289,
			-604,
			493,
			0,
			425,
			-806,
			-808,
			-21,
			-274,
			214,
			-854,
			-214,
			-629,
			-353,
			-84,
			830,
			313,
			-664,
			273,
			-970,
			-567,
			296,
			-869,
			905,
			907,
			29,
			264,
			-708,
			967,
			-364,
			-164,
			905,
			-581,
			-722,
			199,
			-79,
			466,
			947,
			571,
			-999,
			183,
			-955,
			384,
			538,
			-932,
			56,
			-994,
			-900,
			54,
			578,
			780,
			-602,
			-522,
			814,
			447,
			-931,
			-597,
			-292,
			103,
			183,
			-821,
			461,
			901,
			-509,
			771,
			-828,
			-507,
			708,
			-254,
			440,
			-55,
			-348,
			130,
			-37,
			510,
			-119,
			958,
			-887,
			402,
			634,
			-406,
			734,
			-315,
			44,
			-853,
			-384,
			248,
			-313,
			866,
			-21,
			-847,
			-556,
			647,
			334,
			-28,
			459,
			850,
			-362,
			-59,
			-371,
			-124,
			-45,
			520,
			-217,
			216,
			-870,
			495,
			-252,
			770,
			52,
			-47,
			679,
			47,
			888,
			997,
			-936,
			-247,
			371,
			-469,
			-852,
			249,
			-970,
			29,
			672,
			1,
			503,
			-945,
			-480,
			756,
			-80,
			730,
			-93,
			97,
			730,
			680,
			745,
			354,
			-613,
			325,
			-567,
			-827,
			-830,
			-292,
			-939,
			-567,
			-218,
			-274,
			533,
			711,
			638,
			-463,
			-254,
			-440,
			-936,
			-295,
			-951,
			-759,
			-591,
			-447,
			-500,
			-417,
			-314,
			-770,
			873,
			-115,
			-35,
			-281,
			-61,
			562,
			53,
			634,
			934,
			-131,
			776,
			280,
			84,
			-802,
			976,
			918,
			-129,
			604,
			-555,
			408,
			2,
			-847,
			611,
			162,
			-502,
			-784,
			-996,
			370,
			-821,
			982,
			521,
			-775,
			-158,
			-49,
			-836,
			832,
			-180,
			121,
			-270,
			270,
			41,
			657,
			211,
			-86,
			-901,
			255,
			955,
			863,
			890,
			700,
			-635,
			225,
			-427,
			305,
			215,
			625,
			290,
			488,
			-21,
			-862,
			527,
			757,
			211,
			-97,
			-178,
			-573,
			484,
			-758,
			267,
			-825,
			-938,
			-365,
			574,
			775,
			241,
			-594,
			-408,
			-680,
			78,
			-118,
			638,
			-200,
			746,
			-566,
			647,
			-329,
			350,
			-522,
			328,
			146,
			431,
			798,
			-688,
			486,
			-922,
			815,
			-3,
			-695,
			553,
			790,
			-138,
			-229,
			-477,
			449,
			51,
			377,
			954,
			897,
			473,
			561,
			659,
			828,
			-161,
			281,
			-945,
			-366,
			605,
			-587,
			-766,
			702,
			-738,
			241,
			-159,
			777,
			-813,
			983,
			385,
			375,
			824,
			-724,
			-55,
			-190,
			-531,
			608,
			428,
			500,
			438,
			714,
			-763,
			-615,
			908,
			-923,
			974,
			653,
			293,
			-793,
			979,
			148,
			138,
			-920,
			-607,
			-705,
			-174,
			-96,
			-972,
			-763,
			491,
			-476,
			-323,
			162,
			355,
			-603,
			-405,
			724,
			870,
			-919,
			-218,
			81,
			831,
			-98,
			-290,
			-657,
			328,
			253,
			-442})
	}
}