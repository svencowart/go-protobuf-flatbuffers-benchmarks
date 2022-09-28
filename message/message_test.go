package message

import (
	"testing"

	"github.com/elastiflow/codex/message/flatmsg"
	"github.com/elastiflow/codex/message/protomsg"
	flatbuffers "github.com/google/flatbuffers/go"
	"google.golang.org/protobuf/proto"
)

func NewProtoMessageNest() protomsg.MessageNest {
	return protomsg.MessageNest{
		Field_1: 1024,
		Field_2: &protomsg.MessageNestField2{
			Field_3: "Bench",
			Field_4: "Mark",
		},
	}
}

func NewProtoMessageFlat() protomsg.MessageFlat {
	return protomsg.MessageFlat{
		Field_1: 1024,
		Field_2: 1024,
		Field_3: "Bench",
		Field_4: "Mark",
	}
}

func NewProtoMessageFlatLarge() protomsg.MessageFlatLarge {
	label := "Bench"
	org := "Mark"

	return protomsg.MessageFlatLarge{
		Field_1:   1024,
		Field_2:   1024,
		Field_3:   1024,
		Field_4:   1024,
		Field_5:   1024,
		Field_6:   1024,
		Field_7:   true,
		Field_8:   label,
		Field_9:   org,
		Field_10:  true,
		Field_11:  1024,
		Field_12:  1024,
		Field_13:  1024,
		Field_14:  1024,
		Field_15:  1024,
		Field_16:  1024,
		Field_17:  true,
		Field_18:  label,
		Field_19:  org,
		Field_20:  true,
		Field_21:  1024,
		Field_22:  1024,
		Field_23:  1024,
		Field_24:  1024,
		Field_25:  1024,
		Field_26:  1024,
		Field_27:  true,
		Field_28:  label,
		Field_29:  org,
		Field_40:  true,
		Field_41:  1024,
		Field_42:  1024,
		Field_43:  1024,
		Field_44:  1024,
		Field_45:  1024,
		Field_46:  1024,
		Field_47:  true,
		Field_48:  label,
		Field_49:  org,
		Field_50:  true,
		Field_51:  1024,
		Field_52:  1024,
		Field_53:  1024,
		Field_54:  1024,
		Field_55:  1024,
		Field_56:  1024,
		Field_57:  true,
		Field_58:  label,
		Field_59:  org,
		Field_60:  true,
		Field_61:  1024,
		Field_62:  1024,
		Field_63:  1024,
		Field_64:  1024,
		Field_65:  1024,
		Field_66:  1024,
		Field_67:  true,
		Field_68:  label,
		Field_69:  org,
		Field_70:  true,
		Field_71:  1024,
		Field_72:  1024,
		Field_73:  1024,
		Field_74:  1024,
		Field_75:  1024,
		Field_76:  1024,
		Field_77:  true,
		Field_78:  label,
		Field_79:  org,
		Field_80:  true,
		Field_81:  1024,
		Field_82:  1024,
		Field_83:  1024,
		Field_84:  1024,
		Field_85:  1024,
		Field_86:  1024,
		Field_87:  true,
		Field_88:  label,
		Field_89:  org,
		Field_90:  true,
		Field_91:  1024,
		Field_92:  1024,
		Field_93:  1024,
		Field_94:  1024,
		Field_95:  1024,
		Field_96:  1024,
		Field_97:  true,
		Field_98:  label,
		Field_99:  org,
		Field_100: true,
		Field_101: 1024,
		Field_102: 1024,
		Field_103: 1024,
		Field_104: 1024,
		Field_105: 1024,
		Field_106: 1024,
		Field_107: true,
		Field_108: label,
		Field_109: org,
		Field_110: true,
		Field_111: 1024,
		Field_112: 1024,
		Field_113: 1024,
		Field_114: 1024,
		Field_115: 1024,
		Field_116: 1024,
		Field_117: true,
		Field_118: label,
		Field_119: org,
	}
}

func NewProtoMessageFlatXLarge() protomsg.MessageFlatXLarge {
	label := "Bench"
	org := "Mark"

	return protomsg.MessageFlatXLarge{
		Field_1:   1024,
		Field_2:   1024,
		Field_3:   1024,
		Field_4:   1024,
		Field_5:   1024,
		Field_6:   1024,
		Field_7:   true,
		Field_8:   label,
		Field_9:   org,
		Field_10:  true,
		Field_11:  1024,
		Field_12:  1024,
		Field_13:  1024,
		Field_14:  1024,
		Field_15:  1024,
		Field_16:  1024,
		Field_17:  true,
		Field_18:  label,
		Field_19:  org,
		Field_20:  true,
		Field_21:  1024,
		Field_22:  1024,
		Field_23:  1024,
		Field_24:  1024,
		Field_25:  1024,
		Field_26:  1024,
		Field_27:  true,
		Field_28:  label,
		Field_29:  org,
		Field_40:  true,
		Field_41:  1024,
		Field_42:  1024,
		Field_43:  1024,
		Field_44:  1024,
		Field_45:  1024,
		Field_46:  1024,
		Field_47:  true,
		Field_48:  label,
		Field_49:  org,
		Field_50:  true,
		Field_51:  1024,
		Field_52:  1024,
		Field_53:  1024,
		Field_54:  1024,
		Field_55:  1024,
		Field_56:  1024,
		Field_57:  true,
		Field_58:  label,
		Field_59:  org,
		Field_60:  true,
		Field_61:  1024,
		Field_62:  1024,
		Field_63:  1024,
		Field_64:  1024,
		Field_65:  1024,
		Field_66:  1024,
		Field_67:  true,
		Field_68:  label,
		Field_69:  org,
		Field_70:  true,
		Field_71:  1024,
		Field_72:  1024,
		Field_73:  1024,
		Field_74:  1024,
		Field_75:  1024,
		Field_76:  1024,
		Field_77:  true,
		Field_78:  label,
		Field_79:  org,
		Field_80:  true,
		Field_81:  1024,
		Field_82:  1024,
		Field_83:  1024,
		Field_84:  1024,
		Field_85:  1024,
		Field_86:  1024,
		Field_87:  true,
		Field_88:  label,
		Field_89:  org,
		Field_90:  true,
		Field_91:  1024,
		Field_92:  1024,
		Field_93:  1024,
		Field_94:  1024,
		Field_95:  1024,
		Field_96:  1024,
		Field_97:  true,
		Field_98:  label,
		Field_99:  org,
		Field_100: true,
		Field_101: 1024,
		Field_102: 1024,
		Field_103: 1024,
		Field_104: 1024,
		Field_105: 1024,
		Field_106: 1024,
		Field_107: true,
		Field_108: label,
		Field_109: org,
		Field_110: true,
		Field_111: 1024,
		Field_112: 1024,
		Field_113: 1024,
		Field_114: 1024,
		Field_115: 1024,
		Field_116: 1024,
		Field_117: true,
		Field_118: label,
		Field_119: org,
		Field_120: true,
		Field_121: 1024,
		Field_122: 1024,
		Field_123: 1024,
		Field_124: 1024,
		Field_125: 1024,
		Field_126: 1024,
		Field_127: true,
		Field_128: label,
		Field_129: org,
		Field_130: true,
		Field_131: 1024,
		Field_132: 1024,
		Field_133: 1024,
		Field_134: 1024,
		Field_135: 1024,
		Field_136: 1024,
		Field_137: true,
		Field_138: label,
		Field_139: org,
		Field_140: true,
		Field_141: 1024,
		Field_142: 1024,
		Field_143: 1024,
		Field_144: 1024,
		Field_145: 1024,
		Field_146: 1024,
		Field_147: true,
		Field_148: label,
		Field_149: org,
		Field_150: true,
		Field_151: 1024,
		Field_152: 1024,
		Field_153: 1024,
		Field_154: 1024,
		Field_155: 1024,
		Field_156: 1024,
		Field_157: true,
		Field_158: label,
		Field_159: org,
		Field_160: true,
		Field_161: 1024,
		Field_162: 1024,
		Field_163: 1024,
		Field_164: 1024,
		Field_165: 1024,
		Field_166: 1024,
		Field_167: true,
		Field_168: label,
		Field_169: org,
		Field_170: true,
		Field_171: 1024,
		Field_172: 1024,
		Field_173: 1024,
		Field_174: 1024,
		Field_175: 1024,
		Field_176: 1024,
		Field_177: true,
		Field_178: label,
		Field_179: org,
		Field_180: true,
		Field_181: 1024,
		Field_182: 1024,
		Field_183: 1024,
		Field_184: 1024,
		Field_185: 1024,
		Field_186: 1024,
		Field_187: true,
		Field_188: label,
		Field_189: org,
		Field_190: true,
		Field_191: 1024,
		Field_192: 1024,
		Field_193: 1024,
		Field_194: 1024,
		Field_195: 1024,
		Field_196: 1024,
		Field_197: true,
		Field_198: label,
		Field_199: org,
		Field_200: true,
		Field_201: 1024,
		Field_202: 1024,
		Field_203: 1024,
		Field_204: 1024,
		Field_205: 1024,
		Field_206: 1024,
		Field_207: true,
		Field_208: label,
		Field_209: org,
		Field_210: true,
		Field_211: 1024,
		Field_212: 1024,
		Field_213: 1024,
		Field_214: 1024,
		Field_215: 1024,
		Field_216: 1024,
		Field_217: true,
		Field_218: label,
		Field_219: org,
		Field_220: true,
		Field_221: 1024,
		Field_222: 1024,
		Field_223: 1024,
		Field_224: 1024,
		Field_225: 1024,
		Field_226: 1024,
		Field_227: true,
		Field_228: label,
		Field_229: org,
		Field_230: true,
		Field_231: 1024,
		Field_232: 1024,
		Field_233: 1024,
		Field_234: 1024,
		Field_235: 1024,
		Field_236: 1024,
		Field_237: true,
		Field_238: label,
		Field_239: org,
		Field_240: true,
		Field_241: 1024,
		Field_242: 1024,
		Field_243: 1024,
		Field_244: 1024,
		Field_245: 1024,
		Field_246: 1024,
		Field_247: true,
		Field_248: label,
		Field_249: org,
		Field_250: true,
		Field_251: 1024,
		Field_252: 1024,
		Field_253: 1024,
		Field_254: 1024,
		Field_255: 1024,
		Field_256: 1024,
		Field_257: true,
		Field_258: label,
		Field_259: org,
		Field_260: true,
		Field_261: 1024,
		Field_262: 1024,
		Field_263: 1024,
		Field_264: 1024,
		Field_265: 1024,
		Field_266: 1024,
		Field_267: true,
		Field_268: label,
		Field_269: org,
		Field_270: true,
		Field_271: 1024,
		Field_272: 1024,
		Field_273: 1024,
		Field_274: 1024,
		Field_275: 1024,
		Field_276: 1024,
		Field_277: true,
		Field_278: label,
		Field_279: org,
		Field_280: true,
		Field_281: 1024,
		Field_282: 1024,
		Field_283: 1024,
		Field_284: 1024,
		Field_285: 1024,
		Field_286: 1024,
		Field_287: true,
		Field_288: label,
		Field_289: org,
		Field_290: true,
		Field_291: 1024,
		Field_292: 1024,
		Field_293: 1024,
		Field_294: 1024,
		Field_295: 1024,
		Field_296: 1024,
		Field_297: true,
		Field_298: label,
		Field_299: org,
		Field_300: true,
		Field_301: 1024,
		Field_302: 1024,
		Field_303: 1024,
		Field_304: 1024,
		Field_305: 1024,
		Field_306: 1024,
		Field_307: true,
		Field_308: label,
		Field_309: org,
		Field_310: true,
		Field_311: 1024,
		Field_312: 1024,
		Field_313: 1024,
		Field_314: 1024,
		Field_315: 1024,
		Field_316: 1024,
		Field_317: true,
		Field_318: label,
		Field_319: org,
		Field_320: true,
		Field_321: 1024,
		Field_322: 1024,
		Field_323: 1024,
		Field_324: 1024,
		Field_325: 1024,
		Field_326: 1024,
		Field_327: true,
		Field_328: label,
		Field_329: org,
		Field_330: true,
		Field_331: 1024,
		Field_332: 1024,
		Field_333: 1024,
		Field_334: 1024,
		Field_335: 1024,
		Field_336: 1024,
		Field_337: true,
		Field_338: label,
		Field_339: org,
		Field_340: true,
		Field_341: 1024,
		Field_342: 1024,
		Field_343: 1024,
		Field_344: 1024,
		Field_345: 1024,
		Field_346: 1024,
		Field_347: true,
		Field_348: label,
		Field_349: org,
		Field_350: true,
		Field_351: 1024,
		Field_352: 1024,
		Field_353: 1024,
		Field_354: 1024,
		Field_355: 1024,
		Field_356: 1024,
		Field_357: true,
		Field_358: label,
		Field_359: org,
		Field_360: true,
		Field_361: 1024,
		Field_362: 1024,
		Field_363: 1024,
		Field_364: 1024,
		Field_365: 1024,
		Field_366: 1024,
		Field_367: true,
		Field_368: label,
		Field_369: org,
		Field_370: true,
		Field_371: 1024,
		Field_372: 1024,
		Field_373: 1024,
		Field_374: 1024,
		Field_375: 1024,
		Field_376: 1024,
		Field_377: true,
		Field_378: label,
		Field_379: org,
		Field_380: true,
		Field_381: 1024,
		Field_382: 1024,
		Field_383: 1024,
		Field_384: 1024,
		Field_385: 1024,
		Field_386: 1024,
		Field_387: true,
		Field_388: label,
		Field_389: org,
		Field_390: true,
		Field_391: 1024,
		Field_392: 1024,
		Field_393: 1024,
		Field_394: 1024,
		Field_395: 1024,
		Field_396: 1024,
		Field_397: true,
		Field_398: label,
		Field_399: org,
		Field_400: true,
		Field_401: 1024,
		Field_402: 1024,
		Field_403: 1024,
		Field_404: 1024,
		Field_405: 1024,
		Field_406: 1024,
		Field_407: true,
		Field_408: label,
		Field_409: org,
		Field_410: true,
		Field_411: 1024,
		Field_412: 1024,
		Field_413: 1024,
		Field_414: 1024,
		Field_415: 1024,
		Field_416: 1024,
		Field_417: true,
		Field_418: label,
		Field_419: org,
		Field_420: true,
		Field_421: 1024,
		Field_422: 1024,
		Field_423: 1024,
		Field_424: 1024,
		Field_425: 1024,
		Field_426: 1024,
		Field_427: true,
		Field_428: label,
		Field_429: org,
		Field_430: true,
		Field_431: 1024,
		Field_432: 1024,
		Field_433: 1024,
		Field_434: 1024,
		Field_435: 1024,
		Field_436: 1024,
		Field_437: true,
		Field_438: label,
		Field_439: org,
		Field_440: true,
		Field_441: 1024,
		Field_442: 1024,
		Field_443: 1024,
		Field_444: 1024,
		Field_445: 1024,
		Field_446: 1024,
		Field_447: true,
		Field_448: label,
		Field_449: org,
		Field_450: true,
		Field_451: 1024,
		Field_452: 1024,
		Field_453: 1024,
		Field_454: 1024,
		Field_455: 1024,
		Field_456: 1024,
		Field_457: true,
		Field_458: label,
		Field_459: org,
		Field_460: true,
		Field_461: 1024,
		Field_462: 1024,
		Field_463: 1024,
		Field_464: 1024,
		Field_465: 1024,
		Field_466: 1024,
		Field_467: true,
		Field_468: label,
		Field_469: org,
		Field_470: true,
		Field_471: 1024,
		Field_472: 1024,
		Field_473: 1024,
		Field_474: 1024,
		Field_475: 1024,
		Field_476: 1024,
		Field_477: true,
		Field_478: label,
		Field_479: org,
		Field_480: true,
		Field_481: 1024,
		Field_482: 1024,
		Field_483: 1024,
		Field_484: 1024,
		Field_485: 1024,
		Field_486: 1024,
		Field_487: true,
		Field_488: label,
		Field_489: org,
		Field_490: true,
		Field_491: 1024,
		Field_492: 1024,
		Field_493: 1024,
		Field_494: 1024,
		Field_495: 1024,
		Field_496: 1024,
		Field_497: true,
		Field_498: label,
		Field_499: org,
		Field_500: true,
		Field_501: 1024,
		Field_502: 1024,
		Field_503: 1024,
		Field_504: 1024,
		Field_505: 1024,
		Field_506: 1024,
		Field_507: true,
		Field_508: label,
		Field_509: org,
		Field_510: true,
		Field_511: 1024,
		Field_512: 1024,
		Field_513: 1024,
		Field_514: 1024,
		Field_515: 1024,
		Field_516: 1024,
		Field_517: true,
		Field_518: label,
		Field_519: org,
		Field_520: true,
		Field_521: 1024,
		Field_522: 1024,
		Field_523: 1024,
		Field_524: 1024,
		Field_525: 1024,
		Field_526: 1024,
		Field_527: true,
		Field_528: label,
		Field_529: org,
		Field_530: true,
		Field_531: 1024,
		Field_532: 1024,
		Field_533: 1024,
		Field_534: 1024,
		Field_535: 1024,
		Field_536: 1024,
		Field_537: true,
		Field_538: label,
		Field_539: org,
		Field_540: true,
		Field_541: 1024,
		Field_542: 1024,
		Field_543: 1024,
		Field_544: 1024,
		Field_545: 1024,
		Field_546: 1024,
		Field_547: true,
		Field_548: label,
		Field_549: org,
		Field_550: true,
		Field_551: 1024,
		Field_552: 1024,
		Field_553: 1024,
		Field_554: 1024,
		Field_555: 1024,
		Field_556: 1024,
		Field_557: true,
		Field_558: label,
		Field_559: org,
		Field_560: true,
		Field_561: 1024,
		Field_562: 1024,
		Field_563: 1024,
		Field_564: 1024,
		Field_565: 1024,
		Field_566: 1024,
		Field_567: true,
		Field_568: label,
		Field_569: org,
		Field_570: true,
		Field_571: 1024,
		Field_572: 1024,
		Field_573: 1024,
		Field_574: 1024,
		Field_575: 1024,
		Field_576: 1024,
		Field_577: true,
		Field_578: label,
		Field_579: org,
		Field_580: true,
		Field_581: 1024,
		Field_582: 1024,
		Field_583: 1024,
		Field_584: 1024,
		Field_585: 1024,
		Field_586: 1024,
		Field_587: true,
		Field_588: label,
		Field_589: org,
		Field_590: true,
		Field_591: 1024,
		Field_592: 1024,
		Field_593: 1024,
		Field_594: 1024,
		Field_595: 1024,
		Field_596: 1024,
		Field_597: true,
		Field_598: label,
		Field_599: org,
	}
}

