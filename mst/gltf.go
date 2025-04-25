package mst

import (
	"bytes"
	"encoding/binary"
	"image/png"
	"io"

	"github.com/qmuntal/gltf/ext/specular"

	"github.com/qmuntal/gltf"
	"pinkey.ltd/xr/go3d/mat4"
)

const GLTF_VERSION = "2.0"

func MstToGltf[T float64 | float32](msts []*Mesh[T]) (*gltf.Document, error) {
	doc := CreateDoc()
	for _, mst := range msts {
		e := BuildGltf(doc, mst, false, true)
		if e != nil {
			return nil, e
		}
	}
	return doc, nil
}

func MstToGltfWithOutline[T float64 | float32](msts []*Mesh[T]) (*gltf.Document, error) {
	doc := CreateDoc()
	for _, mst := range msts {
		e := BuildGltf(doc, mst, true, true)
		if e != nil {
			return nil, e
		}
	}
	return doc, nil
}
func CreateDoc() *gltf.Document {
	doc := &gltf.Document{}
	doc.Asset.Version = GLTF_VERSION
	srcIndex := 0
	doc.Scene = &srcIndex
	doc.Scenes = append(doc.Scenes, &gltf.Scene{})
	doc.Buffers = append(doc.Buffers, &gltf.Buffer{})
	return doc
}

type calcSizeWriter struct {
	writer io.Writer
	Size   int
}

func (w *calcSizeWriter) Write(p []byte) (n int, err error) {
	si := len(p)
	w.writer.Write(p)
	w.Size += int(si)
	return si, nil
}

func (w *calcSizeWriter) Bytes() []byte {
	return w.writer.(*bytes.Buffer).Bytes()
}

func (w *calcSizeWriter) GetSize() int {
	return len(w.Bytes())
}

func newSizeWriter() calcSizeWriter {
	wt := bytes.NewBuffer([]byte{})
	return calcSizeWriter{Size: int(0), writer: wt}
}

func calcPadding(offset, paddingUnit int) int {
	padding := offset % paddingUnit
	if padding != 0 {
		padding = paddingUnit - padding
	}
	return padding
}

func GetGltfBinary(doc *gltf.Document, paddingUnit int) ([]byte, error) {
	w := newSizeWriter()
	enc := gltf.NewEncoder(w.writer)
	enc.AsBinary = true
	if err := enc.Encode(doc); err != nil {
		return nil, err
	}
	padding := calcPadding(w.Size, paddingUnit)
	if padding == 0 {
		return w.Bytes(), nil
	}
	pad := make([]byte, padding)
	for i := range pad {
		pad[i] = 0x20
	}
	w.Write(pad)
	return w.Bytes(), nil
}

func BuildGltf[T float64 | float32](doc *gltf.Document, mh *Mesh[T], exportOutline, gpu_instance bool) error {
	err := buildGltf(doc, &mh.BaseMesh, nil, exportOutline, gpu_instance)
	if err != nil {
		return err
	}
	for _, inst := range mh.InstanceNode {
		buildGltf(doc, inst.Mesh, inst.Transfors, false, gpu_instance)
	}

	return nil
}

type buildContext struct {
	mtlSize int
	bvIndex int
	bvPos   int
	bvTex   int
	bvNorm  int
}

