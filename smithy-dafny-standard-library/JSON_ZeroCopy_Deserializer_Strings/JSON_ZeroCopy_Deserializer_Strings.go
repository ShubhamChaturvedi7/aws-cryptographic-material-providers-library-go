// Package JSON_ZeroCopy_Deserializer_Strings
// Dafny module JSON_ZeroCopy_Deserializer_Strings compiled into Go

package JSON_ZeroCopy_Deserializer_Strings

import (
	os "os"

	m_Actions "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Actions"
	m_Base64 "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Base64"
	m_Base64Lemmas "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Base64Lemmas"
	m_BoundedInts "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/BoundedInts"
	m_DivInternals "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/DivInternals"
	m_DivInternalsNonlinear "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/DivInternalsNonlinear"
	m_DivMod "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/DivMod"
	m_FileIO "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/FileIO"
	m_FloatCompare "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/FloatCompare"
	m_Functions "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Functions"
	m_GeneralInternals "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/GeneralInternals"
	m_GetOpt "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/GetOpt"
	m_HexStrings "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/HexStrings"
	m_JSON_ConcreteSyntax_Spec "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_ConcreteSyntax_Spec"
	m_JSON_ConcreteSyntax_SpecProperties "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_ConcreteSyntax_SpecProperties"
	m_JSON_Deserializer "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Deserializer"
	m_JSON_Deserializer_ByteStrConversion "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Deserializer_ByteStrConversion"
	m_JSON_Deserializer_Uint16StrConversion "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Deserializer_Uint16StrConversion"
	m_JSON_Errors "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Errors"
	m_JSON_Grammar "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Grammar"
	m_JSON_Serializer "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Serializer"
	m_JSON_Serializer_ByteStrConversion "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Serializer_ByteStrConversion"
	m_JSON_Spec "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Spec"
	m_JSON_Utils_Cursors "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Utils_Cursors"
	m_JSON_Utils_Lexers_Core "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Utils_Lexers_Core"
	m_JSON_Utils_Lexers_Strings "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Utils_Lexers_Strings"
	m_JSON_Utils_Parsers "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Utils_Parsers"
	m_JSON_Utils_Seq "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Utils_Seq"
	m_JSON_Utils_Str "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Utils_Str"
	m_JSON_Utils_Str_CharStrConversion "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Utils_Str_CharStrConversion"
	m_JSON_Utils_Str_CharStrEscaping "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Utils_Str_CharStrEscaping"
	m_JSON_Utils_Vectors "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Utils_Vectors"
	m_JSON_Utils_Views_Core "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Utils_Views_Core"
	m_JSON_Utils_Views_Writers "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Utils_Views_Writers"
	m_JSON_Values "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_Values"
	m_JSON_ZeroCopy_Deserializer_Core "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_ZeroCopy_Deserializer_Core"
	m_JSON_ZeroCopy_Serializer "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/JSON_ZeroCopy_Serializer"
	m_Logarithm "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Logarithm"
	m__Math "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Math_"
	m_ModInternals "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/ModInternals"
	m_ModInternalsNonlinear "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/ModInternalsNonlinear"
	m_Mul "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Mul"
	m_MulInternals "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/MulInternals"
	m_MulInternalsNonlinear "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/MulInternalsNonlinear"
	m_Power "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Power"
	m_Relations "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Relations"
	m_Seq "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Seq"
	m_Seq_MergeSort "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Seq_MergeSort"
	m_Sorting "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Sorting"
	m_StandardLibrary "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/StandardLibrary"
	m_StandardLibraryInterop "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/StandardLibraryInterop"
	m_StandardLibrary_Sequence "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/StandardLibrary_Sequence"
	m_StandardLibrary_String "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/StandardLibrary_String"
	m_StandardLibrary_UInt "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/StandardLibrary_UInt"
	m_Streams "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Streams"
	m_UnicodeStrings "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/UnicodeStrings"
	m__Unicode "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Unicode_"
	m_Utf16EncodingForm "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Utf16EncodingForm"
	m_Utf8EncodingForm "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Utf8EncodingForm"
	m_Wrappers "github.com/ShubhamChaturvedi7/aws-cryptographic-material-providers-library-go/smithy-dafny-standard-library/Wrappers"
	m__System "github.com/dafny-lang/DafnyRuntimeGo/v4/System_"
	_dafny "github.com/dafny-lang/DafnyRuntimeGo/v4/dafny"
)