func NewProtoMessageFlatXLargePartial() protomsg.MessageFlatXLarge {
	label := "Bench"

	return protomsg.MessageFlatXLarge{
		Field_1: 1024,
		Field_2: 1024,
		Field_3: 1024,
		Field_4: 1024,
		Field_5: 1024,
		Field_6: 1024,
		Field_7: true,
		Field_8: label,
	}
}

func NewProtoMessageFlatXLargePartial64() protomsg.MessageFlatXLarge {
	label := "Bench"
	org := "Mark"

	return protomsg.MessageFlatXLarge{
		Field_1:  1024,
		Field_2:  1024,
		Field_3:  1024,
		Field_4:  1024,
		Field_5:  1024,
		Field_6:  1024,
		Field_7:  true,
		Field_8:  label,
		Field_9:  org,
		Field_10: true,
		Field_11: 1024,
		Field_12: 1024,
		Field_13: 1024,
		Field_14: 1024,
		Field_15: 1024,
		Field_16: 1024,
		Field_17: true,
		Field_18: label,
		Field_19: org,
		Field_20: true,
		Field_21: 1024,
		Field_22: 1024,
		Field_23: 1024,
		Field_24: 1024,
		Field_25: 1024,
		Field_26: 1024,
		Field_27: true,
		Field_28: label,
		Field_29: org,
		Field_40: true,
		Field_41: 1024,
		Field_42: 1024,
		Field_43: 1024,
		Field_44: 1024,
		Field_45: 1024,
		Field_46: 1024,
		Field_47: true,
		Field_48: label,
		Field_49: org,
		Field_50: true,
		Field_51: 1024,
		Field_52: 1024,
		Field_53: 1024,
		Field_54: 1024,
		Field_55: 1024,
		Field_56: 1024,
		Field_57: true,
		Field_58: label,
		Field_59: org,
		Field_60: true,
		Field_61: 1024,
		Field_62: 1024,
		Field_63: 1024,
		Field_64: 1024,
	}
}

func NewProtoMessageGroup() protomsg.MessageGroup {
	label := "Bench"
	org := "Mark"

	return protomsg.MessageGroup{
		Field_1: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_2: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_3: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_4: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_5: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_6: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_7: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_8: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_9: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_10: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_11: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_12: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_13: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_14: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_15: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_16: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_17: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_18: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_19: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_20: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_21: &protomsg.MessageGroupInner100{
			Field_1:   1024,
			Field_2:   1024,
			Field_3:   1024,
			Field_4:   1024,
			Field_5:   1024,
			Field_6:   1024,
			Field_7:   true,
			Field_8:   label,
			Field_9:   org,
			Field_10:  true,
			Field_11:  1024,
			Field_12:  1024,
			Field_13:  1024,
			Field_14:  1024,
			Field_15:  1024,
			Field_16:  1024,
			Field_17:  true,
			Field_18:  label,
			Field_19:  org,
			Field_20:  true,
			Field_21:  1024,
			Field_22:  1024,
			Field_23:  1024,
			Field_24:  1024,
			Field_25:  1024,
			Field_26:  1024,
			Field_27:  true,
			Field_28:  label,
			Field_29:  org,
			Field_40:  true,
			Field_41:  1024,
			Field_42:  1024,
			Field_43:  1024,
			Field_44:  1024,
			Field_45:  1024,
			Field_46:  1024,
			Field_47:  true,
			Field_48:  label,
			Field_49:  org,
			Field_50:  true,
			Field_51:  1024,
			Field_52:  1024,
			Field_53:  1024,
			Field_54:  1024,
			Field_55:  1024,
			Field_56:  1024,
			Field_57:  true,
			Field_58:  label,
			Field_59:  org,
			Field_60:  true,
			Field_61:  1024,
			Field_62:  1024,
			Field_63:  1024,
			Field_64:  1024,
			Field_65:  1024,
			Field_66:  1024,
			Field_67:  true,
			Field_68:  label,
			Field_69:  org,
			Field_70:  true,
			Field_71:  1024,
			Field_72:  1024,
			Field_73:  1024,
			Field_74:  1024,
			Field_75:  1024,
			Field_76:  1024,
			Field_77:  true,
			Field_78:  label,
			Field_79:  org,
			Field_80:  true,
			Field_81:  1024,
			Field_82:  1024,
			Field_83:  1024,
			Field_84:  1024,
			Field_85:  1024,
			Field_86:  1024,
			Field_87:  true,
			Field_88:  label,
			Field_89:  org,
			Field_90:  true,
			Field_91:  1024,
			Field_92:  1024,
			Field_93:  1024,
			Field_94:  1024,
			Field_95:  1024,
			Field_96:  1024,
			Field_97:  true,
			Field_98:  label,
			Field_99:  org,
			Field_100: true,
		},
		Field_22: &protomsg.MessageGroupInner100{
			Field_1:   1024,
			Field_2:   1024,
			Field_3:   1024,
			Field_4:   1024,
			Field_5:   1024,
			Field_6:   1024,
			Field_7:   true,
			Field_8:   label,
			Field_9:   org,
			Field_10:  true,
			Field_11:  1024,
			Field_12:  1024,
			Field_13:  1024,
			Field_14:  1024,
			Field_15:  1024,
			Field_16:  1024,
			Field_17:  true,
			Field_18:  label,
			Field_19:  org,
			Field_20:  true,
			Field_21:  1024,
			Field_22:  1024,
			Field_23:  1024,
			Field_24:  1024,
			Field_25:  1024,
			Field_26:  1024,
			Field_27:  true,
			Field_28:  label,
			Field_29:  org,
			Field_40:  true,
			Field_41:  1024,
			Field_42:  1024,
			Field_43:  1024,
			Field_44:  1024,
			Field_45:  1024,
			Field_46:  1024,
			Field_47:  true,
			Field_48:  label,
			Field_49:  org,
			Field_50:  true,
			Field_51:  1024,
			Field_52:  1024,
			Field_53:  1024,
			Field_54:  1024,
			Field_55:  1024,
			Field_56:  1024,
			Field_57:  true,
			Field_58:  label,
			Field_59:  org,
			Field_60:  true,
			Field_61:  1024,
			Field_62:  1024,
			Field_63:  1024,
			Field_64:  1024,
			Field_65:  1024,
			Field_66:  1024,
			Field_67:  true,
			Field_68:  label,
			Field_69:  org,
			Field_70:  true,
			Field_71:  1024,
			Field_72:  1024,
			Field_73:  1024,
			Field_74:  1024,
			Field_75:  1024,
			Field_76:  1024,
			Field_77:  true,
			Field_78:  label,
			Field_79:  org,
			Field_80:  true,
			Field_81:  1024,
			Field_82:  1024,
			Field_83:  1024,
			Field_84:  1024,
			Field_85:  1024,
			Field_86:  1024,
			Field_87:  true,
			Field_88:  label,
			Field_89:  org,
			Field_90:  true,
			Field_91:  1024,
			Field_92:  1024,
			Field_93:  1024,
			Field_94:  1024,
			Field_95:  1024,
			Field_96:  1024,
			Field_97:  true,
			Field_98:  label,
			Field_99:  org,
			Field_100: true,
		},
	}
}

