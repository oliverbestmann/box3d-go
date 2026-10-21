package b3

import "unsafe"

type (
	// DrawShape draws a shape. Get the geometry with the getters of Shape, like Shape.GetHull.
	DrawShape     func(shape Shape, transform WorldTransform, color HexColor)
	DrawSegment   func(p1, p2 Pos, color HexColor)
	DrawTransform func(transform WorldTransform)
	DrawPoint     func(p Pos, size float32, color HexColor)
	DrawSphere    func(p Pos, radius float32, color HexColor, alpha float32)
	DrawCapsule   func(p1, p2 Pos, radius float32, color HexColor, alpha float32)
	DrawBounds    func(aabb AABB, color HexColor)
	DrawBox       func(extents Vec3, transform WorldTransform, color HexColor)
	DrawString    func(p Pos, s string, color HexColor)
)

// DebugDraw holds the callbacks to draw a world. Start with DefaultDebugDraw,
// box3d draws nothing outside of DrawingBounds.
type DebugDraw struct {
	DrawShape     DrawShape
	DrawSegment   DrawSegment
	DrawTransform DrawTransform
	DrawPoint     DrawPoint
	DrawSphere    DrawSphere
	DrawCapsule   DrawCapsule
	DrawBounds    DrawBounds
	DrawBox       DrawBox
	DrawString    DrawString

	DrawingBounds AABB

	// ForceScale scales the drawn forces, JointScale the drawn joints.
	ForceScale float32
	JointScale float32

	DrawShapes          bool
	DrawJoints          bool
	DrawJointExtras     bool
	DrawShapeBounds     bool
	DrawMass            bool
	DrawSleep           bool
	DrawBodyNames       bool
	DrawContacts        bool
	DrawAnchorA         bool
	DrawGraphColors     bool
	DrawContactFeatures bool
	DrawContactNormals  bool
	DrawContactForces   bool
	DrawIslands         bool
}

// DefaultDebugDraw returns the defaults of b3DefaultDebugDraw, without callbacks.
func DefaultDebugDraw() DebugDraw {
	dd := b3DefaultDebugDraw(theStack)

	return DebugDraw{
		DrawingBounds:       dd.DrawingBounds,
		ForceScale:          dd.ForceScale,
		JointScale:          dd.JointScale,
		DrawShapes:          dd.DrawShapes != 0,
		DrawJoints:          dd.DrawJoints != 0,
		DrawJointExtras:     dd.DrawJointExtras != 0,
		DrawShapeBounds:     dd.DrawBounds != 0,
		DrawMass:            dd.DrawMass != 0,
		DrawSleep:           dd.DrawSleep != 0,
		DrawBodyNames:       dd.DrawBodyNames != 0,
		DrawContacts:        dd.DrawContacts != 0,
		DrawAnchorA:         dd.DrawAnchorA != 0,
		DrawGraphColors:     dd.DrawGraphColors != 0,
		DrawContactFeatures: dd.DrawContactFeatures != 0,
		DrawContactNormals:  dd.DrawContactNormals != 0,
		DrawContactForces:   dd.DrawContactForces != 0,
		DrawIslands:         dd.DrawIslands != 0,
	}
}

func (b World) Draw(draw DebugDraw, maskBits uint64) {
	// box3d passes the context back to the callbacks
	context := newHandle(&draw)
	defer handles.Delete(context)

	dd := b3DefaultDebugDraw(theStack)
	dd.DrawShapeFcn = __ccgo_fp(drawShapeCallback(cbDrawShape))
	dd.DrawSegmentFcn = __ccgo_fp(drawSegmentCallback(cbDrawSegment))
	dd.DrawTransformFcn = __ccgo_fp(drawTransformCallback(cbDrawTransform))
	dd.DrawPointFcn = __ccgo_fp(drawPointCallback(cbDrawPoint))
	dd.DrawSphereFcn = __ccgo_fp(drawSphereCallback(cbDrawSphere))
	dd.DrawCapsuleFcn = __ccgo_fp(drawCapsuleCallback(cbDrawCapsule))
	dd.DrawBoundsFcn = __ccgo_fp(drawBoundsCallback(cbDrawBounds))
	dd.DrawBoxFcn = __ccgo_fp(drawBoxCallback(cbDrawBox))
	dd.DrawStringFcn = __ccgo_fp(drawStringCallback(cbDrawString))

	dd.DrawingBounds = draw.DrawingBounds
	dd.ForceScale = draw.ForceScale
	dd.JointScale = draw.JointScale

	dd.DrawShapes = boolUint8(draw.DrawShapes)
	dd.DrawJoints = boolUint8(draw.DrawJoints)
	dd.DrawJointExtras = boolUint8(draw.DrawJointExtras)
	dd.DrawBounds = boolUint8(draw.DrawShapeBounds)
	dd.DrawMass = boolUint8(draw.DrawMass)
	dd.DrawSleep = boolUint8(draw.DrawSleep)
	dd.DrawBodyNames = boolUint8(draw.DrawBodyNames)
	dd.DrawContacts = boolUint8(draw.DrawContacts)
	dd.DrawAnchorA = boolUint8(draw.DrawAnchorA)
	dd.DrawGraphColors = boolUint8(draw.DrawGraphColors)
	dd.DrawContactFeatures = boolUint8(draw.DrawContactFeatures)
	dd.DrawContactNormals = boolUint8(draw.DrawContactNormals)
	dd.DrawContactForces = boolUint8(draw.DrawContactForces)
	dd.DrawIslands = boolUint8(draw.DrawIslands)

	dd.Context = context

	ddStack := copyToStack(theStack, dd)
	defer ddStack.Free()

	b3World_Draw(theStack, b.Id, ddStack.Addr, maskBits)
}