var _ = os.Args
var _ _dafny.Dummy__
var _ m__System.Dummy__
var _ m_Wrappers.Dummy__
var _ m_Relations.Dummy__
var _ m_Seq_MergeSort.Dummy__
var _ m__Math.Dummy__
var _ m_Seq.Dummy__
var _ m_BoundedInts.Dummy__
var _ m__Unicode.Dummy__
var _ m_Functions.Dummy__
var _ m_Utf8EncodingForm.Dummy__
var _ m_Utf16EncodingForm.Dummy__
var _ m_UnicodeStrings.Dummy__
var _ m_FileIO.Dummy__
var _ m_GeneralInternals.Dummy__
var _ m_MulInternalsNonlinear.Dummy__
var _ m_MulInternals.Dummy__
var _ m_Mul.Dummy__
var _ m_ModInternalsNonlinear.Dummy__
var _ m_DivInternalsNonlinear.Dummy__
var _ m_ModInternals.Dummy__
var _ m_DivInternals.Dummy__
var _ m_DivMod.Dummy__
var _ m_Power.Dummy__
var _ m_Logarithm.Dummy__
var _ m_StandardLibraryInterop.Dummy__
var _ m_StandardLibrary_UInt.Dummy__
var _ m_StandardLibrary_Sequence.Dummy__
var _ m_StandardLibrary_String.Dummy__
var _ m_StandardLibrary.Dummy__
var _ m_Streams.Dummy__
var _ m_Sorting.Dummy__
var _ m_HexStrings.Dummy__
var _ m_GetOpt.Dummy__
var _ m_FloatCompare.Dummy__
var _ m_Base64.Dummy__
var _ m_Base64Lemmas.Dummy__
var _ m_Actions.Dummy__
var _ m_JSON_Utils_Views_Core.Dummy__
var _ m_JSON_Utils_Views_Writers.Dummy__
var _ m_JSON_Utils_Lexers_Core.Dummy__
var _ m_JSON_Utils_Lexers_Strings.Dummy__
var _ m_JSON_Utils_Cursors.Dummy__
var _ m_JSON_Utils_Parsers.Dummy__
var _ m_JSON_Utils_Str_CharStrConversion.Dummy__
var _ m_JSON_Utils_Str_CharStrEscaping.Dummy__
var _ m_JSON_Utils_Str.Dummy__
var _ m_JSON_Utils_Seq.Dummy__
var _ m_JSON_Utils_Vectors.Dummy__
var _ m_JSON_Errors.Dummy__
var _ m_JSON_Values.Dummy__
var _ m_JSON_Spec.Dummy__
var _ m_JSON_Grammar.Dummy__
var _ m_JSON_Serializer_ByteStrConversion.Dummy__
var _ m_JSON_Serializer.Dummy__
var _ m_JSON_Deserializer_Uint16StrConversion.Dummy__
var _ m_JSON_Deserializer_ByteStrConversion.Dummy__
var _ m_JSON_Deserializer.Dummy__
var _ m_JSON_ConcreteSyntax_Spec.Dummy__
var _ m_JSON_ConcreteSyntax_SpecProperties.Dummy__
var _ m_JSON_ZeroCopy_Serializer.Dummy__
var _ m_JSON_ZeroCopy_Deserializer_Core.Dummy__

type Dummy__ struct{}

// Definition of class Default__
type Default__ struct {
	dummy byte
}

func New_Default___() *Default__ {
	_this := Default__{}

	return &_this
}

type CompanionStruct_Default___ struct {
}

var Companion_Default___ = CompanionStruct_Default___{}

func (_this *Default__) Equals(other *Default__) bool {
	return _this == other
}

func (_this *Default__) EqualsGeneric(x interface{}) bool {
	other, ok := x.(*Default__)
	return ok && _this.Equals(other)
}

func (*Default__) String() string {
	return "JSON_ZeroCopy_Deserializer_Strings.Default__"
}
func (_this *Default__) ParentTraits_() []*_dafny.TraitID {
	return [](*_dafny.TraitID){}
}

var _ _dafny.TraitOffspring = &Default__{}