func NewProtoMessageGroupPartial() protomsg.MessageGroup {
	label := "Bench"
	org := "Mark"

	return protomsg.MessageGroup{
		Field_1: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_2: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_3: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_4: &protomsg.MessageGroupInner20{
			Field_1: 1024,
			Field_2: 1024,
			Field_3: 1024,
			Field_4: 1024,
		},
	}
}

func NewProtoMessageGroupLarge() protomsg.MessageGroupLarge {
	label := "Bench"
	org := "Mark"

	return protomsg.MessageGroupLarge{
		Field_1: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_2: &protomsg.MessageGroupInner20{
			Field_1:  1024,
			Field_2:  1024,
			Field_3:  1024,
			Field_4:  1024,
			Field_5:  1024,
			Field_6:  1024,
			Field_7:  true,
			Field_8:  label,
			Field_9:  org,
			Field_10: true,
			Field_11: 1024,
			Field_12: 1024,
			Field_13: 1024,
			Field_14: 1024,
			Field_15: 1024,
			Field_16: 1024,
			Field_17: true,
			Field_18: label,
			Field_19: org,
			Field_20: true,
		},
		Field_95: &protomsg.MessageGroupInner100{
			Field_1:   1024,
			Field_2:   1024,
			Field_3:   1024,
			Field_4:   1024,
			Field_5:   1024,
			Field_6:   1024,
			Field_7:   true,
			Field_8:   label,
			Field_9:   org,
			Field_10:  true,
			Field_11:  1024,
			Field_12:  1024,
			Field_13:  1024,
			Field_14:  1024,
			Field_15:  1024,
			Field_16:  1024,
			Field_17:  true,
			Field_18:  label,
			Field_19:  org,
			Field_20:  true,
			Field_21:  1024,
			Field_22:  1024,
			Field_23:  1024,
			Field_24:  1024,
			Field_25:  1024,
			Field_26:  1024,
			Field_27:  true,
			Field_28:  label,
			Field_29:  org,
			Field_40:  true,
			Field_41:  1024,
			Field_42:  1024,
			Field_43:  1024,
			Field_44:  1024,
			Field_45:  1024,
			Field_46:  1024,
			Field_47:  true,
			Field_48:  label,
			Field_49:  org,
			Field_50:  true,
			Field_51:  1024,
			Field_52:  1024,
			Field_53:  1024,
			Field_54:  1024,
			Field_55:  1024,
			Field_56:  1024,
			Field_57:  true,
			Field_58:  label,
			Field_59:  org,
			Field_60:  true,
			Field_61:  1024,
			Field_62:  1024,
			Field_63:  1024,
			Field_64:  1024,
			Field_65:  1024,
			Field_66:  1024,
			Field_67:  true,
			Field_68:  label,
			Field_69:  org,
			Field_70:  true,
			Field_71:  1024,
			Field_72:  1024,
			Field_73:  1024,
			Field_74:  1024,
			Field_75:  1024,
			Field_76:  1024,
			Field_77:  true,
			Field_78:  label,
			Field_79:  org,
			Field_80:  true,
			Field_81:  1024,
			Field_82:  1024,
			Field_83:  1024,
			Field_84:  1024,
			Field_85:  1024,
			Field_86:  1024,
			Field_87:  true,
			Field_88:  label,
			Field_89:  org,
			Field_90:  true,
			Field_91:  1024,
			Field_92:  1024,
			Field_93:  1024,
			Field_94:  1024,
			Field_95:  1024,
			Field_96:  1024,
			Field_97:  true,
			Field_98:  label,
			Field_99:  org,
			Field_100: true,
		},
	}
}

func NewFlatMessageNest(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	label := builder.CreateString("Bench")
	org := builder.CreateString("Mark")

	flatmsg.MessageNestField2Start(builder)
	flatmsg.MessageNestField2AddField3(builder, label)
	flatmsg.MessageNestField2AddField4(builder, org)
	field2 := flatmsg.MessageNestField2End(builder)
	flatmsg.MessageNestStart(builder)
	flatmsg.MessageNestAddField1(builder, 1024)
	flatmsg.MessageNestAddField2(builder, field2)
	return flatmsg.MessageNestEnd(builder)
}

func NewFlatMessageFlat(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	label := builder.CreateString("Bench")
	org := builder.CreateString("Mark")

	flatmsg.MessageFlatStart(builder)
	flatmsg.MessageFlatAddField1(builder, 1024)
	flatmsg.MessageFlatAddField2(builder, 1024)
	flatmsg.MessageFlatAddField3(builder, label)
	flatmsg.MessageFlatAddField4(builder, org)
	return flatmsg.MessageFlatEnd(builder)
}

func NewFlatMessageFlatLarge(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	label := builder.CreateString("Bench")
	org := builder.CreateString("Mark")

	flatmsg.MessageFlatLargeStart(builder)
	flatmsg.MessageFlatLargeAddField1(builder, 1024)
	flatmsg.MessageFlatLargeAddField2(builder, 1024)
	flatmsg.MessageFlatLargeAddField3(builder, 1024)
	flatmsg.MessageFlatLargeAddField4(builder, 1024)
	flatmsg.MessageFlatLargeAddField5(builder, 1024)
	flatmsg.MessageFlatLargeAddField6(builder, 1024)
	flatmsg.MessageFlatLargeAddField7(builder, true)
	flatmsg.MessageFlatLargeAddField8(builder, label)
	flatmsg.MessageFlatLargeAddField9(builder, org)
	flatmsg.MessageFlatLargeAddField10(builder, true)
	flatmsg.MessageFlatLargeAddField11(builder, 1024)
	flatmsg.MessageFlatLargeAddField12(builder, 1024)
	flatmsg.MessageFlatLargeAddField13(builder, 1024)
	flatmsg.MessageFlatLargeAddField14(builder, 1024)
	flatmsg.MessageFlatLargeAddField15(builder, 1024)
	flatmsg.MessageFlatLargeAddField16(builder, 1024)
	flatmsg.MessageFlatLargeAddField17(builder, true)
	flatmsg.MessageFlatLargeAddField18(builder, label)
	flatmsg.MessageFlatLargeAddField19(builder, org)
	flatmsg.MessageFlatLargeAddField20(builder, true)
	flatmsg.MessageFlatLargeAddField21(builder, 1024)
	flatmsg.MessageFlatLargeAddField22(builder, 1024)
	flatmsg.MessageFlatLargeAddField23(builder, 1024)
	flatmsg.MessageFlatLargeAddField24(builder, 1024)
	flatmsg.MessageFlatLargeAddField25(builder, 1024)
	flatmsg.MessageFlatLargeAddField26(builder, 1024)
	flatmsg.MessageFlatLargeAddField27(builder, true)
	flatmsg.MessageFlatLargeAddField28(builder, label)
	flatmsg.MessageFlatLargeAddField29(builder, org)
	flatmsg.MessageFlatLargeAddField30(builder, true)
	flatmsg.MessageFlatLargeAddField31(builder, 1024)
	flatmsg.MessageFlatLargeAddField32(builder, 1024)
	flatmsg.MessageFlatLargeAddField33(builder, 1024)
	flatmsg.MessageFlatLargeAddField34(builder, 1024)
	flatmsg.MessageFlatLargeAddField35(builder, 1024)
	flatmsg.MessageFlatLargeAddField36(builder, 1024)
	flatmsg.MessageFlatLargeAddField37(builder, true)
	flatmsg.MessageFlatLargeAddField38(builder, label)
	flatmsg.MessageFlatLargeAddField39(builder, org)
	flatmsg.MessageFlatLargeAddField40(builder, true)
	flatmsg.MessageFlatLargeAddField41(builder, 1024)
	flatmsg.MessageFlatLargeAddField42(builder, 1024)
	flatmsg.MessageFlatLargeAddField43(builder, 1024)
	flatmsg.MessageFlatLargeAddField44(builder, 1024)
	flatmsg.MessageFlatLargeAddField45(builder, 1024)
	flatmsg.MessageFlatLargeAddField46(builder, 1024)
	flatmsg.MessageFlatLargeAddField47(builder, true)
	flatmsg.MessageFlatLargeAddField48(builder, label)
	flatmsg.MessageFlatLargeAddField49(builder, org)
	flatmsg.MessageFlatLargeAddField50(builder, true)
	flatmsg.MessageFlatLargeAddField51(builder, 1024)
	flatmsg.MessageFlatLargeAddField52(builder, 1024)
	flatmsg.MessageFlatLargeAddField53(builder, 1024)
	flatmsg.MessageFlatLargeAddField54(builder, 1024)
	flatmsg.MessageFlatLargeAddField55(builder, 1024)
	flatmsg.MessageFlatLargeAddField56(builder, 1024)
	flatmsg.MessageFlatLargeAddField57(builder, true)
	flatmsg.MessageFlatLargeAddField58(builder, label)
	flatmsg.MessageFlatLargeAddField59(builder, org)
	flatmsg.MessageFlatLargeAddField60(builder, true)
	flatmsg.MessageFlatLargeAddField61(builder, 1024)
	flatmsg.MessageFlatLargeAddField62(builder, 1024)
	flatmsg.MessageFlatLargeAddField63(builder, 1024)
	flatmsg.MessageFlatLargeAddField64(builder, 1024)
	flatmsg.MessageFlatLargeAddField65(builder, 1024)
	flatmsg.MessageFlatLargeAddField66(builder, 1024)
	flatmsg.MessageFlatLargeAddField67(builder, true)
	flatmsg.MessageFlatLargeAddField68(builder, label)
	flatmsg.MessageFlatLargeAddField69(builder, org)
	flatmsg.MessageFlatLargeAddField70(builder, true)
	flatmsg.MessageFlatLargeAddField71(builder, 1024)
	flatmsg.MessageFlatLargeAddField72(builder, 1024)
	flatmsg.MessageFlatLargeAddField73(builder, 1024)
	flatmsg.MessageFlatLargeAddField74(builder, 1024)
	flatmsg.MessageFlatLargeAddField75(builder, 1024)
	flatmsg.MessageFlatLargeAddField76(builder, 1024)
	flatmsg.MessageFlatLargeAddField77(builder, true)
	flatmsg.MessageFlatLargeAddField78(builder, label)
	flatmsg.MessageFlatLargeAddField79(builder, org)
	flatmsg.MessageFlatLargeAddField80(builder, true)
	flatmsg.MessageFlatLargeAddField81(builder, 1024)
	flatmsg.MessageFlatLargeAddField82(builder, 1024)
	flatmsg.MessageFlatLargeAddField83(builder, 1024)
	flatmsg.MessageFlatLargeAddField84(builder, 1024)
	flatmsg.MessageFlatLargeAddField85(builder, 1024)
	flatmsg.MessageFlatLargeAddField86(builder, 1024)
	flatmsg.MessageFlatLargeAddField87(builder, true)
	flatmsg.MessageFlatLargeAddField88(builder, label)
	flatmsg.MessageFlatLargeAddField89(builder, org)
	flatmsg.MessageFlatLargeAddField90(builder, true)
	flatmsg.MessageFlatLargeAddField91(builder, 1024)
	flatmsg.MessageFlatLargeAddField92(builder, 1024)
	flatmsg.MessageFlatLargeAddField93(builder, 1024)
	flatmsg.MessageFlatLargeAddField94(builder, 1024)
	flatmsg.MessageFlatLargeAddField95(builder, 1024)
	flatmsg.MessageFlatLargeAddField96(builder, 1024)
	flatmsg.MessageFlatLargeAddField97(builder, true)
	flatmsg.MessageFlatLargeAddField98(builder, label)
	flatmsg.MessageFlatLargeAddField99(builder, org)
	flatmsg.MessageFlatLargeAddField100(builder, true)
	flatmsg.MessageFlatLargeAddField101(builder, 1024)
	flatmsg.MessageFlatLargeAddField102(builder, 1024)
	flatmsg.MessageFlatLargeAddField103(builder, 1024)
	flatmsg.MessageFlatLargeAddField104(builder, 1024)
	flatmsg.MessageFlatLargeAddField105(builder, 1024)
	flatmsg.MessageFlatLargeAddField106(builder, 1024)
	flatmsg.MessageFlatLargeAddField107(builder, true)
	flatmsg.MessageFlatLargeAddField108(builder, label)
	flatmsg.MessageFlatLargeAddField109(builder, org)
	flatmsg.MessageFlatLargeAddField110(builder, true)
	flatmsg.MessageFlatLargeAddField111(builder, 1024)
	flatmsg.MessageFlatLargeAddField112(builder, 1024)
	flatmsg.MessageFlatLargeAddField113(builder, 1024)
	flatmsg.MessageFlatLargeAddField114(builder, 1024)
	flatmsg.MessageFlatLargeAddField115(builder, 1024)
	flatmsg.MessageFlatLargeAddField116(builder, 1024)
	flatmsg.MessageFlatLargeAddField117(builder, true)
	flatmsg.MessageFlatLargeAddField118(builder, label)
	flatmsg.MessageFlatLargeAddField119(builder, org)
	return flatmsg.MessageFlatLargeEnd(builder)
}

