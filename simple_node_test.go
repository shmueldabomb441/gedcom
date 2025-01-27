package gedcom_test

import (
	"testing"

	"github.com/elliotchance/gedcom"
	"github.com/elliotchance/tf"
	"github.com/stretchr/testify/assert"
)

func TestSimpleNode_ChildNodes(t *testing.T) {
	node := gedcom.NewNode(gedcom.TagText, "", "")

	assert.Len(t, node.Nodes(), 0)
}

func TestIsNil(t *testing.T) {
	IsNil := tf.Function(t, gedcom.IsNil)

	IsNil((*gedcom.SimpleNode)(nil)).Returns(true)
	IsNil(gedcom.NewBirthNode("")).Returns(false)
	IsNil((*gedcom.NameNode)(nil)).Returns(true)
	IsNil(gedcom.NewNameNode("")).Returns(false)

	// Untyped nil is a special case that cannot be tested above.
	assert.True(t, gedcom.IsNil(nil))
}

func TestSimpleNode_Equals(t *testing.T) {
	Equals := tf.NamedFunction(t, "SimpleNode_Equals", (*gedcom.SimpleNode).Equals)

	left := []*gedcom.SimpleNode{
		(*gedcom.SimpleNode)(nil),
		gedcom.NewNode(gedcom.TagVersion, "", "").(*gedcom.SimpleNode),
		gedcom.NewNode(gedcom.TagVersion, "a", "").(*gedcom.SimpleNode),
		gedcom.NewNode(gedcom.TagVersion, "", "b").(*gedcom.SimpleNode),
		gedcom.NewNode(gedcom.TagVersion, "a", "b").(*gedcom.SimpleNode),
	}

	right := gedcom.Nodes{
		(*gedcom.SimpleNode)(nil),
		gedcom.NewNode(gedcom.TagVersion, "", "").(*gedcom.SimpleNode),
		gedcom.NewNode(gedcom.TagVersion, "a", "").(*gedcom.SimpleNode),
		gedcom.NewNode(gedcom.TagVersion, "", "b").(*gedcom.SimpleNode),
		gedcom.NewNode(gedcom.TagVersion, "a", "b").(*gedcom.SimpleNode),
		gedcom.NewNameNode(""),
	}

	const N = false
	const Y = true

	expected := [][]bool{
		{N, N, N, N, N, N},
		{N, Y, N, N, N, N},
		{N, N, Y, N, N, N},
		{N, N, N, Y, N, N},
		{N, N, N, N, Y, N},
	}

	for i, l := range left {
		for j, r := range right {
			Equals(l, r).Returns(expected[i][j])
		}
	}
}

func TestAncestryNode_Equals(t *testing.T) {
	//test when source ids are the same
	original := GetAncestryIndividual("@S291470533@", "@S291470520@", "@S291470520@", "@S291470520@", "@S291470520@")
	assert.True(t, gedcom.DeepEqualNodes(gedcom.newDocumentFromString(original).Nodes(), gedcom.newDocumentFromString(original).Nodes()))

	//test when source ids are different, but _apid stays the same
	left := gedcom.newDocumentFromString(GetAncestryIndividual("@S222222222@", "@S444444444@", "@S666666666@", "@S888888888@", "@S111111111@"))
	right := gedcom.newDocumentFromString(GetAncestryIndividual("@S333333333@", "@S555555555@", "@S777777777@", "@S999999999@", "@S000000000@"))
	assert.True(t, gedcom.DeepEqualNodes(left.Nodes(), right.Nodes()))
}