// box3d asks for a user shape once per shape and passes it to DrawShapeFcn. The user shape is
// the shape id, the ids are 8 bytes like a uintptr on the 64-bit platforms that ccgo targets.
var _ [unsafe.Sizeof(uintptr(0)) - unsafe.Sizeof(ShapeId{})]byte

func cbCreateDebugShape(_ *_Stack, debugShape uintptr, _ uintptr) uintptr {
	// never 0, Index1 starts at 1
	return *(*uintptr)(unsafe.Pointer(&(*b3DebugShape)(unsafe.Pointer(debugShape)).ShapeId))
}

func cbDestroyDebugShape(_ *_Stack, _ uintptr, _ uintptr) {}

func cbDrawShape(_ *_Stack, userShape uintptr, transform WorldTransform, color HexColor, context uintptr) {
	ctx := handleOf[*DebugDraw](context)
	if ctx.DrawShape != nil {
		ctx.DrawShape(Shape{Id: *(*ShapeId)(unsafe.Pointer(&userShape))}, transform, color)
	}
}

func cbDrawSegment(_ *_Stack, p1, p2 Pos, color HexColor, context uintptr) {
	ctx := handleOf[*DebugDraw](context)
	if ctx.DrawSegment != nil {
		ctx.DrawSegment(p1, p2, color)
	}
}

func cbDrawTransform(_ *_Stack, transform WorldTransform, context uintptr) {
	ctx := handleOf[*DebugDraw](context)
	if ctx.DrawTransform != nil {
		ctx.DrawTransform(transform)
	}
}

func cbDrawPoint(_ *_Stack, p Pos, size float32, color HexColor, context uintptr) {
	ctx := handleOf[*DebugDraw](context)
	if ctx.DrawPoint != nil {
		ctx.DrawPoint(p, size, color)
	}
}

func cbDrawSphere(_ *_Stack, p Pos, radius float32, color HexColor, alpha float32, context uintptr) {
	ctx := handleOf[*DebugDraw](context)
	if ctx.DrawSphere != nil {
		ctx.DrawSphere(p, radius, color, alpha)
	}
}

func cbDrawCapsule(_ *_Stack, p1, p2 Pos, radius float32, color HexColor, alpha float32, context uintptr) {
	ctx := handleOf[*DebugDraw](context)
	if ctx.DrawCapsule != nil {
		ctx.DrawCapsule(p1, p2, radius, color, alpha)
	}
}

func cbDrawBounds(_ *_Stack, aabb AABB, color HexColor, context uintptr) {
	ctx := handleOf[*DebugDraw](context)
	if ctx.DrawBounds != nil {
		ctx.DrawBounds(aabb, color)
	}
}

func cbDrawBox(_ *_Stack, extents Vec3, transform WorldTransform, color HexColor, context uintptr) {
	ctx := handleOf[*DebugDraw](context)
	if ctx.DrawBox != nil {
		ctx.DrawBox(extents, transform, color)
	}
}

func cbDrawString(_ *_Stack, p Pos, s uintptr, color HexColor, context uintptr) {
	ctx := handleOf[*DebugDraw](context)
	if ctx.DrawString != nil {
		ctx.DrawString(p, toString(s), color)
	}
}