func NewFlatMessageFlatXLarge(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	label := builder.CreateString("Bench")
	org := builder.CreateString("Mark")

	flatmsg.MessageFlatXLargeStart(builder)
	flatmsg.MessageFlatXLargeAddField1(builder, 1024)
	flatmsg.MessageFlatXLargeAddField2(builder, 1024)
	flatmsg.MessageFlatXLargeAddField3(builder, 1024)
	flatmsg.MessageFlatXLargeAddField4(builder, 1024)
	flatmsg.MessageFlatXLargeAddField5(builder, 1024)
	flatmsg.MessageFlatXLargeAddField6(builder, 1024)
	flatmsg.MessageFlatXLargeAddField7(builder, true)
	flatmsg.MessageFlatXLargeAddField8(builder, label)
	flatmsg.MessageFlatXLargeAddField9(builder, org)
	flatmsg.MessageFlatXLargeAddField10(builder, true)
	flatmsg.MessageFlatXLargeAddField11(builder, 1024)
	flatmsg.MessageFlatXLargeAddField12(builder, 1024)
	flatmsg.MessageFlatXLargeAddField13(builder, 1024)
	flatmsg.MessageFlatXLargeAddField14(builder, 1024)
	flatmsg.MessageFlatXLargeAddField15(builder, 1024)
	flatmsg.MessageFlatXLargeAddField16(builder, 1024)
	flatmsg.MessageFlatXLargeAddField17(builder, true)
	flatmsg.MessageFlatXLargeAddField18(builder, label)
	flatmsg.MessageFlatXLargeAddField19(builder, org)
	flatmsg.MessageFlatXLargeAddField20(builder, true)
	flatmsg.MessageFlatXLargeAddField21(builder, 1024)
	flatmsg.MessageFlatXLargeAddField22(builder, 1024)
	flatmsg.MessageFlatXLargeAddField23(builder, 1024)
	flatmsg.MessageFlatXLargeAddField24(builder, 1024)
	flatmsg.MessageFlatXLargeAddField25(builder, 1024)
	flatmsg.MessageFlatXLargeAddField26(builder, 1024)
	flatmsg.MessageFlatXLargeAddField27(builder, true)
	flatmsg.MessageFlatXLargeAddField28(builder, label)
	flatmsg.MessageFlatXLargeAddField29(builder, org)
	flatmsg.MessageFlatXLargeAddField30(builder, true)
	flatmsg.MessageFlatXLargeAddField31(builder, 1024)
	flatmsg.MessageFlatXLargeAddField32(builder, 1024)
	flatmsg.MessageFlatXLargeAddField33(builder, 1024)
	flatmsg.MessageFlatXLargeAddField34(builder, 1024)
	flatmsg.MessageFlatXLargeAddField35(builder, 1024)
	flatmsg.MessageFlatXLargeAddField36(builder, 1024)
	flatmsg.MessageFlatXLargeAddField37(builder, true)
	flatmsg.MessageFlatXLargeAddField38(builder, label)
	flatmsg.MessageFlatXLargeAddField39(builder, org)
	flatmsg.MessageFlatXLargeAddField40(builder, true)
	flatmsg.MessageFlatXLargeAddField41(builder, 1024)
	flatmsg.MessageFlatXLargeAddField42(builder, 1024)
	flatmsg.MessageFlatXLargeAddField43(builder, 1024)
	flatmsg.MessageFlatXLargeAddField44(builder, 1024)
	flatmsg.MessageFlatXLargeAddField45(builder, 1024)
	flatmsg.MessageFlatXLargeAddField46(builder, 1024)
	flatmsg.MessageFlatXLargeAddField47(builder, true)
	flatmsg.MessageFlatXLargeAddField48(builder, label)
	flatmsg.MessageFlatXLargeAddField49(builder, org)
	flatmsg.MessageFlatXLargeAddField50(builder, true)
	flatmsg.MessageFlatXLargeAddField51(builder, 1024)
	flatmsg.MessageFlatXLargeAddField52(builder, 1024)
	flatmsg.MessageFlatXLargeAddField53(builder, 1024)
	flatmsg.MessageFlatXLargeAddField54(builder, 1024)
	flatmsg.MessageFlatXLargeAddField55(builder, 1024)
	flatmsg.MessageFlatXLargeAddField56(builder, 1024)
	flatmsg.MessageFlatXLargeAddField57(builder, true)
	flatmsg.MessageFlatXLargeAddField58(builder, label)
	flatmsg.MessageFlatXLargeAddField59(builder, org)
	flatmsg.MessageFlatXLargeAddField60(builder, true)
	flatmsg.MessageFlatXLargeAddField61(builder, 1024)
	flatmsg.MessageFlatXLargeAddField62(builder, 1024)
	flatmsg.MessageFlatXLargeAddField63(builder, 1024)
	flatmsg.MessageFlatXLargeAddField64(builder, 1024)
	flatmsg.MessageFlatXLargeAddField65(builder, 1024)
	flatmsg.MessageFlatXLargeAddField66(builder, 1024)
	flatmsg.MessageFlatXLargeAddField67(builder, true)
	flatmsg.MessageFlatXLargeAddField68(builder, label)
	flatmsg.MessageFlatXLargeAddField69(builder, org)
	flatmsg.MessageFlatXLargeAddField70(builder, true)
	flatmsg.MessageFlatXLargeAddField71(builder, 1024)
	flatmsg.MessageFlatXLargeAddField72(builder, 1024)
	flatmsg.MessageFlatXLargeAddField73(builder, 1024)
	flatmsg.MessageFlatXLargeAddField74(builder, 1024)
	flatmsg.MessageFlatXLargeAddField75(builder, 1024)
	flatmsg.MessageFlatXLargeAddField76(builder, 1024)
	flatmsg.MessageFlatXLargeAddField77(builder, true)
	flatmsg.MessageFlatXLargeAddField78(builder, label)
	flatmsg.MessageFlatXLargeAddField79(builder, org)
	flatmsg.MessageFlatXLargeAddField80(builder, true)
	flatmsg.MessageFlatXLargeAddField81(builder, 1024)
	flatmsg.MessageFlatXLargeAddField82(builder, 1024)
	flatmsg.MessageFlatXLargeAddField83(builder, 1024)
	flatmsg.MessageFlatXLargeAddField84(builder, 1024)
	flatmsg.MessageFlatXLargeAddField85(builder, 1024)
	flatmsg.MessageFlatXLargeAddField86(builder, 1024)
	flatmsg.MessageFlatXLargeAddField87(builder, true)
	flatmsg.MessageFlatXLargeAddField88(builder, label)
	flatmsg.MessageFlatXLargeAddField89(builder, org)
	flatmsg.MessageFlatXLargeAddField90(builder, true)
	flatmsg.MessageFlatXLargeAddField91(builder, 1024)
	flatmsg.MessageFlatXLargeAddField92(builder, 1024)
	flatmsg.MessageFlatXLargeAddField93(builder, 1024)
	flatmsg.MessageFlatXLargeAddField94(builder, 1024)
	flatmsg.MessageFlatXLargeAddField95(builder, 1024)
	flatmsg.MessageFlatXLargeAddField96(builder, 1024)
	flatmsg.MessageFlatXLargeAddField97(builder, true)
	flatmsg.MessageFlatXLargeAddField98(builder, label)
	flatmsg.MessageFlatXLargeAddField99(builder, org)
	flatmsg.MessageFlatXLargeAddField100(builder, true)
	flatmsg.MessageFlatXLargeAddField101(builder, 1024)
	flatmsg.MessageFlatXLargeAddField102(builder, 1024)
	flatmsg.MessageFlatXLargeAddField103(builder, 1024)
	flatmsg.MessageFlatXLargeAddField104(builder, 1024)
	flatmsg.MessageFlatXLargeAddField105(builder, 1024)
	flatmsg.MessageFlatXLargeAddField106(builder, 1024)
	flatmsg.MessageFlatXLargeAddField107(builder, true)
	flatmsg.MessageFlatXLargeAddField108(builder, label)
	flatmsg.MessageFlatXLargeAddField109(builder, org)
	flatmsg.MessageFlatXLargeAddField110(builder, true)
	flatmsg.MessageFlatXLargeAddField111(builder, 1024)
	flatmsg.MessageFlatXLargeAddField112(builder, 1024)
	flatmsg.MessageFlatXLargeAddField113(builder, 1024)
	flatmsg.MessageFlatXLargeAddField114(builder, 1024)
	flatmsg.MessageFlatXLargeAddField115(builder, 1024)
	flatmsg.MessageFlatXLargeAddField116(builder, 1024)
	flatmsg.MessageFlatXLargeAddField117(builder, true)
	flatmsg.MessageFlatXLargeAddField118(builder, label)
	flatmsg.MessageFlatXLargeAddField119(builder, org)
	flatmsg.MessageFlatXLargeAddField120(builder, true)
	flatmsg.MessageFlatXLargeAddField121(builder, 1024)
	flatmsg.MessageFlatXLargeAddField122(builder, 1024)
	flatmsg.MessageFlatXLargeAddField123(builder, 1024)
	flatmsg.MessageFlatXLargeAddField124(builder, 1024)
	flatmsg.MessageFlatXLargeAddField125(builder, 1024)
	flatmsg.MessageFlatXLargeAddField126(builder, 1024)
	flatmsg.MessageFlatXLargeAddField127(builder, true)
	flatmsg.MessageFlatXLargeAddField128(builder, label)
	flatmsg.MessageFlatXLargeAddField129(builder, org)
	flatmsg.MessageFlatXLargeAddField130(builder, true)
	flatmsg.MessageFlatXLargeAddField131(builder, 1024)
	flatmsg.MessageFlatXLargeAddField132(builder, 1024)
	flatmsg.MessageFlatXLargeAddField133(builder, 1024)
	flatmsg.MessageFlatXLargeAddField134(builder, 1024)
	flatmsg.MessageFlatXLargeAddField135(builder, 1024)
	flatmsg.MessageFlatXLargeAddField136(builder, 1024)
	flatmsg.MessageFlatXLargeAddField137(builder, true)
	flatmsg.MessageFlatXLargeAddField138(builder, label)
	flatmsg.MessageFlatXLargeAddField139(builder, org)
	flatmsg.MessageFlatXLargeAddField140(builder, true)
	flatmsg.MessageFlatXLargeAddField141(builder, 1024)
	flatmsg.MessageFlatXLargeAddField142(builder, 1024)
	flatmsg.MessageFlatXLargeAddField143(builder, 1024)
	flatmsg.MessageFlatXLargeAddField144(builder, 1024)
	flatmsg.MessageFlatXLargeAddField145(builder, 1024)
	flatmsg.MessageFlatXLargeAddField146(builder, 1024)
	flatmsg.MessageFlatXLargeAddField147(builder, true)
	flatmsg.MessageFlatXLargeAddField148(builder, label)
	flatmsg.MessageFlatXLargeAddField149(builder, org)
	flatmsg.MessageFlatXLargeAddField150(builder, true)
	flatmsg.MessageFlatXLargeAddField151(builder, 1024)
	flatmsg.MessageFlatXLargeAddField152(builder, 1024)
	flatmsg.MessageFlatXLargeAddField153(builder, 1024)
	flatmsg.MessageFlatXLargeAddField154(builder, 1024)
	flatmsg.MessageFlatXLargeAddField155(builder, 1024)
	flatmsg.MessageFlatXLargeAddField156(builder, 1024)
	flatmsg.MessageFlatXLargeAddField157(builder, true)
	flatmsg.MessageFlatXLargeAddField158(builder, label)
	flatmsg.MessageFlatXLargeAddField159(builder, org)
	flatmsg.MessageFlatXLargeAddField160(builder, true)
	flatmsg.MessageFlatXLargeAddField161(builder, 1024)
	flatmsg.MessageFlatXLargeAddField162(builder, 1024)
	flatmsg.MessageFlatXLargeAddField163(builder, 1024)
	flatmsg.MessageFlatXLargeAddField164(builder, 1024)
	flatmsg.MessageFlatXLargeAddField165(builder, 1024)
	flatmsg.MessageFlatXLargeAddField166(builder, 1024)
	flatmsg.MessageFlatXLargeAddField167(builder, true)
	flatmsg.MessageFlatXLargeAddField168(builder, label)
	flatmsg.MessageFlatXLargeAddField169(builder, org)
	flatmsg.MessageFlatXLargeAddField170(builder, true)
	flatmsg.MessageFlatXLargeAddField171(builder, 1024)
	flatmsg.MessageFlatXLargeAddField172(builder, 1024)
	flatmsg.MessageFlatXLargeAddField173(builder, 1024)
	flatmsg.MessageFlatXLargeAddField174(builder, 1024)
	flatmsg.MessageFlatXLargeAddField175(builder, 1024)
	flatmsg.MessageFlatXLargeAddField176(builder, 1024)
	flatmsg.MessageFlatXLargeAddField177(builder, true)
	flatmsg.MessageFlatXLargeAddField178(builder, label)
	flatmsg.MessageFlatXLargeAddField179(builder, org)
	flatmsg.MessageFlatXLargeAddField180(builder, true)
	flatmsg.MessageFlatXLargeAddField181(builder, 1024)
	flatmsg.MessageFlatXLargeAddField182(builder, 1024)
	flatmsg.MessageFlatXLargeAddField183(builder, 1024)
	flatmsg.MessageFlatXLargeAddField184(builder, 1024)
	flatmsg.MessageFlatXLargeAddField185(builder, 1024)
	flatmsg.MessageFlatXLargeAddField186(builder, 1024)
	flatmsg.MessageFlatXLargeAddField187(builder, true)
	flatmsg.MessageFlatXLargeAddField188(builder, label)
	flatmsg.MessageFlatXLargeAddField189(builder, org)
	flatmsg.MessageFlatXLargeAddField190(builder, true)
	flatmsg.MessageFlatXLargeAddField191(builder, 1024)
	flatmsg.MessageFlatXLargeAddField192(builder, 1024)
	flatmsg.MessageFlatXLargeAddField193(builder, 1024)
	flatmsg.MessageFlatXLargeAddField194(builder, 1024)
	flatmsg.MessageFlatXLargeAddField195(builder, 1024)
	flatmsg.MessageFlatXLargeAddField196(builder, 1024)
	flatmsg.MessageFlatXLargeAddField197(builder, true)
	flatmsg.MessageFlatXLargeAddField198(builder, label)
	flatmsg.MessageFlatXLargeAddField199(builder, org)
	flatmsg.MessageFlatXLargeAddField200(builder, true)
	flatmsg.MessageFlatXLargeAddField201(builder, 1024)
	flatmsg.MessageFlatXLargeAddField202(builder, 1024)
	flatmsg.MessageFlatXLargeAddField203(builder, 1024)
	flatmsg.MessageFlatXLargeAddField204(builder, 1024)
	flatmsg.MessageFlatXLargeAddField205(builder, 1024)
	flatmsg.MessageFlatXLargeAddField206(builder, 1024)
	flatmsg.MessageFlatXLargeAddField207(builder, true)
	flatmsg.MessageFlatXLargeAddField208(builder, label)
	flatmsg.MessageFlatXLargeAddField209(builder, org)
	flatmsg.MessageFlatXLargeAddField210(builder, true)
	flatmsg.MessageFlatXLargeAddField211(builder, 1024)
	flatmsg.MessageFlatXLargeAddField212(builder, 1024)
	flatmsg.MessageFlatXLargeAddField213(builder, 1024)
	flatmsg.MessageFlatXLargeAddField214(builder, 1024)
	flatmsg.MessageFlatXLargeAddField215(builder, 1024)
	flatmsg.MessageFlatXLargeAddField216(builder, 1024)
	flatmsg.MessageFlatXLargeAddField217(builder, true)
	flatmsg.MessageFlatXLargeAddField218(builder, label)
	flatmsg.MessageFlatXLargeAddField219(builder, org)
	flatmsg.MessageFlatXLargeAddField220(builder, true)
	flatmsg.MessageFlatXLargeAddField221(builder, 1024)
	flatmsg.MessageFlatXLargeAddField222(builder, 1024)
	flatmsg.MessageFlatXLargeAddField223(builder, 1024)
	flatmsg.MessageFlatXLargeAddField224(builder, 1024)
	flatmsg.MessageFlatXLargeAddField225(builder, 1024)
	flatmsg.MessageFlatXLargeAddField226(builder, 1024)
	flatmsg.MessageFlatXLargeAddField227(builder, true)
	flatmsg.MessageFlatXLargeAddField228(builder, label)
	flatmsg.MessageFlatXLargeAddField229(builder, org)
	flatmsg.MessageFlatXLargeAddField230(builder, true)
	flatmsg.MessageFlatXLargeAddField231(builder, 1024)
	flatmsg.MessageFlatXLargeAddField232(builder, 1024)
	flatmsg.MessageFlatXLargeAddField233(builder, 1024)
	flatmsg.MessageFlatXLargeAddField234(builder, 1024)
	flatmsg.MessageFlatXLargeAddField235(builder, 1024)
	flatmsg.MessageFlatXLargeAddField236(builder, 1024)
	flatmsg.MessageFlatXLargeAddField237(builder, true)
	flatmsg.MessageFlatXLargeAddField238(builder, label)
	flatmsg.MessageFlatXLargeAddField239(builder, org)
	flatmsg.MessageFlatXLargeAddField240(builder, true)
	flatmsg.MessageFlatXLargeAddField241(builder, 1024)
	flatmsg.MessageFlatXLargeAddField242(builder, 1024)
	flatmsg.MessageFlatXLargeAddField243(builder, 1024)
	flatmsg.MessageFlatXLargeAddField244(builder, 1024)
	flatmsg.MessageFlatXLargeAddField245(builder, 1024)
	flatmsg.MessageFlatXLargeAddField246(builder, 1024)
	flatmsg.MessageFlatXLargeAddField247(builder, true)
	flatmsg.MessageFlatXLargeAddField248(builder, label)
	flatmsg.MessageFlatXLargeAddField249(builder, org)
	flatmsg.MessageFlatXLargeAddField250(builder, true)
	flatmsg.MessageFlatXLargeAddField251(builder, 1024)
	flatmsg.MessageFlatXLargeAddField252(builder, 1024)
	flatmsg.MessageFlatXLargeAddField253(builder, 1024)
	flatmsg.MessageFlatXLargeAddField254(builder, 1024)
	flatmsg.MessageFlatXLargeAddField255(builder, 1024)
	flatmsg.MessageFlatXLargeAddField256(builder, 1024)
	flatmsg.MessageFlatXLargeAddField257(builder, true)
	flatmsg.MessageFlatXLargeAddField258(builder, label)
	flatmsg.MessageFlatXLargeAddField259(builder, org)
	flatmsg.MessageFlatXLargeAddField260(builder, true)
	flatmsg.MessageFlatXLargeAddField261(builder, 1024)
	flatmsg.MessageFlatXLargeAddField262(builder, 1024)
	flatmsg.MessageFlatXLargeAddField263(builder, 1024)
	flatmsg.MessageFlatXLargeAddField264(builder, 1024)
	flatmsg.MessageFlatXLargeAddField265(builder, 1024)
	flatmsg.MessageFlatXLargeAddField266(builder, 1024)
	flatmsg.MessageFlatXLargeAddField267(builder, true)
	flatmsg.MessageFlatXLargeAddField268(builder, label)
	flatmsg.MessageFlatXLargeAddField269(builder, org)
	flatmsg.MessageFlatXLargeAddField270(builder, true)
	flatmsg.MessageFlatXLargeAddField271(builder, 1024)
	flatmsg.MessageFlatXLargeAddField272(builder, 1024)
	flatmsg.MessageFlatXLargeAddField273(builder, 1024)
	flatmsg.MessageFlatXLargeAddField274(builder, 1024)
	flatmsg.MessageFlatXLargeAddField275(builder, 1024)
	flatmsg.MessageFlatXLargeAddField276(builder, 1024)
	flatmsg.MessageFlatXLargeAddField277(builder, true)
	flatmsg.MessageFlatXLargeAddField278(builder, label)
	flatmsg.MessageFlatXLargeAddField279(builder, org)
	flatmsg.MessageFlatXLargeAddField280(builder, true)
	flatmsg.MessageFlatXLargeAddField281(builder, 1024)
	flatmsg.MessageFlatXLargeAddField282(builder, 1024)
	flatmsg.MessageFlatXLargeAddField283(builder, 1024)
	flatmsg.MessageFlatXLargeAddField284(builder, 1024)
	flatmsg.MessageFlatXLargeAddField285(builder, 1024)
	flatmsg.MessageFlatXLargeAddField286(builder, 1024)
	flatmsg.MessageFlatXLargeAddField287(builder, true)
	flatmsg.MessageFlatXLargeAddField288(builder, label)
	flatmsg.MessageFlatXLargeAddField289(builder, org)
	flatmsg.MessageFlatXLargeAddField290(builder, true)
	flatmsg.MessageFlatXLargeAddField291(builder, 1024)
	flatmsg.MessageFlatXLargeAddField292(builder, 1024)
	flatmsg.MessageFlatXLargeAddField293(builder, 1024)
	flatmsg.MessageFlatXLargeAddField294(builder, 1024)
	flatmsg.MessageFlatXLargeAddField295(builder, 1024)
	flatmsg.MessageFlatXLargeAddField296(builder, 1024)
	flatmsg.MessageFlatXLargeAddField297(builder, true)
	flatmsg.MessageFlatXLargeAddField298(builder, label)
	flatmsg.MessageFlatXLargeAddField299(builder, org)
	flatmsg.MessageFlatXLargeAddField300(builder, true)
	flatmsg.MessageFlatXLargeAddField301(builder, 1024)
	flatmsg.MessageFlatXLargeAddField302(builder, 1024)
	flatmsg.MessageFlatXLargeAddField303(builder, 1024)
	flatmsg.MessageFlatXLargeAddField304(builder, 1024)
	flatmsg.MessageFlatXLargeAddField305(builder, 1024)
	flatmsg.MessageFlatXLargeAddField306(builder, 1024)
	flatmsg.MessageFlatXLargeAddField307(builder, true)
	flatmsg.MessageFlatXLargeAddField308(builder, label)
	flatmsg.MessageFlatXLargeAddField309(builder, org)
	flatmsg.MessageFlatXLargeAddField310(builder, true)
	flatmsg.MessageFlatXLargeAddField311(builder, 1024)
	flatmsg.MessageFlatXLargeAddField312(builder, 1024)
	flatmsg.MessageFlatXLargeAddField313(builder, 1024)
	flatmsg.MessageFlatXLargeAddField314(builder, 1024)
	flatmsg.MessageFlatXLargeAddField315(builder, 1024)
	flatmsg.MessageFlatXLargeAddField316(builder, 1024)
	flatmsg.MessageFlatXLargeAddField317(builder, true)
	flatmsg.MessageFlatXLargeAddField318(builder, label)
	flatmsg.MessageFlatXLargeAddField319(builder, org)
	flatmsg.MessageFlatXLargeAddField320(builder, true)
	flatmsg.MessageFlatXLargeAddField321(builder, 1024)
	flatmsg.MessageFlatXLargeAddField322(builder, 1024)
	flatmsg.MessageFlatXLargeAddField323(builder, 1024)
	flatmsg.MessageFlatXLargeAddField324(builder, 1024)
	flatmsg.MessageFlatXLargeAddField325(builder, 1024)
	flatmsg.MessageFlatXLargeAddField326(builder, 1024)
	flatmsg.MessageFlatXLargeAddField327(builder, true)
	flatmsg.MessageFlatXLargeAddField328(builder, label)
	flatmsg.MessageFlatXLargeAddField329(builder, org)
	flatmsg.MessageFlatXLargeAddField330(builder, true)
	flatmsg.MessageFlatXLargeAddField331(builder, 1024)
	flatmsg.MessageFlatXLargeAddField332(builder, 1024)
	flatmsg.MessageFlatXLargeAddField333(builder, 1024)
	flatmsg.MessageFlatXLargeAddField334(builder, 1024)
	flatmsg.MessageFlatXLargeAddField335(builder, 1024)
	flatmsg.MessageFlatXLargeAddField336(builder, 1024)
	flatmsg.MessageFlatXLargeAddField337(builder, true)
	flatmsg.MessageFlatXLargeAddField338(builder, label)
	flatmsg.MessageFlatXLargeAddField339(builder, org)
	flatmsg.MessageFlatXLargeAddField340(builder, true)
	flatmsg.MessageFlatXLargeAddField341(builder, 1024)
	flatmsg.MessageFlatXLargeAddField342(builder, 1024)
	flatmsg.MessageFlatXLargeAddField343(builder, 1024)
	flatmsg.MessageFlatXLargeAddField344(builder, 1024)
	flatmsg.MessageFlatXLargeAddField345(builder, 1024)
	flatmsg.MessageFlatXLargeAddField346(builder, 1024)
	flatmsg.MessageFlatXLargeAddField347(builder, true)
	flatmsg.MessageFlatXLargeAddField348(builder, label)
	flatmsg.MessageFlatXLargeAddField349(builder, org)
	flatmsg.MessageFlatXLargeAddField350(builder, true)
	flatmsg.MessageFlatXLargeAddField351(builder, 1024)
	flatmsg.MessageFlatXLargeAddField352(builder, 1024)
	flatmsg.MessageFlatXLargeAddField353(builder, 1024)
	flatmsg.MessageFlatXLargeAddField354(builder, 1024)
	flatmsg.MessageFlatXLargeAddField355(builder, 1024)
	flatmsg.MessageFlatXLargeAddField356(builder, 1024)
	flatmsg.MessageFlatXLargeAddField357(builder, true)
	flatmsg.MessageFlatXLargeAddField358(builder, label)
	flatmsg.MessageFlatXLargeAddField359(builder, org)
	flatmsg.MessageFlatXLargeAddField360(builder, true)
	flatmsg.MessageFlatXLargeAddField361(builder, 1024)
	flatmsg.MessageFlatXLargeAddField362(builder, 1024)
	flatmsg.MessageFlatXLargeAddField363(builder, 1024)
	flatmsg.MessageFlatXLargeAddField364(builder, 1024)
	flatmsg.MessageFlatXLargeAddField365(builder, 1024)
	flatmsg.MessageFlatXLargeAddField366(builder, 1024)
	flatmsg.MessageFlatXLargeAddField367(builder, true)
	flatmsg.MessageFlatXLargeAddField368(builder, label)
	flatmsg.MessageFlatXLargeAddField369(builder, org)
	flatmsg.MessageFlatXLargeAddField370(builder, true)
	flatmsg.MessageFlatXLargeAddField371(builder, 1024)
	flatmsg.MessageFlatXLargeAddField372(builder, 1024)
	flatmsg.MessageFlatXLargeAddField373(builder, 1024)
	flatmsg.MessageFlatXLargeAddField374(builder, 1024)
	flatmsg.MessageFlatXLargeAddField375(builder, 1024)
	flatmsg.MessageFlatXLargeAddField376(builder, 1024)
	flatmsg.MessageFlatXLargeAddField377(builder, true)
	flatmsg.MessageFlatXLargeAddField378(builder, label)
	flatmsg.MessageFlatXLargeAddField379(builder, org)
	flatmsg.MessageFlatXLargeAddField380(builder, true)
	flatmsg.MessageFlatXLargeAddField381(builder, 1024)
	flatmsg.MessageFlatXLargeAddField382(builder, 1024)
	flatmsg.MessageFlatXLargeAddField383(builder, 1024)
	flatmsg.MessageFlatXLargeAddField384(builder, 1024)
	flatmsg.MessageFlatXLargeAddField385(builder, 1024)
	flatmsg.MessageFlatXLargeAddField386(builder, 1024)
	flatmsg.MessageFlatXLargeAddField387(builder, true)
	flatmsg.MessageFlatXLargeAddField388(builder, label)
	flatmsg.MessageFlatXLargeAddField389(builder, org)
	flatmsg.MessageFlatXLargeAddField390(builder, true)
	flatmsg.MessageFlatXLargeAddField391(builder, 1024)
	flatmsg.MessageFlatXLargeAddField392(builder, 1024)
	flatmsg.MessageFlatXLargeAddField393(builder, 1024)
	flatmsg.MessageFlatXLargeAddField394(builder, 1024)
	flatmsg.MessageFlatXLargeAddField395(builder, 1024)
	flatmsg.MessageFlatXLargeAddField396(builder, 1024)
	flatmsg.MessageFlatXLargeAddField397(builder, true)
	flatmsg.MessageFlatXLargeAddField398(builder, label)
	flatmsg.MessageFlatXLargeAddField399(builder, org)
	flatmsg.MessageFlatXLargeAddField400(builder, true)
	flatmsg.MessageFlatXLargeAddField401(builder, 1024)
	flatmsg.MessageFlatXLargeAddField402(builder, 1024)
	flatmsg.MessageFlatXLargeAddField403(builder, 1024)
	flatmsg.MessageFlatXLargeAddField404(builder, 1024)
	flatmsg.MessageFlatXLargeAddField405(builder, 1024)
	flatmsg.MessageFlatXLargeAddField406(builder, 1024)
	flatmsg.MessageFlatXLargeAddField407(builder, true)
	flatmsg.MessageFlatXLargeAddField408(builder, label)
	flatmsg.MessageFlatXLargeAddField409(builder, org)
	flatmsg.MessageFlatXLargeAddField410(builder, true)
	flatmsg.MessageFlatXLargeAddField411(builder, 1024)
	flatmsg.MessageFlatXLargeAddField412(builder, 1024)
	flatmsg.MessageFlatXLargeAddField413(builder, 1024)
	flatmsg.MessageFlatXLargeAddField414(builder, 1024)
	flatmsg.MessageFlatXLargeAddField415(builder, 1024)
	flatmsg.MessageFlatXLargeAddField416(builder, 1024)
	flatmsg.MessageFlatXLargeAddField417(builder, true)
	flatmsg.MessageFlatXLargeAddField418(builder, label)
	flatmsg.MessageFlatXLargeAddField419(builder, org)
	flatmsg.MessageFlatXLargeAddField420(builder, true)
	flatmsg.MessageFlatXLargeAddField421(builder, 1024)
	flatmsg.MessageFlatXLargeAddField422(builder, 1024)
	flatmsg.MessageFlatXLargeAddField423(builder, 1024)
	flatmsg.MessageFlatXLargeAddField424(builder, 1024)
	flatmsg.MessageFlatXLargeAddField425(builder, 1024)
	flatmsg.MessageFlatXLargeAddField426(builder, 1024)
	flatmsg.MessageFlatXLargeAddField427(builder, true)
	flatmsg.MessageFlatXLargeAddField428(builder, label)
	flatmsg.MessageFlatXLargeAddField429(builder, org)
	flatmsg.MessageFlatXLargeAddField430(builder, true)
	flatmsg.MessageFlatXLargeAddField431(builder, 1024)
	flatmsg.MessageFlatXLargeAddField432(builder, 1024)
	flatmsg.MessageFlatXLargeAddField433(builder, 1024)
	flatmsg.MessageFlatXLargeAddField434(builder, 1024)
	flatmsg.MessageFlatXLargeAddField435(builder, 1024)
	flatmsg.MessageFlatXLargeAddField436(builder, 1024)
	flatmsg.MessageFlatXLargeAddField437(builder, true)
	flatmsg.MessageFlatXLargeAddField438(builder, label)
	flatmsg.MessageFlatXLargeAddField439(builder, org)
	flatmsg.MessageFlatXLargeAddField440(builder, true)
	flatmsg.MessageFlatXLargeAddField441(builder, 1024)
	flatmsg.MessageFlatXLargeAddField442(builder, 1024)
	flatmsg.MessageFlatXLargeAddField443(builder, 1024)
	flatmsg.MessageFlatXLargeAddField444(builder, 1024)
	flatmsg.MessageFlatXLargeAddField445(builder, 1024)
	flatmsg.MessageFlatXLargeAddField446(builder, 1024)
	flatmsg.MessageFlatXLargeAddField447(builder, true)
	flatmsg.MessageFlatXLargeAddField448(builder, label)
	flatmsg.MessageFlatXLargeAddField449(builder, org)
	flatmsg.MessageFlatXLargeAddField450(builder, true)
	flatmsg.MessageFlatXLargeAddField451(builder, 1024)
	flatmsg.MessageFlatXLargeAddField452(builder, 1024)
	flatmsg.MessageFlatXLargeAddField453(builder, 1024)
	flatmsg.MessageFlatXLargeAddField454(builder, 1024)
	flatmsg.MessageFlatXLargeAddField455(builder, 1024)
	flatmsg.MessageFlatXLargeAddField456(builder, 1024)
	flatmsg.MessageFlatXLargeAddField457(builder, true)
	flatmsg.MessageFlatXLargeAddField458(builder, label)
	flatmsg.MessageFlatXLargeAddField459(builder, org)
	flatmsg.MessageFlatXLargeAddField460(builder, true)
	flatmsg.MessageFlatXLargeAddField461(builder, 1024)
	flatmsg.MessageFlatXLargeAddField462(builder, 1024)
	flatmsg.MessageFlatXLargeAddField463(builder, 1024)
	flatmsg.MessageFlatXLargeAddField464(builder, 1024)
	flatmsg.MessageFlatXLargeAddField465(builder, 1024)
	flatmsg.MessageFlatXLargeAddField466(builder, 1024)
	flatmsg.MessageFlatXLargeAddField467(builder, true)
	flatmsg.MessageFlatXLargeAddField468(builder, label)
	flatmsg.MessageFlatXLargeAddField469(builder, org)
	flatmsg.MessageFlatXLargeAddField470(builder, true)
	flatmsg.MessageFlatXLargeAddField471(builder, 1024)
	flatmsg.MessageFlatXLargeAddField472(builder, 1024)
	flatmsg.MessageFlatXLargeAddField473(builder, 1024)
	flatmsg.MessageFlatXLargeAddField474(builder, 1024)
	flatmsg.MessageFlatXLargeAddField475(builder, 1024)
	flatmsg.MessageFlatXLargeAddField476(builder, 1024)
	flatmsg.MessageFlatXLargeAddField477(builder, true)
	flatmsg.MessageFlatXLargeAddField478(builder, label)
	flatmsg.MessageFlatXLargeAddField479(builder, org)
	flatmsg.MessageFlatXLargeAddField480(builder, true)
	flatmsg.MessageFlatXLargeAddField481(builder, 1024)
	flatmsg.MessageFlatXLargeAddField482(builder, 1024)
	flatmsg.MessageFlatXLargeAddField483(builder, 1024)
	flatmsg.MessageFlatXLargeAddField484(builder, 1024)
	flatmsg.MessageFlatXLargeAddField485(builder, 1024)
	flatmsg.MessageFlatXLargeAddField486(builder, 1024)
	flatmsg.MessageFlatXLargeAddField487(builder, true)
	flatmsg.MessageFlatXLargeAddField488(builder, label)
	flatmsg.MessageFlatXLargeAddField489(builder, org)
	flatmsg.MessageFlatXLargeAddField490(builder, true)
	flatmsg.MessageFlatXLargeAddField491(builder, 1024)
	flatmsg.MessageFlatXLargeAddField492(builder, 1024)
	flatmsg.MessageFlatXLargeAddField493(builder, 1024)
	flatmsg.MessageFlatXLargeAddField494(builder, 1024)
	flatmsg.MessageFlatXLargeAddField495(builder, 1024)
	flatmsg.MessageFlatXLargeAddField496(builder, 1024)
	flatmsg.MessageFlatXLargeAddField497(builder, true)
	flatmsg.MessageFlatXLargeAddField498(builder, label)
	flatmsg.MessageFlatXLargeAddField499(builder, org)
	flatmsg.MessageFlatXLargeAddField500(builder, true)
	flatmsg.MessageFlatXLargeAddField501(builder, 1024)
	flatmsg.MessageFlatXLargeAddField502(builder, 1024)
	flatmsg.MessageFlatXLargeAddField503(builder, 1024)
	flatmsg.MessageFlatXLargeAddField504(builder, 1024)
	flatmsg.MessageFlatXLargeAddField505(builder, 1024)
	flatmsg.MessageFlatXLargeAddField506(builder, 1024)
	flatmsg.MessageFlatXLargeAddField507(builder, true)
	flatmsg.MessageFlatXLargeAddField508(builder, label)
	flatmsg.MessageFlatXLargeAddField509(builder, org)
	flatmsg.MessageFlatXLargeAddField510(builder, true)
	flatmsg.MessageFlatXLargeAddField511(builder, 1024)
	flatmsg.MessageFlatXLargeAddField512(builder, 1024)
	flatmsg.MessageFlatXLargeAddField513(builder, 1024)
	flatmsg.MessageFlatXLargeAddField514(builder, 1024)
	flatmsg.MessageFlatXLargeAddField515(builder, 1024)
	flatmsg.MessageFlatXLargeAddField516(builder, 1024)
	flatmsg.MessageFlatXLargeAddField517(builder, true)
	flatmsg.MessageFlatXLargeAddField518(builder, label)
	flatmsg.MessageFlatXLargeAddField519(builder, org)
	flatmsg.MessageFlatXLargeAddField520(builder, true)
	flatmsg.MessageFlatXLargeAddField521(builder, 1024)
	flatmsg.MessageFlatXLargeAddField522(builder, 1024)
	flatmsg.MessageFlatXLargeAddField523(builder, 1024)
	flatmsg.MessageFlatXLargeAddField524(builder, 1024)
	flatmsg.MessageFlatXLargeAddField525(builder, 1024)
	flatmsg.MessageFlatXLargeAddField526(builder, 1024)
	flatmsg.MessageFlatXLargeAddField527(builder, true)
	flatmsg.MessageFlatXLargeAddField528(builder, label)
	flatmsg.MessageFlatXLargeAddField529(builder, org)
	flatmsg.MessageFlatXLargeAddField530(builder, true)
	flatmsg.MessageFlatXLargeAddField531(builder, 1024)
	flatmsg.MessageFlatXLargeAddField532(builder, 1024)
	flatmsg.MessageFlatXLargeAddField533(builder, 1024)
	flatmsg.MessageFlatXLargeAddField534(builder, 1024)
	flatmsg.MessageFlatXLargeAddField535(builder, 1024)
	flatmsg.MessageFlatXLargeAddField536(builder, 1024)
	flatmsg.MessageFlatXLargeAddField537(builder, true)
	flatmsg.MessageFlatXLargeAddField538(builder, label)
	flatmsg.MessageFlatXLargeAddField539(builder, org)
	flatmsg.MessageFlatXLargeAddField540(builder, true)
	flatmsg.MessageFlatXLargeAddField541(builder, 1024)
	flatmsg.MessageFlatXLargeAddField542(builder, 1024)
	flatmsg.MessageFlatXLargeAddField543(builder, 1024)
	flatmsg.MessageFlatXLargeAddField544(builder, 1024)
	flatmsg.MessageFlatXLargeAddField545(builder, 1024)
	flatmsg.MessageFlatXLargeAddField546(builder, 1024)
	flatmsg.MessageFlatXLargeAddField547(builder, true)
	flatmsg.MessageFlatXLargeAddField548(builder, label)
	flatmsg.MessageFlatXLargeAddField549(builder, org)
	flatmsg.MessageFlatXLargeAddField550(builder, true)
	flatmsg.MessageFlatXLargeAddField551(builder, 1024)
	flatmsg.MessageFlatXLargeAddField552(builder, 1024)
	flatmsg.MessageFlatXLargeAddField553(builder, 1024)
	flatmsg.MessageFlatXLargeAddField554(builder, 1024)
	flatmsg.MessageFlatXLargeAddField555(builder, 1024)
	flatmsg.MessageFlatXLargeAddField556(builder, 1024)
	flatmsg.MessageFlatXLargeAddField557(builder, true)
	flatmsg.MessageFlatXLargeAddField558(builder, label)
	flatmsg.MessageFlatXLargeAddField559(builder, org)
	flatmsg.MessageFlatXLargeAddField560(builder, true)
	flatmsg.MessageFlatXLargeAddField561(builder, 1024)
	flatmsg.MessageFlatXLargeAddField562(builder, 1024)
	flatmsg.MessageFlatXLargeAddField563(builder, 1024)
	flatmsg.MessageFlatXLargeAddField564(builder, 1024)
	flatmsg.MessageFlatXLargeAddField565(builder, 1024)
	flatmsg.MessageFlatXLargeAddField566(builder, 1024)
	flatmsg.MessageFlatXLargeAddField567(builder, true)
	flatmsg.MessageFlatXLargeAddField568(builder, label)
	flatmsg.MessageFlatXLargeAddField569(builder, org)
	flatmsg.MessageFlatXLargeAddField570(builder, true)
	flatmsg.MessageFlatXLargeAddField571(builder, 1024)
	flatmsg.MessageFlatXLargeAddField572(builder, 1024)
	flatmsg.MessageFlatXLargeAddField573(builder, 1024)
	flatmsg.MessageFlatXLargeAddField574(builder, 1024)
	flatmsg.MessageFlatXLargeAddField575(builder, 1024)
	flatmsg.MessageFlatXLargeAddField576(builder, 1024)
	flatmsg.MessageFlatXLargeAddField577(builder, true)
	flatmsg.MessageFlatXLargeAddField578(builder, label)
	flatmsg.MessageFlatXLargeAddField579(builder, org)
	flatmsg.MessageFlatXLargeAddField580(builder, true)
	flatmsg.MessageFlatXLargeAddField581(builder, 1024)
	flatmsg.MessageFlatXLargeAddField582(builder, 1024)
	flatmsg.MessageFlatXLargeAddField583(builder, 1024)
	flatmsg.MessageFlatXLargeAddField584(builder, 1024)
	flatmsg.MessageFlatXLargeAddField585(builder, 1024)
	flatmsg.MessageFlatXLargeAddField586(builder, 1024)
	flatmsg.MessageFlatXLargeAddField587(builder, true)
	flatmsg.MessageFlatXLargeAddField588(builder, label)
	flatmsg.MessageFlatXLargeAddField589(builder, org)
	flatmsg.MessageFlatXLargeAddField590(builder, true)
	flatmsg.MessageFlatXLargeAddField591(builder, 1024)
	flatmsg.MessageFlatXLargeAddField592(builder, 1024)
	flatmsg.MessageFlatXLargeAddField593(builder, 1024)
	flatmsg.MessageFlatXLargeAddField594(builder, 1024)
	flatmsg.MessageFlatXLargeAddField595(builder, 1024)
	flatmsg.MessageFlatXLargeAddField596(builder, 1024)
	flatmsg.MessageFlatXLargeAddField597(builder, true)
	flatmsg.MessageFlatXLargeAddField598(builder, label)
	flatmsg.MessageFlatXLargeAddField599(builder, org)
	return flatmsg.MessageFlatXLargeEnd(builder)
}