func buildMeshBuffer[T float64 | float32](ctx *buildContext, buffer *gltf.Buffer, bufferViews []*gltf.BufferView, nd *MeshNode[T]) []*gltf.BufferView {
	var bt []byte
	buf := bytes.NewBuffer(bt)
	ctx.bvIndex = len(bufferViews)
	indecs := &gltf.BufferView{}
	startLen := buffer.ByteLength
	indecs.ByteOffset = startLen
	for _, g := range nd.FaceGroup {
		for _, f := range g.Faces {
			binary.Write(buf, binary.LittleEndian, f.Vertex)
		}
	}
	indecs.ByteLength = buf.Len()
	indecs.Buffer = 0
	bufferViews = append(bufferViews, indecs)

	postions := &gltf.BufferView{}
	postions.ByteOffset = (buf.Len()) + startLen
	binary.Write(buf, binary.LittleEndian, nd.Vertices)
	postions.ByteLength = (buf.Len()) - postions.ByteOffset + startLen
	postions.Buffer = 0
	ctx.bvPos = len(bufferViews)
	bufferViews = append(bufferViews, postions)

	texcood := &gltf.BufferView{}
	ctx.bvTex = len(bufferViews)
	if len(nd.TexCoords) > 0 {
		texcood.ByteOffset = (buf.Len()) + startLen
		binary.Write(buf, binary.LittleEndian, nd.TexCoords)
		texcood.ByteLength = (buf.Len()) - texcood.ByteOffset + startLen
		texcood.Buffer = 0
		bufferViews = append(bufferViews, texcood)
	}

	normalView := &gltf.BufferView{}
	ctx.bvNorm = len(bufferViews)
	if len(nd.Normals) > 0 {
		normalView.ByteOffset = (buf.Len()) + startLen
		binary.Write(buf, binary.LittleEndian, nd.Normals)
		normalView.ByteLength = (buf.Len()) - normalView.ByteOffset + startLen
		normalView.Buffer = 0
		bufferViews = append(bufferViews, normalView)
	}
	buffer.ByteLength += (buf.Len())
	buffer.Data = append(buffer.Data, buf.Bytes()...)

	return bufferViews
}

func buildOutlineBuffer[T float64 | float32](ctx *buildContext, buffer *gltf.Buffer, bufferViews []*gltf.BufferView, nd *MeshNode[T]) []*gltf.BufferView {
	var bt []byte
	buf := bytes.NewBuffer(bt)
	ctx.bvIndex = len(bufferViews)
	indecs := &gltf.BufferView{}
	startLen := buffer.ByteLength
	indecs.ByteOffset = startLen
	for _, g := range nd.EdgeGroup {
		for _, f := range g.Edges {
			binary.Write(buf, binary.LittleEndian, f)
		}
	}
	indecs.ByteLength = (buf.Len())
	indecs.Buffer = 0
	bufferViews = append(bufferViews, indecs)

	postions := &gltf.BufferView{}
	postions.ByteOffset = (buf.Len()) + startLen
	binary.Write(buf, binary.LittleEndian, nd.Vertices)
	postions.ByteLength = (buf.Len()) - postions.ByteOffset + startLen
	postions.Buffer = 0
	ctx.bvPos = len(bufferViews)
	bufferViews = append(bufferViews, postions)

	buffer.ByteLength += buf.Len()
	buffer.Data = append(buffer.Data, buf.Bytes()...)

	return bufferViews
}

func buildOutline[T float64 | float32](ctx *buildContext, accessors []*gltf.Accessor, nd *MeshNode[T]) (*gltf.Mesh, []*gltf.Accessor) {
	mesh := &gltf.Mesh{}
	aftIndices := len(nd.EdgeGroup)
	idx := len(accessors)
	indexPos := aftIndices + idx
	var start int = 0
	for i := range nd.EdgeGroup {
		patch := nd.EdgeGroup[i]
		batchId := patch.Batchid
		if batchId < 0 {
			batchId = 0
		}
		mtlId := batchId + ctx.mtlSize

		ps := &gltf.Primitive{}
		ps.Material = &mtlId
		if ps.Attributes == nil {
			ps.Attributes = make(gltf.PrimitiveAttributes)
		}
		index := i + idx
		ps.Indices = &index

		ps.Attributes["POSITION"] = indexPos

		ps.Mode = gltf.PrimitiveLineStrip
		mesh.Primitives = append(mesh.Primitives, ps)

		indexacc := &gltf.Accessor{}
		indexacc.ComponentType = gltf.ComponentUint

		indexacc.ByteOffset = start * 8
		indexacc.Count = (len(patch.Edges)) * 2

		start += len(patch.Edges)
		bfindex := ctx.bvIndex
		indexacc.BufferView = &bfindex
		accessors = append(accessors, indexacc)
	}

	posacc := &gltf.Accessor{}
	posacc.ComponentType = gltf.ComponentFloat
	posacc.Type = gltf.AccessorVec3
	posacc.Count = len(nd.Vertices)

	posacc.BufferView = &ctx.bvPos
	box := nd.GetBoundbox()
	posacc.Min = []float64{box[0], box[1], box[2]}
	posacc.Max = []float64{box[3], box[4], box[5]}
	accessors = append(accessors, posacc)

	return mesh, accessors
}