func (_static *CompanionStruct_Default___) StringBody(cs m_JSON_Utils_Cursors.Cursor__) m_Wrappers.Result {
	var pr m_Wrappers.Result = m_Wrappers.Companion_Result_.Default(m_JSON_Utils_Cursors.Companion_Cursor_.Witness())
	_ = pr
	var _0_escaped bool
	_ = _0_escaped
	_0_escaped = false
	var _hi0 uint32 = (cs).Dtor_end()
	_ = _hi0
	for _1_point_k := (cs).Dtor_point(); _1_point_k < _hi0; _1_point_k++ {
		var _2_byte uint8
		_ = _2_byte
		_2_byte = ((cs).Dtor_s()).Select(uint32(_1_point_k)).(uint8)
		if ((_2_byte) == (uint8(_dafny.Char('"')))) && (!(_0_escaped)) {
			pr = m_Wrappers.Companion_Result_.Create_Success_(m_JSON_Utils_Cursors.Companion_Cursor___.Create_Cursor_((cs).Dtor_s(), (cs).Dtor_beg(), _1_point_k, (cs).Dtor_end()))
			return pr
		} else if (_2_byte) == (uint8(_dafny.Char('\\'))) {
			_0_escaped = !(_0_escaped)
		} else {
			_0_escaped = false
		}
	}
	pr = m_Wrappers.Companion_Result_.Create_Failure_(m_JSON_Utils_Cursors.Companion_CursorError_.Create_EOF_())
	return pr
	return pr
}
func (_static *CompanionStruct_Default___) Quote(cs m_JSON_Utils_Cursors.Cursor__) m_Wrappers.Result {
	var _0_valueOrError0 m_Wrappers.Result = (cs).AssertChar(_dafny.Char('"'))
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _1_cs m_JSON_Utils_Cursors.Cursor__ = (_0_valueOrError0).Extract().(m_JSON_Utils_Cursors.Cursor__)
		_ = _1_cs
		return m_Wrappers.Companion_Result_.Create_Success_((_1_cs).Split())
	}
}
func (_static *CompanionStruct_Default___) String(cs m_JSON_Utils_Cursors.Cursor__) m_Wrappers.Result {
	var _0_valueOrError0 m_Wrappers.Result = Companion_Default___.Quote(cs)
	_ = _0_valueOrError0
	if (_0_valueOrError0).IsFailure() {
		return (_0_valueOrError0).PropagateFailure()
	} else {
		var _let_tmp_rhs0 m_JSON_Utils_Cursors.Split = (_0_valueOrError0).Extract().(m_JSON_Utils_Cursors.Split)
		_ = _let_tmp_rhs0
		var _1_lq m_JSON_Utils_Views_Core.View__ = _let_tmp_rhs0.Get_().(m_JSON_Utils_Cursors.Split_SP).T.(m_JSON_Utils_Views_Core.View__)
		_ = _1_lq
		var _2_cs m_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs0.Get_().(m_JSON_Utils_Cursors.Split_SP).Cs
		_ = _2_cs
		var _3_valueOrError1 m_Wrappers.Result = Companion_Default___.StringBody(_2_cs)
		_ = _3_valueOrError1
		if (_3_valueOrError1).IsFailure() {
			return (_3_valueOrError1).PropagateFailure()
		} else {
			var _4_contents m_JSON_Utils_Cursors.Cursor__ = (_3_valueOrError1).Extract().(m_JSON_Utils_Cursors.Cursor__)
			_ = _4_contents
			var _let_tmp_rhs1 m_JSON_Utils_Cursors.Split = (_4_contents).Split()
			_ = _let_tmp_rhs1
			var _5_contents m_JSON_Utils_Views_Core.View__ = _let_tmp_rhs1.Get_().(m_JSON_Utils_Cursors.Split_SP).T.(m_JSON_Utils_Views_Core.View__)
			_ = _5_contents
			var _6_cs m_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs1.Get_().(m_JSON_Utils_Cursors.Split_SP).Cs
			_ = _6_cs
			var _7_valueOrError2 m_Wrappers.Result = Companion_Default___.Quote(_6_cs)
			_ = _7_valueOrError2
			if (_7_valueOrError2).IsFailure() {
				return (_7_valueOrError2).PropagateFailure()
			} else {
				var _let_tmp_rhs2 m_JSON_Utils_Cursors.Split = (_7_valueOrError2).Extract().(m_JSON_Utils_Cursors.Split)
				_ = _let_tmp_rhs2
				var _8_rq m_JSON_Utils_Views_Core.View__ = _let_tmp_rhs2.Get_().(m_JSON_Utils_Cursors.Split_SP).T.(m_JSON_Utils_Views_Core.View__)
				_ = _8_rq
				var _9_cs m_JSON_Utils_Cursors.Cursor__ = _let_tmp_rhs2.Get_().(m_JSON_Utils_Cursors.Split_SP).Cs
				_ = _9_cs
				return m_Wrappers.Companion_Result_.Create_Success_(m_JSON_Utils_Cursors.Companion_Split_.Create_SP_(m_JSON_Grammar.Companion_Jstring_.Create_JString_(_1_lq, _5_contents, _8_rq), _9_cs))
			}
		}
	}
}

// End of class Default__