func NewFlatMessageFlatXLargePartial(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	label := builder.CreateString("Bench")

	flatmsg.MessageFlatXLargeStart(builder)
	flatmsg.MessageFlatXLargeAddField1(builder, 1024)
	flatmsg.MessageFlatXLargeAddField2(builder, 1024)
	flatmsg.MessageFlatXLargeAddField3(builder, 1024)
	flatmsg.MessageFlatXLargeAddField4(builder, 1024)
	flatmsg.MessageFlatXLargeAddField5(builder, 1024)
	flatmsg.MessageFlatXLargeAddField6(builder, 1024)
	flatmsg.MessageFlatXLargeAddField7(builder, true)
	flatmsg.MessageFlatXLargeAddField8(builder, label)
	return flatmsg.MessageFlatXLargeEnd(builder)
}

func NewFlatMessageFlatXLargePartial64(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	label := builder.CreateString("Bench")
	org := builder.CreateString("Mark")

	flatmsg.MessageFlatXLargeStart(builder)
	flatmsg.MessageFlatXLargeAddField1(builder, 1024)
	flatmsg.MessageFlatXLargeAddField2(builder, 1024)
	flatmsg.MessageFlatXLargeAddField3(builder, 1024)
	flatmsg.MessageFlatXLargeAddField4(builder, 1024)
	flatmsg.MessageFlatXLargeAddField5(builder, 1024)
	flatmsg.MessageFlatXLargeAddField6(builder, 1024)
	flatmsg.MessageFlatXLargeAddField7(builder, true)
	flatmsg.MessageFlatXLargeAddField8(builder, label)
	flatmsg.MessageFlatXLargeAddField9(builder, org)
	flatmsg.MessageFlatXLargeAddField10(builder, true)
	flatmsg.MessageFlatXLargeAddField11(builder, 1024)
	flatmsg.MessageFlatXLargeAddField12(builder, 1024)
	flatmsg.MessageFlatXLargeAddField13(builder, 1024)
	flatmsg.MessageFlatXLargeAddField14(builder, 1024)
	flatmsg.MessageFlatXLargeAddField15(builder, 1024)
	flatmsg.MessageFlatXLargeAddField16(builder, 1024)
	flatmsg.MessageFlatXLargeAddField17(builder, true)
	flatmsg.MessageFlatXLargeAddField18(builder, label)
	flatmsg.MessageFlatXLargeAddField19(builder, org)
	flatmsg.MessageFlatXLargeAddField20(builder, true)
	flatmsg.MessageFlatXLargeAddField21(builder, 1024)
	flatmsg.MessageFlatXLargeAddField22(builder, 1024)
	flatmsg.MessageFlatXLargeAddField23(builder, 1024)
	flatmsg.MessageFlatXLargeAddField24(builder, 1024)
	flatmsg.MessageFlatXLargeAddField25(builder, 1024)
	flatmsg.MessageFlatXLargeAddField26(builder, 1024)
	flatmsg.MessageFlatXLargeAddField27(builder, true)
	flatmsg.MessageFlatXLargeAddField28(builder, label)
	flatmsg.MessageFlatXLargeAddField29(builder, org)
	flatmsg.MessageFlatXLargeAddField30(builder, true)
	flatmsg.MessageFlatXLargeAddField31(builder, 1024)
	flatmsg.MessageFlatXLargeAddField32(builder, 1024)
	flatmsg.MessageFlatXLargeAddField33(builder, 1024)
	flatmsg.MessageFlatXLargeAddField34(builder, 1024)
	flatmsg.MessageFlatXLargeAddField35(builder, 1024)
	flatmsg.MessageFlatXLargeAddField36(builder, 1024)
	flatmsg.MessageFlatXLargeAddField37(builder, true)
	flatmsg.MessageFlatXLargeAddField38(builder, label)
	flatmsg.MessageFlatXLargeAddField39(builder, org)
	flatmsg.MessageFlatXLargeAddField40(builder, true)
	flatmsg.MessageFlatXLargeAddField41(builder, 1024)
	flatmsg.MessageFlatXLargeAddField42(builder, 1024)
	flatmsg.MessageFlatXLargeAddField43(builder, 1024)
	flatmsg.MessageFlatXLargeAddField44(builder, 1024)
	flatmsg.MessageFlatXLargeAddField45(builder, 1024)
	flatmsg.MessageFlatXLargeAddField46(builder, 1024)
	flatmsg.MessageFlatXLargeAddField47(builder, true)
	flatmsg.MessageFlatXLargeAddField48(builder, label)
	flatmsg.MessageFlatXLargeAddField49(builder, org)
	flatmsg.MessageFlatXLargeAddField50(builder, true)
	flatmsg.MessageFlatXLargeAddField51(builder, 1024)
	flatmsg.MessageFlatXLargeAddField52(builder, 1024)
	flatmsg.MessageFlatXLargeAddField53(builder, 1024)
	flatmsg.MessageFlatXLargeAddField54(builder, 1024)
	flatmsg.MessageFlatXLargeAddField55(builder, 1024)
	flatmsg.MessageFlatXLargeAddField56(builder, 1024)
	flatmsg.MessageFlatXLargeAddField57(builder, true)
	flatmsg.MessageFlatXLargeAddField58(builder, label)
	flatmsg.MessageFlatXLargeAddField59(builder, org)
	flatmsg.MessageFlatXLargeAddField60(builder, true)
	flatmsg.MessageFlatXLargeAddField61(builder, 1024)
	flatmsg.MessageFlatXLargeAddField62(builder, 1024)
	flatmsg.MessageFlatXLargeAddField63(builder, 1024)
	flatmsg.MessageFlatXLargeAddField64(builder, 1024)
	return flatmsg.MessageFlatXLargeEnd(builder)
}