func buildMesh[T float64 | float32](ctx *buildContext, accessors []*gltf.Accessor, nd *MeshNode[T]) (*gltf.Mesh, []*gltf.Accessor) {
	mesh := &gltf.Mesh{}
	aftIndices := len(nd.FaceGroup)
	idx := len(accessors)
	indexPos := aftIndices + idx
	var start = 0

	for i := range nd.FaceGroup {
		tmp := indexPos
		patch := nd.FaceGroup[i]
		batchId := patch.Batchid
		if batchId < 0 {
			batchId = 0
		}
		mtl_id := batchId + ctx.mtlSize

		ps := &gltf.Primitive{}
		ps.Material = &mtl_id
		if ps.Attributes == nil {
			ps.Attributes = make(gltf.PrimitiveAttributes)
		}
		index := i + idx
		ps.Indices = &index

		ps.Attributes["POSITION"] = indexPos
		if len(nd.TexCoords) > 0 {
			tmp++
			ps.Attributes["TEXCOORD_0"] = tmp
		}
		if len(nd.Normals) > 0 {
			tmp++
			ps.Attributes["NORMAL"] = tmp
		}
		ps.Mode = gltf.PrimitiveTriangles
		mesh.Primitives = append(mesh.Primitives, ps)

		indexacc := &gltf.Accessor{}
		indexacc.ComponentType = gltf.ComponentUint
		indexacc.ByteOffset = start * 12
		indexacc.Count = len(patch.Faces) * 3
		start += len(patch.Faces)
		bfindex := ctx.bvIndex
		indexacc.BufferView = &bfindex
		accessors = append(accessors, indexacc)
	}

	posacc := &gltf.Accessor{}
	posacc.ComponentType = gltf.ComponentFloat
	posacc.Type = gltf.AccessorVec3
	posacc.Count = len(nd.Vertices)

	bvPos := ctx.bvPos
	posacc.BufferView = &bvPos
	box := nd.GetBoundbox()
	posacc.Min = []float64{box[0], box[1], box[2]}
	posacc.Max = []float64{box[3], box[4], box[5]}
	accessors = append(accessors, posacc)

	if len(nd.TexCoords) > 0 {
		texacc := &gltf.Accessor{}
		texacc.ComponentType = gltf.ComponentFloat
		texacc.Type = gltf.AccessorVec2
		texacc.Count = len(nd.TexCoords)
		bvTex := ctx.bvTex
		texacc.BufferView = &bvTex
		accessors = append(accessors, texacc)
	}

	if len(nd.Normals) > 0 {
		nlacc := &gltf.Accessor{}
		nlacc.ComponentType = gltf.ComponentFloat
		nlacc.Type = gltf.AccessorVec3
		nlacc.Count = len(nd.Normals)
		bvNorm := ctx.bvNorm
		nlacc.BufferView = &bvNorm
		accessors = append(accessors, nlacc)
	}
	return mesh, accessors
}

