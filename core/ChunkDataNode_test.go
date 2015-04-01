package core

import (
	"github.com/inkyblackness/res"
	"github.com/inkyblackness/res/chunk"

	check "gopkg.in/check.v1"
)

type ChunkDataNodeSuite struct {
	parentNode  DataNode
	chunkHolder *TestingChunkProvider

	chunkDataNode DataNode
}

var _ = check.Suite(&ChunkDataNodeSuite{})

func (suite *ChunkDataNodeSuite) SetUpTest(c *check.C) {
	suite.chunkHolder = NewTestingChunkProvider()
}

func (suite *ChunkDataNodeSuite) TestInfoReturnsListOfAvailableBlockCountAndContentType(c *check.C) {
	holder := chunk.NewBlockHolder(chunk.BasicChunkType, res.Data, [][]byte{[]byte{}, []byte{}})
	suite.chunkDataNode = newChunkDataNode(suite.parentNode, res.ResourceID(0x0200), holder)

	result := suite.chunkDataNode.Info()

	c.Check(result, check.Equals, "Available blocks: 2\nContent type: 0x00")
}

func (suite *ChunkDataNodeSuite) TestResolveReturnsDataNodeForKnownID(c *check.C) {
	holder := chunk.NewBlockHolder(chunk.BasicChunkType, res.Data, [][]byte{[]byte{}, []byte{}})
	suite.chunkDataNode = newChunkDataNode(suite.parentNode, res.ResourceID(0x0200), holder)

	result := suite.chunkDataNode.Resolve("1")

	c.Assert(result, check.NotNil)
	c.Check(result.ID(), check.Equals, "1")
}