// This is an actual gedcom entry, hope that's ok.
func GetAncestryIndividual(source1 string, source2 string, source3 string, source4 string, source5 string) string {
	return "0 @I152151456706@ INDI" +
		"\n1 NAME Jacob /Yourow/" +
		"\n2 GIVN Jacob" +
		"\n2 SURN Yourow" +
		"\n2 SOUR " + source1 +
		"\n3 PAGE New York City Municipal Archives; New York, New York; Borough: Manhattan; Volume Number: 13" +
		"\n3 _APID 1,61406::6159341" +
		"\n2 SOUR " + source2 +
		"\n3 PAGE Year: 1930; Census Place: Bronx, Bronx, New York; Page: 42A; Enumeration District: 0430; FHL microfilm: 2341213" +
		"\n3 _APID 1,6224::30826480" +
		"\n1 SEX M" +
		"\n1 FAMS @F89@" +
		"\n1 BIRT" +
		"\n2 DATE abt 1888" +
		"\n2 PLAC Russia" +
		"\n2 SOUR " + source3 +
		"\n3 PAGE Year: 1930; Census Place: Bronx, Bronx, New York; Page: 42A; Enumeration District: 0430; FHL microfilm: 2341213" +
		"\n3 _APID 1,6224::30826480" +
		"\n1 EVEN" +
		"\n2 TYPE Arrival" +
		"\n2 DATE 1905" +
		"\n2 SOUR " + source4 +
		"\n3 PAGE Year: 1930; Census Place: Bronx, Bronx, New York; Page: 42A; Enumeration District: 0430; FHL microfilm: 2341213" +
		"\n3 _APID 1,6224::30826480" +
		"\n1 RESI Marital Status: Married; Relation to Head: Head" +
		"\n2 DATE 1930" +
		"\n2 PLAC Bronx, Bronx, New York, USA" +
		"\n2 SOUR " + source5 +
		"\n3 PAGE Year: 1930; Census Place: Bronx, Bronx, New York; Page: 42A; Enumeration District: 0430; FHL microfilm: 2341213" +
		"\n3 _APID 1,6224::30826480"
}

func TestSimpleNode_Tag(t *testing.T) {
	Tag := tf.Function(t, (*gedcom.SimpleNode).Tag)

	Tag((*gedcom.SimpleNode)(nil)).Returns(gedcom.Tag{})
}

func TestSimpleNode_Value(t *testing.T) {
	Value := tf.Function(t, (*gedcom.SimpleNode).Value)

	Value((*gedcom.SimpleNode)(nil)).Returns("")
}

func TestSimpleNode_Pointer(t *testing.T) {
	Pointer := tf.Function(t, (*gedcom.SimpleNode).Pointer)

	Pointer((*gedcom.SimpleNode)(nil)).Returns("")
}

func TestSimpleNode_Nodes(t *testing.T) {
	Nodes := tf.Function(t, (*gedcom.SimpleNode).Nodes)

	Nodes((*gedcom.SimpleNode)(nil)).Returns((gedcom.Nodes)(nil))
}

func TestSimpleNode_String(t *testing.T) {
	String := tf.Function(t, (*gedcom.SimpleNode).String)

	String((*gedcom.SimpleNode)(nil)).Returns("")
}

func TestSimpleNode_GEDCOMString(t *testing.T) {
	root := gedcom.NewDocument().AddIndividual("P1",
		gedcom.NewNameNode("Elliot /Chance/"),
		gedcom.NewBirthNode("",
			gedcom.NewDateNode("6 MAY 1989"),
		),
	)

	assert.Equal(t, root.GEDCOMString(0), `0 @P1@ INDI
1 NAME Elliot /Chance/
1 BIRT
2 DATE 6 MAY 1989
`)
}

func TestSimpleNode_GEDCOMLine(t *testing.T) {
	GEDCOMLine := tf.NamedFunction(t, "SimpleNode_GEDCOMLine",
		(*gedcom.SimpleNode).GEDCOMLine)

	GEDCOMLine(gedcom.NewNode(gedcom.TagBirth, "foo", "72").(*gedcom.BirthNode).SimpleNode, 0).
		Returns("0 @72@ BIRT foo")

	GEDCOMLine(gedcom.NewNode(gedcom.TagDeath, "bar", "baz").(*gedcom.DeathNode).SimpleNode, 3).Returns("3 @baz@ DEAT bar")

	GEDCOMLine(gedcom.NewDateNode("3 SEP 1945").SimpleNode, 2).
		Returns("2 DATE 3 SEP 1945")

	GEDCOMLine(gedcom.NewNode(gedcom.TagBirth, "foo", "72").(*gedcom.BirthNode).SimpleNode, -1).
		Returns("@72@ BIRT foo")
}

func TestSimpleNode_SetNodes(t *testing.T) {
	birth := gedcom.NewBirthNode("foo")
	assert.Nil(t, birth.Nodes())

	birth.SetNodes(gedcom.Nodes{
		gedcom.NewDateNode("3 SEP 1945"),
	})
	assert.Equal(t,
		gedcom.Nodes{gedcom.NewDateNode("3 SEP 1945")}, birth.Nodes())

	birth.SetNodes(nil)
	assert.Nil(t, birth.Nodes())
}