func buildGltf[T float64 | float32](doc *gltf.Document, mh *BaseMesh[T], trans []*mat4.Mat[T], exportOutline bool, gpu_instance bool) error {
	ctx := &buildContext{}
	ctx.mtlSize = len(doc.Materials)

	for _, mstNd := range mh.Nodes {
		l := len(doc.Meshes)
		if exportOutline && len(mstNd.EdgeGroup) > 0 {
			doc.BufferViews = buildOutlineBuffer(ctx, doc.Buffers[0], doc.BufferViews, mstNd)

			var mesh *gltf.Mesh
			mesh, doc.Accessors = buildOutline(ctx, doc.Accessors, mstNd)
			doc.Meshes = append(doc.Meshes, mesh)
		} else {
			doc.BufferViews = buildMeshBuffer(ctx, doc.Buffers[0], doc.BufferViews, mstNd)

			var mesh *gltf.Mesh
			mesh, doc.Accessors = buildMesh(ctx, doc.Accessors, mstNd)
			doc.Meshes = append(doc.Meshes, mesh)
		}

		if trans == nil {
			doc.Scenes[0].Nodes = append(doc.Scenes[0].Nodes, len(doc.Nodes))
			node := &gltf.Node{}
			node.Mesh = &l
			doc.Nodes = append(doc.Nodes, node)
		} else {
			if gpu_instance {
				buildInstance(doc, l, trans)
			} else {
				for _, mt := range trans {
					position, quat, scale := mat4.Decompose(mt)
					nd := gltf.Node{
						Mesh:        &l,
						Translation: [3]float64{float64(float32(position[0])), float64(float32(position[1])), float64(float32(position[2]))},
						Rotation:    [4]float64{float64(float32(quat[0])), float64(float32(quat[1])), float64(float32(quat[2])), float64(float32(quat[3]))},
						Scale:       [3]float64{float64(float32(scale[0])), float64(float32(scale[1])), float64(float32(scale[2]))},
					}
					doc.Nodes = append(doc.Nodes, &nd)
					doc.Scenes[0].Nodes = append(doc.Scenes[0].Nodes, len(doc.Nodes)-1)
				}
			}
		}

	}

	err := fillMaterials[T](doc, mh.Materials)
	if err != nil {
		return err
	}

	return nil
}

func buildInstance[T float64 | float32](doc *gltf.Document, l int, trans []*mat4.Mat[T]) {
	bvIdx := len(doc.BufferViews)
	accInx := len(doc.Accessors)
	buf := bytes.NewBuffer([]byte{})
	startBytte := doc.Buffers[0].ByteLength
	for i, mt := range trans {
		position, quat, scale := mat4.Decompose(mt)
		pos := [3]float32{float32(position[0]), float32(position[1]), float32(position[2])}
		rot := [4]float32{float32(quat[0]), float32(quat[1]), float32(quat[2]), float32(quat[3])}
		scl := [3]float32{float32(scale[0]), float32(scale[1]), float32(scale[2])}

		binary.Write(buf, binary.LittleEndian, pos)
		binary.Write(buf, binary.LittleEndian, scl)
		binary.Write(buf, binary.LittleEndian, rot)

		posAcc := &gltf.Accessor{}
		posAcc.ComponentType = gltf.ComponentFloat
		posAcc.Type = gltf.AccessorVec3
		posAcc.Count = 1
		posAcc.BufferView = &bvIdx
		posAcc.ByteOffset = i * 40
		doc.Accessors = append(doc.Accessors, posAcc)

		sclAcc := &gltf.Accessor{}
		sclAcc.ComponentType = gltf.ComponentFloat
		sclAcc.Type = gltf.AccessorVec3
		sclAcc.Count = 1
		sclAcc.BufferView = &bvIdx
		sclAcc.ByteOffset = posAcc.ByteOffset + 12
		doc.Accessors = append(doc.Accessors, sclAcc)

		rotAcc := &gltf.Accessor{}
		rotAcc.ComponentType = gltf.ComponentFloat
		rotAcc.Type = gltf.AccessorVec4
		rotAcc.Count = 1
		rotAcc.BufferView = &bvIdx
		rotAcc.ByteOffset = sclAcc.ByteOffset + 12
		doc.Accessors = append(doc.Accessors, rotAcc)

		nd := gltf.Node{
			Mesh: &l,
			Extensions: map[string]interface{}{"EXT_mesh_gpu_instancing": map[string]interface{}{
				"attributes": map[string]interface{}{
					"TRANSLATION": accInx,
					"SCALE":       accInx + 1,
					"ROTATION":    accInx + 2,
				},
			}},
		}
		accInx += 3
		doc.Nodes = append(doc.Nodes, &nd)
		doc.Scenes[0].Nodes = append(doc.Scenes[0].Nodes, len(doc.Nodes)-1)
	}

	bv := &gltf.BufferView{}
	bv.Buffer = 0
	bv.ByteOffset = startBytte
	bv.ByteLength = buf.Len()
	doc.BufferViews = append(doc.BufferViews, bv)
	doc.Buffers[0].Data = append(doc.Buffers[0].Data, buf.Bytes()...)
	doc.Buffers[0].ByteLength += bv.ByteLength
}

