package b3

import (
	"unsafe"
	"runtime"
)

var _ unsafe.Pointer
var _ runtime.MemStats

type Body struct {
	Id BodyId
}
type Shape struct {
	Id ShapeId
}
type World struct {
	Id WorldId
}
type Joint struct {
	Id JointId
}
type Contact struct {
	Id ContactId
}
type ParallelJoint struct {
	Joint
}
type DistanceJoint struct {
	Joint
}
type FilterJoint struct {
	Joint
}
type MotorJoint struct {
	Joint
}
type PrismaticJoint struct {
	Joint
}
type RevoluteJoint struct {
	Joint
}
type SphericalJoint struct {
	Joint
}
type WeldJoint struct {
	Joint
}
type WheelJoint struct {
	Joint
}

func (b World) CreateBody(def BodyDef) (r Body) {
	defer runtime.KeepAlive(def)
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b3CreateBody(theStack, b.Id, defStack.Addr)
	return
}
func (b Body) DestroyBody()                  { b3DestroyBody(theStack, b.Id) }
func (b Body) GetContactCapacity() (r int32) { r = b3Body_GetContactCapacity(theStack, b.Id); return }
func (b Body) ComputeAABB() (r AABB)         { r = b3Body_ComputeAABB(theStack, b.Id); return }
func (b Body) CastRay(origin Pos, translation Vec3, _filter QueryFilter, maxFraction float32, bodyTransform WorldTransform) (r1 BodyCastResult) {
	r1 = b3Body_CastRay(theStack, b.Id, origin, translation, _filter, maxFraction, bodyTransform)
	return
}
func (b Body) CastShape(origin Pos, proxy ShapeProxy, translation Vec3, _filter QueryFilter, maxFraction float32, canEncroach uint8, bodyTransform WorldTransform) (r1 BodyCastResult) {
	defer runtime.KeepAlive(proxy)
	proxyStack := copyToStackAs[cShapeProxy](theStack, proxy)
	defer proxyStack.Free()
	r1 = b3Body_CastShape(theStack, b.Id, origin, proxyStack.Addr, translation, _filter, maxFraction, canEncroach, bodyTransform)
	return
}
func (b Body) OverlapShape(origin Pos, proxy ShapeProxy, _filter QueryFilter, bodyTransform WorldTransform) (r1 uint8) {
	defer runtime.KeepAlive(proxy)
	proxyStack := copyToStackAs[cShapeProxy](theStack, proxy)
	defer proxyStack.Free()
	r1 = b3Body_OverlapShape(theStack, b.Id, origin, proxyStack.Addr, _filter, bodyTransform)
	return
}
func (b Body) TimeOfImpactMover(origin Pos, mover Capsule, moverTranslation Vec3, _filter QueryFilter, bodyTransform1 WorldTransform, bodyTransform2 WorldTransform) (r1 BodyTOIResult) {
	defer runtime.KeepAlive(mover)
	moverStack := copyToStack(theStack, mover)
	defer moverStack.Free()
	r1 = b3Body_TimeOfImpactMover(theStack, b.Id, origin, moverStack.Addr, moverTranslation, _filter, bodyTransform1, bodyTransform2)
	return
}
func (b Body) GetPosition() (r Pos)             { r = b3Body_GetPosition(theStack, b.Id); return }
func (b Body) GetRotation() (r Quat)            { r = b3Body_GetRotation(theStack, b.Id); return }
func (b Body) GetTransform() (r WorldTransform) { r = b3Body_GetTransform(theStack, b.Id); return }
func (b Body) GetLocalPoint(worldPoint Pos) (r Vec3) {
	r = b3Body_GetLocalPoint(theStack, b.Id, worldPoint)
	return
}
func (b Body) GetWorldPoint(localPoint Vec3) (r1 Pos) {
	r1 = b3Body_GetWorldPoint(theStack, b.Id, localPoint)
	return
}
func (b Body) GetLocalVector(worldVector Vec3) (r Vec3) {
	r = b3Body_GetLocalVector(theStack, b.Id, worldVector)
	return
}
func (b Body) GetWorldVector(localVector Vec3) (r Vec3) {
	r = b3Body_GetWorldVector(theStack, b.Id, localVector)
	return
}
func (b Body) SetTransform(position Pos, rotation Quat) {
	b3Body_SetTransform(theStack, b.Id, position, rotation)
}
func (b Body) GetLinearVelocity() (r Vec3)  { r = b3Body_GetLinearVelocity(theStack, b.Id); return }
func (b Body) GetAngularVelocity() (r Vec3) { r = b3Body_GetAngularVelocity(theStack, b.Id); return }
func (b Body) SetLinearVelocity(linearVelocity Vec3) {
	b3Body_SetLinearVelocity(theStack, b.Id, linearVelocity)
}
func (b Body) SetAngularVelocity(angularVelocity Vec3) {
	b3Body_SetAngularVelocity(theStack, b.Id, angularVelocity)
}
func (b Body) SetTargetTransform(target WorldTransform, timeStep float32, wake uint8) {
	b3Body_SetTargetTransform(theStack, b.Id, target, timeStep, wake)
}
func (b Body) GetLocalPointVelocity(localPoint Vec3) (r1 Vec3) {
	r1 = b3Body_GetLocalPointVelocity(theStack, b.Id, localPoint)
	return
}
func (b Body) GetWorldPointVelocity(worldPoint Pos) (r1 Vec3) {
	r1 = b3Body_GetWorldPointVelocity(theStack, b.Id, worldPoint)
	return
}
func (b Body) ApplyForce(force Vec3, point Pos, wake uint8) {
	b3Body_ApplyForce(theStack, b.Id, force, point, wake)
}
func (b Body) ApplyForceToCenter(force Vec3, wake uint8) {
	b3Body_ApplyForceToCenter(theStack, b.Id, force, wake)
}
func (b Body) ApplyTorque(torque Vec3, wake uint8) { b3Body_ApplyTorque(theStack, b.Id, torque, wake) }
func (b Body) ApplyLinearImpulse(impulse Vec3, point Pos, wake uint8) {
	b3Body_ApplyLinearImpulse(theStack, b.Id, impulse, point, wake)
}
func (b Body) ApplyLinearImpulseToCenter(impulse Vec3, wake uint8) {
	b3Body_ApplyLinearImpulseToCenter(theStack, b.Id, impulse, wake)
}
func (b Body) ApplyAngularImpulse(impulse Vec3, wake uint8) {
	b3Body_ApplyAngularImpulse(theStack, b.Id, impulse, wake)
}
func (b Body) GetType() (r BodyType)  { r = b3Body_GetType(theStack, b.Id); return }
func (b Body) SetType(type1 BodyType) { b3Body_SetType(theStack, b.Id, type1) }
func (b Body) SetName(name string) {
	nameStack := theStack.toCString(name)
	defer nameStack.Free()
	b3Body_SetName(theStack, b.Id, nameStack.Addr)
}
func (b Body) GetName() (r string)         { r = toString(b3Body_GetName(theStack, b.Id)); return }
func (b Body) SetUserData(userData uint64) { b3Body_SetUserData(theStack, b.Id, uintptr(userData)) }
func (b Body) GetUserData() (r uint64)     { r = uint64(b3Body_GetUserData(theStack, b.Id)); return }
func (b Body) GetMass() (r float32)        { r = b3Body_GetMass(theStack, b.Id); return }
func (b Body) GetLocalRotationalInertia() (r Matrix3) {
	r = b3Body_GetLocalRotationalInertia(theStack, b.Id)
	return
}
func (b Body) GetInverseMass() (r float32) { r = b3Body_GetInverseMass(theStack, b.Id); return }
func (b Body) GetWorldInverseRotationalInertia() (r Matrix3) {
	r = b3Body_GetWorldInverseRotationalInertia(theStack, b.Id)
	return
}
func (b Body) GetLocalCenter() (r Vec3)      { r = b3Body_GetLocalCenter(theStack, b.Id); return }
func (b Body) GetWorldCenter() (r Pos)       { r = b3Body_GetWorldCenter(theStack, b.Id); return }
func (b Body) SetMassData(massData MassData) { b3Body_SetMassData(theStack, b.Id, massData) }
func (b Body) GetMassData() (r MassData)     { r = b3Body_GetMassData(theStack, b.Id); return }
func (b Body) ApplyMassFromShapes()          { b3Body_ApplyMassFromShapes(theStack, b.Id) }
func (b Body) SetLinearDamping(linearDamping float32) {
	b3Body_SetLinearDamping(theStack, b.Id, linearDamping)
}
func (b Body) GetLinearDamping() (r float32) { r = b3Body_GetLinearDamping(theStack, b.Id); return }
func (b Body) SetAngularDamping(angularDamping float32) {
	b3Body_SetAngularDamping(theStack, b.Id, angularDamping)
}
func (b Body) GetAngularDamping() (r float32) { r = b3Body_GetAngularDamping(theStack, b.Id); return }
func (b Body) SetGravityScale(gravityScale float32) {
	b3Body_SetGravityScale(theStack, b.Id, gravityScale)
}
func (b Body) GetGravityScale() (r float32) { r = b3Body_GetGravityScale(theStack, b.Id); return }
func (b Body) IsAwake() (r uint8)           { r = b3Body_IsAwake(theStack, b.Id); return }
func (b Body) SetAwake(awake uint8)         { b3Body_SetAwake(theStack, b.Id, awake) }
func (b Body) IsEnabled() (r uint8)         { r = b3Body_IsEnabled(theStack, b.Id); return }
func (b Body) IsSleepEnabled() (r uint8)    { r = b3Body_IsSleepEnabled(theStack, b.Id); return }
func (b Body) SetSleepThreshold(sleepThreshold float32) {
	b3Body_SetSleepThreshold(theStack, b.Id, sleepThreshold)
}
func (b Body) GetSleepThreshold() (r float32) { r = b3Body_GetSleepThreshold(theStack, b.Id); return }
func (b Body) SetSafetyFactor(safetyFactor float32) {
	b3Body_SetSafetyFactor(theStack, b.Id, safetyFactor)
}
func (b Body) GetSafetyFactor() (r float32)     { r = b3Body_GetSafetyFactor(theStack, b.Id); return }
func (b Body) EnableSleep(enableSleep uint8)    { b3Body_EnableSleep(theStack, b.Id, enableSleep) }
func (b Body) Disable()                         { b3Body_Disable(theStack, b.Id) }
func (b Body) Enable()                          { b3Body_Enable(theStack, b.Id) }
func (b Body) SetMotionLocks(locks MotionLocks) { b3Body_SetMotionLocks(theStack, b.Id, locks) }
func (b Body) GetMotionLocks() (r MotionLocks)  { r = b3Body_GetMotionLocks(theStack, b.Id); return }
func (b Body) SetBullet(flag uint8)             { b3Body_SetBullet(theStack, b.Id, flag) }
func (b Body) IsBullet() (r uint8)              { r = b3Body_IsBullet(theStack, b.Id); return }
func (b Body) AllowFastRotation(flag uint8)     { b3Body_AllowFastRotation(theStack, b.Id, flag) }
func (b Body) IsFastRotationAllowed() (r uint8) {
	r = b3Body_IsFastRotationAllowed(theStack, b.Id)
	return
}
func (b Body) EnableContactRecycling(flag uint8) { b3Body_EnableContactRecycling(theStack, b.Id, flag) }
func (b Body) IsContactRecyclingEnabled() (r uint8) {
	r = b3Body_IsContactRecyclingEnabled(theStack, b.Id)
	return
}
func (b Body) EnableHitEvents(flag uint8)   { b3Body_EnableHitEvents(theStack, b.Id, flag) }
func (b Body) GetWorld() (r World)          { r.Id = b3Body_GetWorld(theStack, b.Id); return }
func (b Body) GetShapeCount() (r int32)     { r = b3Body_GetShapeCount(theStack, b.Id); return }
func (b Body) GetJointCount() (r int32)     { r = b3Body_GetJointCount(theStack, b.Id); return }
func (b Body) GetMinExtent() (r float32)    { r = b3Body_GetMinExtent(theStack, b.Id); return }
func (b Body) GetMaxExtent() (r Vec3)       { r = b3Body_GetMaxExtent(theStack, b.Id); return }
func (b Body) GetMaxExtentOrigin() (r Vec3) { r = b3Body_GetMaxExtentOrigin(theStack, b.Id); return }
func ComputeCapsuleMass(shape Capsule, density float32) (r1 MassData) {
	defer runtime.KeepAlive(shape)
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r1 = b3ComputeCapsuleMass(theStack, shapeStack.Addr, density)
	return
}
func ComputeCapsuleAABB(shape Capsule, transform Transform) (r1 AABB) {
	defer runtime.KeepAlive(shape)
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r1 = b3ComputeCapsuleAABB(theStack, shapeStack.Addr, transform)
	return
}
func OverlapCapsule(shape Capsule, shapeTransform Transform, proxy ShapeProxy) (r uint8) {
	defer runtime.KeepAlive(proxy)
	defer runtime.KeepAlive(shape)
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	proxyStack := copyToStackAs[cShapeProxy](theStack, proxy)
	defer proxyStack.Free()
	r = b3OverlapCapsule(theStack, shapeStack.Addr, shapeTransform, proxyStack.Addr)
	return
}
func RayCastCapsule(shape Capsule, input RayCastInput) (r1 CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(shape)
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r1 = b3RayCastCapsule(theStack, shapeStack.Addr, inputStack.Addr)
	return
}
func ShapeCastCapsule(capsule Capsule, input ShapeCastInput) (r CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(capsule)
	capsuleStack := copyToStack(theStack, capsule)
	defer capsuleStack.Free()
	inputStack := copyToStackAs[cShapeCastInput](theStack, input)
	defer inputStack.Free()
	r = b3ShapeCastCapsule(theStack, capsuleStack.Addr, inputStack.Addr)
	return
}
func GetCompoundCapsule(compound *CompoundData, index int32) (r CompoundCapsule) {
	defer runtime.KeepAlive(compound)
	escapes(compound)
	r = b3GetCompoundCapsule(theStack, uintptr(unsafe.Pointer(compound)), index)
	return
}
func GetCompoundHull(compound *CompoundData, index int32) (r CompoundHull) {
	defer runtime.KeepAlive(compound)
	escapes(compound)
	r = b3GetCompoundHull(theStack, uintptr(unsafe.Pointer(compound)), index)
	return
}
func GetCompoundMesh(compound *CompoundData, index int32) (r CompoundMesh) {
	defer runtime.KeepAlive(compound)
	escapes(compound)
	r = b3GetCompoundMesh(theStack, uintptr(unsafe.Pointer(compound)), index)
	return
}
func GetCompoundSphere(compound *CompoundData, index int32) (r CompoundSphere) {
	defer runtime.KeepAlive(compound)
	escapes(compound)
	r = b3GetCompoundSphere(theStack, uintptr(unsafe.Pointer(compound)), index)
	return
}
func GetCompoundChild(compound *CompoundData, childIndex int32) (r ChildShape) {
	defer runtime.KeepAlive(compound)
	escapes(compound)
	r = b3GetCompoundChild(theStack, uintptr(unsafe.Pointer(compound)), childIndex)
	return
}
func destroyCompound(compound *CompoundData) {
	defer runtime.KeepAlive(compound)
	escapes(compound)
	b3DestroyCompound(theStack, uintptr(unsafe.Pointer(compound)))
}
func ComputeCompoundAABB(shape *CompoundData, transform1 Transform) (r AABB) {
	defer runtime.KeepAlive(shape)
	escapes(shape)
	r = b3ComputeCompoundAABB(theStack, uintptr(unsafe.Pointer(shape)), transform1)
	return
}
func OverlapCompound(shape *CompoundData, shapeTransform Transform, proxy ShapeProxy) (r uint8) {
	defer runtime.KeepAlive(proxy)
	defer runtime.KeepAlive(shape)
	escapes(shape)
	proxyStack := copyToStackAs[cShapeProxy](theStack, proxy)
	defer proxyStack.Free()
	r = b3OverlapCompound(theStack, uintptr(unsafe.Pointer(shape)), shapeTransform, proxyStack.Addr)
	return
}
func RayCastCompound(shape *CompoundData, input RayCastInput) (r CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(shape)
	escapes(shape)
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b3RayCastCompound(theStack, uintptr(unsafe.Pointer(shape)), inputStack.Addr)
	return
}
func ShapeCastCompound(shape *CompoundData, input ShapeCastInput) (r1 CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(shape)
	escapes(shape)
	inputStack := copyToStackAs[cShapeCastInput](theStack, input)
	defer inputStack.Free()
	r1 = b3ShapeCastCompound(theStack, uintptr(unsafe.Pointer(shape)), inputStack.Addr)
	return
}
func GetGraphColor(index int32) (r HexColor)     { r = b3GetGraphColor(theStack, index); return }
func (b Contact) GetData() (r ContactData)       { r = b3Contact_GetData(theStack, b.Id); return }
func SetLengthUnitsPerMeter(lengthUnits float32) { b3SetLengthUnitsPerMeter(theStack, lengthUnits) }
func GetLengthUnitsPerMeter() (r float32)        { r = b3GetLengthUnitsPerMeter(theStack); return }
func GetVersion() (r Version)                    { r = b3GetVersion(theStack); return }
func IsDoublePrecision() (r uint8)               { r = b3IsDoublePrecision(theStack); return }
func GetMaxManifoldPoints() (r int32)            { r = b3GetMaxManifoldPoints(theStack); return }
func GetByteCount() (r int64)                    { r = b3GetByteCount(theStack); return }
func ShapeCast(input ShapeCastPairInput) (r CastOutput) {
	defer runtime.KeepAlive(input)
	inputStack := copyToStackAs[cShapeCastPairInput](theStack, input)
	defer inputStack.Free()
	r = b3ShapeCast(theStack, inputStack.Addr)
	return
}
func GetSweepTransform(sweep Sweep, time float32) (r Transform) {
	defer runtime.KeepAlive(sweep)
	sweepStack := copyToStack(theStack, sweep)
	defer sweepStack.Free()
	r = b3GetSweepTransform(theStack, sweepStack.Addr, time)
	return
}
func TimeOfImpact(input TOIInput) (r TOIOutput) {
	defer runtime.KeepAlive(input)
	inputStack := copyToStackAs[cTOIInput](theStack, input)
	defer inputStack.Free()
	r = b3TimeOfImpact(theStack, inputStack.Addr)
	return
}
func (b DistanceJoint) SetLength(length float32) { b3DistanceJoint_SetLength(theStack, b.Id, length) }
func (b DistanceJoint) GetLength() (r float32)   { r = b3DistanceJoint_GetLength(theStack, b.Id); return }
func (b DistanceJoint) EnableLimit(enableLimit uint8) {
	b3DistanceJoint_EnableLimit(theStack, b.Id, enableLimit)
}
func (b DistanceJoint) IsLimitEnabled() (r uint8) {
	r = b3DistanceJoint_IsLimitEnabled(theStack, b.Id)
	return
}
func (b DistanceJoint) SetLengthRange(minLength float32, maxLength float32) {
	b3DistanceJoint_SetLengthRange(theStack, b.Id, minLength, maxLength)
}
func (b DistanceJoint) GetMinLength() (r float32) {
	r = b3DistanceJoint_GetMinLength(theStack, b.Id)
	return
}
func (b DistanceJoint) GetMaxLength() (r float32) {
	r = b3DistanceJoint_GetMaxLength(theStack, b.Id)
	return
}
func (b DistanceJoint) GetCurrentLength() (r1 float32) {
	r1 = b3DistanceJoint_GetCurrentLength(theStack, b.Id)
	return
}
func (b DistanceJoint) EnableSpring(enableSpring uint8) {
	b3DistanceJoint_EnableSpring(theStack, b.Id, enableSpring)
}
func (b DistanceJoint) IsSpringEnabled() (r uint8) {
	r = b3DistanceJoint_IsSpringEnabled(theStack, b.Id)
	return
}
func (b DistanceJoint) SetSpringForceRange(lowerForce float32, upperForce float32) {
	b3DistanceJoint_SetSpringForceRange(theStack, b.Id, lowerForce, upperForce)
}
func (b DistanceJoint) SetSpringHertz(hertz float32) {
	b3DistanceJoint_SetSpringHertz(theStack, b.Id, hertz)
}
func (b DistanceJoint) SetSpringDampingRatio(dampingRatio float32) {
	b3DistanceJoint_SetSpringDampingRatio(theStack, b.Id, dampingRatio)
}
func (b DistanceJoint) GetSpringHertz() (r float32) {
	r = b3DistanceJoint_GetSpringHertz(theStack, b.Id)
	return
}
func (b DistanceJoint) GetSpringDampingRatio() (r float32) {
	r = b3DistanceJoint_GetSpringDampingRatio(theStack, b.Id)
	return
}
func (b DistanceJoint) EnableMotor(enableMotor uint8) {
	b3DistanceJoint_EnableMotor(theStack, b.Id, enableMotor)
}
func (b DistanceJoint) IsMotorEnabled() (r uint8) {
	r = b3DistanceJoint_IsMotorEnabled(theStack, b.Id)
	return
}
func (b DistanceJoint) SetMotorSpeed(motorSpeed float32) {
	b3DistanceJoint_SetMotorSpeed(theStack, b.Id, motorSpeed)
}
func (b DistanceJoint) GetMotorSpeed() (r float32) {
	r = b3DistanceJoint_GetMotorSpeed(theStack, b.Id)
	return
}
func (b DistanceJoint) GetMotorForce() (r float32) {
	r = b3DistanceJoint_GetMotorForce(theStack, b.Id)
	return
}
func (b DistanceJoint) SetMaxMotorForce(force float32) {
	b3DistanceJoint_SetMaxMotorForce(theStack, b.Id, force)
}
func (b DistanceJoint) GetMaxMotorForce() (r float32) {
	r = b3DistanceJoint_GetMaxMotorForce(theStack, b.Id)
	return
}
func ComputeHeightFieldAABB(shape *HeightFieldData, transform1 Transform) (r AABB) {
	defer runtime.KeepAlive(shape)
	escapes(shape)
	r = b3ComputeHeightFieldAABB(theStack, uintptr(unsafe.Pointer(shape)), transform1)
	return
}
func RayCastHeightField(heightField *HeightFieldData, input RayCastInput) (r CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(heightField)
	escapes(heightField)
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b3RayCastHeightField(theStack, uintptr(unsafe.Pointer(heightField)), inputStack.Addr)
	return
}
func ShapeCastHeightField(heightField *HeightFieldData, input ShapeCastInput) (r1 CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(heightField)
	escapes(heightField)
	inputStack := copyToStackAs[cShapeCastInput](theStack, input)
	defer inputStack.Free()
	r1 = b3ShapeCastHeightField(theStack, uintptr(unsafe.Pointer(heightField)), inputStack.Addr)
	return
}
func OverlapHeightField(shape *HeightFieldData, shapeTransform Transform, proxy ShapeProxy) (r uint8) {
	defer runtime.KeepAlive(proxy)
	defer runtime.KeepAlive(shape)
	escapes(shape)
	proxyStack := copyToStackAs[cShapeProxy](theStack, proxy)
	defer proxyStack.Free()
	r = b3OverlapHeightField(theStack, uintptr(unsafe.Pointer(shape)), shapeTransform, proxyStack.Addr)
	return
}
func CreateGrid(rowCount int32, columnCount int32, scale Vec3, makeHoles uint8) (r OwnedHeightField) {
	r = OwnedHeightField{(*HeightFieldData)(unsafe.Pointer(b3CreateGrid(theStack, rowCount, columnCount, scale, makeHoles)))}
	return
}
func CreateWave(rowCount int32, columnCount int32, scale Vec3, rowFrequency float32, columnFrequency float32, makeHoles uint8) (r OwnedHeightField) {
	r = OwnedHeightField{(*HeightFieldData)(unsafe.Pointer(b3CreateWave(theStack, rowCount, columnCount, scale, rowFrequency, columnFrequency, makeHoles)))}
	return
}
func destroyHeightField(heightField *HeightFieldData) {
	defer runtime.KeepAlive(heightField)
	escapes(heightField)
	b3DestroyHeightField(theStack, uintptr(unsafe.Pointer(heightField)))
}
func CreateCylinder(height float32, radius float32, yOffset float32, sides int32) (r OwnedHull) {
	r = OwnedHull{(*HullData)(unsafe.Pointer(b3CreateCylinder(theStack, height, radius, yOffset, sides)))}
	return
}
func CreateCone(height float32, radius1 float32, radius2 float32, slices int32) (r OwnedHull) {
	r = OwnedHull{(*HullData)(unsafe.Pointer(b3CreateCone(theStack, height, radius1, radius2, slices)))}
	return
}
func CreateRock(radius float32) (r OwnedHull) {
	r = OwnedHull{(*HullData)(unsafe.Pointer(b3CreateRock(theStack, radius)))}
	return
}
func CreateComplexHull(radius float32) (r OwnedHull) {
	r = OwnedHull{(*HullData)(unsafe.Pointer(b3CreateComplexHull(theStack, radius)))}
	return
}
func CloneHull(hull *HullData) (r OwnedHull) {
	defer runtime.KeepAlive(hull)
	escapes(hull)
	r = OwnedHull{(*HullData)(unsafe.Pointer(b3CloneHull(theStack, uintptr(unsafe.Pointer(hull)))))}
	return
}
func CloneAndTransformHull(original *HullData, transform Transform, scale Vec3) (r OwnedHull) {
	defer runtime.KeepAlive(original)
	escapes(original)
	r = OwnedHull{(*HullData)(unsafe.Pointer(b3CloneAndTransformHull(theStack, uintptr(unsafe.Pointer(original)), transform, scale)))}
	return
}
func destroyHull(hull *HullData) {
	defer runtime.KeepAlive(hull)
	escapes(hull)
	b3DestroyHull(theStack, uintptr(unsafe.Pointer(hull)))
}
func ComputeHullMass(shape *HullData, density float32) (r MassData) {
	defer runtime.KeepAlive(shape)
	escapes(shape)
	r = b3ComputeHullMass(theStack, uintptr(unsafe.Pointer(shape)), density)
	return
}
func ComputeHullAABB(shape *HullData, transform1 Transform) (r AABB) {
	defer runtime.KeepAlive(shape)
	escapes(shape)
	r = b3ComputeHullAABB(theStack, uintptr(unsafe.Pointer(shape)), transform1)
	return
}
func OverlapHull(shape *HullData, shapeTransform Transform, proxy ShapeProxy) (r uint8) {
	defer runtime.KeepAlive(proxy)
	defer runtime.KeepAlive(shape)
	escapes(shape)
	proxyStack := copyToStackAs[cShapeProxy](theStack, proxy)
	defer proxyStack.Free()
	r = b3OverlapHull(theStack, uintptr(unsafe.Pointer(shape)), shapeTransform, proxyStack.Addr)
	return
}
func RayCastHull(shape *HullData, input RayCastInput) (r CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(shape)
	escapes(shape)
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b3RayCastHull(theStack, uintptr(unsafe.Pointer(shape)), inputStack.Addr)
	return
}
func ShapeCastHull(shape *HullData, input ShapeCastInput) (r CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(shape)
	escapes(shape)
	inputStack := copyToStackAs[cShapeCastInput](theStack, input)
	defer inputStack.Free()
	r = b3ShapeCastHull(theStack, uintptr(unsafe.Pointer(shape)), inputStack.Addr)
	return
}
func MakeTransformedBoxHull(hx float32, hy float32, hz float32, transform2 Transform) (r BoxHull) {
	r = b3MakeTransformedBoxHull(theStack, hx, hy, hz, transform2)
	return
}
func MakeCubeHull(halfWidth float32) (r BoxHull) { r = b3MakeCubeHull(theStack, halfWidth); return }
func MakeOffsetBoxHull(hx float32, hy float32, hz float32, offset Vec3) (r BoxHull) {
	r = b3MakeOffsetBoxHull(theStack, hx, hy, hz, offset)
	return
}
func MakeBoxHull(hx float32, hy float32, hz float32) (r BoxHull) {
	r = b3MakeBoxHull(theStack, hx, hy, hz)
	return
}
func MakeScaledBoxHull(halfWidths Vec3, transform Transform, postScale Vec3) (r BoxHull) {
	r = b3MakeScaledBoxHull(theStack, halfWidths, transform, postScale)
	return
}
func DefaultParallelJointDef() (r ParallelJointDef) { r = b3DefaultParallelJointDef(theStack); return }
func DefaultDistanceJointDef() (r DistanceJointDef) { r = b3DefaultDistanceJointDef(theStack); return }
func DefaultMotorJointDef() (r MotorJointDef)       { r = b3DefaultMotorJointDef(theStack); return }
func DefaultFilterJointDef() (r FilterJointDef)     { r = b3DefaultFilterJointDef(theStack); return }
func DefaultPrismaticJointDef() (r PrismaticJointDef) {
	r = b3DefaultPrismaticJointDef(theStack)
	return
}
func DefaultRevoluteJointDef() (r RevoluteJointDef) { r = b3DefaultRevoluteJointDef(theStack); return }
func DefaultSphericalJointDef() (r SphericalJointDef) {
	r = b3DefaultSphericalJointDef(theStack)
	return
}
func DefaultWeldJointDef() (r WeldJointDef)   { r = b3DefaultWeldJointDef(theStack); return }
func DefaultWheelJointDef() (r WheelJointDef) { r = b3DefaultWheelJointDef(theStack); return }
func DefaultExplosionDef() (r ExplosionDef)   { r = b3DefaultExplosionDef(theStack); return }
func (b Joint) SetConstraintTuning(hertz float32, dampingRatio float32) {
	b3Joint_SetConstraintTuning(theStack, b.Id, hertz, dampingRatio)
}
func (b Joint) SetForceThreshold(threshold float32) {
	b3Joint_SetForceThreshold(theStack, b.Id, threshold)
}
func (b Joint) GetForceThreshold() (r float32) { r = b3Joint_GetForceThreshold(theStack, b.Id); return }
func (b Joint) SetTorqueThreshold(threshold float32) {
	b3Joint_SetTorqueThreshold(theStack, b.Id, threshold)
}
func (b Joint) GetTorqueThreshold() (r float32) {
	r = b3Joint_GetTorqueThreshold(theStack, b.Id)
	return
}
func (b World) CreateDistanceJoint(def DistanceJointDef) (r DistanceJoint) {
	defer runtime.KeepAlive(def)
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b3CreateDistanceJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateMotorJoint(def MotorJointDef) (r MotorJoint) {
	defer runtime.KeepAlive(def)
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b3CreateMotorJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateFilterJoint(def FilterJointDef) (r FilterJoint) {
	defer runtime.KeepAlive(def)
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b3CreateFilterJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateParallelJoint(def ParallelJointDef) (r ParallelJoint) {
	defer runtime.KeepAlive(def)
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b3CreateParallelJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreatePrismaticJoint(def PrismaticJointDef) (r PrismaticJoint) {
	defer runtime.KeepAlive(def)
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b3CreatePrismaticJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateRevoluteJoint(def RevoluteJointDef) (r RevoluteJoint) {
	defer runtime.KeepAlive(def)
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b3CreateRevoluteJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateSphericalJoint(def SphericalJointDef) (r SphericalJoint) {
	defer runtime.KeepAlive(def)
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b3CreateSphericalJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateWeldJoint(def WeldJointDef) (r WeldJoint) {
	defer runtime.KeepAlive(def)
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b3CreateWeldJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b World) CreateWheelJoint(def WheelJointDef) (r WheelJoint) {
	defer runtime.KeepAlive(def)
	defStack := copyToStack(theStack, def)
	defer defStack.Free()
	r.Id = b3CreateWheelJoint(theStack, b.Id, defStack.Addr)
	return
}
func (b Joint) DestroyJoint(wakeAttached uint8) { b3DestroyJoint(theStack, b.Id, wakeAttached) }
func (b Joint) GetType() (r JointType)          { r = b3Joint_GetType(theStack, b.Id); return }
func (b Joint) GetBodyA() (r Body)              { r.Id = b3Joint_GetBodyA(theStack, b.Id); return }
func (b Joint) GetBodyB() (r Body)              { r.Id = b3Joint_GetBodyB(theStack, b.Id); return }
func (b Joint) GetWorld() (r World)             { r.Id = b3Joint_GetWorld(theStack, b.Id); return }
func (b Joint) SetLocalFrameA(localFrame Transform) {
	b3Joint_SetLocalFrameA(theStack, b.Id, localFrame)
}
func (b Joint) GetLocalFrameA() (r Transform) { r = b3Joint_GetLocalFrameA(theStack, b.Id); return }
func (b Joint) SetLocalFrameB(localFrame Transform) {
	b3Joint_SetLocalFrameB(theStack, b.Id, localFrame)
}
func (b Joint) GetLocalFrameB() (r Transform) { r = b3Joint_GetLocalFrameB(theStack, b.Id); return }
func (b Joint) SetCollideConnected(shouldCollide uint8) {
	b3Joint_SetCollideConnected(theStack, b.Id, shouldCollide)
}
func (b Joint) GetCollideConnected() (r uint8) {
	r = b3Joint_GetCollideConnected(theStack, b.Id)
	return
}
func (b Joint) SetUserData(userData uint64)  { b3Joint_SetUserData(theStack, b.Id, uintptr(userData)) }
func (b Joint) GetUserData() (r uint64)      { r = uint64(b3Joint_GetUserData(theStack, b.Id)); return }
func (b Joint) WakeBodies()                  { b3Joint_WakeBodies(theStack, b.Id) }
func (b Joint) IsAwake() (r uint8)           { r = b3Joint_IsAwake(theStack, b.Id); return }
func (b Joint) GetConstraintForce() (r Vec3) { r = b3Joint_GetConstraintForce(theStack, b.Id); return }
func (b Joint) GetConstraintTorque() (r Vec3) {
	r = b3Joint_GetConstraintTorque(theStack, b.Id)
	return
}
func (b Joint) GetLinearSeparation() (r1 float32) {
	r1 = b3Joint_GetLinearSeparation(theStack, b.Id)
	return
}
func (b Joint) GetAngularSeparation() (r float32) {
	r = b3Joint_GetAngularSeparation(theStack, b.Id)
	return
}
func IsValidFloat(a float32) (r uint8)       { r = b3IsValidFloat(theStack, a); return }
func IsValidVec3(a Vec3) (r uint8)           { r = b3IsValidVec3(theStack, a); return }
func IsValidQuat(a Quat) (r uint8)           { r = b3IsValidQuat(theStack, a); return }
func IsValidTransform(a Transform) (r uint8) { r = b3IsValidTransform(theStack, a); return }
func IsValidMatrix3(a Matrix3) (r uint8)     { r = b3IsValidMatrix3(theStack, a); return }
func IsValidAABB(a AABB) (r uint8)           { r = b3IsValidAABB(theStack, a); return }
func IsBoundedAABB(a AABB) (r uint8)         { r = b3IsBoundedAABB(theStack, a); return }
func IsSaneAABB(a AABB) (r uint8)            { r = b3IsSaneAABB(theStack, a); return }
func IsValidPlane(a3 Plane) (r uint8)        { r = b3IsValidPlane(theStack, a3); return }
func IsValidPosition(p Pos) (r uint8)        { r = b3IsValidPosition(theStack, p); return }
func IsValidWorldTransform(t WorldTransform) (r uint8) {
	r = b3IsValidWorldTransform(theStack, t)
	return
}
func Atan2(y float32, x float32) (r1 float32)   { r1 = b3Atan2(theStack, y, x); return }
func ComputeCosSin(radians1 float32) (r CosSin) { r = b3ComputeCosSin(theStack, radians1); return }
func MakeQuatFromMatrix(m Matrix3) (r Quat) {
	defer runtime.KeepAlive(m)
	mStack := copyToStack(theStack, m)
	defer mStack.Free()
	r = b3MakeQuatFromMatrix(theStack, mStack.Addr)
	return
}
func ComputeQuatBetweenUnitVectors(v1 Vec3, v2 Vec3) (r Quat) {
	r = b3ComputeQuatBetweenUnitVectors(theStack, v1, v2)
	return
}
func LineDistance(p1 Vec3, d1 Vec3, p2 Vec3, d2 Vec3) (r SegmentDistanceResult) {
	r = b3LineDistance(theStack, p1, d1, p2, d2)
	return
}
func SegmentDistance(p1 Vec3, q1 Vec3, p2 Vec3, q2 Vec3) (r1 SegmentDistanceResult) {
	r1 = b3SegmentDistance(theStack, p1, q1, p2, q2)
	return
}
func PointToSegmentDistance(a3 Vec3, b3 Vec3, q Vec3) (r Vec3) {
	r = b3PointToSegmentDistance(theStack, a3, b3, q)
	return
}
func Steiner(mass float32, origin Vec3) (r Matrix3) { r = b3Steiner(theStack, mass, origin); return }
func IsValidRay(input RayCastInput) (r uint8) {
	defer runtime.KeepAlive(input)
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b3IsValidRay(theStack, inputStack.Addr)
	return
}
func GetHeight(mesh *MeshData) (r int32) {
	defer runtime.KeepAlive(mesh)
	escapes(mesh)
	r = b3GetHeight(theStack, uintptr(unsafe.Pointer(mesh)))
	return
}
func CreateGridMesh(xCount int32, zCount int32, cellWidth float32, materialCount int32, identifyEdges uint8) (r OwnedMesh) {
	r = OwnedMesh{(*MeshData)(unsafe.Pointer(b3CreateGridMesh(theStack, xCount, zCount, cellWidth, materialCount, identifyEdges)))}
	return
}
func CreateWaveMesh(xCount int32, zCount int32, cellWidth float32, amplitude float32, rowFrequency float32, columnFrequency float32) (r OwnedMesh) {
	r = OwnedMesh{(*MeshData)(unsafe.Pointer(b3CreateWaveMesh(theStack, xCount, zCount, cellWidth, amplitude, rowFrequency, columnFrequency)))}
	return
}
func CreateTorusMesh(radialResolution int32, tubularResolution int32, radius float32, thickness float32) (r OwnedMesh) {
	r = OwnedMesh{(*MeshData)(unsafe.Pointer(b3CreateTorusMesh(theStack, radialResolution, tubularResolution, radius, thickness)))}
	return
}
func CreateBoxMesh(center Vec3, extent Vec3, identifyEdges uint8) (r OwnedMesh) {
	r = OwnedMesh{(*MeshData)(unsafe.Pointer(b3CreateBoxMesh(theStack, center, extent, identifyEdges)))}
	return
}
func CreateHollowBoxMesh(center Vec3, extent Vec3) (r OwnedMesh) {
	r = OwnedMesh{(*MeshData)(unsafe.Pointer(b3CreateHollowBoxMesh(theStack, center, extent)))}
	return
}
func CreatePlatformMesh(center Vec3, height float32, topWidth float32, bottomWidth float32) (r OwnedMesh) {
	r = OwnedMesh{(*MeshData)(unsafe.Pointer(b3CreatePlatformMesh(theStack, center, height, topWidth, bottomWidth)))}
	return
}
func destroyMesh(mesh *MeshData) {
	defer runtime.KeepAlive(mesh)
	escapes(mesh)
	b3DestroyMesh(theStack, uintptr(unsafe.Pointer(mesh)))
}
func OverlapMesh(shape Mesh, shapeTransform Transform, proxy ShapeProxy) (r uint8) {
	defer runtime.KeepAlive(proxy)
	defer runtime.KeepAlive(shape)
	shapeStack := copyToStackAs[cMesh](theStack, shape)
	defer shapeStack.Free()
	proxyStack := copyToStackAs[cShapeProxy](theStack, proxy)
	defer proxyStack.Free()
	r = b3OverlapMesh(theStack, shapeStack.Addr, shapeTransform, proxyStack.Addr)
	return
}
func ComputeMeshAABB(shape *MeshData, transform1 Transform, scale Vec3) (r AABB) {
	defer runtime.KeepAlive(shape)
	escapes(shape)
	r = b3ComputeMeshAABB(theStack, uintptr(unsafe.Pointer(shape)), transform1, scale)
	return
}
func RayCastMesh(mesh3 Mesh, input RayCastInput) (r CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(mesh3)
	mesh3Stack := copyToStackAs[cMesh](theStack, mesh3)
	defer mesh3Stack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r = b3RayCastMesh(theStack, mesh3Stack.Addr, inputStack.Addr)
	return
}
func ShapeCastMesh(mesh3 Mesh, input ShapeCastInput) (r1 CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(mesh3)
	mesh3Stack := copyToStackAs[cMesh](theStack, mesh3)
	defer mesh3Stack.Free()
	inputStack := copyToStackAs[cShapeCastInput](theStack, input)
	defer inputStack.Free()
	r1 = b3ShapeCastMesh(theStack, mesh3Stack.Addr, inputStack.Addr)
	return
}
func (b MotorJoint) SetLinearVelocity(velocity Vec3) {
	b3MotorJoint_SetLinearVelocity(theStack, b.Id, velocity)
}
func (b MotorJoint) GetLinearVelocity() (r Vec3) {
	r = b3MotorJoint_GetLinearVelocity(theStack, b.Id)
	return
}
func (b MotorJoint) SetAngularVelocity(velocity Vec3) {
	b3MotorJoint_SetAngularVelocity(theStack, b.Id, velocity)
}
func (b MotorJoint) GetAngularVelocity() (r Vec3) {
	r = b3MotorJoint_GetAngularVelocity(theStack, b.Id)
	return
}
func (b MotorJoint) SetMaxVelocityTorque(maxTorque float32) {
	b3MotorJoint_SetMaxVelocityTorque(theStack, b.Id, maxTorque)
}
func (b MotorJoint) GetMaxVelocityTorque() (r float32) {
	r = b3MotorJoint_GetMaxVelocityTorque(theStack, b.Id)
	return
}
func (b MotorJoint) SetMaxVelocityForce(maxForce float32) {
	b3MotorJoint_SetMaxVelocityForce(theStack, b.Id, maxForce)
}
func (b MotorJoint) GetMaxVelocityForce() (r float32) {
	r = b3MotorJoint_GetMaxVelocityForce(theStack, b.Id)
	return
}
func (b MotorJoint) SetLinearHertz(hertz float32) { b3MotorJoint_SetLinearHertz(theStack, b.Id, hertz) }
func (b MotorJoint) GetLinearHertz() (r float32) {
	r = b3MotorJoint_GetLinearHertz(theStack, b.Id)
	return
}
func (b MotorJoint) SetLinearDampingRatio(damping float32) {
	b3MotorJoint_SetLinearDampingRatio(theStack, b.Id, damping)
}
func (b MotorJoint) GetLinearDampingRatio() (r float32) {
	r = b3MotorJoint_GetLinearDampingRatio(theStack, b.Id)
	return
}
func (b MotorJoint) SetAngularHertz(hertz float32) {
	b3MotorJoint_SetAngularHertz(theStack, b.Id, hertz)
}
func (b MotorJoint) GetAngularHertz() (r float32) {
	r = b3MotorJoint_GetAngularHertz(theStack, b.Id)
	return
}
func (b MotorJoint) SetAngularDampingRatio(damping float32) {
	b3MotorJoint_SetAngularDampingRatio(theStack, b.Id, damping)
}
func (b MotorJoint) GetAngularDampingRatio() (r float32) {
	r = b3MotorJoint_GetAngularDampingRatio(theStack, b.Id)
	return
}
func (b MotorJoint) SetMaxSpringForce(maxForce float32) {
	b3MotorJoint_SetMaxSpringForce(theStack, b.Id, maxForce)
}
func (b MotorJoint) GetMaxSpringForce() (r float32) {
	r = b3MotorJoint_GetMaxSpringForce(theStack, b.Id)
	return
}
func (b MotorJoint) SetMaxSpringTorque(maxTorque float32) {
	b3MotorJoint_SetMaxSpringTorque(theStack, b.Id, maxTorque)
}
func (b MotorJoint) GetMaxSpringTorque() (r float32) {
	r = b3MotorJoint_GetMaxSpringTorque(theStack, b.Id)
	return
}
func (b ParallelJoint) SetSpringHertz(hertz float32) {
	b3ParallelJoint_SetSpringHertz(theStack, b.Id, hertz)
}
func (b ParallelJoint) GetSpringHertz() (r float32) {
	r = b3ParallelJoint_GetSpringHertz(theStack, b.Id)
	return
}
func (b ParallelJoint) SetSpringDampingRatio(dampingRatio float32) {
	b3ParallelJoint_SetSpringDampingRatio(theStack, b.Id, dampingRatio)
}
func (b ParallelJoint) GetSpringDampingRatio() (r float32) {
	r = b3ParallelJoint_GetSpringDampingRatio(theStack, b.Id)
	return
}
func (b ParallelJoint) SetMaxTorque(maxForce float32) {
	b3ParallelJoint_SetMaxTorque(theStack, b.Id, maxForce)
}
func (b ParallelJoint) GetMaxTorque() (r float32) {
	r = b3ParallelJoint_GetMaxTorque(theStack, b.Id)
	return
}
func GetWorldCount() (r int32)    { r = b3GetWorldCount(theStack); return }
func GetMaxWorldCount() (r int32) { r = b3GetMaxWorldCount(theStack); return }
func (b World) Step(timeStep float32, subStepCount int32) {
	b3World_Step(theStack, b.Id, timeStep, subStepCount)
}
func (b World) GetBounds() (r AABB)           { r = b3World_GetBounds(theStack, b.Id); return }
func (b World) IsValid() (r uint8)            { r = b3World_IsValid(theStack, b.Id); return }
func (b Body) IsValid() (r uint8)             { r = b3Body_IsValid(theStack, b.Id); return }
func (b Shape) IsValid() (r uint8)            { r = b3Shape_IsValid(theStack, b.Id); return }
func (b Joint) IsValid() (r uint8)            { r = b3Joint_IsValid(theStack, b.Id); return }
func (b Contact) IsValid() (r uint8)          { r = b3Contact_IsValid(theStack, b.Id); return }
func (b World) EnableSleeping(flag uint8)     { b3World_EnableSleeping(theStack, b.Id, flag) }
func (b World) IsSleepingEnabled() (r uint8)  { r = b3World_IsSleepingEnabled(theStack, b.Id); return }
func (b World) EnableWarmStarting(flag uint8) { b3World_EnableWarmStarting(theStack, b.Id, flag) }
func (b World) IsWarmStartingEnabled() (r uint8) {
	r = b3World_IsWarmStartingEnabled(theStack, b.Id)
	return
}
func (b World) EnableSSE2Fallback(flag uint8) { b3World_EnableSSE2Fallback(theStack, b.Id, flag) }
func (b World) GetAwakeBodyCount() (r int32)  { r = b3World_GetAwakeBodyCount(theStack, b.Id); return }
func (b World) EnableContinuous(flag uint8)   { b3World_EnableContinuous(theStack, b.Id, flag) }
func (b World) IsContinuousEnabled() (r uint8) {
	r = b3World_IsContinuousEnabled(theStack, b.Id)
	return
}
func (b World) SetRestitutionThreshold(value float32) {
	b3World_SetRestitutionThreshold(theStack, b.Id, value)
}
func (b World) GetRestitutionThreshold() (r float32) {
	r = b3World_GetRestitutionThreshold(theStack, b.Id)
	return
}
func (b World) SetRestitutionIterations(iterations int32) {
	b3World_SetRestitutionIterations(theStack, b.Id, iterations)
}
func (b World) GetRestitutionIterations() (r int32) {
	r = b3World_GetRestitutionIterations(theStack, b.Id)
	return
}
func (b World) EnableRestitutionPropagation(flag uint8) {
	b3World_EnableRestitutionPropagation(theStack, b.Id, flag)
}
func (b World) IsRestitutionPropagationEnabled() (r uint8) {
	r = b3World_IsRestitutionPropagationEnabled(theStack, b.Id)
	return
}
func (b World) SetHitEventThreshold(value float32) {
	b3World_SetHitEventThreshold(theStack, b.Id, value)
}
func (b World) GetHitEventThreshold() (r float32) {
	r = b3World_GetHitEventThreshold(theStack, b.Id)
	return
}
func (b World) SetContactTuning(hertz float32, dampingRatio float32, contactSpeed float32) {
	b3World_SetContactTuning(theStack, b.Id, hertz, dampingRatio, contactSpeed)
}
func (b World) SetContactRecycleDistance(recycleDistance float32) {
	b3World_SetContactRecycleDistance(theStack, b.Id, recycleDistance)
}
func (b World) GetContactRecycleDistance() (r float32) {
	r = b3World_GetContactRecycleDistance(theStack, b.Id)
	return
}
func (b World) SetMaximumLinearSpeed(maximumLinearSpeed float32) {
	b3World_SetMaximumLinearSpeed(theStack, b.Id, maximumLinearSpeed)
}
func (b World) GetMaximumLinearSpeed() (r float32) {
	r = b3World_GetMaximumLinearSpeed(theStack, b.Id)
	return
}
func (b World) SetIsotropicInertiaTolerance(tolerance float32) {
	b3World_SetIsotropicInertiaTolerance(theStack, b.Id, tolerance)
}
func (b World) GetIsotropicInertiaTolerance() (r float32) {
	r = b3World_GetIsotropicInertiaTolerance(theStack, b.Id)
	return
}
func (b World) GetProfile() (r Profile)      { r = b3World_GetProfile(theStack, b.Id); return }
func (b World) GetMaxCapacity() (r Capacity) { r = b3World_GetMaxCapacity(theStack, b.Id); return }
func (b World) SetUserData(userData uint64)  { b3World_SetUserData(theStack, b.Id, uintptr(userData)) }
func (b World) GetUserData() (r uint64)      { r = uint64(b3World_GetUserData(theStack, b.Id)); return }
func (b World) SetWorkerCount(count int32)   { b3World_SetWorkerCount(theStack, b.Id, count) }
func (b World) GetWorkerCount() (r int32)    { r = b3World_GetWorkerCount(theStack, b.Id); return }
func (b World) GetCounters() (r Counters)    { r = b3World_GetCounters(theStack, b.Id); return }
func (b World) CastRayClosest(origin Pos, translation Vec3, filter QueryFilter) (r RayResult) {
	r = b3World_CastRayClosest(theStack, b.Id, origin, translation, filter)
	return
}
func (b World) SetGravity(gravity Vec3) { b3World_SetGravity(theStack, b.Id, gravity) }
func (b World) GetGravity() (r Vec3)    { r = b3World_GetGravity(theStack, b.Id); return }
func (b World) Explode(explosionDef ExplosionDef) {
	defer runtime.KeepAlive(explosionDef)
	explosionDefStack := copyToStack(theStack, explosionDef)
	defer explosionDefStack.Free()
	b3World_Explode(theStack, b.Id, explosionDefStack.Addr)
}
func (b World) EnableSpeculative(flag uint8) { b3World_EnableSpeculative(theStack, b.Id, flag) }
func (b PrismaticJoint) EnableLimit(enableLimit uint8) {
	b3PrismaticJoint_EnableLimit(theStack, b.Id, enableLimit)
}
func (b PrismaticJoint) IsLimitEnabled() (r uint8) {
	r = b3PrismaticJoint_IsLimitEnabled(theStack, b.Id)
	return
}
func (b PrismaticJoint) GetLowerLimit() (r float32) {
	r = b3PrismaticJoint_GetLowerLimit(theStack, b.Id)
	return
}
func (b PrismaticJoint) GetUpperLimit() (r float32) {
	r = b3PrismaticJoint_GetUpperLimit(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetLimits(lower float32, upper float32) {
	b3PrismaticJoint_SetLimits(theStack, b.Id, lower, upper)
}
func (b PrismaticJoint) GetTranslation() (r float32) {
	r = b3PrismaticJoint_GetTranslation(theStack, b.Id)
	return
}
func (b PrismaticJoint) EnableSpring(enableSpring uint8) {
	b3PrismaticJoint_EnableSpring(theStack, b.Id, enableSpring)
}
func (b PrismaticJoint) IsSpringEnabled() (r uint8) {
	r = b3PrismaticJoint_IsSpringEnabled(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetTargetTranslation(targetTranslation float32) {
	b3PrismaticJoint_SetTargetTranslation(theStack, b.Id, targetTranslation)
}
func (b PrismaticJoint) GetTargetTranslation() (r float32) {
	r = b3PrismaticJoint_GetTargetTranslation(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetSpringHertz(hertz float32) {
	b3PrismaticJoint_SetSpringHertz(theStack, b.Id, hertz)
}
func (b PrismaticJoint) GetSpringHertz() (r float32) {
	r = b3PrismaticJoint_GetSpringHertz(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetSpringDampingRatio(dampingRatio float32) {
	b3PrismaticJoint_SetSpringDampingRatio(theStack, b.Id, dampingRatio)
}
func (b PrismaticJoint) GetSpringDampingRatio() (r float32) {
	r = b3PrismaticJoint_GetSpringDampingRatio(theStack, b.Id)
	return
}
func (b PrismaticJoint) EnableMotor(enableMotor uint8) {
	b3PrismaticJoint_EnableMotor(theStack, b.Id, enableMotor)
}
func (b PrismaticJoint) IsMotorEnabled() (r uint8) {
	r = b3PrismaticJoint_IsMotorEnabled(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetMotorSpeed(motorSpeed float32) {
	b3PrismaticJoint_SetMotorSpeed(theStack, b.Id, motorSpeed)
}
func (b PrismaticJoint) GetMotorSpeed() (r float32) {
	r = b3PrismaticJoint_GetMotorSpeed(theStack, b.Id)
	return
}
func (b PrismaticJoint) SetMaxMotorForce(maxForce float32) {
	b3PrismaticJoint_SetMaxMotorForce(theStack, b.Id, maxForce)
}
func (b PrismaticJoint) GetMaxMotorForce() (r float32) {
	r = b3PrismaticJoint_GetMaxMotorForce(theStack, b.Id)
	return
}
func (b PrismaticJoint) GetMotorForce() (r float32) {
	r = b3PrismaticJoint_GetMotorForce(theStack, b.Id)
	return
}
func (b PrismaticJoint) GetSpeed() (r float32) { r = b3PrismaticJoint_GetSpeed(theStack, b.Id); return }
func (b RevoluteJoint) EnableLimit(enableLimit uint8) {
	b3RevoluteJoint_EnableLimit(theStack, b.Id, enableLimit)
}
func (b RevoluteJoint) IsLimitEnabled() (r uint8) {
	r = b3RevoluteJoint_IsLimitEnabled(theStack, b.Id)
	return
}
func (b RevoluteJoint) GetLowerLimit() (r float32) {
	r = b3RevoluteJoint_GetLowerLimit(theStack, b.Id)
	return
}
func (b RevoluteJoint) GetUpperLimit() (r float32) {
	r = b3RevoluteJoint_GetUpperLimit(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetLimits(lowerLimitRadians float32, upperLimitRadians float32) {
	b3RevoluteJoint_SetLimits(theStack, b.Id, lowerLimitRadians, upperLimitRadians)
}
func (b RevoluteJoint) GetAngle() (r float32) { r = b3RevoluteJoint_GetAngle(theStack, b.Id); return }
func (b RevoluteJoint) EnableSpring(enableSpring uint8) {
	b3RevoluteJoint_EnableSpring(theStack, b.Id, enableSpring)
}
func (b RevoluteJoint) IsSpringEnabled() (r uint8) {
	r = b3RevoluteJoint_IsSpringEnabled(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetTargetAngle(targetRadians float32) {
	b3RevoluteJoint_SetTargetAngle(theStack, b.Id, targetRadians)
}
func (b RevoluteJoint) GetTargetAngle() (r float32) {
	r = b3RevoluteJoint_GetTargetAngle(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetSpringHertz(hertz float32) {
	b3RevoluteJoint_SetSpringHertz(theStack, b.Id, hertz)
}
func (b RevoluteJoint) GetSpringHertz() (r float32) {
	r = b3RevoluteJoint_GetSpringHertz(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetSpringDampingRatio(dampingRatio float32) {
	b3RevoluteJoint_SetSpringDampingRatio(theStack, b.Id, dampingRatio)
}
func (b RevoluteJoint) GetSpringDampingRatio() (r float32) {
	r = b3RevoluteJoint_GetSpringDampingRatio(theStack, b.Id)
	return
}
func (b RevoluteJoint) EnableMotor(enableMotor uint8) {
	b3RevoluteJoint_EnableMotor(theStack, b.Id, enableMotor)
}
func (b RevoluteJoint) IsMotorEnabled() (r uint8) {
	r = b3RevoluteJoint_IsMotorEnabled(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetMotorSpeed(motorSpeed float32) {
	b3RevoluteJoint_SetMotorSpeed(theStack, b.Id, motorSpeed)
}
func (b RevoluteJoint) GetMotorSpeed() (r float32) {
	r = b3RevoluteJoint_GetMotorSpeed(theStack, b.Id)
	return
}
func (b RevoluteJoint) SetMaxMotorTorque(maxForce float32) {
	b3RevoluteJoint_SetMaxMotorTorque(theStack, b.Id, maxForce)
}
func (b RevoluteJoint) GetMaxMotorTorque() (r float32) {
	r = b3RevoluteJoint_GetMaxMotorTorque(theStack, b.Id)
	return
}
func (b RevoluteJoint) GetMotorTorque() (r float32) {
	r = b3RevoluteJoint_GetMotorTorque(theStack, b.Id)
	return
}
func (b Body) CreateSphereShape(def ShapeDef, sphere Sphere) (r Shape) {
	defer runtime.KeepAlive(sphere)
	defer runtime.KeepAlive(def)
	defStack := copyToStackAs[cShapeDef](theStack, def)
	defer defStack.Free()
	sphereStack := copyToStack(theStack, sphere)
	defer sphereStack.Free()
	r.Id = b3CreateSphereShape(theStack, b.Id, defStack.Addr, sphereStack.Addr)
	return
}
func (b Body) CreateCapsuleShape(def ShapeDef, capsule Capsule) (r Shape) {
	defer runtime.KeepAlive(capsule)
	defer runtime.KeepAlive(def)
	defStack := copyToStackAs[cShapeDef](theStack, def)
	defer defStack.Free()
	capsuleStack := copyToStack(theStack, capsule)
	defer capsuleStack.Free()
	r.Id = b3CreateCapsuleShape(theStack, b.Id, defStack.Addr, capsuleStack.Addr)
	return
}
func (b Body) CreateHullShape(def ShapeDef, hull *HullData) (r Shape) {
	defer runtime.KeepAlive(hull)
	defer runtime.KeepAlive(def)
	defStack := copyToStackAs[cShapeDef](theStack, def)
	defer defStack.Free()
	escapes(hull)
	r.Id = b3CreateHullShape(theStack, b.Id, defStack.Addr, uintptr(unsafe.Pointer(hull)))
	return
}
func (b Body) CreateTransformedHullShape(def ShapeDef, hull *HullData, transform Transform, scale Vec3) (r Shape) {
	defer runtime.KeepAlive(hull)
	defer runtime.KeepAlive(def)
	defStack := copyToStackAs[cShapeDef](theStack, def)
	defer defStack.Free()
	escapes(hull)
	r.Id = b3CreateTransformedHullShape(theStack, b.Id, defStack.Addr, uintptr(unsafe.Pointer(hull)), transform, scale)
	return
}
func (b Body) CreateMeshShape(def ShapeDef, mesh *MeshData, scale Vec3) (r Shape) {
	defer runtime.KeepAlive(mesh)
	defer runtime.KeepAlive(def)
	defStack := copyToStackAs[cShapeDef](theStack, def)
	defer defStack.Free()
	escapes(mesh)
	r.Id = b3CreateMeshShape(theStack, b.Id, defStack.Addr, uintptr(unsafe.Pointer(mesh)), scale)
	return
}
func (b Body) CreateHeightFieldShape(def ShapeDef, heightField *HeightFieldData) (r Shape) {
	defer runtime.KeepAlive(heightField)
	defer runtime.KeepAlive(def)
	defStack := copyToStackAs[cShapeDef](theStack, def)
	defer defStack.Free()
	escapes(heightField)
	r.Id = b3CreateHeightFieldShape(theStack, b.Id, defStack.Addr, uintptr(unsafe.Pointer(heightField)))
	return
}
func (b Body) CreateBakedCompoundShape(def ShapeDef, compound *CompoundData) (r Shape) {
	defer runtime.KeepAlive(compound)
	defer runtime.KeepAlive(def)
	defStack := copyToStackAs[cShapeDef](theStack, def)
	defer defStack.Free()
	escapes(compound)
	r.Id = b3CreateBakedCompoundShape(theStack, b.Id, defStack.Addr, uintptr(unsafe.Pointer(compound)))
	return
}
func (b Shape) DestroyShape(updateBodyMass uint8) { b3DestroyShape(theStack, b.Id, updateBodyMass) }
func (b Shape) GetBody() (r Body)                 { r.Id = b3Shape_GetBody(theStack, b.Id); return }
func (b Shape) GetWorld() (r World)               { r.Id = b3Shape_GetWorld(theStack, b.Id); return }
func (b Shape) SetUserData(userData uint64)       { b3Shape_SetUserData(theStack, b.Id, uintptr(userData)) }
func (b Shape) GetUserData() (r uint64)           { r = uint64(b3Shape_GetUserData(theStack, b.Id)); return }
func (b Shape) SetName(name string) {
	nameStack := theStack.toCString(name)
	defer nameStack.Free()
	b3Shape_SetName(theStack, b.Id, nameStack.Addr)
}
func (b Shape) GetName() (r string) { r = toString(b3Shape_GetName(theStack, b.Id)); return }
func (b Shape) IsSensor() (r uint8) { r = b3Shape_IsSensor(theStack, b.Id); return }
func (b Shape) RayCast(origin Pos, translation Vec3) (r1 WorldCastOutput) {
	r1 = b3Shape_RayCast(theStack, b.Id, origin, translation)
	return
}
func (b Shape) SetDensity(density float32, updateBodyMass uint8) {
	b3Shape_SetDensity(theStack, b.Id, density, updateBodyMass)
}
func (b Shape) GetDensity() (r float32)      { r = b3Shape_GetDensity(theStack, b.Id); return }
func (b Shape) SetFriction(friction float32) { b3Shape_SetFriction(theStack, b.Id, friction) }
func (b Shape) GetFriction() (r float32)     { r = b3Shape_GetFriction(theStack, b.Id); return }
func (b Shape) SetRestitution(restitution float32) {
	b3Shape_SetRestitution(theStack, b.Id, restitution)
}
func (b Shape) GetRestitution() (r float32) { r = b3Shape_GetRestitution(theStack, b.Id); return }
func (b Shape) SetSurfaceMaterial(surfaceMaterial SurfaceMaterial) {
	b3Shape_SetSurfaceMaterial(theStack, b.Id, surfaceMaterial)
}
func (b Shape) GetSurfaceMaterial() (r SurfaceMaterial) {
	r = b3Shape_GetSurfaceMaterial(theStack, b.Id)
	return
}
func (b Shape) GetMeshMaterialCount() (r int32) {
	r = b3Shape_GetMeshMaterialCount(theStack, b.Id)
	return
}
func (b Shape) SetMeshMaterial(surfaceMaterial SurfaceMaterial, index int32) {
	b3Shape_SetMeshMaterial(theStack, b.Id, surfaceMaterial, index)
}
func (b Shape) GetMeshSurfaceMaterial(index int32) (r SurfaceMaterial) {
	r = b3Shape_GetMeshSurfaceMaterial(theStack, b.Id, index)
	return
}
func (b Shape) GetFilter() (r Filter) { r = b3Shape_GetFilter(theStack, b.Id); return }
func (b Shape) SetFilter(filter Filter, invokeContacts uint8) {
	b3Shape_SetFilter(theStack, b.Id, filter, invokeContacts)
}
func (b Shape) EnableSensorEvents(flag uint8) { b3Shape_EnableSensorEvents(theStack, b.Id, flag) }
func (b Shape) AreSensorEventsEnabled() (r uint8) {
	r = b3Shape_AreSensorEventsEnabled(theStack, b.Id)
	return
}
func (b Shape) EnableContactEvents(flag uint8) { b3Shape_EnableContactEvents(theStack, b.Id, flag) }
func (b Shape) AreContactEventsEnabled() (r uint8) {
	r = b3Shape_AreContactEventsEnabled(theStack, b.Id)
	return
}
func (b Shape) EnablePreSolveEvents(flag uint8) { b3Shape_EnablePreSolveEvents(theStack, b.Id, flag) }
func (b Shape) ArePreSolveEventsEnabled() (r uint8) {
	r = b3Shape_ArePreSolveEventsEnabled(theStack, b.Id)
	return
}
func (b Shape) EnableHitEvents(flag uint8) { b3Shape_EnableHitEvents(theStack, b.Id, flag) }
func (b Shape) AreHitEventsEnabled() (r uint8) {
	r = b3Shape_AreHitEventsEnabled(theStack, b.Id)
	return
}
func (b Shape) GetType() (r ShapeType)  { r = b3Shape_GetType(theStack, b.Id); return }
func (b Shape) GetSphere() (r Sphere)   { r = b3Shape_GetSphere(theStack, b.Id); return }
func (b Shape) GetCapsule() (r Capsule) { r = b3Shape_GetCapsule(theStack, b.Id); return }
func (b Shape) GetHull() (r *HullData) {
	r = (*HullData)(unsafe.Pointer(b3Shape_GetHull(theStack, b.Id)))
	return
}
func (b Shape) GetMesh() (r Mesh) { r = b3Shape_GetMesh(theStack, b.Id); return }
func (b Shape) GetHeightField() (r *HeightFieldData) {
	r = (*HeightFieldData)(unsafe.Pointer(b3Shape_GetHeightField(theStack, b.Id)))
	return
}
func (b Shape) GetCompound() (r *CompoundData) {
	r = (*CompoundData)(unsafe.Pointer(b3Shape_GetCompound(theStack, b.Id)))
	return
}
func (b Shape) SetSphere(sphere Sphere) {
	defer runtime.KeepAlive(sphere)
	sphereStack := copyToStack(theStack, sphere)
	defer sphereStack.Free()
	b3Shape_SetSphere(theStack, b.Id, sphereStack.Addr)
}
func (b Shape) SetCapsule(capsule Capsule) {
	defer runtime.KeepAlive(capsule)
	capsuleStack := copyToStack(theStack, capsule)
	defer capsuleStack.Free()
	b3Shape_SetCapsule(theStack, b.Id, capsuleStack.Addr)
}
func (b Shape) SetHull(hull *HullData) {
	defer runtime.KeepAlive(hull)
	escapes(hull)
	b3Shape_SetHull(theStack, b.Id, uintptr(unsafe.Pointer(hull)))
}
func (b Shape) SetMesh(meshData *MeshData, scale Vec3) {
	defer runtime.KeepAlive(meshData)
	escapes(meshData)
	b3Shape_SetMesh(theStack, b.Id, uintptr(unsafe.Pointer(meshData)), scale)
}
func (b Shape) GetContactCapacity() (r int32) { r = b3Shape_GetContactCapacity(theStack, b.Id); return }
func (b Shape) GetSensorCapacity() (r int32)  { r = b3Shape_GetSensorCapacity(theStack, b.Id); return }
func (b Shape) GetAABB() (r AABB)             { r = b3Shape_GetAABB(theStack, b.Id); return }
func (b Shape) ComputeMassData() (r MassData) { r = b3Shape_ComputeMassData(theStack, b.Id); return }
func (b Shape) GetClosestPoint(_target Vec3) (r1 Vec3) {
	r1 = b3Shape_GetClosestPoint(theStack, b.Id, _target)
	return
}
func (b Shape) ApplyWind(wind Vec3, drag float32, lift float32, maxSpeed float32, wake uint8) {
	b3Shape_ApplyWind(theStack, b.Id, wind, drag, lift, maxSpeed, wake)
}
func IsAVX2Available() (r uint8) { r = b3IsAVX2Available(theStack); return }
func ComputeSphereMass(shape Sphere, density float32) (r MassData) {
	defer runtime.KeepAlive(shape)
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r = b3ComputeSphereMass(theStack, shapeStack.Addr, density)
	return
}
func ComputeSphereAABB(shape Sphere, transform Transform) (r AABB) {
	defer runtime.KeepAlive(shape)
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	r = b3ComputeSphereAABB(theStack, shapeStack.Addr, transform)
	return
}
func OverlapSphere(shape Sphere, shapeTransform Transform, proxy ShapeProxy) (r uint8) {
	defer runtime.KeepAlive(proxy)
	defer runtime.KeepAlive(shape)
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	proxyStack := copyToStackAs[cShapeProxy](theStack, proxy)
	defer proxyStack.Free()
	r = b3OverlapSphere(theStack, shapeStack.Addr, shapeTransform, proxyStack.Addr)
	return
}
func RayCastSphere(shape Sphere, input RayCastInput) (r1 CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(shape)
	shapeStack := copyToStack(theStack, shape)
	defer shapeStack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r1 = b3RayCastSphere(theStack, shapeStack.Addr, inputStack.Addr)
	return
}
func RayCastHollowSphere(sphere Sphere, input RayCastInput) (r1 CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(sphere)
	sphereStack := copyToStack(theStack, sphere)
	defer sphereStack.Free()
	inputStack := copyToStack(theStack, input)
	defer inputStack.Free()
	r1 = b3RayCastHollowSphere(theStack, sphereStack.Addr, inputStack.Addr)
	return
}
func ShapeCastSphere(sphere Sphere, input ShapeCastInput) (r CastOutput) {
	defer runtime.KeepAlive(input)
	defer runtime.KeepAlive(sphere)
	sphereStack := copyToStack(theStack, sphere)
	defer sphereStack.Free()
	inputStack := copyToStackAs[cShapeCastInput](theStack, input)
	defer inputStack.Free()
	r = b3ShapeCastSphere(theStack, sphereStack.Addr, inputStack.Addr)
	return
}
func (b SphericalJoint) EnableConeLimit(enableLimit uint8) {
	b3SphericalJoint_EnableConeLimit(theStack, b.Id, enableLimit)
}
func (b SphericalJoint) IsConeLimitEnabled() (r uint8) {
	r = b3SphericalJoint_IsConeLimitEnabled(theStack, b.Id)
	return
}
func (b SphericalJoint) GetConeLimit() (r float32) {
	r = b3SphericalJoint_GetConeLimit(theStack, b.Id)
	return
}
func (b SphericalJoint) SetConeLimit(angleRadians float32) {
	b3SphericalJoint_SetConeLimit(theStack, b.Id, angleRadians)
}
func (b SphericalJoint) GetConeAngle() (r float32) {
	r = b3SphericalJoint_GetConeAngle(theStack, b.Id)
	return
}
func (b SphericalJoint) EnableTwistLimit(enableLimit uint8) {
	b3SphericalJoint_EnableTwistLimit(theStack, b.Id, enableLimit)
}
func (b SphericalJoint) IsTwistLimitEnabled() (r uint8) {
	r = b3SphericalJoint_IsTwistLimitEnabled(theStack, b.Id)
	return
}
func (b SphericalJoint) GetLowerTwistLimit() (r float32) {
	r = b3SphericalJoint_GetLowerTwistLimit(theStack, b.Id)
	return
}
func (b SphericalJoint) GetUpperTwistLimit() (r float32) {
	r = b3SphericalJoint_GetUpperTwistLimit(theStack, b.Id)
	return
}
func (b SphericalJoint) SetTwistLimits(lowerLimitRadians float32, upperLimitRadians float32) {
	b3SphericalJoint_SetTwistLimits(theStack, b.Id, lowerLimitRadians, upperLimitRadians)
}
func (b SphericalJoint) GetTwistAngle() (r float32) {
	r = b3SphericalJoint_GetTwistAngle(theStack, b.Id)
	return
}
func (b SphericalJoint) EnableSpring(enableSpring uint8) {
	b3SphericalJoint_EnableSpring(theStack, b.Id, enableSpring)
}
func (b SphericalJoint) IsSpringEnabled() (r uint8) {
	r = b3SphericalJoint_IsSpringEnabled(theStack, b.Id)
	return
}
func (b SphericalJoint) SetTargetRotation(targetRotation Quat) {
	b3SphericalJoint_SetTargetRotation(theStack, b.Id, targetRotation)
}
func (b SphericalJoint) GetTargetRotation() (r Quat) {
	r = b3SphericalJoint_GetTargetRotation(theStack, b.Id)
	return
}
func (b SphericalJoint) SetSpringHertz(hertz float32) {
	b3SphericalJoint_SetSpringHertz(theStack, b.Id, hertz)
}
func (b SphericalJoint) GetSpringHertz() (r float32) {
	r = b3SphericalJoint_GetSpringHertz(theStack, b.Id)
	return
}
func (b SphericalJoint) SetSpringDampingRatio(dampingRatio float32) {
	b3SphericalJoint_SetSpringDampingRatio(theStack, b.Id, dampingRatio)
}
func (b SphericalJoint) GetSpringDampingRatio() (r float32) {
	r = b3SphericalJoint_GetSpringDampingRatio(theStack, b.Id)
	return
}
func (b SphericalJoint) EnableMotor(enableMotor uint8) {
	b3SphericalJoint_EnableMotor(theStack, b.Id, enableMotor)
}
func (b SphericalJoint) IsMotorEnabled() (r uint8) {
	r = b3SphericalJoint_IsMotorEnabled(theStack, b.Id)
	return
}
func (b SphericalJoint) SetMotorVelocity(motorVelocity Vec3) {
	b3SphericalJoint_SetMotorVelocity(theStack, b.Id, motorVelocity)
}
func (b SphericalJoint) GetMotorVelocity() (r Vec3) {
	r = b3SphericalJoint_GetMotorVelocity(theStack, b.Id)
	return
}
func (b SphericalJoint) SetMaxMotorTorque(maxForce float32) {
	b3SphericalJoint_SetMaxMotorTorque(theStack, b.Id, maxForce)
}
func (b SphericalJoint) GetMaxMotorTorque() (r float32) {
	r = b3SphericalJoint_GetMaxMotorTorque(theStack, b.Id)
	return
}
func (b SphericalJoint) GetMotorTorque() (r Vec3) {
	r = b3SphericalJoint_GetMotorTorque(theStack, b.Id)
	return
}
func GetTicks() (r uint64)                        { r = b3GetTicks(theStack); return }
func GetMilliseconds(ticks uint64) (r float32)    { r = b3GetMilliseconds(theStack, ticks); return }
func Yield()                                      { b3Yield(theStack) }
func Sleep(milliseconds int32)                    { b3Sleep(theStack, milliseconds) }
func DefaultWorldDef() (r WorldDef)               { r = b3DefaultWorldDef(theStack); return }
func DefaultBodyDef() (r BodyDef)                 { r = b3DefaultBodyDef(theStack); return }
func DefaultFilter() (r Filter)                   { r = b3DefaultFilter(theStack); return }
func DefaultQueryFilter() (r QueryFilter)         { r = b3DefaultQueryFilter(theStack); return }
func DefaultSurfaceMaterial() (r SurfaceMaterial) { r = b3DefaultSurfaceMaterial(theStack); return }
func DefaultShapeDef() (r ShapeDef)               { r = b3DefaultShapeDef(theStack); return }
func (b WeldJoint) SetLinearHertz(hertz float32)  { b3WeldJoint_SetLinearHertz(theStack, b.Id, hertz) }
func (b WeldJoint) GetLinearHertz() (r float32) {
	r = b3WeldJoint_GetLinearHertz(theStack, b.Id)
	return
}
func (b WeldJoint) SetLinearDampingRatio(dampingRatio float32) {
	b3WeldJoint_SetLinearDampingRatio(theStack, b.Id, dampingRatio)
}
func (b WeldJoint) GetLinearDampingRatio() (r float32) {
	r = b3WeldJoint_GetLinearDampingRatio(theStack, b.Id)
	return
}
func (b WeldJoint) SetAngularHertz(hertz float32) { b3WeldJoint_SetAngularHertz(theStack, b.Id, hertz) }
func (b WeldJoint) GetAngularHertz() (r float32) {
	r = b3WeldJoint_GetAngularHertz(theStack, b.Id)
	return
}
func (b WeldJoint) SetAngularDampingRatio(dampingRatio float32) {
	b3WeldJoint_SetAngularDampingRatio(theStack, b.Id, dampingRatio)
}
func (b WeldJoint) GetAngularDampingRatio() (r float32) {
	r = b3WeldJoint_GetAngularDampingRatio(theStack, b.Id)
	return
}
func (b WheelJoint) EnableSuspension(enableSpring uint8) {
	b3WheelJoint_EnableSuspension(theStack, b.Id, enableSpring)
}
func (b WheelJoint) IsSuspensionEnabled() (r uint8) {
	r = b3WheelJoint_IsSuspensionEnabled(theStack, b.Id)
	return
}
func (b WheelJoint) SetSuspensionHertz(hertz float32) {
	b3WheelJoint_SetSuspensionHertz(theStack, b.Id, hertz)
}
func (b WheelJoint) GetSuspensionHertz() (r float32) {
	r = b3WheelJoint_GetSuspensionHertz(theStack, b.Id)
	return
}
func (b WheelJoint) SetSuspensionDampingRatio(dampingRatio float32) {
	b3WheelJoint_SetSuspensionDampingRatio(theStack, b.Id, dampingRatio)
}
func (b WheelJoint) GetSuspensionDampingRatio() (r float32) {
	r = b3WheelJoint_GetSuspensionDampingRatio(theStack, b.Id)
	return
}
func (b WheelJoint) EnableSuspensionLimit(enableLimit uint8) {
	b3WheelJoint_EnableSuspensionLimit(theStack, b.Id, enableLimit)
}
func (b WheelJoint) IsSuspensionLimitEnabled() (r uint8) {
	r = b3WheelJoint_IsSuspensionLimitEnabled(theStack, b.Id)
	return
}
func (b WheelJoint) GetLowerSuspensionLimit() (r float32) {
	r = b3WheelJoint_GetLowerSuspensionLimit(theStack, b.Id)
	return
}
func (b WheelJoint) GetUpperSuspensionLimit() (r float32) {
	r = b3WheelJoint_GetUpperSuspensionLimit(theStack, b.Id)
	return
}
func (b WheelJoint) SetSuspensionLimits(lower float32, upper float32) {
	b3WheelJoint_SetSuspensionLimits(theStack, b.Id, lower, upper)
}
func (b WheelJoint) EnableSpinMotor(enableMotor uint8) {
	b3WheelJoint_EnableSpinMotor(theStack, b.Id, enableMotor)
}
func (b WheelJoint) IsSpinMotorEnabled() (r uint8) {
	r = b3WheelJoint_IsSpinMotorEnabled(theStack, b.Id)
	return
}
func (b WheelJoint) SetSpinMotorSpeed(motorSpeed float32) {
	b3WheelJoint_SetSpinMotorSpeed(theStack, b.Id, motorSpeed)
}
func (b WheelJoint) GetSpinMotorSpeed() (r float32) {
	r = b3WheelJoint_GetSpinMotorSpeed(theStack, b.Id)
	return
}
func (b WheelJoint) SetMaxSpinTorque(torque float32) {
	b3WheelJoint_SetMaxSpinTorque(theStack, b.Id, torque)
}
func (b WheelJoint) GetMaxSpinTorque() (r float32) {
	r = b3WheelJoint_GetMaxSpinTorque(theStack, b.Id)
	return
}
func (b WheelJoint) EnableSteering(flag uint8) { b3WheelJoint_EnableSteering(theStack, b.Id, flag) }
func (b WheelJoint) IsSteeringEnabled() (r uint8) {
	r = b3WheelJoint_IsSteeringEnabled(theStack, b.Id)
	return
}
func (b WheelJoint) SetSteeringHertz(hertz float32) {
	b3WheelJoint_SetSteeringHertz(theStack, b.Id, hertz)
}
func (b WheelJoint) GetSteeringHertz() (r float32) {
	r = b3WheelJoint_GetSteeringHertz(theStack, b.Id)
	return
}
func (b WheelJoint) SetSteeringDampingRatio(dampingRatio float32) {
	b3WheelJoint_SetSteeringDampingRatio(theStack, b.Id, dampingRatio)
}
func (b WheelJoint) GetSteeringDampingRatio() (r float32) {
	r = b3WheelJoint_GetSteeringDampingRatio(theStack, b.Id)
	return
}
func (b WheelJoint) SetMaxSteeringTorque(maxTorque float32) {
	b3WheelJoint_SetMaxSteeringTorque(theStack, b.Id, maxTorque)
}
func (b WheelJoint) GetMaxSteeringTorque() (r float32) {
	r = b3WheelJoint_GetMaxSteeringTorque(theStack, b.Id)
	return
}
func (b WheelJoint) EnableSteeringLimit(flag uint8) {
	b3WheelJoint_EnableSteeringLimit(theStack, b.Id, flag)
}
func (b WheelJoint) IsSteeringLimitEnabled() (r uint8) {
	r = b3WheelJoint_IsSteeringLimitEnabled(theStack, b.Id)
	return
}
func (b WheelJoint) GetLowerSteeringLimit() (r float32) {
	r = b3WheelJoint_GetLowerSteeringLimit(theStack, b.Id)
	return
}
func (b WheelJoint) GetUpperSteeringLimit() (r float32) {
	r = b3WheelJoint_GetUpperSteeringLimit(theStack, b.Id)
	return
}
func (b WheelJoint) SetSteeringLimits(lowerRadians float32, upperRadians float32) {
	b3WheelJoint_SetSteeringLimits(theStack, b.Id, lowerRadians, upperRadians)
}
func (b WheelJoint) SetTargetSteeringAngle(radians float32) {
	b3WheelJoint_SetTargetSteeringAngle(theStack, b.Id, radians)
}
func (b WheelJoint) GetTargetSteeringAngle() (r float32) {
	r = b3WheelJoint_GetTargetSteeringAngle(theStack, b.Id)
	return
}
func (b WheelJoint) GetSpinSpeed() (r float32) { r = b3WheelJoint_GetSpinSpeed(theStack, b.Id); return }
func (b WheelJoint) GetSpinTorque() (r float32) {
	r = b3WheelJoint_GetSpinTorque(theStack, b.Id)
	return
}
func (b WheelJoint) GetSteeringAngle() (r float32) {
	r = b3WheelJoint_GetSteeringAngle(theStack, b.Id)
	return
}
func (b WheelJoint) GetSteeringTorque() (r float32) {
	r = b3WheelJoint_GetSteeringTorque(theStack, b.Id)
	return
}
func MinInt(a1 int32, b1 int32) (r int32) { r = b3go_inline_b3MinInt(theStack, a1, b1); return }
func MaxInt(a1 int32, b1 int32) (r int32) { r = b3go_inline_b3MaxInt(theStack, a1, b1); return }
func ClampInt(a1 int32, lower1 int32, upper1 int32) (r int32) {
	r = b3go_inline_b3ClampInt(theStack, a1, lower1, upper1)
	return
}
func AbsFloat(a1 float32) (r float32) { r = b3go_inline_b3AbsFloat(theStack, a1); return }
func MinFloat(a1 float32, b1 float32) (r float32) {
	r = b3go_inline_b3MinFloat(theStack, a1, b1)
	return
}
func MaxFloat(a1 float32, b1 float32) (r float32) {
	r = b3go_inline_b3MaxFloat(theStack, a1, b1)
	return
}
func ClampFloat(a1 float32, lower1 float32, upper1 float32) (r float32) {
	r = b3go_inline_b3ClampFloat(theStack, a1, lower1, upper1)
	return
}
func LerpFloat(a1 float32, b1 float32, alpha1 float32) (r float32) {
	r = b3go_inline_b3LerpFloat(theStack, a1, b1, alpha1)
	return
}
func Sin(radians1 float32) (r float32) { r = b3go_inline_b3Sin(theStack, radians1); return }
func Cos(radians1 float32) (r float32) { r = b3go_inline_b3Cos(theStack, radians1); return }
func UnwindAngle(radians1 float32) (r float32) {
	r = b3go_inline_b3UnwindAngle(theStack, radians1)
	return
}
func Add(a1 Vec3, b1 Vec3) (r Vec3)         { r = b3go_inline_b3Add(theStack, a1, b1); return }
func Sub(a1 Vec3, b1 Vec3) (r Vec3)         { r = b3go_inline_b3Sub(theStack, a1, b1); return }
func Mul(a1 Vec3, b1 Vec3) (r Vec3)         { r = b3go_inline_b3Mul(theStack, a1, b1); return }
func Neg(a1 Vec3) (r Vec3)                  { r = b3go_inline_b3Neg(theStack, a1); return }
func Dot(a1 Vec3, b1 Vec3) (r float32)      { r = b3go_inline_b3Dot(theStack, a1, b1); return }
func Length(v1 Vec3) (r float32)            { r = b3go_inline_b3Length(theStack, v1); return }
func LengthSquared(a1 Vec3) (r float32)     { r = b3go_inline_b3LengthSquared(theStack, a1); return }
func Distance(a2 Vec3, b2 Vec3) (r float32) { r = b3go_inline_b3Distance(theStack, a2, b2); return }
func DistanceSquared(a1 Vec3, b1 Vec3) (r float32) {
	r = b3go_inline_b3DistanceSquared(theStack, a1, b1)
	return
}
func Normalize(a1 Vec3) (r Vec3)     { r = b3go_inline_b3Normalize(theStack, a1); return }
func Perp(a2 Vec3) (r Vec3)          { r = b3go_inline_b3Perp(theStack, a2); return }
func IsNormalized(a3 Vec3) (r uint8) { r = b3go_inline_b3IsNormalized(theStack, a3); return }
func MulAdd(a1 Vec3, s1 float32, b1 Vec3) (r Vec3) {
	r = b3go_inline_b3MulAdd(theStack, a1, s1, b1)
	return
}
func MulSub(a1 Vec3, s1 float32, b1 Vec3) (r Vec3) {
	r = b3go_inline_b3MulSub(theStack, a1, s1, b1)
	return
}
func MulSV(s1 float32, a1 Vec3) (r Vec3) { r = b3go_inline_b3MulSV(theStack, s1, a1); return }
func Cross(a1 Vec3, b1 Vec3) (r Vec3)    { r = b3go_inline_b3Cross(theStack, a1, b1); return }
func Lerp(a1 Vec3, b1 Vec3, alpha1 float32) (r Vec3) {
	r = b3go_inline_b3Lerp(theStack, a1, b1, alpha1)
	return
}
func Blend2(s1 float32, a1 Vec3, t1 float32, b1 Vec3) (r Vec3) {
	r = b3go_inline_b3Blend2(theStack, s1, a1, t1, b1)
	return
}
func Abs(a2 Vec3) (r Vec3)          { r = b3go_inline_b3Abs(theStack, a2); return }
func Sign(a1 Vec3) (r Vec3)         { r = b3go_inline_b3Sign(theStack, a1); return }
func Min(a2 Vec3, b2 Vec3) (r Vec3) { r = b3go_inline_b3Min(theStack, a2, b2); return }
func Max(a2 Vec3, b2 Vec3) (r Vec3) { r = b3go_inline_b3Max(theStack, a2, b2); return }
func Clamp(a2 Vec3, lower2 Vec3, upper2 Vec3) (r Vec3) {
	r = b3go_inline_b3Clamp(theStack, a2, lower2, upper2)
	return
}
func SafeScale(a7 Vec3) (r Vec3)         { r = b3go_inline_b3SafeScale(theStack, a7); return }
func IsNormalizedQuat(q1 Quat) (r uint8) { r = b3go_inline_b3IsNormalizedQuat(theStack, q1); return }
func RotateVector(q1 Quat, v1 Vec3) (r Vec3) {
	r = b3go_inline_b3RotateVector(theStack, q1, v1)
	return
}
func InvRotateVector(q1 Quat, v1 Vec3) (r Vec3) {
	r = b3go_inline_b3InvRotateVector(theStack, q1, v1)
	return
}
func DotQuat(a1 Quat, b1 Quat) (r float32) { r = b3go_inline_b3DotQuat(theStack, a1, b1); return }
func MulQuat(q11 Quat, q21 Quat) (r Quat)  { r = b3go_inline_b3MulQuat(theStack, q11, q21); return }
func InvMulQuat(q11 Quat, q21 Quat) (r Quat) {
	r = b3go_inline_b3InvMulQuat(theStack, q11, q21)
	return
}
func Conjugate(q1 Quat) (r Quat)     { r = b3go_inline_b3Conjugate(theStack, q1); return }
func NegateQuat(q1 Quat) (r Quat)    { r = b3go_inline_b3NegateQuat(theStack, q1); return }
func NormalizeQuat(q1 Quat) (r Quat) { r = b3go_inline_b3NormalizeQuat(theStack, q1); return }
func MakeQuatFromAxisAngle(axis1 Vec3, radians1 float32) (r Quat) {
	r = b3go_inline_b3MakeQuatFromAxisAngle(theStack, axis1, radians1)
	return
}
func GetQuatAngle(q1 Quat) (r float32)  { r = b3go_inline_b3GetQuatAngle(theStack, q1); return }
func GetTwistAngle(q1 Quat) (r float32) { r = b3go_inline_b3GetTwistAngle(theStack, q1); return }
func GetSwingAngle(q1 Quat) (r float32) { r = b3go_inline_b3GetSwingAngle(theStack, q1); return }
func NLerp(q12 Quat, q21 Quat, alpha2 float32) (r Quat) {
	r = b3go_inline_b3NLerp(theStack, q12, q21, alpha2)
	return
}
func MulTransforms(a5 Transform, b5 Transform) (r Transform) {
	r = b3go_inline_b3MulTransforms(theStack, a5, b5)
	return
}
func InvMulTransforms(a6 Transform, b6 Transform) (r Transform) {
	r = b3go_inline_b3InvMulTransforms(theStack, a6, b6)
	return
}
func InvertTransform(t1 Transform) (r Transform) {
	r = b3go_inline_b3InvertTransform(theStack, t1)
	return
}
func TransformPoint(t1 Transform, v2 Vec3) (r Vec3) {
	r = b3go_inline_b3TransformPoint(theStack, t1, v2)
	return
}
func InvTransformPoint(t1 Transform, v2 Vec3) (r Vec3) {
	r = b3go_inline_b3InvTransformPoint(theStack, t1, v2)
	return
}
func ToPos(v1 Vec3) (r Pos)                 { r = b3go_inline_b3ToPos(theStack, v1); return }
func ToVec3(p1 Pos) (r Vec3)                { r = b3go_inline_b3ToVec3(theStack, p1); return }
func RoundDownFloat(x1 float64) (r float32) { r = b3go_inline_b3RoundDownFloat(theStack, x1); return }
func RoundUpFloat(x1 float64) (r float32)   { r = b3go_inline_b3RoundUpFloat(theStack, x1); return }
func SubPos(a1 Pos, b1 Pos) (r Vec3)        { r = b3go_inline_b3SubPos(theStack, a1, b1); return }
func OffsetPos(p1 Pos, d1 Vec3) (r Pos)     { r = b3go_inline_b3OffsetPos(theStack, p1, d1); return }
func LerpPosition(a1 Pos, b1 Pos, t1 float32) (r Pos) {
	r = b3go_inline_b3LerpPosition(theStack, a1, b1, t1)
	return
}
func TransformWorldPoint(t1 WorldTransform, p1 Vec3) (r1 Pos) {
	r1 = b3go_inline_b3TransformWorldPoint(theStack, t1, p1)
	return
}
func InvTransformWorldPoint(t1 WorldTransform, p1 Pos) (r Vec3) {
	r = b3go_inline_b3InvTransformWorldPoint(theStack, t1, p1)
	return
}
func MulWorldTransforms(A1 WorldTransform, B1 Transform) (r1 WorldTransform) {
	r1 = b3go_inline_b3MulWorldTransforms(theStack, A1, B1)
	return
}
func InvMulWorldTransforms(A1 WorldTransform, B1 WorldTransform) (r Transform) {
	r = b3go_inline_b3InvMulWorldTransforms(theStack, A1, B1)
	return
}
func ToRelativeTransform(t1 WorldTransform, base1 Pos) (r1 Transform) {
	r1 = b3go_inline_b3ToRelativeTransform(theStack, t1, base1)
	return
}
func MakeWorldTransform(t1 Transform) (r WorldTransform) {
	r = b3go_inline_b3MakeWorldTransform(theStack, t1)
	return
}
func OffsetAABB(localBox1 AABB, origin1 Pos) (r AABB) {
	r = b3go_inline_b3OffsetAABB(theStack, localBox1, origin1)
	return
}
func Det(m1 Matrix3) (r float32)               { r = b3go_inline_b3Det(theStack, m1); return }
func MulMV(m1 Matrix3, a1 Vec3) (r Vec3)       { r = b3go_inline_b3MulMV(theStack, m1, a1); return }
func NegateMat3(a1 Matrix3) (r Matrix3)        { r = b3go_inline_b3NegateMat3(theStack, a1); return }
func AddMM(a1 Matrix3, b1 Matrix3) (r Matrix3) { r = b3go_inline_b3AddMM(theStack, a1, b1); return }
func SubMM(a1 Matrix3, b1 Matrix3) (r Matrix3) { r = b3go_inline_b3SubMM(theStack, a1, b1); return }
func MulSM(s1 float32, a1 Matrix3) (r Matrix3) { r = b3go_inline_b3MulSM(theStack, s1, a1); return }
func MulMM(a2 Matrix3, b2 Matrix3) (r Matrix3) { r = b3go_inline_b3MulMM(theStack, a2, b2); return }
func Transpose(m1 Matrix3) (r Matrix3)         { r = b3go_inline_b3Transpose(theStack, m1); return }
func InvertMatrix(m3 Matrix3) (r Matrix3)      { r = b3go_inline_b3InvertMatrix(theStack, m3); return }
func Solve3(m2 Matrix3, a4 Vec3) (r Vec3)      { r = b3go_inline_b3Solve3(theStack, m2, a4); return }
func InvertT(m2 Matrix3) (r Matrix3)           { r = b3go_inline_b3InvertT(theStack, m2); return }
func AbsMatrix3(m1 Matrix3) (r Matrix3)        { r = b3go_inline_b3AbsMatrix3(theStack, m1); return }
func MakeMatrixFromQuat(q1 Quat) (r Matrix3) {
	r = b3go_inline_b3MakeMatrixFromQuat(theStack, q1)
	return
}
func AABBContains(a1 AABB, b1 AABB) (r uint8) {
	r = b3go_inline_b3AABB_Contains(theStack, a1, b1)
	return
}
func AABBArea(a2 AABB) (r float32)        { r = b3go_inline_b3AABB_Area(theStack, a2); return }
func AABBCenter(a3 AABB) (r Vec3)         { r = b3go_inline_b3AABB_Center(theStack, a3); return }
func AABBExtents(a3 AABB) (r Vec3)        { r = b3go_inline_b3AABB_Extents(theStack, a3); return }
func AABBUnion(a5 AABB, b5 AABB) (r AABB) { r = b3go_inline_b3AABB_Union(theStack, a5, b5); return }
func AABBInflate(a3 AABB, extension1 float32) (r AABB) {
	r = b3go_inline_b3AABB_Inflate(theStack, a3, extension1)
	return
}
func AABBOverlaps(a1 AABB, b1 AABB) (r uint8) {
	r = b3go_inline_b3AABB_Overlaps(theStack, a1, b1)
	return
}
func AABBTransform(transform1 Transform, a11 AABB) (r AABB) {
	r = b3go_inline_b3AABB_Transform(theStack, transform1, a11)
	return
}
func ClosestPointToAABB(point1 Vec3, a3 AABB) (r Vec3) {
	r = b3go_inline_b3ClosestPointToAABB(theStack, point1, a3)
	return
}
func MakeDebugColor(rgb1 HexColor, material1 DebugMaterial) (r uint32) {
	r = b3go_inline_b3MakeDebugColor(theStack, rgb1, material1)
	return
}

const CapsuleShape ShapeType = b3_capsuleShape
const CompoundShape ShapeType = b3_compoundShape
const HeightShape ShapeType = b3_heightShape
const HullShape ShapeType = b3_hullShape
const MeshShape ShapeType = b3_meshShape
const SphereShape ShapeType = b3_sphereShape
const StaticBody BodyType = b3_staticBody
const KinematicBody BodyType = b3_kinematicBody
const DynamicBody BodyType = b3_dynamicBody
const JointTypeParallel JointType = b3_parallelJoint
const JointTypeDistance JointType = b3_distanceJoint
const JointTypeFilter JointType = b3_filterJoint
const JointTypeMotor JointType = b3_motorJoint
const JointTypePrismatic JointType = b3_prismaticJoint
const JointTypeRevolute JointType = b3_revoluteJoint
const JointTypeSpherical JointType = b3_sphericalJoint
const JointTypeWeld JointType = b3_weldJoint
const JointTypeWheel JointType = b3_wheelJoint
const ToiStateUnknown TOIState = b3_toiStateUnknown
const ToiStateFailed TOIState = b3_toiStateFailed
const ToiStateOverlapped TOIState = b3_toiStateOverlapped
const ToiStateHit TOIState = b3_toiStateHit
const ToiStateSeparated TOIState = b3_toiStateSeparated
const ConcaveEdge1 MeshEdgeFlags = b3_concaveEdge1
const ConcaveEdge2 MeshEdgeFlags = b3_concaveEdge2
const ConcaveEdge3 MeshEdgeFlags = b3_concaveEdge3
const InverseConcaveEdge1 MeshEdgeFlags = b3_inverseConcaveEdge1
const InverseConcaveEdge2 MeshEdgeFlags = b3_inverseConcaveEdge2
const InverseConcaveEdge3 MeshEdgeFlags = b3_inverseConcaveEdge3
const AllConcaveEdges MeshEdgeFlags = b3_allConcaveEdges
const FlatEdge1 MeshEdgeFlags = b3_flatEdge1
const FlatEdge2 MeshEdgeFlags = b3_flatEdge2
const FlatEdge3 MeshEdgeFlags = b3_flatEdge3
const AllFlatEdges MeshEdgeFlags = b3_allFlatEdges
const InvalidAxis SeparatingFeature = b3_invalidAxis
const BacksideAxis SeparatingFeature = b3_backsideAxis
const FaceAxisA SeparatingFeature = b3_faceAxisA
const FaceAxisB SeparatingFeature = b3_faceAxisB
const EdgePairAxis SeparatingFeature = b3_edgePairAxis
const ClosestPointsAxis SeparatingFeature = b3_closestPointsAxis
const ManualFaceAxisA SeparatingFeature = b3_manualFaceAxisA
const ManualFaceAxisB SeparatingFeature = b3_manualFaceAxisB
const ManualEdgePairAxis SeparatingFeature = b3_manualEdgePairAxis
const FeatureNone TriangleFeature = b3_featureNone
const FeatureTriangleFace TriangleFeature = b3_featureTriangleFace
const FeatureHullFace TriangleFeature = b3_featureHullFace
const FeatureEdge1 TriangleFeature = b3_featureEdge1
const FeatureEdge2 TriangleFeature = b3_featureEdge2
const FeatureEdge3 TriangleFeature = b3_featureEdge3
const FeatureVertex1 TriangleFeature = b3_featureVertex1
const FeatureVertex2 TriangleFeature = b3_featureVertex2
const FeatureVertex3 TriangleFeature = b3_featureVertex3
const ColorAliceBlue HexColor = b3_colorAliceBlue
const ColorAntiqueWhite HexColor = b3_colorAntiqueWhite
const ColorAqua HexColor = b3_colorAqua
const ColorAquamarine HexColor = b3_colorAquamarine
const ColorAzure HexColor = b3_colorAzure
const ColorBeige HexColor = b3_colorBeige
const ColorBisque HexColor = b3_colorBisque
const ColorBlack HexColor = b3_colorBlack
const ColorBlanchedAlmond HexColor = b3_colorBlanchedAlmond
const ColorBlue HexColor = b3_colorBlue
const ColorBlueViolet HexColor = b3_colorBlueViolet
const ColorBrown HexColor = b3_colorBrown
const ColorBurlywood HexColor = b3_colorBurlywood
const ColorCadetBlue HexColor = b3_colorCadetBlue
const ColorChartreuse HexColor = b3_colorChartreuse
const ColorChocolate HexColor = b3_colorChocolate
const ColorCoral HexColor = b3_colorCoral
const ColorCornflowerBlue HexColor = b3_colorCornflowerBlue
const ColorCornsilk HexColor = b3_colorCornsilk
const ColorCrimson HexColor = b3_colorCrimson
const ColorCyan HexColor = b3_colorCyan
const ColorDarkBlue HexColor = b3_colorDarkBlue
const ColorDarkCyan HexColor = b3_colorDarkCyan
const ColorDarkGoldenRod HexColor = b3_colorDarkGoldenRod
const ColorDarkGray HexColor = b3_colorDarkGray
const ColorDarkGreen HexColor = b3_colorDarkGreen
const ColorDarkKhaki HexColor = b3_colorDarkKhaki
const ColorDarkMagenta HexColor = b3_colorDarkMagenta
const ColorDarkOliveGreen HexColor = b3_colorDarkOliveGreen
const ColorDarkOrange HexColor = b3_colorDarkOrange
const ColorDarkOrchid HexColor = b3_colorDarkOrchid
const ColorDarkRed HexColor = b3_colorDarkRed
const ColorDarkSalmon HexColor = b3_colorDarkSalmon
const ColorDarkSeaGreen HexColor = b3_colorDarkSeaGreen
const ColorDarkSlateBlue HexColor = b3_colorDarkSlateBlue
const ColorDarkSlateGray HexColor = b3_colorDarkSlateGray
const ColorDarkTurquoise HexColor = b3_colorDarkTurquoise
const ColorDarkViolet HexColor = b3_colorDarkViolet
const ColorDeepPink HexColor = b3_colorDeepPink
const ColorDeepSkyBlue HexColor = b3_colorDeepSkyBlue
const ColorDimGray HexColor = b3_colorDimGray
const ColorDodgerBlue HexColor = b3_colorDodgerBlue
const ColorFireBrick HexColor = b3_colorFireBrick
const ColorFloralWhite HexColor = b3_colorFloralWhite
const ColorForestGreen HexColor = b3_colorForestGreen
const ColorFuchsia HexColor = b3_colorFuchsia
const ColorGainsboro HexColor = b3_colorGainsboro
const ColorGhostWhite HexColor = b3_colorGhostWhite
const ColorGold HexColor = b3_colorGold
const ColorGoldenRod HexColor = b3_colorGoldenRod
const ColorGray HexColor = b3_colorGray
const ColorGreen HexColor = b3_colorGreen
const ColorGreenYellow HexColor = b3_colorGreenYellow
const ColorHoneyDew HexColor = b3_colorHoneyDew
const ColorHotPink HexColor = b3_colorHotPink
const ColorIndianRed HexColor = b3_colorIndianRed
const ColorIndigo HexColor = b3_colorIndigo
const ColorIvory HexColor = b3_colorIvory
const ColorKhaki HexColor = b3_colorKhaki
const ColorLavender HexColor = b3_colorLavender
const ColorLavenderBlush HexColor = b3_colorLavenderBlush
const ColorLawnGreen HexColor = b3_colorLawnGreen
const ColorLemonChiffon HexColor = b3_colorLemonChiffon
const ColorLightBlue HexColor = b3_colorLightBlue
const ColorLightCoral HexColor = b3_colorLightCoral
const ColorLightCyan HexColor = b3_colorLightCyan
const ColorLightGoldenRodYellow HexColor = b3_colorLightGoldenRodYellow
const ColorLightGray HexColor = b3_colorLightGray
const ColorLightGreen HexColor = b3_colorLightGreen
const ColorLightPink HexColor = b3_colorLightPink
const ColorLightSalmon HexColor = b3_colorLightSalmon
const ColorLightSeaGreen HexColor = b3_colorLightSeaGreen
const ColorLightSkyBlue HexColor = b3_colorLightSkyBlue
const ColorLightSlateGray HexColor = b3_colorLightSlateGray
const ColorLightSteelBlue HexColor = b3_colorLightSteelBlue
const ColorLightYellow HexColor = b3_colorLightYellow
const ColorLime HexColor = b3_colorLime
const ColorLimeGreen HexColor = b3_colorLimeGreen
const ColorLinen HexColor = b3_colorLinen
const ColorMagenta HexColor = b3_colorMagenta
const ColorMaroon HexColor = b3_colorMaroon
const ColorMediumAquaMarine HexColor = b3_colorMediumAquaMarine
const ColorMediumBlue HexColor = b3_colorMediumBlue
const ColorMediumOrchid HexColor = b3_colorMediumOrchid
const ColorMediumPurple HexColor = b3_colorMediumPurple
const ColorMediumSeaGreen HexColor = b3_colorMediumSeaGreen
const ColorMediumSlateBlue HexColor = b3_colorMediumSlateBlue
const ColorMediumSpringGreen HexColor = b3_colorMediumSpringGreen
const ColorMediumTurquoise HexColor = b3_colorMediumTurquoise
const ColorMediumVioletRed HexColor = b3_colorMediumVioletRed
const ColorMidnightBlue HexColor = b3_colorMidnightBlue
const ColorMintCream HexColor = b3_colorMintCream
const ColorMistyRose HexColor = b3_colorMistyRose
const ColorMoccasin HexColor = b3_colorMoccasin
const ColorNavajoWhite HexColor = b3_colorNavajoWhite
const ColorNavy HexColor = b3_colorNavy
const ColorOldLace HexColor = b3_colorOldLace
const ColorOlive HexColor = b3_colorOlive
const ColorOliveDrab HexColor = b3_colorOliveDrab
const ColorOrange HexColor = b3_colorOrange
const ColorOrangeRed HexColor = b3_colorOrangeRed
const ColorOrchid HexColor = b3_colorOrchid
const ColorPaleGoldenRod HexColor = b3_colorPaleGoldenRod
const ColorPaleGreen HexColor = b3_colorPaleGreen
const ColorPaleTurquoise HexColor = b3_colorPaleTurquoise
const ColorPaleVioletRed HexColor = b3_colorPaleVioletRed
const ColorPapayaWhip HexColor = b3_colorPapayaWhip
const ColorPeachPuff HexColor = b3_colorPeachPuff
const ColorPeru HexColor = b3_colorPeru
const ColorPink HexColor = b3_colorPink
const ColorPlum HexColor = b3_colorPlum
const ColorPowderBlue HexColor = b3_colorPowderBlue
const ColorPurple HexColor = b3_colorPurple
const ColorRebeccaPurple HexColor = b3_colorRebeccaPurple
const ColorRed HexColor = b3_colorRed
const ColorRosyBrown HexColor = b3_colorRosyBrown
const ColorRoyalBlue HexColor = b3_colorRoyalBlue
const ColorSaddleBrown HexColor = b3_colorSaddleBrown
const ColorSalmon HexColor = b3_colorSalmon
const ColorSandyBrown HexColor = b3_colorSandyBrown
const ColorSeaGreen HexColor = b3_colorSeaGreen
const ColorSeaShell HexColor = b3_colorSeaShell
const ColorSienna HexColor = b3_colorSienna
const ColorSilver HexColor = b3_colorSilver
const ColorSkyBlue HexColor = b3_colorSkyBlue
const ColorSlateBlue HexColor = b3_colorSlateBlue
const ColorSlateGray HexColor = b3_colorSlateGray
const ColorSnow HexColor = b3_colorSnow
const ColorSpringGreen HexColor = b3_colorSpringGreen
const ColorSteelBlue HexColor = b3_colorSteelBlue
const ColorTan HexColor = b3_colorTan
const ColorTeal HexColor = b3_colorTeal
const ColorThistle HexColor = b3_colorThistle
const ColorTomato HexColor = b3_colorTomato
const ColorTurquoise HexColor = b3_colorTurquoise
const ColorViolet HexColor = b3_colorViolet
const ColorWheat HexColor = b3_colorWheat
const ColorWhite HexColor = b3_colorWhite
const ColorWhiteSmoke HexColor = b3_colorWhiteSmoke
const ColorYellow HexColor = b3_colorYellow
const ColorYellowGreen HexColor = b3_colorYellowGreen
const ColorBox2DRed HexColor = b3_colorBox2DRed
const ColorBox2DBlue HexColor = b3_colorBox2DBlue
const ColorBox2DGreen HexColor = b3_colorBox2DGreen
const ColorBox2DYellow HexColor = b3_colorBox2DYellow
const DebugMaterialDefault DebugMaterial = b3_debugMaterialDefault
const DebugMaterialMatte DebugMaterial = b3_debugMaterialMatte
const DebugMaterialSoft DebugMaterial = b3_debugMaterialSoft
const DebugMaterialDead DebugMaterial = b3_debugMaterialDead
const DebugMaterialGlossy DebugMaterial = b3_debugMaterialGlossy
const DebugMaterialMetallic DebugMaterial = b3_debugMaterialMetallic