/*
PROTOBUF LENGTHS
*/

func TestNestProtobuf_Size(t *testing.T) {
	message := NewProtoMessageNest()

	msg, _ := proto.Marshal(&message)
	t.Log(len(msg))
}

func TestFlatProtobuf_Size(t *testing.T) {
	message := NewProtoMessageFlat()

	msg, _ := proto.Marshal(&message)
	t.Log(len(msg))
}

func TestLargeProtobuf_Size(t *testing.T) {
	message := NewProtoMessageFlatLarge()

	msg, _ := proto.Marshal(&message)
	t.Log(len(msg))
}

func TestXLargeProtobuf_Size(t *testing.T) {
	message := NewProtoMessageFlatXLarge()

	msg, _ := proto.Marshal(&message)
	t.Log(len(msg))
}

func TestXLargePartialProtobuf_Size(t *testing.T) {
	message := NewProtoMessageFlatXLargePartial()

	msg, _ := proto.Marshal(&message)
	t.Log(len(msg))
}

func TestXLargePartial64Protobuf_Size(t *testing.T) {
	message := NewProtoMessageFlatXLargePartial64()

	msg, _ := proto.Marshal(&message)
	t.Log(len(msg))
}

func TestGroupProtobuf_Size(t *testing.T) {
	message := NewProtoMessageGroup()

	msg, _ := proto.Marshal(&message)
	t.Log(len(msg))
}