func buildTextureBuffer(doc *gltf.Document, buffer *gltf.Buffer, texture *Texture) (*gltf.Texture, error) {
	spCount := len(doc.Samplers)
	imCount := len(doc.Images)

	tx := &gltf.Texture{Sampler: &spCount, Source: &imCount}

	gimg := &gltf.Image{}
	gimg.MimeType = "image/png"
	imgIndex := len(doc.BufferViews)
	gimg.BufferView = &imgIndex

	img, e := LoadTexture(texture, true)
	if e != nil {
		return nil, e
	}
	var bt []byte
	buf := bytes.NewBuffer(bt)
	png.Encode(buf, img)

	imgBuffView := &gltf.BufferView{}
	imgBuffView.ByteOffset = buffer.ByteLength
	imgBuffView.ByteLength = buf.Len()
	imgBuffView.Buffer = 0
	buffer.ByteLength += buf.Len()
	buffer.Data = append(buffer.Data, buf.Bytes()...)

	doc.BufferViews = append(doc.BufferViews, imgBuffView)
	doc.Images = append(doc.Images, gimg)

	var sp *gltf.Sampler
	if texture.Repeated {
		sp = &gltf.Sampler{WrapS: gltf.WrapRepeat, WrapT: gltf.WrapRepeat}
	} else {
		sp = &gltf.Sampler{WrapS: gltf.WrapClampToEdge, WrapT: gltf.WrapClampToEdge}
	}
	doc.Samplers = append(doc.Samplers, sp)

	return tx, nil
}