func TestGroupPartialProtobuf_Size(t *testing.T) {
	message := NewProtoMessageGroupPartial()

	msg, _ := proto.Marshal(&message)
	t.Log(len(msg))
}

func TestGroupLargeProtobuf_Size(t *testing.T) {
	message := NewProtoMessageGroupLarge()

	msg, _ := proto.Marshal(&message)
	t.Log(len(msg))
}

/*
FLATBUFFERS LENGTHS
*/

func TestNestFlatBuffer_Size(t *testing.T) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageNest(builder)
	builder.Finish(message)
	msg := builder.FinishedBytes()
	t.Log(len(msg))
}

func TestFlatFlatBuffer_Size(t *testing.T) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlat(builder)
	builder.Finish(message)
	msg := builder.FinishedBytes()
	t.Log(len(msg))
}

func TestLargeFlatBuffer_Size(t *testing.T) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlatLarge(builder)
	builder.Finish(message)
	builder.FinishedBytes()
	msg := builder.FinishedBytes()
	t.Log(len(msg))
}

func TestXLargeFlatBuffer_Size(t *testing.T) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlatXLarge(builder)
	builder.Finish(message)
	builder.FinishedBytes()
	msg := builder.FinishedBytes()
	t.Log(len(msg))
}

func TestXLargePartialFlatBuffer_Size(t *testing.T) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlatXLargePartial(builder)
	builder.Finish(message)
	builder.FinishedBytes()
	msg := builder.FinishedBytes()
	t.Log(len(msg))
}

func TestXLargePartial64FlatBuffer_Size(t *testing.T) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlatXLargePartial64(builder)
	builder.Finish(message)
	builder.FinishedBytes()
	msg := builder.FinishedBytes()
	t.Log(len(msg))
}

/*
PROTOBUF BENCHMARKS
*/

func BenchmarkNestProtobuf_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := NewProtoMessageNest()

		proto.Marshal(&message)
	}
}

func BenchmarkNestProtobuf_Unmarshal(b *testing.B) {
	message := NewProtoMessageNest()
	buf, _ := proto.Marshal(&message)
	msg := &protomsg.MessageNest{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(buf, msg)
	}
}

func BenchmarkFlatProtobuf_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := NewProtoMessageFlat()

		proto.Marshal(&message)
	}
}

func BenchmarkFlatProtobuf_Unmarshal(b *testing.B) {
	message := NewProtoMessageFlat()
	buf, _ := proto.Marshal(&message)
	msg := &protomsg.MessageFlat{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(buf, msg)
	}
}

func BenchmarkFlatProtobuf_UnmarshalFieldAccess(b *testing.B) {
	message := NewProtoMessageFlat()
	buf, _ := proto.Marshal(&message)
	msg := &protomsg.MessageFlat{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(buf, msg)
		msg.GetField_1()
		msg.GetField_2()
		msg.GetField_3()
		msg.GetField_4()
	}
}

func BenchmarkFlatProtobuf_E2E(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := NewProtoMessageFlat()
		buf, _ := proto.Marshal(&message)
		msg := &protomsg.MessageFlat{}
		proto.Unmarshal(buf, msg)
		msg.GetField_1()
		msg.GetField_2()
		msg.GetField_3()
		msg.GetField_4()
	}
}

func BenchmarkLargeProtobuf_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := NewProtoMessageFlatLarge()

		proto.Marshal(&message)
	}
}

func BenchmarkLargeProtobuf_Unmarshal(b *testing.B) {
	message := NewProtoMessageFlatLarge()
	buf, _ := proto.Marshal(&message)
	msg := &protomsg.MessageFlatLarge{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(buf, msg)
	}
}

func BenchmarkXLargeProtobuf_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := NewProtoMessageFlatXLarge()

		proto.Marshal(&message)
	}
}

func BenchmarkXLargeProtobuf_Unmarshal(b *testing.B) {
	message := NewProtoMessageFlatXLarge()
	buf, _ := proto.Marshal(&message)
	msg := &protomsg.MessageFlatXLarge{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(buf, msg)
	}
}

func BenchmarkXLargePartialProtobuf_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := NewProtoMessageFlatXLargePartial()

		proto.Marshal(&message)
	}
}

func BenchmarkXLargePartialProtobuf_Unmarshal(b *testing.B) {
	message := NewProtoMessageFlatXLargePartial()
	buf, _ := proto.Marshal(&message)
	msg := &protomsg.MessageFlatXLarge{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(buf, msg)
	}
}

func BenchmarkXLargePartial64Protobuf_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := NewProtoMessageFlatXLargePartial64()

		proto.Marshal(&message)
	}
}

func BenchmarkXLargePartial64Protobuf_Unmarshal(b *testing.B) {
	message := NewProtoMessageFlatXLargePartial64()
	buf, _ := proto.Marshal(&message)
	msg := &protomsg.MessageFlatXLarge{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(buf, msg)
	}
}

func BenchmarkXLargePartial64Protobuf_FieldAccess(b *testing.B) {
	message := NewProtoMessageFlatXLargePartial64()
	buf, _ := proto.Marshal(&message)
	msg := &protomsg.MessageFlatXLarge{}
	proto.Unmarshal(buf, msg)

	for i := 0; i < b.N; i++ {
		msg.GetField_1()
	}
}

func BenchmarkGroupProtobuf_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := NewProtoMessageGroup()

		proto.Marshal(&message)
	}
}

func BenchmarkGroupProtobuf_Unmarshal(b *testing.B) {
	message := NewProtoMessageGroup()
	buf, _ := proto.Marshal(&message)
	msg := &protomsg.MessageGroup{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(buf, msg)
	}
}

func BenchmarkGroupPartialProtobuf_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := NewProtoMessageGroupPartial()

		proto.Marshal(&message)
	}
}

func BenchmarkGroupPartialProtobuf_Unmarshal(b *testing.B) {
	message := NewProtoMessageGroupPartial()
	buf, _ := proto.Marshal(&message)
	msg := &protomsg.MessageGroup{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(buf, msg)
	}
}

func BenchmarkGroupLargeProtobuf_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message := NewProtoMessageGroupLarge()

		proto.Marshal(&message)
	}
}

func BenchmarkGroupLargeProtobuf_Unmarshal(b *testing.B) {
	message := NewProtoMessageGroupLarge()
	buf, _ := proto.Marshal(&message)
	msg := &protomsg.MessageGroupLarge{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(buf, msg)
	}
}

/*
FLATBUFFERS BENCHMARKS
*/

func BenchmarkNestFlatBuffer_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(1024)
		message := NewFlatMessageNest(builder)
		builder.Finish(message)
		builder.FinishedBytes()
	}
}

func BenchmarkNestFlatBuffer_Unmarshal(b *testing.B) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageNest(builder)
	builder.Finish(message)
	buf := builder.FinishedBytes()
	msg := &flatmsg.MessageNest{}

	for i := 0; i < b.N; i++ {
		msg = flatmsg.GetRootAsMessageNest(buf, 0)
		_ = msg
	}
}

func BenchmarkFlatFlatBuffer_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(1024)
		message := NewFlatMessageFlat(builder)
		builder.Finish(message)
		builder.FinishedBytes()
	}
}

func BenchmarkFlatFlatBuffer_Unmarshal(b *testing.B) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlat(builder)
	builder.Finish(message)
	buf := builder.FinishedBytes()
	msg := &flatmsg.MessageFlat{}

	for i := 0; i < b.N; i++ {
		msg = flatmsg.GetRootAsMessageFlat(buf, 0)
		_ = msg
	}
}

func BenchmarkFlatFlatBuffer_UnmarshalFieldAccess(b *testing.B) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlat(builder)
	builder.Finish(message)
	buf := builder.FinishedBytes()
	msg := &flatmsg.MessageFlat{}

	for i := 0; i < b.N; i++ {
		msg = flatmsg.GetRootAsMessageFlat(buf, 0)
		msg.Field1()
		msg.Field2()
		msg.Field3()
		msg.Field4()
	}
}

func BenchmarkFlatFlatBuffer_E2E(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(1024)
		message := NewFlatMessageFlat(builder)
		builder.Finish(message)
		buf := builder.FinishedBytes()
		msg := &flatmsg.MessageFlat{}
		msg = flatmsg.GetRootAsMessageFlat(buf, 0)
		msg.Field1()
		msg.Field2()
		msg.Field3()
		msg.Field4()
	}
}

func BenchmarkLargeFlatBuffer_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(1024)
		message := NewFlatMessageFlatLarge(builder)
		builder.Finish(message)
		builder.FinishedBytes()
	}
}

func BenchmarkLargeFlatBuffer_Unmarshal(b *testing.B) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlatLarge(builder)
	builder.Finish(message)
	buf := builder.FinishedBytes()
	msg := &flatmsg.MessageFlatLarge{}

	for i := 0; i < b.N; i++ {
		msg = flatmsg.GetRootAsMessageFlatLarge(buf, 0)
		_ = msg
	}
}

func BenchmarkXLargeFlatBuffer_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(1024)
		message := NewFlatMessageFlatXLarge(builder)
		builder.Finish(message)
		builder.FinishedBytes()
	}
}

func BenchmarkXLargeFlatBuffer_Unmarshal(b *testing.B) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlatXLarge(builder)
	builder.Finish(message)
	buf := builder.FinishedBytes()
	msg := &flatmsg.MessageFlatXLarge{}

	for i := 0; i < b.N; i++ {
		msg = flatmsg.GetRootAsMessageFlatXLarge(buf, 0)
		msg.Field1()
		_ = msg
	}
}

func BenchmarkXLargePartialFlatBuffer_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(1024)
		message := NewFlatMessageFlatXLargePartial(builder)
		builder.Finish(message)
		builder.FinishedBytes()
	}
}

func BenchmarkXLargePartialFlatBuffer_Unmarshal(b *testing.B) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlatXLargePartial(builder)
	builder.Finish(message)
	buf := builder.FinishedBytes()
	msg := &flatmsg.MessageFlatXLarge{}

	for i := 0; i < b.N; i++ {
		msg = flatmsg.GetRootAsMessageFlatXLarge(buf, 0)
		_ = msg
	}
}

func BenchmarkXLargePartial64FlatBuffer_Marshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		builder := flatbuffers.NewBuilder(1024)
		message := NewFlatMessageFlatXLargePartial64(builder)
		builder.Finish(message)
		builder.FinishedBytes()
	}
}

func BenchmarkXLargePartial64FlatBuffer_Unmarshal(b *testing.B) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlatXLargePartial64(builder)
	builder.Finish(message)
	buf := builder.FinishedBytes()
	msg := &flatmsg.MessageFlatXLarge{}

	for i := 0; i < b.N; i++ {
		msg = flatmsg.GetRootAsMessageFlatXLarge(buf, 0)
		_ = msg
	}
}

func BenchmarkXLargePartial64FlatBuffer_FieldAccess(b *testing.B) {
	builder := flatbuffers.NewBuilder(1024)
	message := NewFlatMessageFlatXLargePartial64(builder)
	builder.Finish(message)
	buf := builder.FinishedBytes()
	msg := &flatmsg.MessageFlatXLarge{}
	msg = flatmsg.GetRootAsMessageFlatXLarge(buf, 0)

	for i := 0; i < b.N; i++ {
		msg.Field1()
	}
}