func fillMaterials[T float64 | float32](doc *gltf.Document, mts []MeshMaterial) error {
	texMap := make(map[int]int)
	useExtension := false
	for i := range mts {
		mtl := mts[i]

		gm := &gltf.Material{DoubleSided: true, AlphaMode: gltf.AlphaMask}
		gm.PBRMetallicRoughness = &gltf.PBRMetallicRoughness{BaseColorFactor: &[4]float64{1, 1, 1, 1}}
		gm.Extensions = make(map[string]interface{})
		var texMtl *TextureMaterial
		var cl *[4]float64
		switch ml := mtl.(type) {
		case *BaseMaterial:
			cl = &[4]float64{float64(float32(ml.Color[0]) / 255), float64(float32(ml.Color[1]) / 255), float64(float32(ml.Color[2]) / 255), float64(1 - float32(ml.Transparency))}
		case *PbrMaterial[T]:
			cl = &[4]float64{float64(float32(ml.Color[0]) / 255), float64(float32(ml.Color[1]) / 255), float64(float32(ml.Color[2]) / 255), float64(1 - float32(ml.Transparency))}
			mc := float64(ml.Metallic)
			gm.PBRMetallicRoughness.MetallicFactor = &mc
			rs := float64(ml.Roughness)
			gm.PBRMetallicRoughness.RoughnessFactor = &rs
			gm.EmissiveFactor[0] = float64(ml.Emissive[0]) / 255
			gm.EmissiveFactor[1] = float64(ml.Emissive[1]) / 255
			gm.EmissiveFactor[2] = float64(ml.Emissive[2]) / 255
			texMtl = &ml.TextureMaterial
		case *LambertMaterial:
			cl = &[4]float64{float64(float32(ml.Color[0]) / 255), float64(float32(ml.Color[1]) / 255), float64(float32(ml.Color[2]) / 255), float64(1 - float32(ml.Transparency))}
			texMtl = &ml.TextureMaterial

			spmtl := &specular.PBRSpecularGlossiness{
				DiffuseFactor: &[4]float64{float64((ml.Diffuse[0]) / 255), float64((ml.Diffuse[1]) / 255), float64((ml.Diffuse[2]) / 255), 1},
			}

			gm.EmissiveFactor[0] = float64(ml.Emissive[0]) / 255
			gm.EmissiveFactor[1] = float64(ml.Emissive[1]) / 255
			gm.EmissiveFactor[2] = float64(ml.Emissive[2]) / 255

			gm.Extensions[specular.ExtensionName] = spmtl
			useExtension = true
		case *PhongMaterial:
			cl = &[4]float64{float64(float32(ml.Color[0]) / 255), float64(float32(ml.Color[1]) / 255), float64(float32(ml.Color[2]) / 255), float64(1 - float32(ml.Transparency))}
			texMtl = &ml.TextureMaterial

			spmtl := &specular.PBRSpecularGlossiness{
				DiffuseFactor:    &[4]float64{float64(float32(ml.Diffuse[0]) / 255), float64(float32(ml.Diffuse[1]) / 255), float64(float32(ml.Diffuse[2]) / 255), 1},
				SpecularFactor:   &[3]float64{float64(float32(ml.Specular[0]) / 255), float64(float32(ml.Specular[1]) / 255), float64(float32(ml.Specular[2]) / 255)},
				GlossinessFactor: &ml.Shininess,
			}

			gm.EmissiveFactor[0] = float64(float32(ml.Emissive[0]) / 255)
			gm.EmissiveFactor[1] = float64(float32(ml.Emissive[1]) / 255)
			gm.EmissiveFactor[2] = float64(float32(ml.Emissive[2]) / 255)

			gm.Extensions[specular.ExtensionName] = spmtl
			useExtension = true
		case *TextureMaterial:
			texMtl = ml
			cl = &[4]float64{float64(float32(ml.Color[0]) / 255), float64(float32(ml.Color[1]) / 255), float64(float32(ml.Color[2]) / 255), float64(1 - float32(ml.Transparency))}
		}

		if texMtl != nil && texMtl.HasTexture() {
			if idx, ok := texMap[texMtl.Texture.Id]; ok {
				gm.PBRMetallicRoughness.BaseColorTexture = &gltf.TextureInfo{Index: idx}
			} else {
				texIndex := len(doc.Textures)
				texMap[texMtl.Texture.Id] = texIndex
				tex, err := buildTextureBuffer(doc, doc.Buffers[0], texMtl.Texture)

				if err != nil {
					return err
				}

				gm.PBRMetallicRoughness.BaseColorTexture = &gltf.TextureInfo{Index: texIndex}
				doc.Textures = append(doc.Textures, tex)
			}
		}

		if texMtl != nil && texMtl.HasNormalTexture() {
			if idx, ok := texMap[texMtl.Normal.Id]; ok {
				gm.NormalTexture = &gltf.NormalTexture{Index: &idx}
			} else {
				normalTexIndex := len(doc.Textures)
				texMap[texMtl.Normal.Id] = normalTexIndex
				tex, err := buildTextureBuffer(doc, doc.Buffers[0], texMtl.Normal)

				if err != nil {
					return err
				}
				gm.NormalTexture = &gltf.NormalTexture{Index: &normalTexIndex}
				doc.Textures = append(doc.Textures, tex)
			}
		}

		gm.PBRMetallicRoughness.BaseColorFactor = cl

		if gm.PBRMetallicRoughness.MetallicFactor == nil {
			mc := 0.0
			gm.PBRMetallicRoughness.MetallicFactor = &mc
		}

		if gm.PBRMetallicRoughness.RoughnessFactor == nil {
			rg := 1.0
			gm.PBRMetallicRoughness.RoughnessFactor = &rg
		}

		doc.Materials = append(doc.Materials, gm)
	}
	if useExtension {
		has := false
		for _, nm := range doc.ExtensionsUsed {
			if nm == specular.ExtensionName {
				has = true
				break
			}
		}
		if !has {
			doc.ExtensionsUsed = append(doc.ExtensionsUsed, specular.ExtensionName)
		}
	}
	return nil
}
