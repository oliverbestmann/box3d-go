package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

// C documentation
//
//	// Solve contacts: normal, friction and rolling resistance.
func b3SolveContacts_Mesh(tls *_Stack, block b3SolverBlock, context uintptr) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var P, P1, b8, c, d, deltaImpulse2, dp, ds, normal, oldImpulse1, rA, rA1, rB, rB1, t1, t2, t3, tangent1, tangent2, vA, vB, vr, vrA, vrA1, vrB, vrB1, wA, wB, v11, v12, v13, v14, v16, v18, v19, v21, v22, v23, v25, v27, v29, v30, v33, v34, v35, v36, v38, v4, v40, v41, v43, v44, v45, v47, v49, v5, v51, v52, v54, v55, v56, v58, v59, v6, v60, v62, v63, v67, v68, v69, v71, v72, v73, v75, v76, v77, v79, v80, v81, v83, v84, v85, v87, v88 Vec3
	var b10, c1, deltaImpulse3, newImpulse1, tm, vt, v245, v246, v248, v249 Vec2
	var color, constraint, contactConstraint, contactConstraints, cp, stateA, stateB, states, world, v2, v3 uintptr
	var deltaImpulse, deltaImpulse1, friction, inv_h, lengthSquared, mA, mB, magSqr, maxImpulse, maxImpulse1, maxImpulse2, newImpulse, oldImpulse, rollingResistance, s4, scale, totalNormalImpulse, totalTwistLimit, twistSpeed, velocityBias, vn, v17, v28, v39, v50, v64, v66, v89, v91, v92, v93, v95 float32
	var dqA, dqB, v10, v32 Quat
	var endIndex, i, indexA, indexB, j, manifoldCount, pointCount, pointIndex, startIndex int32
	var iA, iB, v109 Matrix3
	var v244 b3Matrix2
	var _ /* dummyState at bp+0 */ b3BodyState
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = P, P1, b10, b8, c, c1, color, constraint, contactConstraint, contactConstraints, cp, d, deltaImpulse, deltaImpulse1, deltaImpulse2, deltaImpulse3, dp, dqA, dqB, ds, endIndex, friction, i, iA, iB, indexA, indexB, inv_h, j, lengthSquared, mA, mB, magSqr, manifoldCount, maxImpulse, maxImpulse1, maxImpulse2, newImpulse, newImpulse1, normal, oldImpulse, oldImpulse1, pointCount, pointIndex, rA, rA1, rB, rB1, rollingResistance, s4, scale, startIndex, stateA, stateB, states, t1, t2, t3, tangent1, tangent2, tm, totalNormalImpulse, totalTwistLimit, twistSpeed, vA, vB, velocityBias, vn, vr, vrA, vrA1, vrB, vrB1, vt, wA, wB, world, v10, v109, v11, v12, v13, v14, v16, v17, v18, v19, v2, v21, v22, v23, v244, v245, v246, v248, v249, v25, v27, v28, v29, v3, v30, v32, v33, v34, v35, v36, v38, v39, v4, v40, v41, v43, v44, v45, v47, v49, v5, v50, v51, v52, v54, v55, v56, v58, v59, v6, v60, v62, v63, v64, v66, v67, v68, v69, v71, v72, v73, v75, v76, v77, v79, v80, v81, v83, v84, v85, v87, v88, v89, v91, v92, v93, v95
	world = (*b3StepContext)(unsafe.Pointer(context)).World
	color = world + 1160 + uintptr(block.ColorIndex)*120
	contactConstraints = (*b3GraphColor)(unsafe.Pointer(color)).ContactConstraints
	states = (*b3StepContext)(unsafe.Pointer(context)).States
	// This is a dummy state to represent a static body because static bodies have no solver body.
	*(*b3BodyState)(unsafe.Add(nil, bp)) = b3_identityBodyState
	// The last block might not be full
	startIndex = block.StartIndex
	endIndex = startIndex + int32FromUint16(block.Count)
	inv_h = (*b3StepContext)(unsafe.Pointer(context)).Inv_h
	i = startIndex
	for {
		if !(i < endIndex) {
			break
		}
		contactConstraint = contactConstraints + uintptr(i)*168
		manifoldCount = (*b3ContactConstraint)(unsafe.Pointer(contactConstraint)).ManifoldCount
		indexA = (*b3ContactConstraint)(unsafe.Pointer(contactConstraint)).IndexA
		indexB = (*b3ContactConstraint)(unsafe.Pointer(contactConstraint)).IndexB
		mA = (*b3ContactConstraint)(unsafe.Pointer(contactConstraint)).InvMassA
		iA = (*b3ContactConstraint)(unsafe.Pointer(contactConstraint)).InvIA
		mB = (*b3ContactConstraint)(unsafe.Pointer(contactConstraint)).InvMassB
		iB = (*b3ContactConstraint)(unsafe.Pointer(contactConstraint)).InvIB
		if indexA == -int32(1) {
			v2 = bp
		} else {
			v2 = states + uintptr(indexA)*64
		}
		stateA = v2
		vA = (*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity
		wA = (*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity
		dqA = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaRotation
		if indexB == -int32(1) {
			v3 = bp
		} else {
			v3 = states + uintptr(indexB)*64
		}
		stateB = v3
		vB = (*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity
		wB = (*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity
		dqB = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaRotation
		v4 = (*b3BodyState)(unsafe.Pointer(stateB)).DeltaPosition
		v5 = (*b3BodyState)(unsafe.Pointer(stateA)).DeltaPosition
		v6 = Vec3{
			X: v4.X - v5.X,
			Y: v4.Y - v5.Y,
			Z: v4.Z - v5.Z,
		}
		goto _7
	_7:
		dp = v6
		friction = (*b3ContactConstraint)(unsafe.Pointer(contactConstraint)).Friction
		rollingResistance = (*b3ContactConstraint)(unsafe.Pointer(contactConstraint)).RollingResistance
		j = 0
		for {
			if !(j < manifoldCount) {
				break
			}
			constraint = (*b3ContactConstraint)(unsafe.Pointer(contactConstraint)).Constraints + uintptr(j)*324
			pointCount = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).PointCount
			normal = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).Normal
			totalNormalImpulse = float32FromFloat32(0)
			totalTwistLimit = float32FromFloat32(0)
			pointIndex = 0
			for {
				if !(pointIndex < pointCount) {
					break
				}
				cp = constraint + uintptr(pointIndex)*52
				// Fixed anchor points for applying impulses
				rA = (*b3ManifoldConstraintPoint)(unsafe.Pointer(cp)).RA
				rB = (*b3ManifoldConstraintPoint)(unsafe.Pointer(cp)).RB
				v10 = dqB
				v11 = rB
				v12 = v10.V
				v13 = v11
				c.X = float32(v12.Y*v13.Z) - float32(v12.Z*v13.Y)
				c.Y = float32(v12.Z*v13.X) - float32(v12.X*v13.Z)
				c.Z = float32(v12.X*v13.Y) - float32(v12.Y*v13.X)
				v14 = c
				goto _15
			_15:
				t1 = v14
				v16 = t1
				v17 = v10.S
				v18 = v11
				v19 = Vec3{
					X: v16.X + float32(v17*v18.X),
					Y: v16.Y + float32(v17*v18.Y),
					Z: v16.Z + float32(v17*v18.Z),
				}
				goto _20
			_20:
				t2 = v19
				v21 = v10.V
				v22 = t2
				c.X = float32(v21.Y*v22.Z) - float32(v21.Z*v22.Y)
				c.Y = float32(v21.Z*v22.X) - float32(v21.X*v22.Z)
				c.Z = float32(v21.X*v22.Y) - float32(v21.Y*v22.X)
				v23 = c
				goto _24
			_24:
				t3 = v23
				v27 = v11
				v28 = float32FromFloat32(2)
				v29 = t3
				v30 = Vec3{
					X: v27.X + float32(v28*v29.X),
					Y: v27.Y + float32(v28*v29.Y),
					Z: v27.Z + float32(v28*v29.Z),
				}
				goto _31
			_31:
				v25 = v30
				goto _26
			_26:
				v32 = dqA
				v33 = rA
				v34 = v32.V
				v35 = v33
				c.X = float32(v34.Y*v35.Z) - float32(v34.Z*v35.Y)
				c.Y = float32(v34.Z*v35.X) - float32(v34.X*v35.Z)
				c.Z = float32(v34.X*v35.Y) - float32(v34.Y*v35.X)
				v36 = c
				goto _37
			_37:
				t1 = v36
				v38 = t1
				v39 = v32.S
				v40 = v33
				v41 = Vec3{
					X: v38.X + float32(v39*v40.X),
					Y: v38.Y + float32(v39*v40.Y),
					Z: v38.Z + float32(v39*v40.Z),
				}
				goto _42
			_42:
				t2 = v41
				v43 = v32.V
				v44 = t2
				c.X = float32(v43.Y*v44.Z) - float32(v43.Z*v44.Y)
				c.Y = float32(v43.Z*v44.X) - float32(v43.X*v44.Z)
				c.Z = float32(v43.X*v44.Y) - float32(v43.Y*v44.X)
				v45 = c
				goto _46
			_46:
				t3 = v45
				v49 = v33
				v50 = float32FromFloat32(2)
				v51 = t3
				v52 = Vec3{
					X: v49.X + float32(v50*v51.X),
					Y: v49.Y + float32(v50*v51.Y),
					Z: v49.Z + float32(v50*v51.Z),
				}
				goto _53
			_53:
				v47 = v52
				goto _48
			_48:
				v54 = v25
				v55 = v47
				v56 = Vec3{
					X: v54.X - v55.X,
					Y: v54.Y - v55.Y,
					Z: v54.Z - v55.Z,
				}
				goto _57
			_57:
				v58 = dp
				v59 = v56
				v60 = Vec3{
					X: v58.X + v59.X,
					Y: v58.Y + v59.Y,
					Z: v58.Z + v59.Z,
				}
				goto _61
			_61:
				// compute current separation
				// this is subject to round-off error if the anchor is far from the body center of mass
				ds = v60
				v62 = ds
				v63 = normal
				v64 = float32(v62.X*v63.X) + float32(v62.Y*v63.Y) + float32(v62.Z*v63.Z)
				goto _65
			_65:
				s4 = v64 + (*b3ManifoldConstraintPoint)(unsafe.Pointer(cp)).BaseSeparation
				if s4 > float32FromFloat32(0) {
					v66 = float32(s4 * inv_h)
				} else {
					v66 = float32FromFloat32(0)
				}
				// speculative bias, zero when overlapped
				velocityBias = v66
				v67 = wA
				v68 = rA
				c.X = float32(v67.Y*v68.Z) - float32(v67.Z*v68.Y)
				c.Y = float32(v67.Z*v68.X) - float32(v67.X*v68.Z)
				c.Z = float32(v67.X*v68.Y) - float32(v67.Y*v68.X)
				v69 = c
				goto _70
			_70:
				v71 = vA
				v72 = v69
				v73 = Vec3{
					X: v71.X + v72.X,
					Y: v71.Y + v72.Y,
					Z: v71.Z + v72.Z,
				}
				goto _74
			_74:
				// relative normal velocity at contact
				vrA = v73
				v75 = wB
				v76 = rB
				c.X = float32(v75.Y*v76.Z) - float32(v75.Z*v76.Y)
				c.Y = float32(v75.Z*v76.X) - float32(v75.X*v76.Z)
				c.Z = float32(v75.X*v76.Y) - float32(v75.Y*v76.X)
				v77 = c
				goto _78
			_78:
				v79 = vB
				v80 = v77
				v81 = Vec3{
					X: v79.X + v80.X,
					Y: v79.Y + v80.Y,
					Z: v79.Z + v80.Z,
				}
				goto _82
			_82:
				vrB = v81
				v83 = vrB
				v84 = vrA
				v85 = Vec3{
					X: v83.X - v84.X,
					Y: v83.Y - v84.Y,
					Z: v83.Z - v84.Z,
				}
				goto _86
			_86:
				v87 = v85
				v88 = normal
				v89 = float32(v87.X*v88.X) + float32(v87.Y*v88.Y) + float32(v87.Z*v88.Z)
				goto _90
			_90:
				vn = v89
				// incremental normal impulse
				deltaImpulse = float32(-(*b3ManifoldConstraintPoint)(unsafe.Pointer(cp)).NormalMass * (vn + velocityBias))
				v91 = (*b3ManifoldConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse + deltaImpulse
				v92 = float32FromFloat32(0)
				if v91 > v92 {
					v95 = v91
				} else {
					v95 = v92
				}
				v93 = v95
				goto _94
			_94:
				// clamp the accumulated impulse
				newImpulse = v93
				deltaImpulse = newImpulse - (*b3ManifoldConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse
				(*b3ManifoldConstraintPoint)(unsafe.Pointer(cp)).NormalImpulse = newImpulse
				*(*float32)(unsafe.Add(nil, cp+36)) += newImpulse
				totalNormalImpulse = totalNormalImpulse + newImpulse
				totalTwistLimit = totalTwistLimit + float32((*b3ManifoldConstraintPoint)(unsafe.Pointer(cp)).LeverArm*newImpulse)
				v17 = deltaImpulse
				v4 = normal
				v5 = Vec3{
					X: float32(v17 * v4.X),
					Y: float32(v17 * v4.Y),
					Z: float32(v17 * v4.Z),
				}
				goto _99
			_99:
				// apply normal impulse
				P = v5
				v6 = vA
				v28 = mA
				v11 = P
				v12 = Vec3{
					X: v6.X - float32(v28*v11.X),
					Y: v6.Y - float32(v28*v11.Y),
					Z: v6.Z - float32(v28*v11.Z),
				}
				goto _104
			_104:
				vA = v12
				v4 = rA
				v5 = P
				c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
				c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
				c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
				v6 = c
				goto _108
			_108:
				v109 = iA
				v11 = v6
				b8 = Vec3{
					X: float32(v109.Cx.X*v11.X) + float32(v109.Cy.X*v11.Y) + float32(v109.Cz.X*v11.Z),
					Y: float32(v109.Cx.Y*v11.X) + float32(v109.Cy.Y*v11.Y) + float32(v109.Cz.Y*v11.Z),
					Z: float32(v109.Cx.Z*v11.X) + float32(v109.Cy.Z*v11.Y) + float32(v109.Cz.Z*v11.Z),
				}
				v12 = b8
				goto _112
			_112:
				v13 = wA
				v14 = v12
				v16 = Vec3{
					X: v13.X - v14.X,
					Y: v13.Y - v14.Y,
					Z: v13.Z - v14.Z,
				}
				goto _116
			_116:
				wA = v16
				v4 = vB
				v17 = mB
				v5 = P
				v6 = Vec3{
					X: v4.X + float32(v17*v5.X),
					Y: v4.Y + float32(v17*v5.Y),
					Z: v4.Z + float32(v17*v5.Z),
				}
				goto _121
			_121:
				vB = v6
				v4 = rB
				v5 = P
				c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
				c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
				c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
				v6 = c
				goto _125
			_125:
				v109 = iB
				v11 = v6
				b8 = Vec3{
					X: float32(v109.Cx.X*v11.X) + float32(v109.Cy.X*v11.Y) + float32(v109.Cz.X*v11.Z),
					Y: float32(v109.Cx.Y*v11.X) + float32(v109.Cy.Y*v11.Y) + float32(v109.Cz.Y*v11.Z),
					Z: float32(v109.Cx.Z*v11.X) + float32(v109.Cy.Z*v11.Y) + float32(v109.Cz.Z*v11.Z),
				}
				v12 = b8
				goto _129
			_129:
				v13 = wB
				v14 = v12
				v16 = Vec3{
					X: v13.X + v14.X,
					Y: v13.Y + v14.Y,
					Z: v13.Z + v14.Z,
				}
				goto _133
			_133:
				wB = v16
				goto _9
			_9:
				;
				pointIndex = pointIndex + 1
			}
			// Central twist friction
			v4 = wB
			v5 = wA
			v6 = Vec3{
				X: v4.X - v5.X,
				Y: v4.Y - v5.Y,
				Z: v4.Z - v5.Z,
			}
			goto _137
		_137:
			v11 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).Normal
			v12 = v6
			v17 = float32(v11.X*v12.X) + float32(v11.Y*v12.Y) + float32(v11.Z*v12.Z)
			goto _141
		_141:
			twistSpeed = v17
			maxImpulse = float32(friction * totalTwistLimit)
			deltaImpulse1 = float32(-(*b3ManifoldConstraint)(unsafe.Pointer(constraint)).TwistMass * twistSpeed)
			oldImpulse = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).TwistImpulse
			v28 = oldImpulse + deltaImpulse1
			v39 = -maxImpulse
			v50 = maxImpulse
			if v28 < v39 {
				v66 = v39
			} else {
				if v50 < v28 {
					v89 = v50
				} else {
					v89 = v28
				}
				v66 = v89
			}
			v64 = v66
			goto _146
		_146:
			(*b3ManifoldConstraint)(unsafe.Pointer(constraint)).TwistImpulse = v64
			deltaImpulse1 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).TwistImpulse - oldImpulse
			v17 = deltaImpulse1
			v4 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).Normal
			v5 = Vec3{
				X: float32(v17 * v4.X),
				Y: float32(v17 * v4.Y),
				Z: float32(v17 * v4.Z),
			}
			goto _152
		_152:
			v109 = iA
			v6 = v5
			b8 = Vec3{
				X: float32(v109.Cx.X*v6.X) + float32(v109.Cy.X*v6.Y) + float32(v109.Cz.X*v6.Z),
				Y: float32(v109.Cx.Y*v6.X) + float32(v109.Cy.Y*v6.Y) + float32(v109.Cz.Y*v6.Z),
				Z: float32(v109.Cx.Z*v6.X) + float32(v109.Cy.Z*v6.Y) + float32(v109.Cz.Z*v6.Z),
			}
			v11 = b8
			goto _156
		_156:
			v12 = wA
			v13 = v11
			v14 = Vec3{
				X: v12.X - v13.X,
				Y: v12.Y - v13.Y,
				Z: v12.Z - v13.Z,
			}
			goto _160
		_160:
			wA = v14
			v17 = deltaImpulse1
			v4 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).Normal
			v5 = Vec3{
				X: float32(v17 * v4.X),
				Y: float32(v17 * v4.Y),
				Z: float32(v17 * v4.Z),
			}
			goto _164
		_164:
			v109 = iB
			v6 = v5
			b8 = Vec3{
				X: float32(v109.Cx.X*v6.X) + float32(v109.Cy.X*v6.Y) + float32(v109.Cz.X*v6.Z),
				Y: float32(v109.Cx.Y*v6.X) + float32(v109.Cy.Y*v6.Y) + float32(v109.Cz.Y*v6.Z),
				Z: float32(v109.Cx.Z*v6.X) + float32(v109.Cy.Z*v6.Y) + float32(v109.Cz.Z*v6.Z),
			}
			v11 = b8
			goto _168
		_168:
			v12 = wB
			v13 = v11
			v14 = Vec3{
				X: v12.X + v13.X,
				Y: v12.Y + v13.Y,
				Z: v12.Z + v13.Z,
			}
			goto _172
		_172:
			wB = v14
			// Rolling resistance
			if rollingResistance > float32FromFloat32(0) {
				v4 = wB
				v5 = wA
				v6 = Vec3{
					X: v4.X - v5.X,
					Y: v4.Y - v5.Y,
					Z: v4.Z - v5.Z,
				}
				goto _176
			_176:
				v109 = (*b3ContactConstraint)(unsafe.Pointer(contactConstraint)).RollingMass
				v11 = v6
				b8 = Vec3{
					X: float32(v109.Cx.X*v11.X) + float32(v109.Cy.X*v11.Y) + float32(v109.Cz.X*v11.Z),
					Y: float32(v109.Cx.Y*v11.X) + float32(v109.Cy.Y*v11.Y) + float32(v109.Cz.Y*v11.Z),
					Z: float32(v109.Cx.Z*v11.X) + float32(v109.Cy.Z*v11.Y) + float32(v109.Cz.Z*v11.Z),
				}
				v12 = b8
				goto _180
			_180:
				v13 = v12
				v14 = Vec3{
					X: -v13.X,
					Y: -v13.Y,
					Z: -v13.Z,
				}
				goto _183
			_183:
				deltaImpulse2 = v14
				oldImpulse1 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).RollingImpulse
				v16 = oldImpulse1
				v18 = deltaImpulse2
				v19 = Vec3{
					X: v16.X + v18.X,
					Y: v16.Y + v18.Y,
					Z: v16.Z + v18.Z,
				}
				goto _187
			_187:
				(*b3ManifoldConstraint)(unsafe.Pointer(constraint)).RollingImpulse = v19
				maxImpulse1 = float32(rollingResistance * totalNormalImpulse)
				v4 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).RollingImpulse
				v5 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).RollingImpulse
				v17 = float32(v4.X*v5.X) + float32(v4.Y*v5.Y) + float32(v4.Z*v5.Z)
				goto _191
			_191:
				magSqr = v17
				if magSqr > float32(maxImpulse1*maxImpulse1)+float32FromFloat32(1.1920928955078125e-07) {
					v28 = maxImpulse1 / sqrtf(tls, magSqr)
					v6 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).RollingImpulse
					v11 = Vec3{
						X: float32(v28 * v6.X),
						Y: float32(v28 * v6.Y),
						Z: float32(v28 * v6.Z),
					}
					goto _195
				_195:
					(*b3ManifoldConstraint)(unsafe.Pointer(constraint)).RollingImpulse = v11
				}
				v4 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).RollingImpulse
				v5 = oldImpulse1
				v6 = Vec3{
					X: v4.X - v5.X,
					Y: v4.Y - v5.Y,
					Z: v4.Z - v5.Z,
				}
				goto _199
			_199:
				deltaImpulse2 = v6
				v109 = iA
				v4 = deltaImpulse2
				b8 = Vec3{
					X: float32(v109.Cx.X*v4.X) + float32(v109.Cy.X*v4.Y) + float32(v109.Cz.X*v4.Z),
					Y: float32(v109.Cx.Y*v4.X) + float32(v109.Cy.Y*v4.Y) + float32(v109.Cz.Y*v4.Z),
					Z: float32(v109.Cx.Z*v4.X) + float32(v109.Cy.Z*v4.Y) + float32(v109.Cz.Z*v4.Z),
				}
				v5 = b8
				goto _203
			_203:
				v6 = wA
				v11 = v5
				v12 = Vec3{
					X: v6.X - v11.X,
					Y: v6.Y - v11.Y,
					Z: v6.Z - v11.Z,
				}
				goto _207
			_207:
				wA = v12
				v109 = iB
				v4 = deltaImpulse2
				b8 = Vec3{
					X: float32(v109.Cx.X*v4.X) + float32(v109.Cy.X*v4.Y) + float32(v109.Cz.X*v4.Z),
					Y: float32(v109.Cx.Y*v4.X) + float32(v109.Cy.Y*v4.Y) + float32(v109.Cz.Y*v4.Z),
					Z: float32(v109.Cx.Z*v4.X) + float32(v109.Cy.Z*v4.Y) + float32(v109.Cz.Z*v4.Z),
				}
				v5 = b8
				goto _211
			_211:
				v6 = wB
				v11 = v5
				v12 = Vec3{
					X: v6.X + v11.X,
					Y: v6.Y + v11.Y,
					Z: v6.Z + v11.Z,
				}
				goto _215
			_215:
				wB = v12
			}
			// Central friction
			tangent1 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).Tangent1
			tangent2 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).Tangent2
			// Fixed anchor points for applying impulses
			rA1 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).CenterA
			rB1 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).CenterB
			v4 = wA
			v5 = rA1
			c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
			c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
			c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
			v6 = c
			goto _219
		_219:
			v11 = vA
			v12 = v6
			v13 = Vec3{
				X: v11.X + v12.X,
				Y: v11.Y + v12.Y,
				Z: v11.Z + v12.Z,
			}
			goto _223
		_223:
			// Relative tangent velocity at contact
			vrA1 = v13
			v14 = wB
			v16 = rB1
			c.X = float32(v14.Y*v16.Z) - float32(v14.Z*v16.Y)
			c.Y = float32(v14.Z*v16.X) - float32(v14.X*v16.Z)
			c.Z = float32(v14.X*v16.Y) - float32(v14.Y*v16.X)
			v18 = c
			goto _227
		_227:
			v19 = vB
			v21 = v18
			v22 = Vec3{
				X: v19.X + v21.X,
				Y: v19.Y + v21.Y,
				Z: v19.Z + v21.Z,
			}
			goto _231
		_231:
			vrB1 = v22
			v23 = vrB1
			v25 = vrA1
			v27 = Vec3{
				X: v23.X - v25.X,
				Y: v23.Y - v25.Y,
				Z: v23.Z - v25.Z,
			}
			goto _235
		_235:
			vr = v27
			v29 = vr
			v30 = tangent1
			v17 = float32(v29.X*v30.X) + float32(v29.Y*v30.Y) + float32(v29.Z*v30.Z)
			goto _239
		_239:
			v33 = vr
			v34 = tangent2
			v28 = float32(v33.X*v34.X) + float32(v33.Y*v34.Y) + float32(v33.Z*v34.Z)
			goto _243
		_243:
			vt = Vec2{
				X: v17 - (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).TangentVelocity1,
				Y: v28 - (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).TangentVelocity2,
			}
			v244 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).TangentMass
			v245 = vt
			b10 = Vec2{
				X: float32(v244.Cx.X*v245.X) + float32(v244.Cy.X*v245.Y),
				Y: float32(v244.Cx.Y*v245.X) + float32(v244.Cy.Y*v245.Y),
			}
			v246 = b10
			goto _247
		_247:
			// Incremental tangent impulse
			tm = v246
			deltaImpulse3 = Vec2{
				X: -tm.X,
				Y: -tm.Y,
			}
			newImpulse1 = Vec2{
				X: (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).FrictionImpulse.X + deltaImpulse3.X,
				Y: (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).FrictionImpulse.Y + deltaImpulse3.Y,
			}
			maxImpulse2 = float32(friction * totalNormalImpulse)
			v248 = newImpulse1
			v249 = newImpulse1
			v39 = float32(v248.X*v249.X) + float32(v248.Y*v249.Y)
			goto _251
		_251:
			// Clamp the accumulated impulse
			lengthSquared = v39
			if lengthSquared > float32(maxImpulse2*maxImpulse2) {
				scale = maxImpulse2 / sqrtf(tls, lengthSquared)
				newImpulse1.X = float32(newImpulse1.X * scale)
				newImpulse1.Y = float32(newImpulse1.Y * scale)
			}
			v245 = newImpulse1
			v246 = (*b3ManifoldConstraint)(unsafe.Pointer(constraint)).FrictionImpulse
			c1 = Vec2{
				X: v245.X - v246.X,
				Y: v245.Y - v246.Y,
			}
			v248 = c1
			goto _255
		_255:
			deltaImpulse3 = v248
			(*b3ManifoldConstraint)(unsafe.Pointer(constraint)).FrictionImpulse = newImpulse1
			v17 = deltaImpulse3.X
			v4 = tangent1
			v28 = deltaImpulse3.Y
			v5 = tangent2
			d = Vec3{
				X: float32(v17*v4.X) + float32(v28*v5.X),
				Y: float32(v17*v4.Y) + float32(v28*v5.Y),
				Z: float32(v17*v4.Z) + float32(v28*v5.Z),
			}
			v6 = d
			goto _261
		_261:
			// Apply delta impulse
			P1 = v6
			v11 = vA
			v39 = mA
			v12 = P1
			v13 = Vec3{
				X: v11.X - float32(v39*v12.X),
				Y: v11.Y - float32(v39*v12.Y),
				Z: v11.Z - float32(v39*v12.Z),
			}
			goto _266
		_266:
			vA = v13
			v4 = rA1
			v5 = P1
			c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
			c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
			c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
			v6 = c
			goto _270
		_270:
			v109 = iA
			v11 = v6
			b8 = Vec3{
				X: float32(v109.Cx.X*v11.X) + float32(v109.Cy.X*v11.Y) + float32(v109.Cz.X*v11.Z),
				Y: float32(v109.Cx.Y*v11.X) + float32(v109.Cy.Y*v11.Y) + float32(v109.Cz.Y*v11.Z),
				Z: float32(v109.Cx.Z*v11.X) + float32(v109.Cy.Z*v11.Y) + float32(v109.Cz.Z*v11.Z),
			}
			v12 = b8
			goto _274
		_274:
			v13 = wA
			v14 = v12
			v16 = Vec3{
				X: v13.X - v14.X,
				Y: v13.Y - v14.Y,
				Z: v13.Z - v14.Z,
			}
			goto _278
		_278:
			wA = v16
			v4 = vB
			v17 = mB
			v5 = P1
			v6 = Vec3{
				X: v4.X + float32(v17*v5.X),
				Y: v4.Y + float32(v17*v5.Y),
				Z: v4.Z + float32(v17*v5.Z),
			}
			goto _283
		_283:
			vB = v6
			v4 = rB1
			v5 = P1
			c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
			c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
			c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
			v6 = c
			goto _287
		_287:
			v109 = iB
			v11 = v6
			b8 = Vec3{
				X: float32(v109.Cx.X*v11.X) + float32(v109.Cy.X*v11.Y) + float32(v109.Cz.X*v11.Z),
				Y: float32(v109.Cx.Y*v11.X) + float32(v109.Cy.Y*v11.Y) + float32(v109.Cz.Y*v11.Z),
				Z: float32(v109.Cx.Z*v11.X) + float32(v109.Cy.Z*v11.Y) + float32(v109.Cz.Z*v11.Z),
			}
			v12 = b8
			goto _291
		_291:
			v13 = wB
			v14 = v12
			v16 = Vec3{
				X: v13.X + v14.X,
				Y: v13.Y + v14.Y,
				Z: v13.Z + v14.Z,
			}
			goto _295
		_295:
			wB = v16
			goto _8
		_8:
			;
			j = j + 1
		}
		if (*b3BodyState)(unsafe.Pointer(stateA)).Flags&uint32(b3_dynamicFlag) != 0 {
			(*b3BodyState)(unsafe.Pointer(stateA)).LinearVelocity = vA
			(*b3BodyState)(unsafe.Pointer(stateA)).AngularVelocity = wA
		}
		if (*b3BodyState)(unsafe.Pointer(stateB)).Flags&uint32(b3_dynamicFlag) != 0 {
			(*b3BodyState)(unsafe.Pointer(stateB)).LinearVelocity = vB
			(*b3BodyState)(unsafe.Pointer(stateB)).AngularVelocity = wB
		}
		goto _1
	_1:
		;
		i = i + 1
	}
}

func b3SolveContacts_Overflow(tls *_Stack, context uintptr) {
	var block b3SolverBlock
	var color, graph uintptr
	var count uint16_t
	_, _, _, _ = block, color, count, graph
	graph = (*b3StepContext)(unsafe.Pointer(context)).Graph
	color = graph + uintptr(int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1))*120
	count = uint16FromInt32((*b3GraphColor)(unsafe.Pointer(color)).Contacts.Count)
	if int32FromUint16(count) == 0 {
		return
	}
	block = b3SolverBlock{
		Count:      count,
		BlockType:  uint8(b3_overflowBlock),
		ColorIndex: uint8FromInt32(int32FromInt32(_B3_GRAPH_COLOR_COUNT) - int32FromInt32(1)),
	}
	b3SolveContacts_Mesh(tls, block, context)
}

func b3SolveContacts_Convex(tls *_Stack, block b3SolverBlock, context uintptr) {
	b3SolveContacts_ConvexW4(tls, block, context)
}

func b3SolveContacts_MeshWide(tls *_Stack, block b3SolverBlock, context uintptr) {
	b3SolveContacts_MeshWideW4(tls, block, context)
}

func b3SolveSimplex2(tls *_Stack, simplex uintptr) (r uint8) {
	var a2, ab, b2, v10, v11, v14, v15, v2, v3, v4, v6, v7 Vec3
	var denominator, divisor, u, v, v12, v16, v8 float32
	var vs uintptr
	var v1 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a2, ab, b2, denominator, divisor, u, v, vs, v1, v10, v11, v12, v14, v15, v16, v2, v3, v4, v6, v7, v8
	vs = simplex
	if v1 = !!((*Simplex)(unsafe.Pointer(simplex)).Count == int32FromInt32(2)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+10647, __ccgo_ts+10532, int32FromInt32(179))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	// Vertex regions
	//float wAB[3];
	a2 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).W
	b2 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).W
	v2 = b2
	v3 = a2
	v4 = Vec3{
		X: v2.X - v3.X,
		Y: v2.Y - v3.Y,
		Z: v2.Z - v3.Z,
	}
	goto _5
_5:
	ab = v4
	v6 = ab
	v7 = ab
	v8 = float32(v6.X*v7.X) + float32(v6.Y*v7.Y) + float32(v6.Z*v7.Z)
	goto _9
_9:
	// Last element is divisor
	divisor = v8
	v10 = b2
	v11 = ab
	v12 = float32(v10.X*v11.X) + float32(v10.Y*v11.Y) + float32(v10.Z*v11.Z)
	goto _13
_13:
	u = v12
	v14 = a2
	v15 = ab
	v16 = float32(v14.X*v15.X) + float32(v14.Y*v15.Y) + float32(v14.Z*v15.Z)
	goto _17
_17:
	v = -v16
	//wAB[2] = divisor;
	// V( A )
	if v <= float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(1)
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = float32FromFloat32(1)
		return boolUint8(true1 != 0)
	}
	// V( B )
	if u <= float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(1)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = *(*SimplexVertex)(unsafe.Add(nil, vs+1*48))
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = float32FromFloat32(1)
		return boolUint8(true1 != 0)
	}
	// Edge region
	if divisor <= float32FromFloat32(0) {
		return boolUint8(false1 != 0)
	}
	// VR( AB )
	denominator = float32FromFloat32(1) / divisor
	(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = float32(denominator * u)
	(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = float32(denominator * v)
	return boolUint8(true1 != 0)
}

func b3SolveSimplex3(tls *_Stack, simplex uintptr) (r uint8) {
	bp := tls.Alloc(64)
	defer tls.Free(64)
	var divisor, divisor1, divisor2, divisor3 float32
	var v1, v2, v3 SimplexVertex
	var vs uintptr
	var v11 bool
	var _ /* wAB at bp+0 */ [3]float32
	var _ /* wABC at bp+36 */ [4]float32
	var _ /* wBC at bp+12 */ [3]float32
	var _ /* wCA at bp+24 */ [3]float32
	_, _, _, _, _, _, _, _, _ = divisor, divisor1, divisor2, divisor3, v1, v2, v3, vs, v11
	vs = simplex
	if v11 = !!((*Simplex)(unsafe.Pointer(simplex)).Count == int32FromInt32(3)); !v11 {
		b3InternalAssert(tls, __ccgo_ts+10667, __ccgo_ts+10532, int32FromInt32(233))
	}
	_ = v11 || bool(int32FromInt32(0) != 0)
	// Get simplex (be aware of aliasing here!)
	v1 = *(*SimplexVertex)(unsafe.Add(nil, vs))
	v2 = *(*SimplexVertex)(unsafe.Add(nil, vs+1*48))
	v3 = *(*SimplexVertex)(unsafe.Add(nil, vs+2*48))
	b3BarycentricCoordsEdge(tls, bp, v1.W, v2.W)
	b3BarycentricCoordsEdge(tls, bp+12, v2.W, v3.W)
	b3BarycentricCoordsEdge(tls, bp+24, v3.W, v1.W)
	// VR( A )
	if (*(*[3]float32)(unsafe.Add(nil, bp)))[int32(1)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+24)))[0] <= float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(1)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = v1
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = float32FromFloat32(1)
		return boolUint8(true1 != 0)
	}
	// VR( B )
	if (*(*[3]float32)(unsafe.Add(nil, bp+12)))[int32(1)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp)))[0] <= float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(1)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = v2
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = float32FromFloat32(1)
		return boolUint8(true1 != 0)
	}
	// VR( C )
	if (*(*[3]float32)(unsafe.Add(nil, bp+24)))[int32(1)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+12)))[0] <= float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(1)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = v3
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = float32FromFloat32(1)
		return boolUint8(true1 != 0)
	}
	b3BarycentricCoordsTri(tls, bp+36, v1.W, v2.W, v3.W)
	// VR( AB )
	if (*(*[4]float32)(unsafe.Add(nil, bp+36)))[int32(2)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp)))[0] > float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp)))[int32(1)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(2)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = v1
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = v2
		// Normalize
		divisor = (*(*[3]float32)(unsafe.Add(nil, bp)))[int32(2)]
		if divisor <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[3]float32)(unsafe.Add(nil, bp)))[0] / divisor
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[3]float32)(unsafe.Add(nil, bp)))[int32(1)] / divisor
		return boolUint8(true1 != 0)
	}
	// VR( BC )
	if (*(*[4]float32)(unsafe.Add(nil, bp+36)))[0] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+12)))[0] > float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+12)))[int32(1)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(2)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = v2
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = v3
		// Normalize
		divisor1 = (*(*[3]float32)(unsafe.Add(nil, bp+12)))[int32(2)]
		if divisor1 <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[3]float32)(unsafe.Add(nil, bp+12)))[0] / divisor1
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[3]float32)(unsafe.Add(nil, bp+12)))[int32(1)] / divisor1
		return boolUint8(true1 != 0)
	}
	// VR( CA )
	if (*(*[4]float32)(unsafe.Add(nil, bp+36)))[int32(1)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+24)))[0] > float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+24)))[int32(1)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(2)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = v3
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = v1
		// Normalize
		divisor2 = (*(*[3]float32)(unsafe.Add(nil, bp+24)))[int32(2)]
		if divisor2 <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[3]float32)(unsafe.Add(nil, bp+24)))[0] / divisor2
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[3]float32)(unsafe.Add(nil, bp+24)))[int32(1)] / divisor2
		return boolUint8(true1 != 0)
	}
	// Face region
	divisor3 = (*(*[4]float32)(unsafe.Add(nil, bp+36)))[int32(3)]
	if divisor3 <= float32FromFloat32(0) {
		return boolUint8(false1 != 0)
	}
	// VR( ABC )
	(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[4]float32)(unsafe.Add(nil, bp+36)))[0] / divisor3
	(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[4]float32)(unsafe.Add(nil, bp+36)))[int32(1)] / divisor3
	(*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).A = (*(*[4]float32)(unsafe.Add(nil, bp+36)))[int32(2)] / divisor3
	return boolUint8(true1 != 0)
}

func b3SolveSimplex4(tls *_Stack, simplex uintptr) (r uint8) {
	bp := tls.Alloc(160)
	defer tls.Free(160)
	var divisor, divisor1, divisor10, divisor2, divisor3, divisor4, divisor5, divisor6, divisor7, divisor8, divisor9 float32
	var vertexA, vertexB, vertexC, vertexD SimplexVertex
	var vs uintptr
	var v1 bool
	var _ /* wAB at bp+0 */ [3]float32
	var _ /* wABCD at bp+136 */ [5]float32
	var _ /* wABD at bp+88 */ [4]float32
	var _ /* wAC at bp+12 */ [3]float32
	var _ /* wACB at bp+72 */ [4]float32
	var _ /* wAD at bp+24 */ [3]float32
	var _ /* wADC at bp+104 */ [4]float32
	var _ /* wBC at bp+36 */ [3]float32
	var _ /* wBCD at bp+120 */ [4]float32
	var _ /* wCD at bp+48 */ [3]float32
	var _ /* wDB at bp+60 */ [3]float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = divisor, divisor1, divisor10, divisor2, divisor3, divisor4, divisor5, divisor6, divisor7, divisor8, divisor9, vertexA, vertexB, vertexC, vertexD, vs, v1
	vs = simplex
	// Get simplex (be aware of aliasing here!)
	if v1 = !!((*Simplex)(unsafe.Pointer(simplex)).Count == int32FromInt32(4)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+10687, __ccgo_ts+10532, int32FromInt32(366))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	vertexA = *(*SimplexVertex)(unsafe.Add(nil, vs))
	vertexB = *(*SimplexVertex)(unsafe.Add(nil, vs+1*48))
	vertexC = *(*SimplexVertex)(unsafe.Add(nil, vs+2*48))
	vertexD = *(*SimplexVertex)(unsafe.Add(nil, vs+3*48))
	b3BarycentricCoordsEdge(tls, bp, vertexA.W, vertexB.W)
	b3BarycentricCoordsEdge(tls, bp+12, vertexA.W, vertexC.W)
	b3BarycentricCoordsEdge(tls, bp+24, vertexA.W, vertexD.W)
	b3BarycentricCoordsEdge(tls, bp+36, vertexB.W, vertexC.W)
	b3BarycentricCoordsEdge(tls, bp+48, vertexC.W, vertexD.W)
	b3BarycentricCoordsEdge(tls, bp+60, vertexD.W, vertexB.W)
	// VR( A )
	if (*(*[3]float32)(unsafe.Add(nil, bp)))[int32(1)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+12)))[int32(1)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+24)))[int32(1)] <= float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(1)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexA
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = float32FromFloat32(1)
		return boolUint8(true1 != 0)
	}
	// VR( B )
	if (*(*[3]float32)(unsafe.Add(nil, bp)))[0] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+60)))[0] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+36)))[int32(1)] <= float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(1)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexB
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = float32FromFloat32(1)
		return boolUint8(true1 != 0)
	}
	// VR( C )
	if (*(*[3]float32)(unsafe.Add(nil, bp+12)))[0] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+36)))[0] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+48)))[int32(1)] <= float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(1)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexC
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = float32FromFloat32(1)
		return boolUint8(true1 != 0)
	}
	// VR( D )
	if (*(*[3]float32)(unsafe.Add(nil, bp+24)))[0] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+48)))[0] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+60)))[int32(1)] <= float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(1)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexD
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = float32FromFloat32(1)
		return boolUint8(true1 != 0)
	}
	b3BarycentricCoordsTri(tls, bp+72, vertexA.W, vertexC.W, vertexB.W)
	b3BarycentricCoordsTri(tls, bp+88, vertexA.W, vertexB.W, vertexD.W)
	b3BarycentricCoordsTri(tls, bp+104, vertexA.W, vertexD.W, vertexC.W)
	b3BarycentricCoordsTri(tls, bp+120, vertexB.W, vertexC.W, vertexD.W)
	// VR( AB )
	if (*(*[4]float32)(unsafe.Add(nil, bp+88)))[int32(2)] <= float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+72)))[int32(1)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp)))[0] > float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp)))[int32(1)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(2)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexA
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = vertexB
		// Normalize
		divisor = (*(*[3]float32)(unsafe.Add(nil, bp)))[int32(2)]
		if divisor <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[3]float32)(unsafe.Add(nil, bp)))[0] / divisor
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[3]float32)(unsafe.Add(nil, bp)))[int32(1)] / divisor
		return boolUint8(true1 != 0)
	}
	// VR( AC )
	if (*(*[4]float32)(unsafe.Add(nil, bp+72)))[int32(2)] <= float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+104)))[int32(1)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+12)))[0] > float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+12)))[int32(1)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(2)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexA
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = vertexC
		// Normalize
		divisor1 = (*(*[3]float32)(unsafe.Add(nil, bp+12)))[int32(2)]
		if divisor1 <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[3]float32)(unsafe.Add(nil, bp+12)))[0] / divisor1
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[3]float32)(unsafe.Add(nil, bp+12)))[int32(1)] / divisor1
		return boolUint8(true1 != 0)
	}
	// VR( AD )
	if (*(*[4]float32)(unsafe.Add(nil, bp+104)))[int32(2)] <= float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+88)))[int32(1)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+24)))[0] > float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+24)))[int32(1)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(2)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexA
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = vertexD
		// Normalize
		divisor2 = (*(*[3]float32)(unsafe.Add(nil, bp+24)))[int32(2)]
		if divisor2 <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[3]float32)(unsafe.Add(nil, bp+24)))[0] / divisor2
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[3]float32)(unsafe.Add(nil, bp+24)))[int32(1)] / divisor2
		return boolUint8(true1 != 0)
	}
	// VR( BC )
	if (*(*[4]float32)(unsafe.Add(nil, bp+72)))[0] <= float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+120)))[int32(2)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+36)))[0] > float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+36)))[int32(1)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(2)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexB
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = vertexC
		// Normalize
		divisor3 = (*(*[3]float32)(unsafe.Add(nil, bp+36)))[int32(2)]
		if divisor3 <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[3]float32)(unsafe.Add(nil, bp+36)))[0] / divisor3
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[3]float32)(unsafe.Add(nil, bp+36)))[int32(1)] / divisor3
		return boolUint8(true1 != 0)
	}
	// VR( CD )
	if (*(*[4]float32)(unsafe.Add(nil, bp+104)))[0] <= float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+120)))[0] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+48)))[0] > float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+48)))[int32(1)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(2)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexC
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = vertexD
		// Normalize
		divisor4 = (*(*[3]float32)(unsafe.Add(nil, bp+48)))[int32(2)]
		if divisor4 <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[3]float32)(unsafe.Add(nil, bp+48)))[0] / divisor4
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[3]float32)(unsafe.Add(nil, bp+48)))[int32(1)] / divisor4
		return boolUint8(true1 != 0)
	}
	// VR( DB )
	if (*(*[4]float32)(unsafe.Add(nil, bp+88)))[0] <= float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+120)))[int32(1)] <= float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+60)))[0] > float32FromFloat32(0) && (*(*[3]float32)(unsafe.Add(nil, bp+60)))[int32(1)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(2)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexD
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = vertexB
		// Normalize
		divisor5 = (*(*[3]float32)(unsafe.Add(nil, bp+60)))[int32(2)]
		if divisor5 <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[3]float32)(unsafe.Add(nil, bp+60)))[0] / divisor5
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[3]float32)(unsafe.Add(nil, bp+60)))[int32(1)] / divisor5
		return boolUint8(true1 != 0)
	}
	b3BarycentricCoordsTet(tls, bp+136, vertexA.W, vertexB.W, vertexC.W, vertexD.W)
	// VR( ACB )
	if (*(*[5]float32)(unsafe.Add(nil, bp+136)))[int32(3)] < float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+72)))[0] > float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+72)))[int32(1)] > float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+72)))[int32(2)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(3)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexA
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = vertexC
		*(*SimplexVertex)(unsafe.Add(nil, vs+2*48)) = vertexB
		// Normalize
		divisor6 = (*(*[4]float32)(unsafe.Add(nil, bp+72)))[int32(3)]
		if divisor6 <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[4]float32)(unsafe.Add(nil, bp+72)))[0] / divisor6
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[4]float32)(unsafe.Add(nil, bp+72)))[int32(1)] / divisor6
		(*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).A = (*(*[4]float32)(unsafe.Add(nil, bp+72)))[int32(2)] / divisor6
		return boolUint8(true1 != 0)
	}
	// VR( ABD )
	if (*(*[5]float32)(unsafe.Add(nil, bp+136)))[int32(2)] < float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+88)))[0] > float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+88)))[int32(1)] > float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+88)))[int32(2)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(3)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexA
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = vertexB
		*(*SimplexVertex)(unsafe.Add(nil, vs+2*48)) = vertexD
		// Normalize
		divisor7 = (*(*[4]float32)(unsafe.Add(nil, bp+88)))[int32(3)]
		if divisor7 <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[4]float32)(unsafe.Add(nil, bp+88)))[0] / divisor7
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[4]float32)(unsafe.Add(nil, bp+88)))[int32(1)] / divisor7
		(*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).A = (*(*[4]float32)(unsafe.Add(nil, bp+88)))[int32(2)] / divisor7
		return boolUint8(true1 != 0)
	}
	// VR( ADC )
	if (*(*[5]float32)(unsafe.Add(nil, bp+136)))[int32(1)] < float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+104)))[0] > float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+104)))[int32(1)] > float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+104)))[int32(2)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(3)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexA
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = vertexD
		*(*SimplexVertex)(unsafe.Add(nil, vs+2*48)) = vertexC
		// Normalize
		divisor8 = (*(*[4]float32)(unsafe.Add(nil, bp+104)))[int32(3)]
		if divisor8 <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[4]float32)(unsafe.Add(nil, bp+104)))[0] / divisor8
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[4]float32)(unsafe.Add(nil, bp+104)))[int32(1)] / divisor8
		(*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).A = (*(*[4]float32)(unsafe.Add(nil, bp+104)))[int32(2)] / divisor8
		return boolUint8(true1 != 0)
	}
	// VR( BCD )
	if (*(*[5]float32)(unsafe.Add(nil, bp+136)))[0] < float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+120)))[0] > float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+120)))[int32(1)] > float32FromFloat32(0) && (*(*[4]float32)(unsafe.Add(nil, bp+120)))[int32(2)] > float32FromFloat32(0) {
		// Reduce simplex
		(*Simplex)(unsafe.Pointer(simplex)).Count = int32(3)
		*(*SimplexVertex)(unsafe.Add(nil, vs)) = vertexB
		*(*SimplexVertex)(unsafe.Add(nil, vs+1*48)) = vertexC
		*(*SimplexVertex)(unsafe.Add(nil, vs+2*48)) = vertexD
		// Normalize
		divisor9 = (*(*[4]float32)(unsafe.Add(nil, bp+120)))[int32(3)]
		if divisor9 <= float32FromFloat32(0) {
			return boolUint8(false1 != 0)
		}
		(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[4]float32)(unsafe.Add(nil, bp+120)))[0] / divisor9
		(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[4]float32)(unsafe.Add(nil, bp+120)))[int32(1)] / divisor9
		(*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).A = (*(*[4]float32)(unsafe.Add(nil, bp+120)))[int32(2)] / divisor9
		return boolUint8(true1 != 0)
	}
	// *** Inside tetrahedron ***
	divisor10 = (*(*[5]float32)(unsafe.Add(nil, bp+136)))[int32(4)]
	if divisor10 <= float32FromFloat32(0) {
		return boolUint8(false1 != 0)
	}
	// VR( ABCD )
	(*(*SimplexVertex)(unsafe.Add(nil, vs))).A = (*(*[5]float32)(unsafe.Add(nil, bp+136)))[0] / divisor10
	(*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A = (*(*[5]float32)(unsafe.Add(nil, bp+136)))[int32(1)] / divisor10
	(*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).A = (*(*[5]float32)(unsafe.Add(nil, bp+136)))[int32(2)] / divisor10
	(*(*SimplexVertex)(unsafe.Add(nil, vs+3*48))).A = (*(*[5]float32)(unsafe.Add(nil, bp+136)))[int32(3)] / divisor10
	return boolUint8(true1 != 0)
}

// Used to detect bad values. In float mode positions greater than about 16km have precision
// problems, so 100km is a safe limit. Large world mode keeps coordinates accurate much farther
// from the origin, so the sanity limit widens to keep valid far-field positions from tripping it.

/// Maximum parallel workers. Used for some fixed size arrays.

/// Maximum number of tasks queued per world step. b3EnqueueTaskCallback will never be called
/// more than this per world step. This is related to B3_MAX_WORKERS. With 32 workers,
/// the maximum observed task count is 130. This allows an external task system to use a fixed
/// size array for Box3D task, which may help with creating stable user task pointers.

// Maximum number of colors in the constraint graph. Constraints that cannot
// find a color are added to the overflow set which are solved single-threaded.
// The compound barrel benchmark has minor overflow with 24 colors

// Number of contact point buckets for counting the number of contact points per
// shape contact pair. This is just for reporting and doesn't affect simulation.

// A small length used as a collision and constraint tolerance. Usually it is
// chosen to be numerically significant, but visually insignificant. In meters.
// @warning modifying this can have a significant impact on stability

/// The minimum length of a capsules. Very short capsules should be created as spheres
/// to avoid numerical problems.

/// Minimum contact point friction weight, lower bound for speculative points. Made small
/// enough to be washed away by weights that hit 1.

/// The distance between shapes where they are considered overlapped. This is needed
/// because GJK may return small positive values for overlapped shapes in degenerate
/// configurations.

/// Maximum number of simultaneous worlds that can be allocated

/// The maximum rotation of a body per time step. This limit is very large and is used
/// to prevent numerical problems. You shouldn't need to adjust this.
/// @warning increasing this to 0.5f * B3_PI or greater will break continuous collision.

/// @warning modifying this can have a significant impact on performance and stability

/// The rest offset is used for mesh contact to reduce ghost collisions and assist with CCD.
/// The rest offset adjusts the contact point separation value, making the solver push the shapes
/// apart by this distance.
/// Must be at least B3_LINEAR_SLOP and less than B3_SPECULATIVE_DISTANCE.

/// The default contact recycling distance.

/// The default contact recycling world angle threshold. For performance this value
/// is cos(angle/2)^2. This value corresponds to 10 degrees.

/// This is used to fatten AABBs in the dynamic tree. This allows proxies
/// to move by a small amount without triggering a tree adjustment. This is in meters.
/// @warning modifying this can have a significant impact on performance

/// Per-shape AABB margin is a fraction of the shape extent (capped by B3_MAX_AABB_MARGIN).
/// Small shapes get small margins; large shapes are clamped to the cap.

/// The time that a body must be still before it will go to sleep. In seconds.

/// The maximum number of restitution iterations. Needed to avoid a solver overflow.

/// The maximum number of contact points between two touching shapes.
/// The default and minimum is 4 and this case uses a fast approximate hull to reduce the
/// point count. A larger value builds a 2D convex hull of the candidate points and then
/// simplifies it to the target by dropping the vertices that contribute the least area.
/// Useful for testing, but I don't recommend shipping with it adjusted.
/// WARNING: if you change this you risk breaking ABI and a corrupted runtime.

/// The number of iterations for gyroscopic torques.

/// The maximum number of convex hull vertices. This is fixed for performance reasons.

/// The maximum number of convex hull faces.

/// The maximum number of convex hull edges. Full edges, not half-edges.

/// Relative tolerance used to determine if two edges are parallel.

/// The maximum number points to use for shape cast proxies (swept point cloud).

/// These generous limits allow for easy hashing. See b3ShapePairKey.

/// This is the limit on how many mesh or heightfield triangles a single convex shape can collide with.
/// Increasing this will increase stack usage, so be careful. I recommend to simplify your collision data
/// before increasing this. For example, using render mesh for collision often leads to poor performance.

func b3SolvePlanes(tls *_Stack, targetDelta Vec3, planes uintptr, count int32) (r PlaneSolverResult) {
	var accumulatedPush, push, separation, tolerance, totalPush, v11, v12, v13, v14, v16, v17, v5, v9 float32
	var delta, v18, v7, v8 Vec3
	var i, iteration, planeIndex int32
	var plane1 uintptr
	var v4 Plane
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = accumulatedPush, delta, i, iteration, plane1, planeIndex, push, separation, tolerance, totalPush, v11, v12, v13, v14, v16, v17, v18, v4, v5, v7, v8, v9
	i = 0
	for {
		if !(i < count) {
			break
		}
		(*(*CollisionPlane)(unsafe.Add(nil, planes+uintptr(i)*28))).Push = float32FromFloat32(0)
		goto _1
	_1:
		;
		i = i + 1
	}
	delta = targetDelta
	tolerance = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	iteration = 0
	for {
		if !(iteration < int32(20)) {
			break
		}
		totalPush = float32FromFloat32(0)
		planeIndex = 0
		for {
			if !(planeIndex < count) {
				break
			}
			plane1 = planes + uintptr(planeIndex)*28
			v4 = (*CollisionPlane)(unsafe.Pointer(plane1)).Plane
			v7 = v4.Normal
			v8 = delta
			v9 = float32(v7.X*v8.X) + float32(v7.Y*v8.Y) + float32(v7.Z*v8.Z)
			goto _10
		_10:
			v5 = v9 - v4.Offset
			goto _6
		_6:
			// Add slop to prevent jitter
			separation = v5 + float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls))
			push = -separation
			// Clamp accumulated push
			accumulatedPush = (*CollisionPlane)(unsafe.Pointer(plane1)).Push
			v11 = (*CollisionPlane)(unsafe.Pointer(plane1)).Push + push
			v12 = float32FromFloat32(0)
			v13 = (*CollisionPlane)(unsafe.Pointer(plane1)).PushLimit
			if v11 < v12 {
				v16 = v12
			} else {
				if v13 < v11 {
					v17 = v13
				} else {
					v17 = v11
				}
				v16 = v17
			}
			v14 = v16
			goto _15
		_15:
			(*CollisionPlane)(unsafe.Pointer(plane1)).Push = v14
			push = (*CollisionPlane)(unsafe.Pointer(plane1)).Push - accumulatedPush
			v7 = delta
			v5 = push
			v8 = (*CollisionPlane)(unsafe.Pointer(plane1)).Plane.Normal
			v18 = Vec3{
				X: v7.X + float32(v5*v8.X),
				Y: v7.Y + float32(v5*v8.Y),
				Z: v7.Z + float32(v5*v8.Z),
			}
			goto _22
		_22:
			delta = v18
			// Track maximum push for convergence
			v5 = push
			if v5 < float32FromInt32(0) {
				v11 = -v5
			} else {
				v11 = v5
			}
			v9 = v11
			goto _25
		_25:
			totalPush = totalPush + v9
			goto _3
		_3:
			;
			planeIndex = planeIndex + 1
		}
		if totalPush < tolerance {
			break
		}
		goto _2
	_2:
		;
		iteration = iteration + 1
	}
	return PlaneSolverResult{
		Delta:          delta,
		IterationCount: iteration,
	}
}

func b3ValidateSolverSets(tls *_Stack, world uintptr) {
	_ = uint64FromInt64(4)
}

// C documentation
//
//	// Continuous collision of dynamic versus static
func b3SolveContinuous(tls *_Stack, world uintptr, bodySimIndex int32, taskContext uintptr, dt float32) {
	bp := tls.Alloc(256)
	defer tls.Free(256)
	var aabb1, box1, box2, out, out1, out2, sweptBox, v156, v170, v172, v173 AABB
	var awakeSet, dynamicTree, event, fastBody, fastBodySim, fastBodyState, fastShape, kinematicTree, shape, shape1, shapeFatAABB, shapeFatAABB1, staticTree, v31 uintptr
	var base1, center, v10, v4, v5, v6, v9 Pos
	var c, c1, c2, dv, origin1, radius, rv, t1, t11, t2, t21, t3, t31, v3, v11, v16, v17, v18, v19, v21, v23, v24, v26, v27, v28, v30, v32, v34, v35, v37, v38, v39, v42, v43, v44, v45, v47, v49, v7 Vec3
	var i, newCapacity, newSize, oldSize, proxyId, shapeId, v306, v324, v326, v327 int32
	var isBullet, isShapeFast, v292 uint8
	var lengthSq, rotationChord, s3, safetyFactor, speculativeScalar, timeLoss, v106, v117, v135, v146, v158, v160, v162, v164, v166, v168, v178, v179, v180, v182, v183, v22, v33, v48, v59, v75, v80, v86, v91, v95 float32
	var proxyType b3BodyType1
	var q1, q3, q4, qn, v128, v15, v214, v215, v217, v41, v68, v69, v88, v99 Quat
	var s4, sweep, v13 Sweep
	var transform WorldTransform
	var xf1, xf2, v98 Transform
	var v1, v2 bool
	var _ /* context at bp+0 */ b3ContinuousContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb1, awakeSet, base1, box1, box2, c, c1, c2, center, dv, dynamicTree, event, fastBody, fastBodySim, fastBodyState, fastShape, i, isBullet, isShapeFast, kinematicTree, lengthSq, newCapacity, newSize, oldSize, origin1, out, out1, out2, proxyId, proxyType, q1, q3, q4, qn, radius, rotationChord, rv, s3, s4, safetyFactor, shape, shape1, shapeFatAABB, shapeFatAABB1, shapeId, speculativeScalar, staticTree, sweep, sweptBox, t1, t11, t2, t21, t3, t31, timeLoss, transform, v3, xf1, xf2, v1, v10, v106, v11, v117, v128, v13, v135, v146, v15, v156, v158, v16, v160, v162, v164, v166, v168, v17, v170, v172, v173, v178, v179, v18, v180, v182, v183, v19, v2, v21, v214, v215, v217, v22, v23, v24, v26, v27, v28, v292, v31, v30, v306, v32, v324, v326, v327, v33, v34, v35, v37, v38, v39, v4, v41, v42, v43, v44, v45, v47, v48, v49, v5, v59, v6, v68, v69, v7, v75, v80, v86, v88, v9, v91, v95, v98, v99
	if v1 = !!(bool(0 <= int32(b3_awakeSet)) && int32(b3_awakeSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+802, __ccgo_ts+31091, int32FromInt32(501))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_awakeSet))*88
	if v2 = !!(0 <= bodySimIndex && bodySimIndex < (*b3SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+31387, __ccgo_ts+31091, int32FromInt32(502))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	fastBodySim = (*b3SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Data + uintptr(bodySimIndex)*252
	// Re-center the sweep on the fast body so the TOI and the swept query stay in float precision
	base1 = (*b3BodySim)(unsafe.Pointer(fastBodySim)).Center0
	v31 = fastBodySim
	v4 = base1
	v5 = (*b3BodySim)(unsafe.Pointer(v31)).Center0
	v6 = v4
	v7 = Vec3{
		X: v5.X - v6.X,
		Y: v5.Y - v6.Y,
		Z: v5.Z - v6.Z,
	}
	goto _8
_8:
	s4.C1 = v7
	v9 = (*b3BodySim)(unsafe.Pointer(v31)).Center
	v10 = v4
	v11 = Vec3{
		X: v9.X - v10.X,
		Y: v9.Y - v10.Y,
		Z: v9.Z - v10.Z,
	}
	goto _12
_12:
	s4.C2 = v11
	s4.Q1 = (*b3BodySim)(unsafe.Pointer(v31)).Rotation0
	s4.Q2 = (*b3BodySim)(unsafe.Pointer(v31)).Transform.Q
	s4.LocalCenter = (*b3BodySim)(unsafe.Pointer(v31)).LocalCenter
	v13 = s4
	goto _14
_14:
	sweep = v13
	xf1.Q = sweep.Q1
	v15 = sweep.Q1
	v7 = sweep.LocalCenter
	v11 = v15.V
	v16 = v7
	c.X = float32(v11.Y*v16.Z) - float32(v11.Z*v16.Y)
	c.Y = float32(v11.Z*v16.X) - float32(v11.X*v16.Z)
	c.Z = float32(v11.X*v16.Y) - float32(v11.Y*v16.X)
	v17 = c
	goto _20
_20:
	t1 = v17
	v18 = t1
	v22 = v15.S
	v19 = v7
	v21 = Vec3{
		X: v18.X + float32(v22*v19.X),
		Y: v18.Y + float32(v22*v19.Y),
		Z: v18.Z + float32(v22*v19.Z),
	}
	goto _25
_25:
	t2 = v21
	v23 = v15.V
	v24 = t2
	c.X = float32(v23.Y*v24.Z) - float32(v23.Z*v24.Y)
	c.Y = float32(v23.Z*v24.X) - float32(v23.X*v24.Z)
	c.Z = float32(v23.X*v24.Y) - float32(v23.Y*v24.X)
	v26 = c
	goto _29
_29:
	t3 = v26
	v28 = v7
	v33 = float32FromFloat32(2)
	v30 = t3
	v32 = Vec3{
		X: v28.X + float32(v33*v30.X),
		Y: v28.Y + float32(v33*v30.Y),
		Z: v28.Z + float32(v33*v30.Z),
	}
	goto _36
_36:
	v27 = v32
	goto _31
_31:
	v34 = sweep.C1
	v35 = v27
	v37 = Vec3{
		X: v34.X - v35.X,
		Y: v34.Y - v35.Y,
		Z: v34.Z - v35.Z,
	}
	goto _40
_40:
	xf1.P = v37
	xf2.Q = sweep.Q2
	v15 = sweep.Q2
	v7 = sweep.LocalCenter
	v11 = v15.V
	v16 = v7
	c.X = float32(v11.Y*v16.Z) - float32(v11.Z*v16.Y)
	c.Y = float32(v11.Z*v16.X) - float32(v11.X*v16.Z)
	c.Z = float32(v11.X*v16.Y) - float32(v11.Y*v16.X)
	v17 = c
	goto _46
_46:
	t1 = v17
	v18 = t1
	v22 = v15.S
	v19 = v7
	v21 = Vec3{
		X: v18.X + float32(v22*v19.X),
		Y: v18.Y + float32(v22*v19.Y),
		Z: v18.Z + float32(v22*v19.Z),
	}
	goto _51
_51:
	t2 = v21
	v23 = v15.V
	v24 = t2
	c.X = float32(v23.Y*v24.Z) - float32(v23.Z*v24.Y)
	c.Y = float32(v23.Z*v24.X) - float32(v23.X*v24.Z)
	c.Z = float32(v23.X*v24.Y) - float32(v23.Y*v24.X)
	v26 = c
	goto _55
_55:
	t3 = v26
	v28 = v7
	v33 = float32FromFloat32(2)
	v30 = t3
	v32 = Vec3{
		X: v28.X + float32(v33*v30.X),
		Y: v28.Y + float32(v33*v30.Y),
		Z: v28.Z + float32(v33*v30.Z),
	}
	goto _62
_62:
	v27 = v32
	goto _57
_57:
	v34 = sweep.C2
	v35 = v27
	v37 = Vec3{
		X: v34.X - v35.X,
		Y: v34.Y - v35.Y,
		Z: v34.Z - v35.Z,
	}
	goto _66
_66:
	xf2.P = v37
	staticTree = world + 800 + uintptr(b3_staticBody)*112
	kinematicTree = world + 800 + uintptr(b3_kinematicBody)*112
	dynamicTree = world + 800 + uintptr(b3_dynamicBody)*112
	if v1 = !!(0 <= (*b3BodySim)(unsafe.Pointer(fastBodySim)).BodyId && (*b3BodySim)(unsafe.Pointer(fastBodySim)).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+31305, __ccgo_ts+31091, int32FromInt32(521))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	fastBody = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3BodySim)(unsafe.Pointer(fastBodySim)).BodyId)*96
	*(*b3ContinuousContext)(unsafe.Add(nil, bp)) = b3ContinuousContext{}
	(*(*b3ContinuousContext)(unsafe.Add(nil, bp))).World = world
	(*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Sweep = sweep
	(*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Base = base1
	(*(*b3ContinuousContext)(unsafe.Add(nil, bp))).FastBodySim = fastBodySim
	(*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Fraction = float32FromFloat32(1)
	isBullet = boolUint8((*b3BodySim)(unsafe.Pointer(fastBodySim)).Flags&uint32(b3_isBullet) != uint32(0))
	safetyFactor = (*b3Body)(unsafe.Pointer(fastBody)).SafetyFactor
	v15 = sweep.Q1
	v41 = sweep.Q2
	v7 = v41.V
	v11 = v15.V
	c.X = float32(v7.Y*v11.Z) - float32(v7.Z*v11.Y)
	c.Y = float32(v7.Z*v11.X) - float32(v7.X*v11.Z)
	c.Z = float32(v7.X*v11.Y) - float32(v7.Y*v11.X)
	v16 = c
	goto _73
_73:
	t11 = v16
	v17 = t11
	v22 = v15.S
	v18 = v41.V
	v19 = Vec3{
		X: v17.X + float32(v22*v18.X),
		Y: v17.Y + float32(v22*v18.Y),
		Z: v17.Z + float32(v22*v18.Z),
	}
	goto _78
_78:
	t21 = v19
	v21 = t21
	v33 = v41.S
	v23 = v15.V
	v24 = Vec3{
		X: v21.X - float32(v33*v23.X),
		Y: v21.Y - float32(v33*v23.Y),
		Z: v21.Z - float32(v33*v23.Z),
	}
	goto _83
_83:
	t31 = v24
	v26 = v15.V
	v27 = v41.V
	v48 = float32(v26.X*v27.X) + float32(v26.Y*v27.Y) + float32(v26.Z*v27.Z)
	goto _87
_87:
	q1 = Quat{
		V: t31,
		S: float32(v15.S*v41.S) + v48,
	}
	v68 = q1
	goto _89
_89:
	v28 = v68.V
	v30 = v28
	v32 = v28
	v75 = float32(v30.X*v32.X) + float32(v30.Y*v32.Y) + float32(v30.Z*v32.Z)
	goto _96
_96:
	v59 = sqrtf(tls, v75)
	goto _92
_92:
	rotationChord = float32(float32FromFloat32(2) * v59)
	shapeId = (*b3Body)(unsafe.Pointer(fastBody)).HeadShapeId
	for shapeId != -int32(1) {
		if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+31091, int32FromInt32(537))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		fastShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
		shapeId = (*b3Shape)(unsafe.Pointer(fastShape)).NextShapeId
		(*(*b3ContinuousContext)(unsafe.Add(nil, bp))).FastShape = fastShape
		v98 = xf1
		v15 = v98.Q
		v7 = (*b3Shape)(unsafe.Pointer(fastShape)).LocalCentroid
		v11 = v15.V
		v16 = v7
		c.X = float32(v11.Y*v16.Z) - float32(v11.Z*v16.Y)
		c.Y = float32(v11.Z*v16.X) - float32(v11.X*v16.Z)
		c.Z = float32(v11.X*v16.Y) - float32(v11.Y*v16.X)
		v17 = c
		goto _104
	_104:
		t1 = v17
		v18 = t1
		v22 = v15.S
		v19 = v7
		v21 = Vec3{
			X: v18.X + float32(v22*v19.X),
			Y: v18.Y + float32(v22*v19.Y),
			Z: v18.Z + float32(v22*v19.Z),
		}
		goto _109
	_109:
		t2 = v21
		v23 = v15.V
		v24 = t2
		c.X = float32(v23.Y*v24.Z) - float32(v23.Z*v24.Y)
		c.Y = float32(v23.Z*v24.X) - float32(v23.X*v24.Z)
		c.Z = float32(v23.X*v24.Y) - float32(v23.Y*v24.X)
		v26 = c
		goto _113
	_113:
		t3 = v26
		v28 = v7
		v33 = float32FromFloat32(2)
		v30 = t3
		v32 = Vec3{
			X: v28.X + float32(v33*v30.X),
			Y: v28.Y + float32(v33*v30.Y),
			Z: v28.Z + float32(v33*v30.Z),
		}
		goto _120
	_120:
		v27 = v32
		goto _115
	_115:
		rv = v27
		v35 = rv
		v37 = v98.P
		v38 = Vec3{
			X: v35.X + v37.X,
			Y: v35.Y + v37.Y,
			Z: v35.Z + v37.Z,
		}
		goto _126
	_126:
		v34 = v38
		goto _122
	_122:
		(*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Centroid1 = v34
		v98 = xf2
		v15 = v98.Q
		v7 = (*b3Shape)(unsafe.Pointer(fastShape)).LocalCentroid
		v11 = v15.V
		v16 = v7
		c.X = float32(v11.Y*v16.Z) - float32(v11.Z*v16.Y)
		c.Y = float32(v11.Z*v16.X) - float32(v11.X*v16.Z)
		c.Z = float32(v11.X*v16.Y) - float32(v11.Y*v16.X)
		v17 = c
		goto _133
	_133:
		t1 = v17
		v18 = t1
		v22 = v15.S
		v19 = v7
		v21 = Vec3{
			X: v18.X + float32(v22*v19.X),
			Y: v18.Y + float32(v22*v19.Y),
			Z: v18.Z + float32(v22*v19.Z),
		}
		goto _138
	_138:
		t2 = v21
		v23 = v15.V
		v24 = t2
		c.X = float32(v23.Y*v24.Z) - float32(v23.Z*v24.Y)
		c.Y = float32(v23.Z*v24.X) - float32(v23.X*v24.Z)
		c.Z = float32(v23.X*v24.Y) - float32(v23.Y*v24.X)
		v26 = c
		goto _142
	_142:
		t3 = v26
		v28 = v7
		v33 = float32FromFloat32(2)
		v30 = t3
		v32 = Vec3{
			X: v28.X + float32(v33*v30.X),
			Y: v28.Y + float32(v33*v30.Y),
			Z: v28.Z + float32(v33*v30.Z),
		}
		goto _149
	_149:
		v27 = v32
		goto _144
	_144:
		rv = v27
		v35 = rv
		v37 = v98.P
		v38 = Vec3{
			X: v35.X + v37.X,
			Y: v35.Y + v37.Y,
			Z: v35.Z + v37.Z,
		}
		goto _155
	_155:
		v34 = v38
		goto _151
	_151:
		(*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Centroid2 = v34
		box1 = (*b3Shape)(unsafe.Pointer(fastShape)).Aabb
		v156 = b3ComputeShapeAABB(tls, fastShape, xf2)
		v4 = base1
		v22 = float32(float64(v4.X + v156.LowerBound.X))
		goto _159
	_159:
		out.LowerBound.X = v22
		v33 = float32(float64(v4.Y + v156.LowerBound.Y))
		goto _161
	_161:
		out.LowerBound.Y = v33
		v48 = float32(float64(v4.Z + v156.LowerBound.Z))
		goto _163
	_163:
		out.LowerBound.Z = v48
		v59 = float32(float64(v4.X + v156.UpperBound.X))
		goto _165
	_165:
		out.UpperBound.X = v59
		v75 = float32(float64(v4.Y + v156.UpperBound.Y))
		goto _167
	_167:
		out.UpperBound.Y = v75
		v80 = float32(float64(v4.Z + v156.UpperBound.Z))
		goto _169
	_169:
		out.UpperBound.Z = v80
		v170 = out
		goto _171
	_171:
		// xf2 is relative to the base, so translate the box back to world space, rounding outward
		box2 = v170
		// Store this to avoid double computation in the case there is no impact event
		(*b3Shape)(unsafe.Pointer(fastShape)).Aabb = box2
		// No continuous collision for meshes
		if (*b3Shape)(unsafe.Pointer(fastShape)).Type1 == int32(b3_meshShape) || (*b3Shape)(unsafe.Pointer(fastShape)).Type1 == int32(b3_heightShape) {
			continue
		}
		// No continuous collision for sensors
		if (*b3Shape)(unsafe.Pointer(fastShape)).SensorIndex != -int32(1) {
			continue
		}
		if int32FromUint8(isBullet) == false1 {
			isShapeFast = b3IsShapeFast(tls, fastShape, (*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Centroid1, (*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Centroid2, rotationChord, safetyFactor)
			if int32FromUint8(isShapeFast) == false1 {
				continue
			}
		}
		v156 = box1
		v170 = box2
		v7 = v156.LowerBound
		v11 = v170.LowerBound
		v22 = v7.X
		v33 = v11.X
		if v22 < v33 {
			v59 = v22
		} else {
			v59 = v33
		}
		v48 = v59
		goto _181
	_181:
		v75 = v7.Y
		v80 = v11.Y
		if v75 < v80 {
			v91 = v75
		} else {
			v91 = v80
		}
		v86 = v91
		goto _186
	_186:
		v95 = v7.Z
		v106 = v11.Z
		if v95 < v106 {
			v135 = v95
		} else {
			v135 = v106
		}
		v117 = v135
		goto _191
	_191:
		v16 = Vec3{
			X: v48,
			Y: v86,
			Z: v117,
		}
		goto _177
	_177:
		out1.LowerBound = v16
		v17 = v156.UpperBound
		v18 = v170.UpperBound
		v146 = v17.X
		v158 = v18.X
		if v146 > v158 {
			v162 = v146
		} else {
			v162 = v158
		}
		v160 = v162
		goto _200
	_200:
		v164 = v17.Y
		v166 = v18.Y
		if v164 > v166 {
			v178 = v164
		} else {
			v178 = v166
		}
		v168 = v178
		goto _205
	_205:
		v179 = v17.Z
		v180 = v18.Z
		if v179 > v180 {
			v183 = v179
		} else {
			v183 = v180
		}
		v182 = v183
		goto _210
	_210:
		v19 = Vec3{
			X: v160,
			Y: v168,
			Z: v182,
		}
		goto _196
	_196:
		out1.UpperBound = v19
		v172 = out1
		goto _213
	_213:
		sweptBox = v172
		b3DynamicTree_Query(tls, staticTree, sweptBox, uint64FromUint64(0xffffffffffffffff), boolUint8(false1 != 0), __ccgo_fp(b3ContinuousQueryCallback), bp)
		if isBullet != 0 {
			b3DynamicTree_Query(tls, kinematicTree, sweptBox, uint64FromUint64(0xffffffffffffffff), boolUint8(false1 != 0), __ccgo_fp(b3ContinuousQueryCallback), bp)
			b3DynamicTree_Query(tls, dynamicTree, sweptBox, uint64FromUint64(0xffffffffffffffff), boolUint8(false1 != 0), __ccgo_fp(b3ContinuousQueryCallback), bp)
		}
	}
	speculativeScalar = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	if (*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Fraction < float32FromFloat32(1) {
		v15 = sweep.Q1
		v41 = sweep.Q2
		v22 = (*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Fraction
		v68 = v15
		v69 = v41
		v33 = float32(v68.V.X*v69.V.X) + float32(v68.V.Y*v69.V.Y) + float32(v68.V.Z*v69.V.Z) + float32(v68.S*v69.S)
		goto _220
	_220:
		if v33 < float32FromFloat32(0) {
			v15 = Quat{
				V: Vec3{
					X: -v15.V.X,
					Y: -v15.V.Y,
					Z: -v15.V.Z,
				},
				S: -v15.S,
			}
		}
		v7 = v15.V
		v11 = v41.V
		v48 = v22
		if v1 = !!(float32FromFloat32(0) <= v48 && v48 <= float32FromFloat32(1)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		c1 = Vec3{
			X: float32((float32FromFloat32(1)-v48)*v7.X) + float32(v48*v11.X),
			Y: float32((float32FromFloat32(1)-v48)*v7.Y) + float32(v48*v11.Y),
			Z: float32((float32FromFloat32(1)-v48)*v7.Z) + float32(v48*v11.Z),
		}
		v16 = c1
		goto _226
	_226:
		q3.V = v16
		q3.S = float32((float32FromFloat32(1)-v22)*v15.S) + float32(v22*v41.S)
		v99 = q3
		v128 = v99
		v214 = v99
		v59 = float32(v128.V.X*v214.V.X) + float32(v128.V.Y*v214.V.Y) + float32(v128.V.Z*v214.V.Z) + float32(v128.S*v214.S)
		goto _233
	_233:
		lengthSq = v59
		if lengthSq > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s3 = float32FromFloat32(1) / sqrtf(tls, lengthSq)
			qn = Quat{
				V: Vec3{
					X: float32(s3 * v99.V.X),
					Y: float32(s3 * v99.V.Y),
					Z: float32(s3 * v99.V.Z),
				},
				S: float32(s3 * v99.S),
			}
			v215 = qn
			goto _235
		}
		v215 = b3Quat_identity
		goto _235
	_235:
		v88 = v215
		goto _228
	_228:
		// Handle time of impact event. The sweep is relative to the base, so re-add the base
		// to return the advanced pose to world space.
		q4 = v88
		v17 = sweep.C1
		v18 = sweep.C2
		v75 = (*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Fraction
		if v2 = !!(float32FromFloat32(0) <= v75 && v75 <= float32FromFloat32(1)); !v2 {
			b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		c1 = Vec3{
			X: float32((float32FromFloat32(1)-v75)*v17.X) + float32(v75*v18.X),
			Y: float32((float32FromFloat32(1)-v75)*v17.Y) + float32(v75*v18.Y),
			Z: float32((float32FromFloat32(1)-v75)*v17.Z) + float32(v75*v18.Z),
		}
		v19 = c1
		goto _241
	_241:
		c2 = v19
		v217 = q4
		v21 = sweep.LocalCenter
		v23 = v217.V
		v24 = v21
		c.X = float32(v23.Y*v24.Z) - float32(v23.Z*v24.Y)
		c.Y = float32(v23.Z*v24.X) - float32(v23.X*v24.Z)
		c.Z = float32(v23.X*v24.Y) - float32(v23.Y*v24.X)
		v26 = c
		goto _247
	_247:
		t1 = v26
		v27 = t1
		v80 = v217.S
		v28 = v21
		v30 = Vec3{
			X: v27.X + float32(v80*v28.X),
			Y: v27.Y + float32(v80*v28.Y),
			Z: v27.Z + float32(v80*v28.Z),
		}
		goto _252
	_252:
		t2 = v30
		v32 = v217.V
		v34 = t2
		c.X = float32(v32.Y*v34.Z) - float32(v32.Z*v34.Y)
		c.Y = float32(v32.Z*v34.X) - float32(v32.X*v34.Z)
		c.Z = float32(v32.X*v34.Y) - float32(v32.Y*v34.X)
		v35 = c
		goto _256
	_256:
		t3 = v35
		v38 = v21
		v86 = float32FromFloat32(2)
		v39 = t3
		v42 = Vec3{
			X: v38.X + float32(v86*v39.X),
			Y: v38.Y + float32(v86*v39.Y),
			Z: v38.Z + float32(v86*v39.Z),
		}
		goto _263
	_263:
		v37 = v42
		goto _258
	_258:
		v43 = c2
		v44 = v37
		v45 = Vec3{
			X: v43.X - v44.X,
			Y: v43.Y - v44.Y,
			Z: v43.Z - v44.Z,
		}
		goto _267
	_267:
		origin1 = v45
		v4 = base1
		v47 = origin1
		v5 = Pos{
			X: v4.X + v47.X,
			Y: v4.Y + v47.Y,
			Z: v4.Z + v47.Z,
		}
		goto _271
	_271:
		// Advance body
		transform = WorldTransform{
			P: v5,
			Q: q4,
		}
		v6 = base1
		v49 = c2
		v9 = Pos{
			X: v6.X + v49.X,
			Y: v6.Y + v49.Y,
			Z: v6.Z + v49.Z,
		}
		goto _275
	_275:
		center = v9
		(*b3BodySim)(unsafe.Pointer(fastBodySim)).Transform = transform
		(*b3BodySim)(unsafe.Pointer(fastBodySim)).Center = center
		(*b3BodySim)(unsafe.Pointer(fastBodySim)).Rotation0 = q4
		(*b3BodySim)(unsafe.Pointer(fastBodySim)).Center0 = center
		*(*uint32_t)(unsafe.Add(nil, fastBodySim+40)) |= uint32(b3_hadTimeOfImpact)
		if v1 = !!(0 <= bodySimIndex && bodySimIndex < (*b3SolverSet)(unsafe.Pointer(awakeSet)).BodyStates.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+31460, __ccgo_ts+31091, int32FromInt32(602))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// Timeloss means there is a lost gravity contribution. Other forces and torques are ignored for now.
		fastBodyState = (*b3SolverSet)(unsafe.Pointer(awakeSet)).BodyStates.Data + uintptr(bodySimIndex)*64
		v3 = (*b3BodyState)(unsafe.Pointer(fastBodyState)).LinearVelocity
		timeLoss = float32((float32FromFloat32(1) - (*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Fraction) * dt)
		v22 = float32(-timeLoss * (*b3BodySim)(unsafe.Pointer(fastBodySim)).GravityScale)
		v7 = (*b3World)(unsafe.Pointer(world)).Gravity
		v11 = Vec3{
			X: float32(v22 * v7.X),
			Y: float32(v22 * v7.Y),
			Z: float32(v22 * v7.Z),
		}
		goto _280
	_280:
		dv = v11
		if (*b3BodyState)(unsafe.Pointer(fastBodyState)).Flags&uint32(b3_lockLinearX) != 0 {
			v33 = float32FromFloat32(0)
		} else {
			v33 = dv.X
		}
		dv.X = v33
		if (*b3BodyState)(unsafe.Pointer(fastBodyState)).Flags&uint32(b3_lockLinearY) != 0 {
			v22 = float32FromFloat32(0)
		} else {
			v22 = dv.Y
		}
		dv.Y = v22
		if (*b3BodyState)(unsafe.Pointer(fastBodyState)).Flags&uint32(b3_lockLinearZ) != 0 {
			v22 = float32FromFloat32(0)
		} else {
			v22 = dv.Z
		}
		dv.Z = v22
		v7 = v3
		v11 = dv
		v16 = Vec3{
			X: v7.X + v11.X,
			Y: v7.Y + v11.Y,
			Z: v7.Z + v11.Z,
		}
		goto _287
	_287:
		(*b3BodyState)(unsafe.Pointer(fastBodyState)).LinearVelocity = v16
		if v1 = !!(0 <= bodySimIndex && bodySimIndex < (*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+31535, __ccgo_ts+31091, int32FromInt32(612))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		// The move event was written before CCD, so correct it with the impact pose
		event = (*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Data + uintptr(bodySimIndex)*48
		(*BodyMoveEvent)(unsafe.Pointer(event)).Transform = (*b3BodySim)(unsafe.Pointer(fastBodySim)).Transform
		// Prepare AABBs for broad-phase.
		// Even though a body is fast, it may not move much. So the AABB may not need enlargement.
		shapeId = (*b3Body)(unsafe.Pointer(fastBody)).HeadShapeId
		for shapeId != -int32(1) {
			if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+31091, int32FromInt32(621))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			shape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
			// Must recompute aabb at the interpolated transform
			aabb1 = b3ComputeFatShapeAABB(tls, shape, transform, speculativeScalar)
			(*b3Shape)(unsafe.Pointer(shape)).Aabb = aabb1
			shapeFatAABB = (*b3World)(unsafe.Pointer(world)).FatAABBs.Data + uintptr(shapeId)*24
			v156 = *(*AABB)(unsafe.Add(nil, shapeFatAABB))
			v170 = aabb1
			if v156.LowerBound.X > v170.LowerBound.X || v170.UpperBound.X > v156.UpperBound.X {
				v292 = boolUint8(false1 != 0)
				goto _293
			}
			if v156.LowerBound.Y > v170.LowerBound.Y || v170.UpperBound.Y > v156.UpperBound.Y {
				v292 = boolUint8(false1 != 0)
				goto _293
			}
			if v156.LowerBound.Z > v170.LowerBound.Z || v170.UpperBound.Z > v156.UpperBound.Z {
				v292 = boolUint8(false1 != 0)
				goto _293
			}
			v292 = boolUint8(true1 != 0)
			goto _293
		_293:
			if int32FromUint8(v292) == false1 {
				v172 = aabb1
				v22 = (*b3Shape)(unsafe.Pointer(shape)).AabbMargin
				radius = Vec3{
					X: v22,
					Y: v22,
					Z: v22,
				}
				v7 = v172.LowerBound
				v11 = radius
				v16 = Vec3{
					X: v7.X - v11.X,
					Y: v7.Y - v11.Y,
					Z: v7.Z - v11.Z,
				}
				goto _299
			_299:
				out2.LowerBound = v16
				v17 = v172.UpperBound
				v18 = radius
				v19 = Vec3{
					X: v17.X + v18.X,
					Y: v17.Y + v18.Y,
					Z: v17.Z + v18.Z,
				}
				goto _303
			_303:
				out2.UpperBound = v19
				v173 = out2
				goto _305
			_305:
				*(*AABB)(unsafe.Add(nil, shapeFatAABB)) = v173
				// Regular bodies mark the hierarchy as moved using atomic operations.
				// Bullets are handled separately at a later stage.
				if int32FromUint8(isBullet) == false1 {
					v306 = (*b3Shape)(unsafe.Pointer(shape)).ProxyKey
					proxyType = v306 & int32FromInt32(3)
					proxyId = v306 >> int32(2)
					b3DynamicTree_MarkProxyMoved(tls, world+800+uintptr(proxyType)*112, proxyId, *(*AABB)(unsafe.Add(nil, shapeFatAABB)))
				} else {
					*(*uint32_t)(unsafe.Add(nil, fastBodySim+40)) |= uint32(b3_enlargeBulletBounds)
				}
			}
			shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
		}
	} else {
		// No time of impact event
		// Advance body
		(*b3BodySim)(unsafe.Pointer(fastBodySim)).Rotation0 = (*b3BodySim)(unsafe.Pointer(fastBodySim)).Transform.Q
		(*b3BodySim)(unsafe.Pointer(fastBodySim)).Center0 = (*b3BodySim)(unsafe.Pointer(fastBodySim)).Center
		// Prepare AABBs for broad-phase
		shapeId = (*b3Body)(unsafe.Pointer(fastBody)).HeadShapeId
		for shapeId != -int32(1) {
			if v1 = !!(0 <= shapeId && shapeId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+1997, __ccgo_ts+31091, int32FromInt32(659))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			shape1 = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(shapeId)*208
			// shape->aabb is still valid from above
			shapeFatAABB1 = (*b3World)(unsafe.Pointer(world)).FatAABBs.Data + uintptr(shapeId)*24
			v156 = *(*AABB)(unsafe.Add(nil, shapeFatAABB1))
			v170 = (*b3Shape)(unsafe.Pointer(shape1)).Aabb
			if v156.LowerBound.X > v170.LowerBound.X || v170.UpperBound.X > v156.UpperBound.X {
				v292 = boolUint8(false1 != 0)
				goto _311
			}
			if v156.LowerBound.Y > v170.LowerBound.Y || v170.UpperBound.Y > v156.UpperBound.Y {
				v292 = boolUint8(false1 != 0)
				goto _311
			}
			if v156.LowerBound.Z > v170.LowerBound.Z || v170.UpperBound.Z > v156.UpperBound.Z {
				v292 = boolUint8(false1 != 0)
				goto _311
			}
			v292 = boolUint8(true1 != 0)
			goto _311
		_311:
			if int32FromUint8(v292) == false1 {
				v172 = (*b3Shape)(unsafe.Pointer(shape1)).Aabb
				v22 = (*b3Shape)(unsafe.Pointer(shape1)).AabbMargin
				radius = Vec3{
					X: v22,
					Y: v22,
					Z: v22,
				}
				v7 = v172.LowerBound
				v11 = radius
				v16 = Vec3{
					X: v7.X - v11.X,
					Y: v7.Y - v11.Y,
					Z: v7.Z - v11.Z,
				}
				goto _317
			_317:
				out2.LowerBound = v16
				v17 = v172.UpperBound
				v18 = radius
				v19 = Vec3{
					X: v17.X + v18.X,
					Y: v17.Y + v18.Y,
					Z: v17.Z + v18.Z,
				}
				goto _321
			_321:
				out2.UpperBound = v19
				v173 = out2
				goto _323
			_323:
				*(*AABB)(unsafe.Add(nil, shapeFatAABB1)) = v173
				if int32FromUint8(isBullet) == false1 {
					v306 = (*b3Shape)(unsafe.Pointer(shape1)).ProxyKey
					proxyType = v306 & int32FromInt32(3)
					proxyId = v306 >> int32(2)
					b3DynamicTree_MarkProxyMoved(tls, world+800+uintptr(proxyType)*112, proxyId, *(*AABB)(unsafe.Add(nil, shapeFatAABB1)))
				} else {
					*(*uint32_t)(unsafe.Add(nil, fastBodySim+40)) |= uint32(b3_enlargeBulletBounds)
				}
			}
			shapeId = (*b3Shape)(unsafe.Pointer(shape1)).NextShapeId
		}
	}
	// Push sensor hits on the the task context for serial processing.
	i = 0
	for {
		if !(i < (*(*b3ContinuousContext)(unsafe.Add(nil, bp))).SensorCount) {
			break
		}
		// Skip any sensor hits that occurred after a solid hit
		if *(*float32)(unsafe.Add(nil, bp+196+uintptr(i)*4)) < (*(*b3ContinuousContext)(unsafe.Add(nil, bp))).Fraction {
			if (*b3TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Count >= (*b3TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity {
				oldSize = int32FromUint64(uint64FromInt32((*b3TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity) * uint64(8))
				if (*b3TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity == 0 {
					v306 = int32(8)
				} else {
					v306 = int32(2) * (*b3TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity
				}
				newCapacity = v306
				newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(8))
				(*b3TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Data = b3GrowAlloc(tls, (*b3TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Data, oldSize, newSize)
				(*b3TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Capacity = newCapacity
			}
			v31 = taskContext + 24 + 8
			v306 = *(*int32)(unsafe.Pointer(v31))
			*(*int32)(unsafe.Pointer(v31)) = *(*int32)(unsafe.Pointer(v31)) + 1
			*(*b3SensorHit)(unsafe.Add(nil, (*b3TaskContext)(unsafe.Pointer(taskContext)).SensorHits.Data+uintptr(v306)*8)) = *(*b3SensorHit)(unsafe.Add(nil, bp+132+uintptr(i)*8))
		}
		goto _325
	_325:
		;
		i = i + 1
	}
	v306 = (*b3TaskContext)(unsafe.Pointer(taskContext)).DistanceIterations
	v324 = (*(*b3ContinuousContext)(unsafe.Add(nil, bp))).DistanceIterations
	if v306 > v324 {
		v327 = v306
	} else {
		v327 = v324
	}
	v326 = v327
	goto _332
_332:
	(*b3TaskContext)(unsafe.Pointer(taskContext)).DistanceIterations = v326
	v306 = (*b3TaskContext)(unsafe.Pointer(taskContext)).PushBackIterations
	v324 = (*(*b3ContinuousContext)(unsafe.Add(nil, bp))).PushBackIterations
	if v306 > v324 {
		v327 = v306
	} else {
		v327 = v324
	}
	v326 = v327
	goto _337
_337:
	(*b3TaskContext)(unsafe.Pointer(taskContext)).PushBackIterations = v326
	v306 = (*b3TaskContext)(unsafe.Pointer(taskContext)).RootIterations
	v324 = (*(*b3ContinuousContext)(unsafe.Add(nil, bp))).RootIterations
	if v306 > v324 {
		v327 = v306
	} else {
		v327 = v324
	}
	v326 = v327
	goto _342
_342:
	(*b3TaskContext)(unsafe.Pointer(taskContext)).RootIterations = v326
}

// C documentation
//
//	// Parallel solver task
func b3SolverTask(tls *_Stack, taskContext uintptr) {
	bp := tls.Alloc(16)
	defer tls.Free(16)
	var activeColorCount, bodySyncIndex, colorIndex, colorIndex1, colorIndex2, colorIndex3, graphSyncIndex, iterationStageIndex, iterationStageIndex1, j, j1, j2, previousSyncIndex, restitutionIterations, spinCount, stageIndex, stageIndex1, subStepCount, subStepIndex, syncIndex, workerIndex, v1 int32
	var context, profile, stage, stages, workerContext uintptr
	var convexSyncIndex, jointSyncIndex, lastSyncBits, meshSyncIndex, syncBits, syncBits1, v27, v28 uint32_t
	var useBias, v2 uint8
	var v4 bool
	var _ /* ticks at bp+8 */ uint64_t
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = activeColorCount, bodySyncIndex, colorIndex, colorIndex1, colorIndex2, colorIndex3, context, convexSyncIndex, graphSyncIndex, iterationStageIndex, iterationStageIndex1, j, j1, j2, jointSyncIndex, lastSyncBits, meshSyncIndex, previousSyncIndex, profile, restitutionIterations, spinCount, stage, stageIndex, stageIndex1, stages, subStepCount, subStepIndex, syncBits, syncBits1, syncIndex, useBias, workerContext, workerIndex, v1, v2, v27, v28, v4
	workerContext = taskContext
	workerIndex = (*b3WorkerContext)(unsafe.Pointer(workerContext)).WorkerIndex
	context = (*b3WorkerContext)(unsafe.Pointer(workerContext)).Context
	activeColorCount = (*b3StepContext)(unsafe.Pointer(context)).ActiveColorCount
	stages = (*b3StepContext)(unsafe.Pointer(context)).Stages
	profile = (*b3StepContext)(unsafe.Pointer(context)).World + 4732
	if workerIndex == 0 {
		// The orchestrator slot is a race. The calling thread of b3World_Step also enters here
		// as worker 0, so progress is guaranteed even if the user's task system schedules tasks
		// out of order, has fewer threads than workerCount, or runs the task synchronously
		// inside enqueueTaskFcn. Whoever wins the CAS becomes the orchestrator; the loser
		// returns and lets the spinner-only path handle workers >0.
		v1 = 0
		*(*int32)(unsafe.Pointer(bp)) = v1
		v2 = __builtin___atomic_compare_exchange_n(tls, context+340, bp, int32(1), boolUint8(false1 != 0), int32(__ATOMIC_SEQ_CST), int32(__ATOMIC_SEQ_CST))
		goto _3
	_3:
		if int32FromUint8(v2) == false1 {
			return
		}
		// Main thread synchronizes the workers and does work itself.
		//
		// This needs to be a task for the main thread because the user's task system may execute
		// the tasks serially and this is the first task. This single task is able to fully
		// complete all work even if all other workers are blocked.
		// Stages are re-used by loops so that I don't need more stages for large substep counts.
		// The sync indices grow monotonically for the body/graph/constraint groupings because they share solver blocks.
		// The stage index and sync indices are combined in to sync bits for atomic synchronization.
		// The workers need to compute the previous sync index for a given stage so that CAS works correctly. This
		// setup makes this easy to do.
		/*
			b3_stagePrepareJoints,
			b3_stagePrepareContacts,
			b3_stageIntegrateVelocities,
			b3_stageWarmStart,
			b3_stageSolve,
			b3_stageIntegratePositions,
			b3_stageRelax,
			b3_stageStoreImpulses
		*/
		*(*uint64_t)(unsafe.Add(nil, bp+8)) = b3GetTicks(tls)
		bodySyncIndex = int32(1)
		stageIndex = 0
		// Prepare joint constraints
		jointSyncIndex = uint32(1)
		syncBits = jointSyncIndex<<int32FromInt32(16) | uint32FromInt32(stageIndex)
		if v4 = !!((*(*b3SolverStage)(unsafe.Add(nil, stages+uintptr(stageIndex)*24))).Type1 == int32(b3_stagePrepareJoints)); !v4 {
			b3InternalAssert(tls, __ccgo_ts+31924, __ccgo_ts+31091, int32FromInt32(1229))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		b3ExecuteMainStage(tls, stages+uintptr(stageIndex)*24, context, syncBits)
		stageIndex = stageIndex + int32(1)
		jointSyncIndex = jointSyncIndex + uint32(1)
		// Prepare convex contact constraints
		convexSyncIndex = uint32(1)
		syncBits = convexSyncIndex<<int32FromInt32(16) | uint32FromInt32(stageIndex)
		if v4 = !!((*(*b3SolverStage)(unsafe.Add(nil, stages+uintptr(stageIndex)*24))).Type1 == int32(b3_stagePrepareWideContacts)); !v4 {
			b3InternalAssert(tls, __ccgo_ts+31973, __ccgo_ts+31091, int32FromInt32(1237))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		b3ExecuteMainStage(tls, stages+uintptr(stageIndex)*24, context, syncBits)
		stageIndex = stageIndex + int32(1)
		convexSyncIndex = convexSyncIndex + uint32(1)
		// Prepare mesh contact constraints
		meshSyncIndex = uint32(1)
		syncBits = meshSyncIndex<<int32FromInt32(16) | uint32FromInt32(stageIndex)
		if v4 = !!((*(*b3SolverStage)(unsafe.Add(nil, stages+uintptr(stageIndex)*24))).Type1 == int32(b3_stagePrepareContacts)); !v4 {
			b3InternalAssert(tls, __ccgo_ts+32028, __ccgo_ts+31091, int32FromInt32(1245))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		b3ExecuteMainStage(tls, stages+uintptr(stageIndex)*24, context, syncBits)
		stageIndex = stageIndex + int32(1)
		meshSyncIndex = meshSyncIndex + uint32(1)
		// Single-threaded overflow work. These constraints don't fit in the graph coloring.
		b3PrepareJoints_Overflow(tls, context)
		b3PrepareContacts_Overflow(tls, context)
		*(*float32)(unsafe.Add(nil, profile+24)) += b3GetMillisecondsAndReset(tls, bp+8)
		graphSyncIndex = int32(1)
		subStepCount = (*b3StepContext)(unsafe.Pointer(context)).SubStepCount
		subStepIndex = 0
		for {
			if !(subStepIndex < subStepCount) {
				break
			}
			// stageIndex restarted each iteration
			// syncBits still increases monotonically because the upper bits increase each iteration
			iterationStageIndex = stageIndex
			// Integrate velocities
			syncBits = uint32FromInt32(bodySyncIndex<<int32(16) | iterationStageIndex)
			if v4 = !!((*(*b3SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex)*24))).Type1 == int32(b3_stageIntegrateVelocities)); !v4 {
				b3InternalAssert(tls, __ccgo_ts+32079, __ccgo_ts+31091, int32FromInt32(1266))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			b3ExecuteMainStage(tls, stages+uintptr(iterationStageIndex)*24, context, syncBits)
			iterationStageIndex = iterationStageIndex + int32(1)
			bodySyncIndex = bodySyncIndex + int32(1)
			*(*float32)(unsafe.Add(nil, profile+28)) += b3GetMillisecondsAndReset(tls, bp+8)
			// Warm start constraints
			b3WarmStartJoints_Overflow(tls, context)
			b3WarmStartContacts_Overflow(tls, context)
			colorIndex = 0
			for {
				if !(colorIndex < activeColorCount) {
					break
				}
				syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterationStageIndex)
				if v4 = !!((*(*b3SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex)*24))).Type1 == int32(b3_stageWarmStart)); !v4 {
					b3InternalAssert(tls, __ccgo_ts+32143, __ccgo_ts+31091, int32FromInt32(1280))
				}
				_ = v4 || bool(int32FromInt32(0) != 0)
				b3ExecuteMainStage(tls, stages+uintptr(iterationStageIndex)*24, context, syncBits)
				iterationStageIndex = iterationStageIndex + int32(1)
				goto _9
			_9:
				;
				colorIndex = colorIndex + 1
			}
			graphSyncIndex = graphSyncIndex + int32(1)
			*(*float32)(unsafe.Add(nil, profile+32)) += b3GetMillisecondsAndReset(tls, bp+8)
			// Solve constraints
			useBias = boolUint8(true1 != 0)
			j = 0
			for {
				if !(j < int32(_ITERATIONS)) {
					break
				}
				// Overflow constraints have lower priority. Typically these are dynamic-vs-dynamic.
				b3SolveJoints_Overflow(tls, context, useBias)
				b3PushContacts_Overflow(tls, context)
				colorIndex1 = 0
				for {
					if !(colorIndex1 < activeColorCount) {
						break
					}
					syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterationStageIndex)
					if v4 = !!((*(*b3SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex)*24))).Type1 == int32(b3_stageSolve)); !v4 {
						b3InternalAssert(tls, __ccgo_ts+32197, __ccgo_ts+31091, int32FromInt32(1299))
					}
					_ = v4 || bool(int32FromInt32(0) != 0)
					b3ExecuteMainStage(tls, stages+uintptr(iterationStageIndex)*24, context, syncBits)
					iterationStageIndex = iterationStageIndex + int32(1)
					goto _12
				_12:
					;
					colorIndex1 = colorIndex1 + 1
				}
				graphSyncIndex = graphSyncIndex + int32(1)
				goto _11
			_11:
				;
				j = j + 1
			}
			*(*float32)(unsafe.Add(nil, profile+36)) += b3GetMillisecondsAndReset(tls, bp+8)
			// Integrate positions
			if v4 = !!((*(*b3SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex)*24))).Type1 == int32(b3_stageIntegratePositions)); !v4 {
				b3InternalAssert(tls, __ccgo_ts+32247, __ccgo_ts+31091, int32FromInt32(1309))
			}
			_ = v4 || bool(int32FromInt32(0) != 0)
			syncBits = uint32FromInt32(bodySyncIndex<<int32(16) | iterationStageIndex)
			b3ExecuteMainStage(tls, stages+uintptr(iterationStageIndex)*24, context, syncBits)
			iterationStageIndex = iterationStageIndex + int32(1)
			bodySyncIndex = bodySyncIndex + int32(1)
			*(*float32)(unsafe.Add(nil, profile+40)) += b3GetMillisecondsAndReset(tls, bp+8)
			// Relax constraints
			useBias = boolUint8(false1 != 0)
			j1 = 0
			for {
				if !(j1 < int32(_RELAX_ITERATIONS)) {
					break
				}
				b3SolveJoints_Overflow(tls, context, useBias)
				b3SolveContacts_Overflow(tls, context)
				colorIndex2 = 0
				for {
					if !(colorIndex2 < activeColorCount) {
						break
					}
					syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterationStageIndex)
					if v4 = !!((*(*b3SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex)*24))).Type1 == int32(b3_stageRelax)); !v4 {
						b3InternalAssert(tls, __ccgo_ts+32310, __ccgo_ts+31091, int32FromInt32(1327))
					}
					_ = v4 || bool(int32FromInt32(0) != 0)
					b3ExecuteMainStage(tls, stages+uintptr(iterationStageIndex)*24, context, syncBits)
					iterationStageIndex = iterationStageIndex + int32(1)
					goto _16
				_16:
					;
					colorIndex2 = colorIndex2 + 1
				}
				graphSyncIndex = graphSyncIndex + int32(1)
				goto _15
			_15:
				;
				j1 = j1 + 1
			}
			*(*float32)(unsafe.Add(nil, profile+44)) += b3GetMillisecondsAndReset(tls, bp+8)
			goto _7
		_7:
			;
			subStepIndex = subStepIndex + 1
		}
		// Advance the stage according to the sub-stepping tasks just completed
		// integrate velocities / warm start / solve / integrate positions / relax
		stageIndex = stageIndex + (int32(1) + activeColorCount + int32(_ITERATIONS)*activeColorCount + int32(1) + int32(_RELAX_ITERATIONS)*activeColorCount)
		restitutionIterations = (*b3World)(unsafe.Pointer((*b3StepContext)(unsafe.Pointer(context)).World)).RestitutionIterations
		if v4 = restitutionIterations > 0; v4 {
			v1 = atomicLoadNInt32(context+108, int32FromInt32(__ATOMIC_SEQ_CST))
			goto _19
		_19:
		}
		if v4 && v1 != 0 {
			j2 = 0
			for {
				if !(j2 < restitutionIterations) {
					break
				}
				b3ApplyRestitution_Overflow(tls, context)
				iterationStageIndex1 = stageIndex
				colorIndex3 = 0
				for {
					if !(colorIndex3 < activeColorCount) {
						break
					}
					syncBits = uint32FromInt32(graphSyncIndex<<int32(16) | iterationStageIndex1)
					if v4 = !!((*(*b3SolverStage)(unsafe.Add(nil, stages+uintptr(iterationStageIndex1)*24))).Type1 == int32(b3_stageRestitution)); !v4 {
						b3InternalAssert(tls, __ccgo_ts+32360, __ccgo_ts+31091, int32FromInt32(1352))
					}
					_ = v4 || bool(int32FromInt32(0) != 0)
					b3ExecuteMainStage(tls, stages+uintptr(iterationStageIndex1)*24, context, syncBits)
					iterationStageIndex1 = iterationStageIndex1 + int32(1)
					goto _22
				_22:
					;
					colorIndex3 = colorIndex3 + 1
				}
				graphSyncIndex = graphSyncIndex + int32(1)
				goto _21
			_21:
				;
				j2 = j2 + 1
			}
			*(*float32)(unsafe.Add(nil, profile+48)) += b3GetMillisecondsAndReset(tls, bp+8)
		}
		stageIndex = stageIndex + activeColorCount
		// Store impulses
		b3StoreImpulses_Overflow(tls, context)
		syncBits = convexSyncIndex<<int32FromInt32(16) | uint32FromInt32(stageIndex)
		if v4 = !!((*(*b3SolverStage)(unsafe.Add(nil, stages+uintptr(stageIndex)*24))).Type1 == int32(b3_stageStoreWideImpulses)); !v4 {
			b3InternalAssert(tls, __ccgo_ts+32416, __ccgo_ts+31091, int32FromInt32(1368))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		b3ExecuteMainStage(tls, stages+uintptr(stageIndex)*24, context, syncBits)
		stageIndex = stageIndex + int32(1)
		syncBits = meshSyncIndex<<int32FromInt32(16) | uint32FromInt32(stageIndex)
		if v4 = !!((*(*b3SolverStage)(unsafe.Add(nil, stages+uintptr(stageIndex)*24))).Type1 == int32(b3_stageStoreImpulses)); !v4 {
			b3InternalAssert(tls, __ccgo_ts+32469, __ccgo_ts+31091, int32FromInt32(1373))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		b3ExecuteMainStage(tls, stages+uintptr(stageIndex)*24, context, syncBits)
		stageIndex = stageIndex + int32(1)
		*(*float32)(unsafe.Add(nil, profile+52)) += b3GetMillisecondsAndReset(tls, bp+8)
		// Signal workers to finish
		atomicStoreNUint32(context+272, uint32(0xffffffff), int32FromInt32(__ATOMIC_SEQ_CST))
		if v4 = !!(stageIndex == (*b3StepContext)(unsafe.Pointer(context)).StageCount); !v4 {
			b3InternalAssert(tls, __ccgo_ts+32518, __ccgo_ts+31091, int32FromInt32(1382))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		return
	}
	// Worker spins and waits for work
	lastSyncBits = uint32(0)
	// uint64_t maxSpinTime = 10;
	for int32(true1) != 0 {
		spinCount = 0
		for {
			v28 = atomicLoadNUint32(context+272, int32FromInt32(__ATOMIC_SEQ_CST))
			goto _29
		_29:
			v27 = v28
			syncBits1 = v27
			if !(v27 == lastSyncBits) {
				break
			}
			if spinCount > int32(5) {
				b3Yield(tls)
				spinCount = 0
			} else {
				// Using the cycle counter helps to account for variation in mm_pause timing across different
				// CPUs. However, this is X64 only.
				// uint64_t prev = __rdtsc();
				// do
				//{
				//	b3Pause();
				//}
				// while ((__rdtsc() - prev) < maxSpinTime);
				// maxSpinTime += 10;
				b3Pause(tls)
				b3Pause(tls)
				spinCount = spinCount + int32(1)
			}
		}
		if syncBits1 == uint32(0xffffffff) {
			// sentinel hit
			break
		}
		stageIndex1 = int32FromUint32(syncBits1 & uint32(0xFFFF))
		if v4 = !!(stageIndex1 < (*b3StepContext)(unsafe.Pointer(context)).StageCount); !v4 {
			b3InternalAssert(tls, __ccgo_ts+32552, __ccgo_ts+31091, int32FromInt32(1427))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		syncIndex = int32FromUint32(syncBits1 >> int32FromInt32(16) & uint32(0xFFFF))
		if v4 = !!(syncIndex > int32FromInt32(0)); !v4 {
			b3InternalAssert(tls, __ccgo_ts+31910, __ccgo_ts+31091, int32FromInt32(1430))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		previousSyncIndex = syncIndex - int32(1)
		stage = stages + uintptr(stageIndex1)*24
		b3ExecuteStage(tls, stage, context, previousSyncIndex, syncIndex, workerIndex)
		lastSyncBits = syncBits1
	}
}

// C documentation
//
//	// Solve with graph coloring
func b3Solve(tls *_Stack, world uintptr, stepContext uintptr) {
	bp := tls.Alloc(2656)
	defer tls.Free(2656)
	var activeColorCount, awakeBodyCount, awakeIslandCount, bodyId, bulletBodyCount, c1, colorContactCount, colorContactCount1, colorContactCount2, colorContactCountW, colorConvexContactCount, colorConvexContactCount1, colorJointCount, colorMeshContactCount, colorMeshGroupCount, colorMeshGroupCount1, colorMeshGroupCount2, colorWideConstraintCount, contactBase, contactId, contactIdCapacity, count, count1, g, graphBlockCount, groupBase, hitCount, i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i2, i3, i4, i5, i6, i7, i8, i9, islandId, islandIndex, j, j1, j2, j3, j4, jointBase, jointCount, jointId, jointIdCapacity, k, k1, manifoldCount, maxBlockCount, maxManifoldCount, meshContactCount, meshGroupBase, meshGroupCount, meshSlotCount, minBodiesPerBlock, minContactsPerBlock, minJointsPerBlock, minRange, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, newSize3, occupancyCount, oldSize, oldSize1, oldSize2, oldSize3, orderBase, overflowCount, overflowManifoldCount, p1, perColorContactCount, perColorJointCount, pointCount, proxyId1, proxyKey, shapeId, simdShift, slotCount, stageCount, start, triangleIndex, wideBase, wideContactByteCount, wideContactCount, wideMeshConstraintByteCount, wideMeshManifoldByteCount, workerCount, workerCount1, v1, v4, v6, v7 int32
	var anyHitEvents, found, v94 uint8
	var approachSpeed, splitSleepTimer, threshold, v67 float32
	var awakeIslandBitSet, awakeSet, baseGraphBlock, bits, bits1, bodyA, bodyArray, bodyB, bodyBlocks, bodySimArray, broadPhase, bucketStarts, bulletBody, bulletBodySim, bulletBodySimIndices, color, color1, color2, color3, colors, contact, contactArray, convexBlocks, dynamicTree, graph, graphBlocks, hitEventBitSet, hits, island, islands, joint, jointArray, jointBlocks, jointStateBitSet, manifold, meshBlocks, meshLaneOrder, meshManifoldStarts, mp, order, overflow, sensor, sensorShape, shape, shapeA, shapeArray, shapeB, simA, simB, specs, specs1, splitIslandTask, stage, stages, taskContext, taskContext1, taskContext2, taskContext3, visitor, wideConstraints, wideMeshConstraints, v21 uintptr
	var blockIndex, ctz, ctz1, k2, k3, wordCount, wordCount1, v47 uint32_t
	var bodyDim, convexPrepareDim, jointPrepareDim, meshPrepareDim b3BlockDim
	var bulletTicks, hitTicks, jointEventTicks, refitTicks, sensorHitTicks, sleepTicks, transformTicks, word, word1 uint64_t
	var c, v72, v73, v76, v79 Vec3
	var colorIndex uint8_t
	var colorJointCounts, colorMeshContactCounts, colorMeshGroupCounts, colorWideContactCounts [24]int32
	var event JointEvent
	var event1 ContactHitEvent
	var graphJointDims, graphMeshContactDims, graphWideContactDims [24]b3BlockDim
	var hit b3SensorHit
	var midCenter, v65, v66, v68 Pos
	var shapeFatAABB, treeAABB, v90, v92, v93 AABB
	var shapeRef b3Visitor
	var worldId, worldIndex0 uint16_t
	var v2, v23, v24, v27 bool
	var _ /* activeColorIndices at bp+8 */ [24]int32
	var _ /* buffer at bp+2600 */ [16]uint8
	var _ /* callerContext at bp+2616 */ b3WorkerContext
	var _ /* constraintTicks at bp+2592 */ uint64_t
	var _ /* graphBlockCounts at bp+1728 */ [24]int32
	var _ /* graphColorBlocks at bp+1536 */ [24]uintptr
	var _ /* jointPrepareSpans at bp+1104 */ [25]b3JointPrepareSpan
	var _ /* meshPrepareSpans at bp+504 */ [25]b3MeshPrepareSpan
	var _ /* overflowSpans at bp+1504 */ [2]b3ContactPrepareSpan
	var _ /* setupTicks at bp+0 */ uint64_t
	var _ /* widePrepareSpans at bp+104 */ [25]b3WidePrepareSpan
	var _ /* workerContext at bp+1824 */ [32]b3WorkerContext
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = activeColorCount, anyHitEvents, approachSpeed, awakeBodyCount, awakeIslandBitSet, awakeIslandCount, awakeSet, baseGraphBlock, bits, bits1, blockIndex, bodyA, bodyArray, bodyB, bodyBlocks, bodyDim, bodyId, bodySimArray, broadPhase, bucketStarts, bulletBody, bulletBodyCount, bulletBodySim, bulletBodySimIndices, bulletTicks, c, c1, color, color1, color2, color3, colorContactCount, colorContactCount1, colorContactCount2, colorContactCountW, colorConvexContactCount, colorConvexContactCount1, colorIndex, colorJointCount, colorJointCounts, colorMeshContactCount, colorMeshContactCounts, colorMeshGroupCount, colorMeshGroupCount1, colorMeshGroupCount2, colorMeshGroupCounts, colorWideConstraintCount, colorWideContactCounts, colors, contact, contactArray, contactBase, contactId, contactIdCapacity, convexBlocks, convexPrepareDim, count, count1, ctz, ctz1, dynamicTree, event, event1, found, g, graph, graphBlockCount, graphBlocks, graphJointDims, graphMeshContactDims, graphWideContactDims, groupBase, hit, hitCount, hitEventBitSet, hitTicks, hits, i, i1, i10, i11, i12, i13, i14, i15, i16, i17, i2, i3, i4, i5, i6, i7, i8, i9, island, islandId, islandIndex, islands, j, j1, j2, j3, j4, joint, jointArray, jointBase, jointBlocks, jointCount, jointEventTicks, jointId, jointIdCapacity, jointPrepareDim, jointStateBitSet, k, k1, k2, k3, manifold, manifoldCount, maxBlockCount, maxManifoldCount, meshBlocks, meshContactCount, meshGroupBase, meshGroupCount, meshLaneOrder, meshManifoldStarts, meshPrepareDim, meshSlotCount, midCenter, minBodiesPerBlock, minContactsPerBlock, minJointsPerBlock, minRange, mp, newCapacity, newCapacity1, newCapacity2, newSize, newSize1, newSize2, newSize3, occupancyCount, oldSize, oldSize1, oldSize2, oldSize3, order, orderBase, overflow, overflowCount, overflowManifoldCount, p1, perColorContactCount, perColorJointCount, pointCount, proxyId1, proxyKey, refitTicks, sensor, sensorHitTicks, sensorShape, shape, shapeA, shapeArray, shapeB, shapeFatAABB, shapeId, shapeRef, simA, simB, simdShift, sleepTicks, slotCount, specs, specs1, splitIslandTask, splitSleepTimer, stage, stageCount, stages, start, taskContext, taskContext1, taskContext2, taskContext3, threshold, transformTicks, treeAABB, triangleIndex, visitor, wideBase, wideConstraints, wideContactByteCount, wideContactCount, wideMeshConstraintByteCount, wideMeshConstraints, wideMeshManifoldByteCount, word, word1, wordCount, wordCount1, workerCount, workerCount1, worldId, worldIndex0, v1, v2, v21, v23, v24, v27, v4, v47, v6, v65, v66, v67, v68, v7, v72, v73, v76, v79, v90, v92, v93, v94
	if (*b3World)(unsafe.Pointer(world)).SimdWidth == int32(8) {
		v1 = int32(3)
	} else {
		v1 = int32(2)
	}
	simdShift = v1
	// Only count steps that advance the simulation
	*(*uint64_t)(unsafe.Add(nil, world+4648)) += uint64(1)
	if v2 = !!(bool(0 <= int32(b3_awakeSet)) && int32(b3_awakeSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+802, __ccgo_ts+31091, int32FromInt32(1467))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_awakeSet))*88
	awakeBodyCount = (*b3SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Count
	if awakeBodyCount == 0 {
		return
	}
	// Solve constraints using graph coloring
	*(*uint64_t)(unsafe.Add(nil, bp)) = b3GetTicks(tls)
	// Prepare buffers for continuous collision (fast bodies)
	atomicStoreNInt32(stepContext+104, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	(*b3StepContext)(unsafe.Pointer(stepContext)).BulletBodies = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(awakeBodyCount)*uint64(4)), __ccgo_ts+32693)
	graph = world + 1160
	colors = graph
	(*b3StepContext)(unsafe.Pointer(stepContext)).Sims = (*b3SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Data
	(*b3StepContext)(unsafe.Pointer(stepContext)).States = (*b3SolverSet)(unsafe.Pointer(awakeSet)).BodyStates.Data
	// count contacts, joints, and colors
	activeColorCount = 0
	i = 0
	for {
		if !(i < int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
			break
		}
		perColorContactCount = (*(*b3GraphColor)(unsafe.Add(nil, colors+uintptr(i)*120))).ConvexContacts.Count + (*(*b3GraphColor)(unsafe.Add(nil, colors+uintptr(i)*120))).Contacts.Count
		perColorJointCount = (*(*b3GraphColor)(unsafe.Add(nil, colors+uintptr(i)*120))).JointSims.Count
		occupancyCount = perColorContactCount + perColorJointCount
		if occupancyCount > 0 {
			v1 = int32(1)
		} else {
			v1 = 0
		}
		activeColorCount = activeColorCount + v1
		goto _3
	_3:
		;
		i = i + 1
	}
	// prepare for move events
	if (*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity < awakeBodyCount {
		oldSize = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity) * uint64(48))
		newSize = int32FromUint64(uint64FromInt32(awakeBodyCount) * uint64(48))
		(*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Data, oldSize, newSize)
		(*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Capacity = awakeBodyCount
	}
	(*b3World)(unsafe.Pointer(world)).BodyMoveEvents.Count = awakeBodyCount
	workerCount = (*b3World)(unsafe.Pointer(world)).WorkerCount
	// Target 4 blocks per worker to allow work stealing
	maxBlockCount = int32(4) * workerCount
	// Body blocks are for parallel iteration over bodies directly (integration, update transforms)
	minBodiesPerBlock = int32(32)
	bodyDim = b3ComputeBlockCount(tls, awakeBodyCount, minBodiesPerBlock, maxBlockCount)
	minContactsPerBlock = int32(4)
	minJointsPerBlock = int32(4)
	graphBlockCount = 0
	// c is the active color index
	wideContactCount = 0
	meshContactCount = 0
	meshGroupCount = 0
	jointCount = 0
	c1 = 0
	i1 = 0
	for {
		if !(i1 < int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1)) {
			break
		}
		color = colors + uintptr(i1)*120
		colorConvexContactCount = (*b3GraphColor)(unsafe.Pointer(color)).ConvexContacts.Count
		colorMeshContactCount = (*b3GraphColor)(unsafe.Pointer(color)).Contacts.Count
		colorJointCount = (*b3GraphColor)(unsafe.Pointer(color)).JointSims.Count
		if colorConvexContactCount+colorMeshContactCount+colorJointCount == 0 {
			goto _5
		}
		(*(*[24]int32)(unsafe.Add(nil, bp+8)))[c1] = i1
		if colorConvexContactCount > 0 {
			v1 = (colorConvexContactCount-int32(1))>>simdShift + int32(1)
		} else {
			v1 = 0
		}
		// Ceiling for wide constraint count
		colorWideConstraintCount = v1
		wideContactCount = wideContactCount + colorWideConstraintCount
		colorWideContactCounts[c1] = colorWideConstraintCount
		colorMeshContactCounts[c1] = colorMeshContactCount
		meshContactCount = meshContactCount + colorMeshContactCount
		if colorMeshContactCount > 0 {
			v1 = (colorMeshContactCount-int32(1))>>simdShift + int32(1)
		} else {
			v1 = 0
		}
		colorMeshGroupCount = v1
		colorMeshGroupCounts[c1] = colorMeshGroupCount
		meshGroupCount = meshGroupCount + colorMeshGroupCount
		colorJointCounts[c1] = colorJointCount
		jointCount = jointCount + colorJointCount
		// Solver block dimensions
		graphWideContactDims[c1] = b3ComputeBlockCount(tls, colorWideConstraintCount, minContactsPerBlock, maxBlockCount)
		graphMeshContactDims[c1] = b3ComputeBlockCount(tls, colorMeshGroupCount, minContactsPerBlock, maxBlockCount)
		graphJointDims[c1] = b3ComputeBlockCount(tls, colorJointCount, minJointsPerBlock, maxBlockCount)
		graphBlockCount = graphBlockCount + (graphWideContactDims[c1].Count + graphMeshContactDims[c1].Count + graphJointDims[c1].Count)
		c1 = c1 + int32(1)
		goto _5
	_5:
		;
		i1 = i1 + 1
	}
	activeColorCount = c1
	// Prepare and store run as one flat parallel-for over the entire wide constraint range,
	// partitioned into uniformly sized blocks. Color info is consulted inside the task via
	// a small span array, so blocks do not need to honor color boundaries here.
	convexPrepareDim = b3ComputeBlockCount(tls, wideContactCount, minContactsPerBlock, maxBlockCount)
	meshPrepareDim = b3ComputeBlockCount(tls, meshGroupCount, minContactsPerBlock, maxBlockCount)
	jointPrepareDim = b3ComputeBlockCount(tls, jointCount, minJointsPerBlock, maxBlockCount)
	wideContactByteCount = b3GetWideContactConstraintByteCount(tls, (*b3World)(unsafe.Pointer(world)).SimdWidth)
	wideConstraints = b3StackAlloc(tls, world, wideContactCount*wideContactByteCount, __ccgo_ts+32707)
	meshLaneOrder = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(meshContactCount)*uint64(4)), __ccgo_ts+32721)
	meshManifoldStarts = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(meshGroupCount+int32FromInt32(1))*uint64(4)), __ccgo_ts+32737)
	// Sort mesh contacts by manifold count so that contacts with more manifolds are solved
	// together keeping the SIMD lanes fuller and less ragged.
	meshSlotCount = 0
	maxManifoldCount = 0
	i2 = 0
	for {
		if !(i2 < activeColorCount) {
			break
		}
		color1 = colors + uintptr((*(*[24]int32)(unsafe.Add(nil, bp+8)))[i2])*120
		specs = (*b3GraphColor)(unsafe.Pointer(color1)).Contacts.Data
		colorContactCount = colorMeshContactCounts[i2]
		j = 0
		for {
			if !(j < colorContactCount) {
				break
			}
			v1 = maxManifoldCount
			v4 = int32FromUint16((*(*b3ContactSpec)(unsafe.Add(nil, specs+uintptr(j)*12))).ManifoldCount)
			if v1 > v4 {
				v7 = v1
			} else {
				v7 = v4
			}
			v6 = v7
			goto _13
		_13:
			maxManifoldCount = v6
			goto _9
		_9:
			;
			j = j + 1
		}
		goto _8
	_8:
		;
		i2 = i2 + 1
	}
	bucketStarts = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(maxManifoldCount+int32FromInt32(1))*uint64(4)), __ccgo_ts+32758)
	orderBase = 0
	groupBase = 0
	i3 = 0
	for {
		if !(i3 < activeColorCount) {
			break
		}
		color2 = colors + uintptr((*(*[24]int32)(unsafe.Add(nil, bp+8)))[i3])*120
		colorContactCount1 = colorMeshContactCounts[i3]
		if colorContactCount1 == 0 {
			goto _15
		}
		specs1 = (*b3GraphColor)(unsafe.Pointer(color2)).Contacts.Data
		order = meshLaneOrder + uintptr(orderBase)*4
		k = 0
		for {
			if !(k <= maxManifoldCount) {
				break
			}
			*(*int32)(unsafe.Add(nil, bucketStarts+uintptr(k)*4)) = 0
			goto _16
		_16:
			;
			k = k + 1
		}
		j1 = 0
		for {
			if !(j1 < colorContactCount1) {
				break
			}
			*(*int32)(unsafe.Add(nil, bucketStarts+uintptr((*(*b3ContactSpec)(unsafe.Add(nil, specs1+uintptr(j1)*12))).ManifoldCount)*4)) += int32(1)
			goto _17
		_17:
			;
			j1 = j1 + 1
		}
		start = 0
		k1 = maxManifoldCount
		for {
			if !(k1 >= 0) {
				break
			}
			count = *(*int32)(unsafe.Add(nil, bucketStarts+uintptr(k1)*4))
			*(*int32)(unsafe.Add(nil, bucketStarts+uintptr(k1)*4)) = start
			start = start + count
			goto _18
		_18:
			;
			k1 = k1 - 1
		}
		j2 = 0
		for {
			if !(j2 < colorContactCount1) {
				break
			}
			v21 = bucketStarts + uintptr((*(*b3ContactSpec)(unsafe.Add(nil, specs1+uintptr(j2)*12))).ManifoldCount)*4
			v1 = *(*int32)(unsafe.Pointer(v21))
			*(*int32)(unsafe.Pointer(v21)) = *(*int32)(unsafe.Pointer(v21)) + 1
			*(*int32)(unsafe.Add(nil, order+uintptr(v1)*4)) = j2
			goto _19
		_19:
			;
			j2 = j2 + 1
		}
		colorMeshGroupCount1 = colorMeshGroupCounts[i3]
		g = 0
		for {
			if !(g < colorMeshGroupCount1) {
				break
			}
			slotCount = int32FromUint16((*(*b3ContactSpec)(unsafe.Add(nil, specs1+uintptr(*(*int32)(unsafe.Add(nil, order+uintptr(g<<simdShift)*4)))*12))).ManifoldCount)
			*(*int32)(unsafe.Add(nil, meshManifoldStarts+uintptr(groupBase+g)*4)) = meshSlotCount
			meshSlotCount = meshSlotCount + slotCount
			goto _22
		_22:
			;
			g = g + 1
		}
		orderBase = orderBase + colorContactCount1
		groupBase = groupBase + colorMeshGroupCount1
		goto _15
	_15:
		;
		i3 = i3 + 1
	}
	if v2 = !!(orderBase == meshContactCount); !v2 {
		b3InternalAssert(tls, __ccgo_ts+32771, __ccgo_ts+31091, int32FromInt32(1651))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	if v2 = !!(groupBase == meshGroupCount); !v2 {
		b3InternalAssert(tls, __ccgo_ts+32801, __ccgo_ts+31091, int32FromInt32(1652))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	*(*int32)(unsafe.Add(nil, meshManifoldStarts+uintptr(meshGroupCount)*4)) = meshSlotCount
	b3StackFree(tls, world, bucketStarts)
	wideMeshConstraintByteCount = b3GetWideMeshConstraintByteCount(tls, (*b3World)(unsafe.Pointer(world)).SimdWidth)
	wideMeshManifoldByteCount = b3GetWideMeshManifoldByteCount(tls, (*b3World)(unsafe.Pointer(world)).SimdWidth)
	wideMeshConstraints = b3StackAlloc(tls, world, meshGroupCount*wideMeshConstraintByteCount+meshSlotCount*wideMeshManifoldByteCount, __ccgo_ts+32829)
	overflow = colors + uintptr(int32FromInt32(_B3_GRAPH_COLOR_COUNT)-int32FromInt32(1))*120
	overflowCount = (*b3GraphColor)(unsafe.Pointer(overflow)).Contacts.Count
	overflowManifoldCount = 0
	i4 = 0
	for {
		if !(i4 < overflowCount) {
			break
		}
		(*(*b3ContactSpec)(unsafe.Add(nil, (*b3GraphColor)(unsafe.Pointer(overflow)).Contacts.Data+uintptr(i4)*12))).ManifoldStart = overflowManifoldCount
		overflowManifoldCount = overflowManifoldCount + int32FromUint16((*(*b3ContactSpec)(unsafe.Add(nil, (*b3GraphColor)(unsafe.Pointer(overflow)).Contacts.Data+uintptr(i4)*12))).ManifoldCount)
		goto _25
	_25:
		;
		i4 = i4 + 1
	}
	(*b3GraphColor)(unsafe.Pointer(overflow)).ContactConstraints = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(overflowCount)*uint64(168)), __ccgo_ts+32851)
	(*b3GraphColor)(unsafe.Pointer(overflow)).ManifoldConstraints = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(overflowManifoldCount)*uint64(324)), __ccgo_ts+32869)
	// Distribute transient constraints to each graph color and prepare spans
	wideBase = 0
	contactBase = 0
	meshGroupBase = 0
	jointBase = 0
	i5 = 0
	for {
		if !(i5 < activeColorCount) {
			break
		}
		j3 = (*(*[24]int32)(unsafe.Add(nil, bp+8)))[i5]
		color3 = colors + uintptr(j3)*120
		colorConvexContactCount1 = (*b3GraphColor)(unsafe.Pointer(color3)).ConvexContacts.Count
		(*(*[25]b3WidePrepareSpan)(unsafe.Add(nil, bp+104)))[i5].Start = wideBase
		(*(*[25]b3WidePrepareSpan)(unsafe.Add(nil, bp+104)))[i5].Count = colorConvexContactCount1
		(*(*[25]b3WidePrepareSpan)(unsafe.Add(nil, bp+104)))[i5].Contacts = (*b3GraphColor)(unsafe.Pointer(color3)).ConvexContacts.Data
		if colorConvexContactCount1 == 0 {
			(*b3GraphColor)(unsafe.Pointer(color3)).WideConstraints = uintptrFromInt32(0)
			(*b3GraphColor)(unsafe.Pointer(color3)).WideConstraintCount = 0
		} else {
			(*b3GraphColor)(unsafe.Pointer(color3)).WideConstraints = wideConstraints + uintptr(wideBase*wideContactByteCount)
			colorContactCountW = (colorConvexContactCount1-int32(1))>>simdShift + int32(1)
			(*b3GraphColor)(unsafe.Pointer(color3)).WideConstraintCount = colorContactCountW
			wideBase = wideBase + colorContactCountW
		}
		colorContactCount2 = (*b3GraphColor)(unsafe.Pointer(color3)).Contacts.Count
		(*(*[25]b3MeshPrepareSpan)(unsafe.Add(nil, bp+504)))[i5].Start = meshGroupBase
		(*(*[25]b3MeshPrepareSpan)(unsafe.Add(nil, bp+504)))[i5].Count = colorContactCount2
		(*(*[25]b3MeshPrepareSpan)(unsafe.Add(nil, bp+504)))[i5].Contacts = (*b3GraphColor)(unsafe.Pointer(color3)).Contacts.Data
		(*(*[25]b3MeshPrepareSpan)(unsafe.Add(nil, bp+504)))[i5].Order = meshLaneOrder + uintptr(contactBase)*4
		if colorContactCount2 == 0 {
			(*b3GraphColor)(unsafe.Pointer(color3)).WideMeshConstraints = uintptrFromInt32(0)
			(*b3GraphColor)(unsafe.Pointer(color3)).WideMeshManifoldStarts = uintptrFromInt32(0)
			(*b3GraphColor)(unsafe.Pointer(color3)).WideMeshConstraintCount = 0
		} else {
			colorMeshGroupCount2 = colorMeshGroupCounts[i5]
			(*b3GraphColor)(unsafe.Pointer(color3)).WideMeshConstraints = wideMeshConstraints + uintptr(meshGroupBase*wideMeshConstraintByteCount) + uintptr(*(*int32)(unsafe.Add(nil, meshManifoldStarts+uintptr(meshGroupBase)*4))*wideMeshManifoldByteCount)
			(*b3GraphColor)(unsafe.Pointer(color3)).WideMeshManifoldStarts = meshManifoldStarts + uintptr(meshGroupBase)*4
			(*b3GraphColor)(unsafe.Pointer(color3)).WideMeshConstraintCount = colorMeshGroupCount2
			meshGroupBase = meshGroupBase + colorMeshGroupCount2
			contactBase = contactBase + colorContactCount2
		}
		(*(*[25]b3JointPrepareSpan)(unsafe.Add(nil, bp+1104)))[i5].Start = jointBase
		(*(*[25]b3JointPrepareSpan)(unsafe.Add(nil, bp+1104)))[i5].Count = (*b3GraphColor)(unsafe.Pointer(color3)).JointSims.Count
		(*(*[25]b3JointPrepareSpan)(unsafe.Add(nil, bp+1104)))[i5].Joints = (*b3GraphColor)(unsafe.Pointer(color3)).JointSims.Data
		jointBase = jointBase + (*b3GraphColor)(unsafe.Pointer(color3)).JointSims.Count
		goto _26
	_26:
		;
		i5 = i5 + 1
	}
	// Sentinels
	(*(*[25]b3WidePrepareSpan)(unsafe.Add(nil, bp+104)))[activeColorCount].Start = wideContactCount
	(*(*[25]b3WidePrepareSpan)(unsafe.Add(nil, bp+104)))[activeColorCount].Count = 0
	(*(*[25]b3WidePrepareSpan)(unsafe.Add(nil, bp+104)))[activeColorCount].Contacts = uintptrFromInt32(0)
	if v2 = !!(wideBase == wideContactCount); !v2 {
		b3InternalAssert(tls, __ccgo_ts+32888, __ccgo_ts+31091, int32FromInt32(1748))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	(*(*[25]b3MeshPrepareSpan)(unsafe.Add(nil, bp+504)))[activeColorCount].Start = meshGroupCount
	(*(*[25]b3MeshPrepareSpan)(unsafe.Add(nil, bp+504)))[activeColorCount].Count = 0
	(*(*[25]b3MeshPrepareSpan)(unsafe.Add(nil, bp+504)))[activeColorCount].Contacts = uintptrFromInt32(0)
	(*(*[25]b3MeshPrepareSpan)(unsafe.Add(nil, bp+504)))[activeColorCount].Order = uintptrFromInt32(0)
	if v2 = !!(contactBase == meshContactCount); !v2 {
		b3InternalAssert(tls, __ccgo_ts+32917, __ccgo_ts+31091, int32FromInt32(1754))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	if v2 = !!(meshGroupBase == meshGroupCount); !v2 {
		b3InternalAssert(tls, __ccgo_ts+32949, __ccgo_ts+31091, int32FromInt32(1755))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	(*(*[25]b3JointPrepareSpan)(unsafe.Add(nil, bp+1104)))[activeColorCount].Start = jointCount
	(*(*[25]b3JointPrepareSpan)(unsafe.Add(nil, bp+1104)))[activeColorCount].Count = 0
	(*(*[25]b3JointPrepareSpan)(unsafe.Add(nil, bp+1104)))[activeColorCount].Joints = uintptrFromInt32(0)
	if v2 = !!(jointBase == jointCount); !v2 {
		b3InternalAssert(tls, __ccgo_ts+32981, __ccgo_ts+31091, int32FromInt32(1760))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	//// Special span for overflow to allow for function re-use
	*(*[2]b3ContactPrepareSpan)(unsafe.Add(nil, bp+1504)) = [2]b3ContactPrepareSpan{}
	(*(*[2]b3ContactPrepareSpan)(unsafe.Add(nil, bp+1504)))[0].Start = 0
	(*(*[2]b3ContactPrepareSpan)(unsafe.Add(nil, bp+1504)))[0].Count = (*b3GraphColor)(unsafe.Pointer(overflow)).Contacts.Count
	(*(*[2]b3ContactPrepareSpan)(unsafe.Add(nil, bp+1504)))[0].Contacts = (*b3GraphColor)(unsafe.Pointer(overflow)).Contacts.Data
	(*(*[2]b3ContactPrepareSpan)(unsafe.Add(nil, bp+1504)))[int32(1)].Start = (*b3GraphColor)(unsafe.Pointer(overflow)).Contacts.Count
	(*(*[2]b3ContactPrepareSpan)(unsafe.Add(nil, bp+1504)))[int32(1)].Count = 0
	(*(*[2]b3ContactPrepareSpan)(unsafe.Add(nil, bp+1504)))[int32(1)].Contacts = uintptrFromInt32(0)
	stageCount = 0
	// b3_stagePrepareJoints
	stageCount = stageCount + int32(1)
	// b3_stagePrepareWideContacts
	stageCount = stageCount + int32(1)
	// b3_stagePrepareContacts
	stageCount = stageCount + int32(1)
	// b3_stageIntegrateVelocities
	stageCount = stageCount + int32(1)
	// b3_stageWarmStart
	stageCount = stageCount + activeColorCount
	// b3_stageSolve
	stageCount = stageCount + int32(_ITERATIONS)*activeColorCount
	// b3_stageIntegratePositions
	stageCount = stageCount + int32(1)
	// b3_stageRelax
	stageCount = stageCount + int32(_RELAX_ITERATIONS)*activeColorCount
	// b3_stageRestitution
	stageCount = stageCount + activeColorCount
	// b3_stageStoreWideImpulses
	stageCount = stageCount + int32(1)
	// b3_stageStoreImpulses
	stageCount = stageCount + int32(1)
	stages = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(stageCount)*uint64(24)), __ccgo_ts+33005)
	bodyBlocks = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(bodyDim.Count)*uint64(12)), __ccgo_ts+33012)
	convexBlocks = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(convexPrepareDim.Count)*uint64(12)), __ccgo_ts+33024)
	meshBlocks = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(meshPrepareDim.Count)*uint64(12)), __ccgo_ts+33038)
	jointBlocks = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(jointPrepareDim.Count)*uint64(12)), __ccgo_ts+33050)
	graphBlocks = b3StackAlloc(tls, world, int32FromUint64(uint64FromInt32(graphBlockCount)*uint64(12)), __ccgo_ts+33063)
	// Split an awake island. This modifies:
	// - stack allocator
	// - world island array and solver set
	// - island indices on bodies, contacts, and joints
	// I'm squeezing this task in here because it may be expensive and this is a safe place to put it.
	// Note: cannot split islands in parallel with FinalizeBodies
	splitIslandTask = uintptrFromInt32(0)
	if (*b3World)(unsafe.Pointer(world)).SplitIslandId != -int32(1) {
		if (*b3World)(unsafe.Pointer(world)).TaskCount < int32(_B3_MAX_TASKS) {
			splitIslandTask = (*(*func(*_Stack, uintptr, uintptr, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b3SplitIslandTask), world, (*b3World)(unsafe.Pointer(world)).UserTaskContext, __ccgo_ts+33076)
			*(*int32)(unsafe.Add(nil, world+4996)) += int32(1)
			if splitIslandTask == uintptrFromInt32(0) {
				v1 = 0
			} else {
				v1 = int32(1)
			}
			*(*int32)(unsafe.Add(nil, world+4992)) += v1
		} else {
			b3SplitIslandTask(tls, world)
		}
	}
	// Prepare body blocks
	b3InitBlocks(tls, bodyBlocks, bodyDim, awakeBodyCount, uint8(b3_bodyBlock), uint8FromInt32(int32FromInt32(_UINT8_MAX7)))
	// Prepare blocks as a single flat parallel-for over the whole constraint range.
	// The task walks spans to decode flat slot indices back to per-color arrays.
	b3InitBlocks(tls, convexBlocks, convexPrepareDim, wideContactCount, uint8(b3_wideContactBlock), uint8FromInt32(int32FromInt32(_UINT8_MAX7)))
	b3InitBlocks(tls, meshBlocks, meshPrepareDim, meshGroupCount, uint8(b3_contactBlock), uint8FromInt32(int32FromInt32(_UINT8_MAX7)))
	b3InitBlocks(tls, jointBlocks, jointPrepareDim, jointCount, uint8(b3_jointBlock), uint8FromInt32(int32FromInt32(_UINT8_MAX7)))
	// Prepare graph work blocks. Each color gets joint blocks followed by contact blocks.
	*(*[24]uintptr)(unsafe.Add(nil, bp+1536)) = [24]uintptr{}
	baseGraphBlock = graphBlocks
	*(*[24]int32)(unsafe.Add(nil, bp+1728)) = [24]int32{}
	i6 = 0
	for {
		if !(i6 < activeColorCount) {
			break
		}
		(*(*[24]uintptr)(unsafe.Add(nil, bp+1536)))[i6] = baseGraphBlock
		colorIndex = uint8FromInt32((*(*[24]int32)(unsafe.Add(nil, bp+8)))[i6])
		b3InitBlocks(tls, baseGraphBlock, graphJointDims[i6], colorJointCounts[i6], uint8(b3_graphJointBlock), colorIndex)
		baseGraphBlock = baseGraphBlock + uintptr(graphJointDims[i6].Count)*12
		b3InitBlocks(tls, baseGraphBlock, graphWideContactDims[i6], colorWideContactCounts[i6], uint8(b3_graphWideContactBlock), colorIndex)
		baseGraphBlock = baseGraphBlock + uintptr(graphWideContactDims[i6].Count)*12
		b3InitBlocks(tls, baseGraphBlock, graphMeshContactDims[i6], colorMeshGroupCounts[i6], uint8(b3_graphContactBlock), colorIndex)
		baseGraphBlock = baseGraphBlock + uintptr(graphMeshContactDims[i6].Count)*12
		(*(*[24]int32)(unsafe.Add(nil, bp+1728)))[i6] = graphJointDims[i6].Count + graphWideContactDims[i6].Count + graphMeshContactDims[i6].Count
		goto _32
	_32:
		;
		i6 = i6 + 1
	}
	if v2 = !!((int64(baseGraphBlock)-int64(graphBlocks))/12 == int64(graphBlockCount)); !v2 {
		b3InternalAssert(tls, __ccgo_ts+33082, __ccgo_ts+31091, int32FromInt32(1857))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	stage = stages
	stage = b3InitStage(tls, stage, int32(b3_stagePrepareJoints), jointBlocks, jointPrepareDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX7)))
	stage = b3InitStage(tls, stage, int32(b3_stagePrepareWideContacts), convexBlocks, convexPrepareDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX7)))
	stage = b3InitStage(tls, stage, int32(b3_stagePrepareContacts), meshBlocks, meshPrepareDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX7)))
	stage = b3InitStage(tls, stage, int32(b3_stageIntegrateVelocities), bodyBlocks, bodyDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX7)))
	stage = b3InitColorStages(tls, stage, int32(b3_stageWarmStart), int32(1), activeColorCount, bp+1536, bp+1728, bp+8)
	stage = b3InitColorStages(tls, stage, int32(b3_stageSolve), int32(_ITERATIONS), activeColorCount, bp+1536, bp+1728, bp+8)
	stage = b3InitStage(tls, stage, int32(b3_stageIntegratePositions), bodyBlocks, bodyDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX7)))
	stage = b3InitColorStages(tls, stage, int32(b3_stageRelax), int32(_RELAX_ITERATIONS), activeColorCount, bp+1536, bp+1728, bp+8)
	stage = b3InitColorStages(tls, stage, int32(b3_stageRestitution), int32(1), activeColorCount, bp+1536, bp+1728, bp+8)
	stage = b3InitStage(tls, stage, int32(b3_stageStoreWideImpulses), convexBlocks, convexPrepareDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX7)))
	stage = b3InitStage(tls, stage, int32(b3_stageStoreImpulses), meshBlocks, meshPrepareDim.Count, uint8FromInt32(int32FromInt32(_UINT8_MAX7)))
	if v2 = !!(int32((int64(stage)-int64(stages))/24) == stageCount); !v2 {
		b3InternalAssert(tls, __ccgo_ts+33145, __ccgo_ts+31091, int32FromInt32(1876))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	if v2 = !!(workerCount <= int32FromInt32(_B3_MAX_WORKERS)); !v2 {
		b3InternalAssert(tls, __ccgo_ts+33183, __ccgo_ts+31091, int32FromInt32(1878))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	(*b3StepContext)(unsafe.Pointer(stepContext)).Graph = graph
	(*b3StepContext)(unsafe.Pointer(stepContext)).ActiveColorCount = activeColorCount
	(*b3StepContext)(unsafe.Pointer(stepContext)).WorkerCount = workerCount
	(*b3StepContext)(unsafe.Pointer(stepContext)).StageCount = stageCount
	(*b3StepContext)(unsafe.Pointer(stepContext)).Stages = stages
	(*b3StepContext)(unsafe.Pointer(stepContext)).WideConstraints = wideConstraints
	(*b3StepContext)(unsafe.Pointer(stepContext)).WidePrepareSpans = bp + 104
	(*b3StepContext)(unsafe.Pointer(stepContext)).WideContactCount = wideContactCount
	(*b3StepContext)(unsafe.Pointer(stepContext)).OverflowSpans = bp + 1504
	(*b3StepContext)(unsafe.Pointer(stepContext)).WideMeshConstraints = wideMeshConstraints
	(*b3StepContext)(unsafe.Pointer(stepContext)).WideMeshManifoldStarts = meshManifoldStarts
	(*b3StepContext)(unsafe.Pointer(stepContext)).MeshPrepareSpans = bp + 504
	(*b3StepContext)(unsafe.Pointer(stepContext)).JointPrepareSpans = bp + 1104
	atomicStoreNUint32(stepContext+272, uint32(0), int32FromInt32(__ATOMIC_SEQ_CST))
	atomicStoreNInt32(stepContext+340, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	(*b3World)(unsafe.Pointer(world)).Profile.SolverSetup = b3GetMillisecondsAndReset(tls, bp)
	*(*uint64_t)(unsafe.Add(nil, bp+2592)) = b3GetTicks(tls)
	v1 = (*b3IdPool)(unsafe.Pointer(world + 4144)).NextIndex
	goto _37
_37:
	jointIdCapacity = v1
	v4 = (*b3IdPool)(unsafe.Pointer(world + 4184)).NextIndex
	goto _39
_39:
	contactIdCapacity = v4
	i7 = 0
	for {
		if !(i7 < workerCount) {
			break
		}
		if v2 = !!(0 <= i7 && i7 < (*b3World)(unsafe.Pointer(world)).TaskContexts.Count); !v2 {
			b3InternalAssert(tls, __ccgo_ts+33213, __ccgo_ts+31091, int32FromInt32(1907))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		taskContext = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i7)*3848
		b3SetBitCountAndClear(tls, taskContext+72, uint32FromInt32(jointIdCapacity))
		b3SetBitCountAndClear(tls, taskContext+88, uint32FromInt32(contactIdCapacity))
		(*b3TaskContext)(unsafe.Pointer(taskContext)).HasHitEvents = boolUint8(false1 != 0)
		(*(*[32]b3WorkerContext)(unsafe.Add(nil, bp+1824)))[i7].Context = stepContext
		(*(*[32]b3WorkerContext)(unsafe.Add(nil, bp+1824)))[i7].WorkerIndex = i7
		if (*b3World)(unsafe.Pointer(world)).TaskCount < int32(_B3_MAX_TASKS) {
			__builtin_snprintf(tls, bp+2600, uint64(16), __ccgo_ts+33265, vaList(bp+2648, i7))
			(*(*[32]b3WorkerContext)(unsafe.Add(nil, bp+1824)))[i7].UserTask = (*(*func(*_Stack, uintptr, uintptr, uintptr, uintptr) uintptr)(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world)).EnqueueTaskFcn})))(tls, __ccgo_fp(b3SolverTask), bp+1824+uintptr(i7)*24, (*b3World)(unsafe.Pointer(world)).UserTaskContext, bp+2600)
			*(*int32)(unsafe.Add(nil, world+4996)) += int32(1)
			if (*(*[32]b3WorkerContext)(unsafe.Add(nil, bp+1824)))[i7].UserTask == uintptrFromInt32(0) {
				v1 = 0
			} else {
				v1 = int32(1)
			}
			*(*int32)(unsafe.Add(nil, world+4992)) += v1
		} else {
			(*(*[32]b3WorkerContext)(unsafe.Add(nil, bp+1824)))[i7].UserTask = uintptrFromInt32(0)
			b3SolverTask(tls, bp+1824+uintptr(i7)*24)
		}
		goto _40
	_40:
		;
		i7 = i7 + 1
	}
	// The calling thread of b3World_Step also enters b3SolverTask as worker 0 and races for the
	// orchestrator slot via the CAS inside. This guarantees progress even when the user's task
	// system can't run the queued worker 0 promptly: it might schedule out of order, have fewer
	// threads than workerCount, or invert priority by parking the calling thread in finishTaskFcn.
	// Whoever wins the CAS becomes the orchestrator; the loser returns and lets the spinner-only
	// path handle workers >0.
	*(*b3WorkerContext)(unsafe.Add(nil, bp+2616)) = b3WorkerContext{
		Context: stepContext,
	}
	b3SolverTask(tls, bp+2616)
	// Finish constraint solve
	i8 = 0
	for {
		if !(i8 < workerCount) {
			break
		}
		if (*(*[32]b3WorkerContext)(unsafe.Add(nil, bp+1824)))[i8].UserTask != uintptrFromInt32(0) {
			(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, (*(*[32]b3WorkerContext)(unsafe.Add(nil, bp+1824)))[i8].UserTask, (*b3World)(unsafe.Pointer(world)).UserTaskContext)
			*(*int32)(unsafe.Add(nil, world+4992)) -= int32(1)
		}
		goto _43
	_43:
		;
		i8 = i8 + 1
	}
	// Finish island split
	if splitIslandTask != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, splitIslandTask, (*b3World)(unsafe.Pointer(world)).UserTaskContext)
		*(*int32)(unsafe.Add(nil, world+4992)) -= int32(1)
	}
	(*b3World)(unsafe.Pointer(world)).SplitIslandId = -int32(1)
	(*b3World)(unsafe.Pointer(world)).Profile.Constraints = b3GetMillisecondsAndReset(tls, bp+2592)
	transformTicks = b3GetTicks(tls)
	// Prepare the island bit set used in body finalization.
	awakeIslandCount = (*b3SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Count
	i9 = 0
	for {
		if !(i9 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		taskContext1 = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i9)*3848
		(*b3TaskContext)(unsafe.Pointer(taskContext1)).SensorHits.Count = 0
		b3SetBitCountAndClear(tls, taskContext1+112, uint32FromInt32(awakeIslandCount))
		(*b3TaskContext)(unsafe.Pointer(taskContext1)).SplitIslandId = -int32(1)
		(*b3TaskContext)(unsafe.Pointer(taskContext1)).SplitSleepTime = float32FromFloat32(0)
		goto _44
	_44:
		;
		i9 = i9 + 1
	}
	// Finish the user tree task that was queued earlier in the time step. This must be complete before touching the
	// broad-phase.
	if (*b3World)(unsafe.Pointer(world)).UserTreeTask != uintptrFromInt32(0) {
		(*(*func(*_Stack, uintptr, uintptr))(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world)).FinishTaskFcn})))(tls, (*b3World)(unsafe.Pointer(world)).UserTreeTask, (*b3World)(unsafe.Pointer(world)).UserTaskContext)
		(*b3World)(unsafe.Pointer(world)).UserTreeTask = uintptrFromInt32(0)
		*(*int32)(unsafe.Add(nil, world+4992)) -= int32(1)
	}
	b3ValidateNoMoved(tls, world+800)
	// Finalize bodies. Must happen after the constraint solver and after island splitting.
	b3ParallelFor(tls, world, __ccgo_fp(b3FinalizeBodiesTask), awakeBodyCount, int32(8), stepContext, __ccgo_ts+33275)
	// Free in reverse order
	b3StackFree(tls, world, graphBlocks)
	b3StackFree(tls, world, jointBlocks)
	b3StackFree(tls, world, meshBlocks)
	b3StackFree(tls, world, convexBlocks)
	b3StackFree(tls, world, bodyBlocks)
	b3StackFree(tls, world, stages)
	b3StackFree(tls, world, (*b3GraphColor)(unsafe.Pointer(overflow)).ManifoldConstraints)
	b3StackFree(tls, world, (*b3GraphColor)(unsafe.Pointer(overflow)).ContactConstraints)
	b3StackFree(tls, world, wideMeshConstraints)
	b3StackFree(tls, world, meshManifoldStarts)
	b3StackFree(tls, world, meshLaneOrder)
	b3StackFree(tls, world, wideConstraints)
	(*b3World)(unsafe.Pointer(world)).Profile.Transforms = b3GetMilliseconds(tls, transformTicks)
	// Report joint events
	jointEventTicks = b3GetTicks(tls)
	// Gather bits for all joints that have force/torque events
	jointStateBitSet = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + 72
	i10 = int32(1)
	for {
		if !(i10 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		b3InPlaceUnion(tls, jointStateBitSet, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i10)*3848+72)
		goto _45
	_45:
		;
		i10 = i10 + 1
	}
	wordCount = (*b3BitSet)(unsafe.Pointer(jointStateBitSet)).BlockCount
	bits = (*b3BitSet)(unsafe.Pointer(jointStateBitSet)).Bits
	jointArray = (*b3World)(unsafe.Pointer(world)).Joints.Data
	worldIndex0 = (*b3World)(unsafe.Pointer(world)).WorldId
	k2 = uint32(0)
	for {
		if !(k2 < wordCount) {
			break
		}
		word = *(*uint64_t)(unsafe.Add(nil, bits+uintptr(k2)*8))
		for word != uint64(0) {
			v47 = uint32FromInt32(__builtin_ctzll(tls, word))
			goto _48
		_48:
			ctz = v47
			jointId = int32FromUint32(uint32FromInt32(64)*k2 + ctz)
			if v2 = !!(jointId < (*b3World)(unsafe.Pointer(world)).Joints.Capacity); !v2 {
				b3InternalAssert(tls, __ccgo_ts+33279, __ccgo_ts+31091, int32FromInt32(2034))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			joint = jointArray + uintptr(jointId)*72
			if v2 = !!((*b3Joint)(unsafe.Pointer(joint)).SetIndex == int32(b3_awakeSet)); !v2 {
				b3InternalAssert(tls, __ccgo_ts+33312, __ccgo_ts+31091, int32FromInt32(2038))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			event = JointEvent{
				JointId: JointId{
					Index1:     jointId + int32(1),
					World0:     worldIndex0,
					Generation: (*b3Joint)(unsafe.Pointer(joint)).Generation,
				},
				UserData: uint64((*b3Joint)(unsafe.Pointer(joint)).UserData),
			}
			if (*b3World)(unsafe.Pointer(world)).JointEvents.Count >= (*b3World)(unsafe.Pointer(world)).JointEvents.Capacity {
				oldSize1 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).JointEvents.Capacity) * uint64(16))
				if (*b3World)(unsafe.Pointer(world)).JointEvents.Capacity == 0 {
					v1 = int32(8)
				} else {
					v1 = int32(2) * (*b3World)(unsafe.Pointer(world)).JointEvents.Capacity
				}
				newCapacity = v1
				newSize1 = int32FromUint64(uint64FromInt32(newCapacity) * uint64(16))
				(*b3World)(unsafe.Pointer(world)).JointEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).JointEvents.Data, oldSize1, newSize1)
				(*b3World)(unsafe.Pointer(world)).JointEvents.Capacity = newCapacity
			}
			v21 = world + 4544 + 8
			v1 = *(*int32)(unsafe.Pointer(v21))
			*(*int32)(unsafe.Pointer(v21)) = *(*int32)(unsafe.Pointer(v21)) + 1
			*(*JointEvent)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).JointEvents.Data+uintptr(v1)*16)) = event
			// Clear the smallest set bit
			word = word & (word - uint64(1))
		}
		goto _46
	_46:
		;
		k2 = k2 + 1
	}
	(*b3World)(unsafe.Pointer(world)).Profile.JointEvents = b3GetMilliseconds(tls, jointEventTicks)
	// Report hit events
	hitTicks = b3GetTicks(tls)
	if v2 = !!((*b3World)(unsafe.Pointer(world)).ContactHitEvents.Count == int32FromInt32(0)); !v2 {
		b3InternalAssert(tls, __ccgo_ts+33343, __ccgo_ts+31091, int32FromInt32(2067))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	// Fast path: if no worker flagged any hit-event candidates during b2StoreImpulsesTask, skip entirely.
	anyHitEvents = boolUint8(false1 != 0)
	i11 = 0
	for {
		if !(i11 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
			break
		}
		if (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i11)*3848))).HasHitEvents != 0 {
			anyHitEvents = boolUint8(true1 != 0)
			break
		}
		goto _55
	_55:
		;
		i11 = i11 + 1
	}
	if anyHitEvents != 0 {
		// Union per-worker bits into worker 0's bit set.
		hitEventBitSet = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + 88
		i12 = int32(1)
		for {
			if !(i12 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
				break
			}
			if (*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i12)*3848))).HasHitEvents != 0 {
				b3InPlaceUnion(tls, hitEventBitSet, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i12)*3848+88)
			}
			goto _56
		_56:
			;
			i12 = i12 + 1
		}
		threshold = (*b3World)(unsafe.Pointer(world)).HitEventThreshold
		contactArray = (*b3World)(unsafe.Pointer(world)).Contacts.Data
		worldId = (*b3World)(unsafe.Pointer(world)).WorldId
		wordCount1 = (*b3BitSet)(unsafe.Pointer(hitEventBitSet)).BlockCount
		bits1 = (*b3BitSet)(unsafe.Pointer(hitEventBitSet)).Bits
		k3 = uint32(0)
		for {
			if !(k3 < wordCount1) {
				break
			}
			word1 = *(*uint64_t)(unsafe.Add(nil, bits1+uintptr(k3)*8))
			for word1 != uint64(0) {
				v47 = uint32FromInt32(__builtin_ctzll(tls, word1))
				goto _59
			_59:
				ctz1 = v47
				contactId = int32FromUint32(uint32FromInt32(64)*k3 + ctz1)
				contact = contactArray + uintptr(contactId)*216
				if v2 = !!((*b3Contact)(unsafe.Pointer(contact)).SetIndex == int32(b3_awakeSet) && (*b3Contact)(unsafe.Pointer(contact)).ColorIndex != -int32(1)); !v2 {
					b3InternalAssert(tls, __ccgo_ts+33378, __ccgo_ts+31091, int32FromInt32(2107))
				}
				_ = v2 || bool(int32FromInt32(0) != 0)
				if v2 = !!(0 <= (*b3Contact)(unsafe.Pointer(contact)).ShapeIdA && (*b3Contact)(unsafe.Pointer(contact)).ShapeIdA < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v2 {
					b3InternalAssert(tls, __ccgo_ts+2313, __ccgo_ts+31091, int32FromInt32(2109))
				}
				_ = v2 || bool(int32FromInt32(0) != 0)
				shapeA = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Contact)(unsafe.Pointer(contact)).ShapeIdA)*208
				if v23 = !!(0 <= (*b3Contact)(unsafe.Pointer(contact)).ShapeIdB && (*b3Contact)(unsafe.Pointer(contact)).ShapeIdB < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v23 {
					b3InternalAssert(tls, __ccgo_ts+2391, __ccgo_ts+31091, int32FromInt32(2110))
				}
				_ = v23 || bool(int32FromInt32(0) != 0)
				shapeB = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr((*b3Contact)(unsafe.Pointer(contact)).ShapeIdB)*208
				if v24 = !!(0 <= (*b3Shape)(unsafe.Pointer(shapeA)).BodyId && (*b3Shape)(unsafe.Pointer(shapeA)).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v24 {
					b3InternalAssert(tls, __ccgo_ts+7540, __ccgo_ts+31091, int32FromInt32(2111))
				}
				_ = v24 || bool(int32FromInt32(0) != 0)
				bodyA = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3Shape)(unsafe.Pointer(shapeA)).BodyId)*96
				if v27 = !!(0 <= (*b3Shape)(unsafe.Pointer(shapeB)).BodyId && (*b3Shape)(unsafe.Pointer(shapeB)).BodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v27 {
					b3InternalAssert(tls, __ccgo_ts+7612, __ccgo_ts+31091, int32FromInt32(2112))
				}
				_ = v27 || bool(int32FromInt32(0) != 0)
				bodyB = (*b3World)(unsafe.Pointer(world)).Bodies.Data + uintptr((*b3Shape)(unsafe.Pointer(shapeB)).BodyId)*96
				simA = b3GetBodySim(tls, world, bodyA)
				simB = b3GetBodySim(tls, world, bodyB)
				v65 = (*b3BodySim)(unsafe.Pointer(simA)).Center
				v66 = (*b3BodySim)(unsafe.Pointer(simB)).Center
				v67 = float32FromFloat32(0.5)
				v68 = Pos{
					X: float32((float32FromFloat32(1)-v67)*v65.X) + float32(v67*v66.X),
					Y: float32((float32FromFloat32(1)-v67)*v65.Y) + float32(v67*v66.Y),
					Z: float32((float32FromFloat32(1)-v67)*v65.Z) + float32(v67*v66.Z),
				}
				goto _69
			_69:
				midCenter = v68
				event1 = ContactHitEvent{}
				event1.ApproachSpeed = threshold
				found = boolUint8(false1 != 0)
				triangleIndex = 0
				manifoldCount = (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount
				i13 = 0
				for {
					if !(i13 < manifoldCount) {
						break
					}
					manifold = (*b3Contact)(unsafe.Pointer(contact)).Manifolds + uintptr(i13)*268
					pointCount = (*Manifold)(unsafe.Pointer(manifold)).PointCount
					p1 = 0
					for {
						if !(p1 < pointCount) {
							break
						}
						mp = manifold + uintptr(p1)*56
						approachSpeed = -(*ManifoldPoint)(unsafe.Pointer(mp)).NormalVelocity
						// Need to check total impulse because the point may be speculative and not colliding
						if approachSpeed > event1.ApproachSpeed && (*ManifoldPoint)(unsafe.Pointer(mp)).TotalNormalImpulse > float32FromFloat32(0) {
							event1.ApproachSpeed = approachSpeed
							v72 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorA
							v73 = (*ManifoldPoint)(unsafe.Pointer(mp)).AnchorB
							v67 = float32FromFloat32(0.5)
							if v2 = !!(float32FromFloat32(0) <= v67 && v67 <= float32FromFloat32(1)); !v2 {
								b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
							}
							_ = v2 || bool(int32FromInt32(0) != 0)
							c = Vec3{
								X: float32((float32FromFloat32(1)-v67)*v72.X) + float32(v67*v73.X),
								Y: float32((float32FromFloat32(1)-v67)*v72.Y) + float32(v67*v73.Y),
								Z: float32((float32FromFloat32(1)-v67)*v72.Z) + float32(v67*v73.Z),
							}
							v76 = c
							goto _77
						_77:
							v65 = midCenter
							v79 = v76
							v66 = Pos{
								X: v65.X + v79.X,
								Y: v65.Y + v79.Y,
								Z: v65.Z + v79.Z,
							}
							goto _81
						_81:
							event1.Point = v66
							event1.Normal = (*Manifold)(unsafe.Pointer(manifold)).Normal
							triangleIndex = (*ManifoldPoint)(unsafe.Pointer(mp)).TriangleIndex
							found = boolUint8(true1 != 0)
						}
						goto _71
					_71:
						;
						p1 = p1 + 1
					}
					goto _70
				_70:
					;
					i13 = i13 + 1
				}
				if int32FromUint8(found) == int32(true1) {
					event1.ShapeIdA = ShapeId{
						Index1:     (*b3Shape)(unsafe.Pointer(shapeA)).Id + int32(1),
						World0:     worldId,
						Generation: (*b3Shape)(unsafe.Pointer(shapeA)).Generation,
					}
					event1.ShapeIdB = ShapeId{
						Index1:     (*b3Shape)(unsafe.Pointer(shapeB)).Id + int32(1),
						World0:     worldId,
						Generation: (*b3Shape)(unsafe.Pointer(shapeB)).Generation,
					}
					event1.ContactId = ContactId{
						Index1:     (*b3Contact)(unsafe.Pointer(contact)).ContactId + int32(1),
						World0:     worldId,
						Generation: (*b3Contact)(unsafe.Pointer(contact)).Generation,
					}
					// shapeB is never a compound today (asserted in b3CreateContact), so the
					// childIndex argument is irrelevant for it. shapeA carries the compound.
					event1.UserMaterialIdA = b3GetShapeUserMaterialId(tls, shapeA, (*b3Contact)(unsafe.Pointer(contact)).ChildIndex, triangleIndex)
					event1.UserMaterialIdB = b3GetShapeUserMaterialId(tls, shapeB, 0, triangleIndex)
					if (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Count >= (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Capacity {
						oldSize2 = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(world)).ContactHitEvents.Capacity) * uint64(72))
						if (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Capacity == 0 {
							v1 = int32(8)
						} else {
							v1 = int32(2) * (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Capacity
						}
						newCapacity1 = v1
						newSize2 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(72))
						(*b3World)(unsafe.Pointer(world)).ContactHitEvents.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Data, oldSize2, newSize2)
						(*b3World)(unsafe.Pointer(world)).ContactHitEvents.Capacity = newCapacity1
					}
					v21 = world + 4528 + 8
					v1 = *(*int32)(unsafe.Pointer(v21))
					*(*int32)(unsafe.Pointer(v21)) = *(*int32)(unsafe.Pointer(v21)) + 1
					*(*ContactHitEvent)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).ContactHitEvents.Data+uintptr(v1)*72)) = event1
				}
				// Clear the smallest set bit
				word1 = word1 & (word1 - uint64(1))
			}
			goto _57
		_57:
			;
			k3 = k3 + 1
		}
	}
	(*b3World)(unsafe.Pointer(world)).Profile.HitEvents = b3GetMilliseconds(tls, hitTicks)
	refitTicks = b3GetTicks(tls)
	b3DynamicTree_Refit(tls, world+800+uintptr(b3_kinematicBody)*112)
	b3DynamicTree_Refit(tls, world+800+uintptr(b3_dynamicBody)*112)
	b3ValidateBroadPhase(tls, world+800)
	(*b3World)(unsafe.Pointer(world)).Profile.Refit = b3GetMilliseconds(tls, refitTicks)
	v1 = atomicLoadNInt32(stepContext+104, int32FromInt32(__ATOMIC_SEQ_CST))
	goto _86
_86:
	bulletBodyCount = v1
	if bulletBodyCount > 0 {
		bulletTicks = b3GetTicks(tls)
		// Fast bullet bodies
		// Note: a bullet body may be moving slow
		minRange = int32(8)
		b3ParallelFor(tls, world, __ccgo_fp(b3BulletBodyTask), bulletBodyCount, minRange, stepContext, __ccgo_ts+33451)
		// Serially enlarge broad-phase proxies for bullet shapes
		broadPhase = world + 800
		dynamicTree = broadPhase + uintptr(b3_dynamicBody)*112
		// Fast array access is important here
		bodyArray = (*b3World)(unsafe.Pointer(world)).Bodies.Data
		bodySimArray = (*b3SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Data
		shapeArray = (*b3World)(unsafe.Pointer(world)).Shapes.Data
		// Serially enlarge broad-phase proxies for bullet shapes
		bulletBodySimIndices = (*b3StepContext)(unsafe.Pointer(stepContext)).BulletBodies
		// This loop has non-deterministic order but it shouldn't affect the result
		i14 = 0
		for {
			if !(i14 < bulletBodyCount) {
				break
			}
			bulletBodySim = bodySimArray + uintptr(*(*int32)(unsafe.Add(nil, bulletBodySimIndices+uintptr(i14)*4)))*252
			if (*b3BodySim)(unsafe.Pointer(bulletBodySim)).Flags&uint32(b3_enlargeBulletBounds) == uint32(0) {
				goto _87
			}
			// Clear flag
			*(*uint32_t)(unsafe.Add(nil, bulletBodySim+40)) &= uint32FromInt32(^int32(b3_enlargeBulletBounds))
			bodyId = (*b3BodySim)(unsafe.Pointer(bulletBodySim)).BodyId
			if v2 = !!(0 <= bodyId && bodyId < (*b3World)(unsafe.Pointer(world)).Bodies.Count); !v2 {
				b3InternalAssert(tls, __ccgo_ts+33459, __ccgo_ts+31091, int32FromInt32(2223))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			bulletBody = bodyArray + uintptr(bodyId)*96
			shapeId = (*b3Body)(unsafe.Pointer(bulletBody)).HeadShapeId
			for shapeId != -int32(1) {
				shape = shapeArray + uintptr(shapeId)*208
				proxyKey = (*b3Shape)(unsafe.Pointer(shape)).ProxyKey
				proxyId1 = proxyKey >> int32(2)
				v21 = dynamicTree
				v90 = (*(*b3TreeNode)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(v21)).Nodes+uintptr((*(*b3TreeProxy)(unsafe.Add(nil, (*b3DynamicTree)(unsafe.Pointer(v21)).Proxies+uintptr(proxyId1)*24))).Node)*32))).Aabb
				goto _91
			_91:
				treeAABB = v90
				shapeFatAABB = *(*AABB)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world)).FatAABBs.Data+uintptr(shapeId)*24))
				v92 = treeAABB
				v93 = shapeFatAABB
				if v92.LowerBound.X > v93.LowerBound.X || v93.UpperBound.X > v92.UpperBound.X {
					v94 = boolUint8(false1 != 0)
					goto _95
				}
				if v92.LowerBound.Y > v93.LowerBound.Y || v93.UpperBound.Y > v92.UpperBound.Y {
					v94 = boolUint8(false1 != 0)
					goto _95
				}
				if v92.LowerBound.Z > v93.LowerBound.Z || v93.UpperBound.Z > v92.UpperBound.Z {
					v94 = boolUint8(false1 != 0)
					goto _95
				}
				v94 = boolUint8(true1 != 0)
				goto _95
			_95:
				if int32FromUint8(v94) == false1 {
					b3DynamicTree_EnlargeProxy(tls, dynamicTree, proxyId1, shapeFatAABB)
				}
				shapeId = (*b3Shape)(unsafe.Pointer(shape)).NextShapeId
			}
			goto _87
		_87:
			;
			i14 = i14 + 1
		}
		(*b3World)(unsafe.Pointer(world)).Profile.Bullets = b3GetMilliseconds(tls, bulletTicks)
	}
	b3StackFree(tls, world, (*b3StepContext)(unsafe.Pointer(stepContext)).BulletBodies)
	(*b3StepContext)(unsafe.Pointer(stepContext)).BulletBodies = uintptrFromInt32(0)
	atomicStoreNInt32(stepContext+104, 0, int32FromInt32(__ATOMIC_SEQ_CST))
	// Report sensor hits. This may include bullets sensor hits.
	sensorHitTicks = b3GetTicks(tls)
	workerCount1 = (*b3World)(unsafe.Pointer(world)).WorkerCount
	if v2 = !!(workerCount1 == (*b3World)(unsafe.Pointer(world)).TaskContexts.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+33503, __ccgo_ts+31091, int32FromInt32(2261))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	i15 = 0
	for {
		if !(i15 < workerCount1) {
			break
		}
		taskContext2 = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i15)*3848
		hitCount = (*b3TaskContext)(unsafe.Pointer(taskContext2)).SensorHits.Count
		hits = (*b3TaskContext)(unsafe.Pointer(taskContext2)).SensorHits.Data
		j4 = 0
		for {
			if !(j4 < hitCount) {
				break
			}
			hit = *(*b3SensorHit)(unsafe.Add(nil, hits+uintptr(j4)*8))
			if v2 = !!(0 <= hit.SensorId && hit.SensorId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v2 {
				b3InternalAssert(tls, __ccgo_ts+33544, __ccgo_ts+31091, int32FromInt32(2272))
			}
			_ = v2 || bool(int32FromInt32(0) != 0)
			sensorShape = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(hit.SensorId)*208
			if v23 = !!(0 <= hit.VisitorId && hit.VisitorId < (*b3World)(unsafe.Pointer(world)).Shapes.Count); !v23 {
				b3InternalAssert(tls, __ccgo_ts+33612, __ccgo_ts+31091, int32FromInt32(2273))
			}
			_ = v23 || bool(int32FromInt32(0) != 0)
			visitor = (*b3World)(unsafe.Pointer(world)).Shapes.Data + uintptr(hit.VisitorId)*208
			if v24 = !!(0 <= (*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex && (*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex < (*b3World)(unsafe.Pointer(world)).Sensors.Count); !v24 {
				b3InternalAssert(tls, __ccgo_ts+28656, __ccgo_ts+31091, int32FromInt32(2275))
			}
			_ = v24 || bool(int32FromInt32(0) != 0)
			sensor = (*b3World)(unsafe.Pointer(world)).Sensors.Data + uintptr((*b3Shape)(unsafe.Pointer(sensorShape)).SensorIndex)*56
			shapeRef = b3Visitor{
				ShapeId:    hit.VisitorId,
				Generation: (*b3Shape)(unsafe.Pointer(visitor)).Generation,
			}
			if (*b3Sensor)(unsafe.Pointer(sensor)).Hits.Count >= (*b3Sensor)(unsafe.Pointer(sensor)).Hits.Capacity {
				oldSize3 = int32FromUint64(uint64FromInt32((*b3Sensor)(unsafe.Pointer(sensor)).Hits.Capacity) * uint64(8))
				if (*b3Sensor)(unsafe.Pointer(sensor)).Hits.Capacity == 0 {
					v1 = int32(8)
				} else {
					v1 = int32(2) * (*b3Sensor)(unsafe.Pointer(sensor)).Hits.Capacity
				}
				newCapacity2 = v1
				newSize3 = int32FromUint64(uint64FromInt32(newCapacity2) * uint64(8))
				(*b3Sensor)(unsafe.Pointer(sensor)).Hits.Data = b3GrowAlloc(tls, (*b3Sensor)(unsafe.Pointer(sensor)).Hits.Data, oldSize3, newSize3)
				(*b3Sensor)(unsafe.Pointer(sensor)).Hits.Capacity = newCapacity2
			}
			v21 = sensor + 8
			v1 = *(*int32)(unsafe.Pointer(v21))
			*(*int32)(unsafe.Pointer(v21)) = *(*int32)(unsafe.Pointer(v21)) + 1
			*(*b3Visitor)(unsafe.Add(nil, (*b3Sensor)(unsafe.Pointer(sensor)).Hits.Data+uintptr(v1)*8)) = shapeRef
			goto _98
		_98:
			;
			j4 = j4 + 1
		}
		goto _97
	_97:
		;
		i15 = i15 + 1
	}
	(*b3World)(unsafe.Pointer(world)).Profile.SensorHits = b3GetMilliseconds(tls, sensorHitTicks)
	// Island sleeping
	// This must be done last because putting islands to sleep invalidates the enlarged body bits.
	// todo_erin figure out how to do this in parallel with tree refit
	if int32FromUint8((*b3World)(unsafe.Pointer(world)).EnableSleep) == int32(true1) {
		sleepTicks = b3GetTicks(tls)
		// Collect split island candidate for the next time step. No need to split if sleeping is disabled.
		if v2 = !!((*b3World)(unsafe.Pointer(world)).SplitIslandId == -int32FromInt32(1)); !v2 {
			b3InternalAssert(tls, __ccgo_ts+33682, __ccgo_ts+31091, int32FromInt32(2297))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		splitSleepTimer = float32FromFloat32(0)
		i16 = 0
		for {
			if !(i16 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
				break
			}
			taskContext3 = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + uintptr(i16)*3848
			if (*b3TaskContext)(unsafe.Pointer(taskContext3)).SplitIslandId != -int32(1) && (*b3TaskContext)(unsafe.Pointer(taskContext3)).SplitSleepTime >= splitSleepTimer {
				if v2 = !!((*b3TaskContext)(unsafe.Pointer(taskContext3)).SplitSleepTime > float32FromFloat32(0)); !v2 {
					b3InternalAssert(tls, __ccgo_ts+33720, __ccgo_ts+31091, int32FromInt32(2304))
				}
				_ = v2 || bool(int32FromInt32(0) != 0)
				// Tie breaking for determinism. Largest island id wins. Needed due to work stealing.
				if (*b3TaskContext)(unsafe.Pointer(taskContext3)).SplitSleepTime == splitSleepTimer && (*b3TaskContext)(unsafe.Pointer(taskContext3)).SplitIslandId < (*b3World)(unsafe.Pointer(world)).SplitIslandId {
					goto _106
				}
				(*b3World)(unsafe.Pointer(world)).SplitIslandId = (*b3TaskContext)(unsafe.Pointer(taskContext3)).SplitIslandId
				splitSleepTimer = (*b3TaskContext)(unsafe.Pointer(taskContext3)).SplitSleepTime
			}
			goto _106
		_106:
			;
			i16 = i16 + 1
		}
		awakeIslandBitSet = (*b3World)(unsafe.Pointer(world)).TaskContexts.Data + 112
		i17 = int32(1)
		for {
			if !(i17 < (*b3World)(unsafe.Pointer(world)).WorkerCount) {
				break
			}
			b3InPlaceUnion(tls, awakeIslandBitSet, (*b3World)(unsafe.Pointer(world)).TaskContexts.Data+uintptr(i17)*3848+112)
			goto _108
		_108:
			;
			i17 = i17 + 1
		}
		// Need to process in reverse because this moves islands to sleeping solver sets.
		islands = (*b3SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Data
		count1 = (*b3SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Count
		islandIndex = count1 - int32(1)
		for {
			if !(islandIndex >= 0) {
				break
			}
			v21 = awakeIslandBitSet
			v47 = uint32FromInt32(islandIndex)
			blockIndex = v47 / uint32(64)
			if blockIndex >= (*b3BitSet)(unsafe.Pointer(v21)).BlockCount {
				v94 = boolUint8(false1 != 0)
				goto _113
			}
			v94 = boolUint8(*(*uint64_t)(unsafe.Add(nil, (*b3BitSet)(unsafe.Pointer(v21)).Bits+uintptr(blockIndex)*8))&(uint64FromInt32(1)<<(v47%uint32FromInt32(64))) != uint64(0))
			goto _113
		_113:
			if int32FromUint8(v94) == int32(true1) {
				// this island is still awake
				goto _109
			}
			island = islands + uintptr(islandIndex)*4
			islandId = (*b3IslandSim)(unsafe.Pointer(island)).IslandId
			b3TrySleepIsland(tls, world, islandId)
			goto _109
		_109:
			;
			islandIndex = islandIndex - int32(1)
		}
		b3ValidateSolverSets(tls, world)
		(*b3World)(unsafe.Pointer(world)).Profile.SleepIslands = b3GetMilliseconds(tls, sleepTicks)
	}
}

func b3DestroySolverSet(tls *_Stack, world uintptr, setIndex int32) {
	var set uintptr
	var v1 bool
	_, _ = set, v1
	if v1 = !!(0 <= setIndex && setIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+7822, __ccgo_ts+33755, int32FromInt32(17))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setIndex)*88
	b3Free(tls, (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data, uint64FromInt32((*b3SolverSet)(unsafe.Pointer(set)).BodySims.Capacity)*uint64(252))
	(*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data = uintptrFromInt32(0)
	(*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count = 0
	(*b3SolverSet)(unsafe.Pointer(set)).BodySims.Capacity = 0
	b3Free(tls, (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Data, uint64FromInt32((*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity)*uint64(64))
	(*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Data = uintptrFromInt32(0)
	(*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Count = 0
	(*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity = 0
	b3Free(tls, (*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Data, uint64FromInt32((*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Capacity)*uint64(4))
	(*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Data = uintptrFromInt32(0)
	(*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Count = 0
	(*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Capacity = 0
	b3Free(tls, (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Data, uint64FromInt32((*b3SolverSet)(unsafe.Pointer(set)).JointSims.Capacity)*uint64(444))
	(*b3SolverSet)(unsafe.Pointer(set)).JointSims.Data = uintptrFromInt32(0)
	(*b3SolverSet)(unsafe.Pointer(set)).JointSims.Count = 0
	(*b3SolverSet)(unsafe.Pointer(set)).JointSims.Capacity = 0
	b3Free(tls, (*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Data, uint64FromInt32((*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity)*uint64(4))
	(*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Data = uintptrFromInt32(0)
	(*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Count = 0
	(*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity = 0
	b3FreeId(tls, world+4104, setIndex)
	*(*b3SolverSet)(unsafe.Add(nil, set)) = b3SolverSet{}
	(*b3SolverSet)(unsafe.Pointer(set)).SetIndex = -int32(1)
}

// C documentation
//
//	// Wake a solver set. Does not merge islands.
//	// Contacts can be in several places:
//	// 1. non-touching contacts in the disabled set
//	// 2. non-touching contacts already in the awake set
//	// 3. touching contacts in the sleeping set
//	// This handles contact types 1 and 3. Type 2 doesn't need any action.
func b3WakeSolverSet(tls *_Stack, world uintptr, setIndex int32) {
	var awakeSet, bodies, body, contact, contact1, disabledSet, island, islandDst, islandSrc, joint, jointSim, movedContact, set, simDst, simSrc, state, v12, v15, v7, v8, v9 uintptr
	var bodyCount, contactCount, contactId, contactIndex, contactKey, edgeIndex, encodedBodySimIndex, i, i1, i2, i3, islandCount, jointCount, localIndex, movedContactIndex, movedLocalIndex, newCapacity, newCapacity1, newSize, newSize1, oldCapacity, oldSize, oldSize1, v10, v11, v14 int32
	var v1, v2, v3, v4 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = awakeSet, bodies, body, bodyCount, contact, contact1, contactCount, contactId, contactIndex, contactKey, disabledSet, edgeIndex, encodedBodySimIndex, i, i1, i2, i3, island, islandCount, islandDst, islandSrc, joint, jointCount, jointSim, localIndex, movedContact, movedContactIndex, movedLocalIndex, newCapacity, newCapacity1, newSize, newSize1, oldCapacity, oldSize, oldSize1, set, simDst, simSrc, state, v1, v10, v11, v12, v14, v15, v2, v3, v4, v7, v8, v9
	if v1 = !!(setIndex >= int32(b3_firstSleepingSet)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+33783, __ccgo_ts+33755, int32FromInt32(38))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= setIndex && setIndex < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+7822, __ccgo_ts+33755, int32FromInt32(39))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setIndex)*88
	if v2 = !!(bool(0 <= int32(b3_awakeSet)) && int32(b3_awakeSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+802, __ccgo_ts+33755, int32FromInt32(40))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	awakeSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_awakeSet))*88
	if v3 = !!(bool(0 <= int32(b3_disabledSet)) && int32(b3_disabledSet) < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v3 {
		b3InternalAssert(tls, __ccgo_ts+4158, __ccgo_ts+33755, int32FromInt32(41))
	}
	_ = v3 || bool(int32FromInt32(0) != 0)
	disabledSet = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(int32(b3_disabledSet))*88
	bodies = (*b3World)(unsafe.Pointer(world)).Bodies.Data
	bodyCount = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count
	i = 0
	for {
		if !(i < bodyCount) {
			break
		}
		simSrc = (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data + uintptr(i)*252
		body = bodies + uintptr((*b3BodySim)(unsafe.Pointer(simSrc)).BodyId)*96
		if v4 = !!((*b3Body)(unsafe.Pointer(body)).SetIndex == setIndex); !v4 {
			b3InternalAssert(tls, __ccgo_ts+33815, __ccgo_ts+33755, int32FromInt32(51))
		}
		_ = v4 || bool(int32FromInt32(0) != 0)
		(*b3Body)(unsafe.Pointer(body)).SetIndex = int32(b3_awakeSet)
		(*b3Body)(unsafe.Pointer(body)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(awakeSet)).BodySims.Count
		// Reset sleep timer
		(*b3Body)(unsafe.Pointer(body)).SleepTime = float32FromFloat32(0)
		v7 = awakeSet
		v8 = awakeSet + 8
		v9 = awakeSet + 12
		v10 = int32(252)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(16)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), oldSize, newSize)
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _13
	_13:
		simDst = v12
		memcpy(tls, simDst, simSrc, uint64(252))
		v7 = awakeSet + 16
		v8 = awakeSet + 16 + 8
		v9 = awakeSet + 16 + 12
		v10 = int32(64)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(16)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), oldSize, newSize)
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _22
	_22:
		state = v12
		*(*b3BodyState)(unsafe.Add(nil, state)) = b3_identityBodyState
		(*b3BodyState)(unsafe.Pointer(state)).Flags = (*b3Body)(unsafe.Pointer(body)).Flags
		// Move non-touching contacts from disabled set to awake set.
		encodedBodySimIndex = (*b3Body)(unsafe.Pointer(body)).LocalIndex
		contactKey = (*b3Body)(unsafe.Pointer(body)).HeadContactKey
		for contactKey != -int32(1) {
			edgeIndex = contactKey & int32(1)
			contactId = contactKey >> int32(1)
			if v1 = !!(0 <= contactId && contactId < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+950, __ccgo_ts+33755, int32FromInt32(73))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			contact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactId)*216
			contactKey = (*(*b3ContactEdge)(unsafe.Add(nil, contact+12+uintptr(edgeIndex)*12))).NextKey
			if (*b3Contact)(unsafe.Pointer(contact)).SetIndex != int32(b3_disabledSet) {
				if v1 = !!((*b3Contact)(unsafe.Pointer(contact)).SetIndex == int32(b3_awakeSet) || (*b3Contact)(unsafe.Pointer(contact)).SetIndex == setIndex); !v1 {
					b3InternalAssert(tls, __ccgo_ts+33842, __ccgo_ts+33755, int32FromInt32(78))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				// If this contact is awake but not touching I need to fixup the body sim index here.
				if (*b3Contact)(unsafe.Pointer(contact)).SetIndex == int32(b3_awakeSet) {
					if v1 = !!((*b3Contact)(unsafe.Pointer(contact)).ColorIndex == -int32FromInt32(1)); !v1 {
						b3InternalAssert(tls, __ccgo_ts+24755, __ccgo_ts+33755, int32FromInt32(83))
					}
					_ = v1 || bool(int32FromInt32(0) != 0)
					if edgeIndex == 0 {
						(*b3Contact)(unsafe.Pointer(contact)).EncodedBodySimA = encodedBodySimIndex
					} else {
						(*b3Contact)(unsafe.Pointer(contact)).EncodedBodySimB = encodedBodySimIndex
					}
				}
				continue
			}
			localIndex = (*b3Contact)(unsafe.Pointer(contact)).LocalIndex
			if v1 = !!(0 <= localIndex && localIndex < (*b3SolverSet)(unsafe.Pointer(disabledSet)).ContactIndices.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+33908, __ccgo_ts+33755, int32FromInt32(99))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			if v1 = !!(*(*int32)(unsafe.Add(nil, (*b3SolverSet)(unsafe.Pointer(disabledSet)).ContactIndices.Data+uintptr(localIndex)*4)) == contactId); !v1 {
				b3InternalAssert(tls, __ccgo_ts+33974, __ccgo_ts+33755, int32FromInt32(100))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			if v1 = !!((*b3Contact)(unsafe.Pointer(contact)).Flags&uint32(b3_contactTouchingFlag) == uint32(0) && (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount == 0); !v1 {
				b3InternalAssert(tls, __ccgo_ts+34032, __ccgo_ts+33755, int32FromInt32(102))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			(*b3Contact)(unsafe.Pointer(contact)).SetIndex = int32(b3_awakeSet)
			(*b3Contact)(unsafe.Pointer(contact)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Count
			if (*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Count >= (*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Capacity {
				oldSize1 = int32FromUint64(uint64FromInt32((*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Capacity) * uint64(4))
				if (*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Capacity == 0 {
					v10 = int32(8)
				} else {
					v10 = int32(2) * (*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Capacity
				}
				newCapacity1 = v10
				newSize1 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(4))
				(*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Data = b3GrowAlloc(tls, (*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Data, oldSize1, newSize1)
				(*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Capacity = newCapacity1
			}
			v7 = awakeSet + 48 + 8
			v10 = *(*int32)(unsafe.Pointer(v7))
			*(*int32)(unsafe.Pointer(v7)) = *(*int32)(unsafe.Pointer(v7)) + 1
			*(*int32)(unsafe.Add(nil, (*b3SolverSet)(unsafe.Pointer(awakeSet)).ContactIndices.Data+uintptr(v10)*4)) = contactId
			if edgeIndex == 0 {
				(*b3Contact)(unsafe.Pointer(contact)).EncodedBodySimA = encodedBodySimIndex
			} else {
				(*b3Contact)(unsafe.Pointer(contact)).EncodedBodySimB = encodedBodySimIndex
			}
			v7 = (*b3SolverSet)(unsafe.Pointer(disabledSet)).ContactIndices.Data
			v8 = disabledSet + 48 + 8
			v10 = localIndex
			v11 = int32(4)
			if v1 = !!(0 <= v10 && v10 < *(*int32)(unsafe.Add(nil, v8)) && bool(__ccgo_ts+2085 != 0)); !v1 {
				b3InternalAssert(tls, __ccgo_ts+2111, __ccgo_ts+2171, int32FromInt32(163))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			*(*int32)(unsafe.Add(nil, v8)) = *(*int32)(unsafe.Add(nil, v8)) - 1
			if v10 != *(*int32)(unsafe.Add(nil, v8)) {
				memcpy(tls, v7+uintptr(v10*v11), v7+uintptr(*(*int32)(unsafe.Add(nil, v8))*v11), uint64FromInt32(v11))
				v14 = *(*int32)(unsafe.Add(nil, v8))
				goto _40
			}
			v14 = -int32(1)
			goto _40
		_40:
			movedLocalIndex = v14
			if movedLocalIndex != -int32(1) {
				// fix moved element
				movedContactIndex = *(*int32)(unsafe.Add(nil, (*b3SolverSet)(unsafe.Pointer(disabledSet)).ContactIndices.Data+uintptr(localIndex)*4))
				if v2 = !!(0 <= movedContactIndex && movedContactIndex < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v2 {
					b3InternalAssert(tls, __ccgo_ts+8778, __ccgo_ts+33755, int32FromInt32(122))
				}
				_ = v2 || bool(int32FromInt32(0) != 0)
				movedContact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(movedContactIndex)*216
				if v3 = !!((*b3Contact)(unsafe.Pointer(movedContact)).LocalIndex == movedLocalIndex); !v3 {
					b3InternalAssert(tls, __ccgo_ts+34112, __ccgo_ts+33755, int32FromInt32(123))
				}
				_ = v3 || bool(int32FromInt32(0) != 0)
				(*b3Contact)(unsafe.Pointer(movedContact)).LocalIndex = localIndex
			}
		}
		goto _5
	_5:
		;
		i = i + 1
	}
	// Transfer touching contacts from sleeping set to constraint graph.
	contactCount = (*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Count
	i1 = 0
	for {
		if !(i1 < contactCount) {
			break
		}
		contactIndex = *(*int32)(unsafe.Add(nil, (*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Data+uintptr(i1)*4))
		if v1 = !!(0 <= contactIndex && contactIndex < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+34156, __ccgo_ts+33755, int32FromInt32(135))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact1 = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactIndex)*216
		if v2 = !!((*b3Contact)(unsafe.Pointer(contact1)).Flags&uint32(b3_contactTouchingFlag) != 0); !v2 {
			b3InternalAssert(tls, __ccgo_ts+6646, __ccgo_ts+33755, int32FromInt32(136))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b3Contact)(unsafe.Pointer(contact1)).Flags&uint32(b3_simTouchingFlag) != 0); !v1 {
			b3InternalAssert(tls, __ccgo_ts+34226, __ccgo_ts+33755, int32FromInt32(137))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!((*b3Contact)(unsafe.Pointer(contact1)).SetIndex == setIndex); !v1 {
			b3InternalAssert(tls, __ccgo_ts+34262, __ccgo_ts+33755, int32FromInt32(138))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		b3AddContactToGraph(tls, world, contact1)
		(*b3Contact)(unsafe.Pointer(contact1)).SetIndex = int32(b3_awakeSet)
		goto _43
	_43:
		;
		i1 = i1 + 1
	}
	// transfer joints from sleeping set to awake set
	jointCount = (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Count
	i2 = 0
	for {
		if !(i2 < jointCount) {
			break
		}
		jointSim = (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Data + uintptr(i2)*444
		if v1 = !!(0 <= (*b3JointSim)(unsafe.Pointer(jointSim)).JointId && (*b3JointSim)(unsafe.Pointer(jointSim)).JointId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+34292, __ccgo_ts+33755, int32FromInt32(150))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b3JointSim)(unsafe.Pointer(jointSim)).JointId)*72
		if v2 = !!((*b3Joint)(unsafe.Pointer(joint)).SetIndex == setIndex); !v2 {
			b3InternalAssert(tls, __ccgo_ts+19247, __ccgo_ts+33755, int32FromInt32(151))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		b3AddJointToGraph(tls, world, jointSim, joint)
		(*b3Joint)(unsafe.Pointer(joint)).SetIndex = int32(b3_awakeSet)
		goto _48
	_48:
		;
		i2 = i2 + 1
	}
	// transfer island from sleeping set to awake set
	// Usually a sleeping set has only one island, but it is possible
	// that joints are created between sleeping islands and they
	// are moved to the same sleeping set.
	islandCount = (*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Count
	i3 = 0
	for {
		if !(i3 < islandCount) {
			break
		}
		islandSrc = (*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Data + uintptr(i3)*4
		if v1 = !!(0 <= (*b3IslandSim)(unsafe.Pointer(islandSrc)).IslandId && (*b3IslandSim)(unsafe.Pointer(islandSrc)).IslandId < (*b3World)(unsafe.Pointer(world)).Islands.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+34370, __ccgo_ts+33755, int32FromInt32(166))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		island = (*b3World)(unsafe.Pointer(world)).Islands.Data + uintptr((*b3IslandSim)(unsafe.Pointer(islandSrc)).IslandId)*64
		(*b3Island)(unsafe.Pointer(island)).SetIndex = int32(b3_awakeSet)
		(*b3Island)(unsafe.Pointer(island)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(awakeSet)).IslandSims.Count
		v7 = awakeSet + 64
		v8 = awakeSet + 64 + 8
		v9 = awakeSet + 64 + 12
		v10 = int32(4)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(16)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), oldSize, newSize)
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _59
	_59:
		islandDst = v12
		memcpy(tls, islandDst, islandSrc, uint64(4))
		goto _51
	_51:
		;
		i3 = i3 + 1
	}
	// destroy the sleeping set
	b3DestroySolverSet(tls, world, setIndex)
}

// C documentation
//
//	// This is called when joints are created between sets. I want to allow the sets
//	// to continue sleeping if both are asleep. Otherwise one set is waked.
//	// Islands will get merge when the set is woke.
func b3MergeSolverSets(tls *_Stack, world uintptr, setId1 int32, setId2 int32) {
	var bodies, body, contact, island, islandDst, islandSrc, joint, jointDst, jointSrc, set1, set2, simDst, simSrc, tempSet, v12, v15, v7, v8, v9 uintptr
	var bodyCount, contactCount, contactIndex, i, i1, i2, i3, islandCount, islandId, jointCount, newCapacity, newCapacity1, newSize, newSize1, oldCapacity, oldSize, oldSize1, tempId, v10, v11, v14 int32
	var v1, v2 bool
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = bodies, body, bodyCount, contact, contactCount, contactIndex, i, i1, i2, i3, island, islandCount, islandDst, islandId, islandSrc, joint, jointCount, jointDst, jointSrc, newCapacity, newCapacity1, newSize, newSize1, oldCapacity, oldSize, oldSize1, set1, set2, simDst, simSrc, tempId, tempSet, v1, v10, v11, v12, v14, v15, v2, v7, v8, v9
	if v1 = !!(setId1 >= int32(b3_firstSleepingSet)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+35327, __ccgo_ts+33755, int32FromInt32(477))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(setId2 >= int32(b3_firstSleepingSet)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+35357, __ccgo_ts+33755, int32FromInt32(478))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(0 <= setId1 && setId1 < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+35387, __ccgo_ts+33755, int32FromInt32(479))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	set1 = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setId1)*88
	if v2 = !!(0 <= setId2 && setId2 < (*b3World)(unsafe.Pointer(world)).SolverSets.Count); !v2 {
		b3InternalAssert(tls, __ccgo_ts+35447, __ccgo_ts+33755, int32FromInt32(480))
	}
	_ = v2 || bool(int32FromInt32(0) != 0)
	set2 = (*b3World)(unsafe.Pointer(world)).SolverSets.Data + uintptr(setId2)*88
	// Move the fewest number of bodies
	if (*b3SolverSet)(unsafe.Pointer(set1)).BodySims.Count < (*b3SolverSet)(unsafe.Pointer(set2)).BodySims.Count {
		tempSet = set1
		set1 = set2
		set2 = tempSet
		tempId = setId1
		setId1 = setId2
		setId2 = tempId
	}
	// transfer bodies
	bodies = (*b3World)(unsafe.Pointer(world)).Bodies.Data
	bodyCount = (*b3SolverSet)(unsafe.Pointer(set2)).BodySims.Count
	i = 0
	for {
		if !(i < bodyCount) {
			break
		}
		simSrc = (*b3SolverSet)(unsafe.Pointer(set2)).BodySims.Data + uintptr(i)*252
		body = bodies + uintptr((*b3BodySim)(unsafe.Pointer(simSrc)).BodyId)*96
		if v1 = !!((*b3Body)(unsafe.Pointer(body)).SetIndex == setId2); !v1 {
			b3InternalAssert(tls, __ccgo_ts+35507, __ccgo_ts+33755, int32FromInt32(503))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		(*b3Body)(unsafe.Pointer(body)).SetIndex = setId1
		(*b3Body)(unsafe.Pointer(body)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(set1)).BodySims.Count
		v7 = set1
		v8 = set1 + 8
		v9 = set1 + 12
		v10 = int32(252)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(16)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), oldSize, newSize)
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _13
	_13:
		simDst = v12
		memcpy(tls, simDst, simSrc, uint64(252))
		goto _5
	_5:
		;
		i = i + 1
	}
	// transfer contacts
	contactCount = (*b3SolverSet)(unsafe.Pointer(set2)).ContactIndices.Count
	i1 = 0
	for {
		if !(i1 < contactCount) {
			break
		}
		contactIndex = *(*int32)(unsafe.Add(nil, (*b3SolverSet)(unsafe.Pointer(set2)).ContactIndices.Data+uintptr(i1)*4))
		if v1 = !!(0 <= contactIndex && contactIndex < (*b3World)(unsafe.Pointer(world)).Contacts.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+34156, __ccgo_ts+33755, int32FromInt32(518))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		contact = (*b3World)(unsafe.Pointer(world)).Contacts.Data + uintptr(contactIndex)*216
		if v2 = !!((*b3Contact)(unsafe.Pointer(contact)).SetIndex == setId2); !v2 {
			b3InternalAssert(tls, __ccgo_ts+35532, __ccgo_ts+33755, int32FromInt32(519))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		(*b3Contact)(unsafe.Pointer(contact)).SetIndex = setId1
		(*b3Contact)(unsafe.Pointer(contact)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(set1)).ContactIndices.Count
		if (*b3SolverSet)(unsafe.Pointer(set1)).ContactIndices.Count >= (*b3SolverSet)(unsafe.Pointer(set1)).ContactIndices.Capacity {
			oldSize1 = int32FromUint64(uint64FromInt32((*b3SolverSet)(unsafe.Pointer(set1)).ContactIndices.Capacity) * uint64(4))
			if (*b3SolverSet)(unsafe.Pointer(set1)).ContactIndices.Capacity == 0 {
				v10 = int32(8)
			} else {
				v10 = int32(2) * (*b3SolverSet)(unsafe.Pointer(set1)).ContactIndices.Capacity
			}
			newCapacity1 = v10
			newSize1 = int32FromUint64(uint64FromInt32(newCapacity1) * uint64(4))
			(*b3SolverSet)(unsafe.Pointer(set1)).ContactIndices.Data = b3GrowAlloc(tls, (*b3SolverSet)(unsafe.Pointer(set1)).ContactIndices.Data, oldSize1, newSize1)
			(*b3SolverSet)(unsafe.Pointer(set1)).ContactIndices.Capacity = newCapacity1
		}
		v7 = set1 + 48 + 8
		v10 = *(*int32)(unsafe.Pointer(v7))
		*(*int32)(unsafe.Pointer(v7)) = *(*int32)(unsafe.Pointer(v7)) + 1
		*(*int32)(unsafe.Add(nil, (*b3SolverSet)(unsafe.Pointer(set1)).ContactIndices.Data+uintptr(v10)*4)) = contactIndex
		goto _16
	_16:
		;
		i1 = i1 + 1
	}
	// transfer joints
	jointCount = (*b3SolverSet)(unsafe.Pointer(set2)).JointSims.Count
	i2 = 0
	for {
		if !(i2 < jointCount) {
			break
		}
		jointSrc = (*b3SolverSet)(unsafe.Pointer(set2)).JointSims.Data + uintptr(i2)*444
		if v1 = !!(0 <= (*b3JointSim)(unsafe.Pointer(jointSrc)).JointId && (*b3JointSim)(unsafe.Pointer(jointSrc)).JointId < (*b3World)(unsafe.Pointer(world)).Joints.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+35560, __ccgo_ts+33755, int32FromInt32(533))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		joint = (*b3World)(unsafe.Pointer(world)).Joints.Data + uintptr((*b3JointSim)(unsafe.Pointer(jointSrc)).JointId)*72
		if v2 = !!((*b3Joint)(unsafe.Pointer(joint)).SetIndex == setId2); !v2 {
			b3InternalAssert(tls, __ccgo_ts+35638, __ccgo_ts+33755, int32FromInt32(534))
		}
		_ = v2 || bool(int32FromInt32(0) != 0)
		(*b3Joint)(unsafe.Pointer(joint)).SetIndex = setId1
		(*b3Joint)(unsafe.Pointer(joint)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(set1)).JointSims.Count
		v7 = set1 + 32
		v8 = set1 + 32 + 8
		v9 = set1 + 32 + 12
		v10 = int32(444)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(16)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), oldSize, newSize)
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _31
	_31:
		jointDst = v12
		memcpy(tls, jointDst, jointSrc, uint64(444))
		goto _22
	_22:
		;
		i2 = i2 + 1
	}
	// transfer islands
	islandCount = (*b3SolverSet)(unsafe.Pointer(set2)).IslandSims.Count
	i3 = 0
	for {
		if !(i3 < islandCount) {
			break
		}
		islandSrc = (*b3SolverSet)(unsafe.Pointer(set2)).IslandSims.Data + uintptr(i3)*4
		islandId = (*b3IslandSim)(unsafe.Pointer(islandSrc)).IslandId
		if v1 = !!(0 <= islandId && islandId < (*b3World)(unsafe.Pointer(world)).Islands.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+1108, __ccgo_ts+33755, int32FromInt32(551))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		island = (*b3World)(unsafe.Pointer(world)).Islands.Data + uintptr(islandId)*64
		(*b3Island)(unsafe.Pointer(island)).SetIndex = setId1
		(*b3Island)(unsafe.Pointer(island)).LocalIndex = (*b3SolverSet)(unsafe.Pointer(set1)).IslandSims.Count
		v7 = set1 + 64
		v8 = set1 + 64 + 8
		v9 = set1 + 64 + 12
		v10 = int32(4)
		if *(*int32)(unsafe.Add(nil, v8)) >= *(*int32)(unsafe.Add(nil, v9)) {
			oldCapacity = *(*int32)(unsafe.Add(nil, v9))
			oldSize = oldCapacity * v10
			if oldCapacity == 0 {
				v11 = int32(16)
			} else {
				v11 = int32(2) * oldCapacity
			}
			newCapacity = v11
			newSize = newCapacity * v10
			*(*uintptr)(unsafe.Add(nil, v7)) = b3GrowAlloc(tls, *(*uintptr)(unsafe.Add(nil, v7)), oldSize, newSize)
			*(*int32)(unsafe.Add(nil, v9)) = newCapacity
		}
		v15 = v8
		v14 = *(*int32)(unsafe.Pointer(v15))
		*(*int32)(unsafe.Pointer(v15)) = *(*int32)(unsafe.Pointer(v15)) + 1
		v12 = *(*uintptr)(unsafe.Add(nil, v7)) + uintptr(v14*v10)
		goto _42
	_42:
		islandDst = v12
		memcpy(tls, islandDst, islandSrc, uint64(4))
		goto _34
	_34:
		;
		i3 = i3 + 1
	}
	// destroy the merged set
	// Warning: need to be careful not to destroy things that got transferred, like triangle caches.
	b3DestroySolverSet(tls, world, setId2)
	b3ValidateSolverSets(tls, world)
}

// C documentation
//
//	// Solve the normal constraint, friction, and rolling resistance.
func b3SolveContactsW(tls *_Stack, block b3SolverBlock, context uintptr, base uintptr, manifoldStarts uintptr, fixedManifoldCount int32) {
	bp := tls.Alloc(416)
	defer tls.Free(416)
	var P, angularRelativeVelocity, deltaImpulse, dp, dv, gA, gB, normalA, normalB, oldImpulse, rA, rA1, rB, rB1, scaledImpulse, tangent1, tangent2, uA, uB b3Vec3W
	var anyRolling, v15 uint8
	var c1, v102, v103, v130, v131, v159, v160, v17, v18, v183, v184, v19, v231, v232, v269, v271, v272, v309, v311, v312, v335, v336, v38, v39, v62, v63, v64, v78, v79 b3Vec3W4
	var c2, cm, cp, manifolds, states uintptr
	var deltaImpulse1, newImpulse1, vt b3Vec2W
	var epsilonW, impulse, impulse1, invMass, inv_h, lengthSquared, lengthSquared1, linearNormalVelocity, linearTangentVelocity1, linearTangentVelocity2, manifoldImpulse, mask1, mask2, maxImpulse, maxImpulse1, maxLambda, negImpulse, newImpulse, normalSeparation, normalize, normalize1, oldImpulse1, oneW, s3, scale, scale1, totalNormalImpulse, totalTwistLimit, twistSpeed, velocityBias, vn, zeroW b3FloatW
	var i, manifoldCount, manifoldIndex, pointIndex, v13 int32
	var r, r1, r2, r3, v10, v100, v104, v106, v107, v108, v110, v111, v112, v114, v115, v116, v118, v119, v120, v122, v123, v124, v126, v127, v128, v132, v134, v135, v136, v138, v139, v14, v140, v142, v143, v144, v146, v147, v148, v150, v151, v152, v154, v155, v156, v161, v163, v164, v165, v167, v168, v169, v171, v172, v173, v175, v176, v177, v179, v180, v181, v185, v187, v188, v189, v191, v192, v193, v195, v196, v197, v199, v2, v200, v201, v203, v204, v205, v207, v208, v209, v21, v211, v212, v213, v215, v216, v217, v219, v22, v220, v221, v223, v224, v229, v23, v233, v234, v235, v237, v238, v239, v241, v242, v243, v245, v246, v247, v249, v25, v250, v251, v253, v254, v255, v257, v258, v259, v26, v261, v262, v263, v265, v266, v267, v27, v273, v274, v275, v277, v278, v279, v281, v282, v283, v285, v286, v287, v289, v29, v290, v291, v293, v294, v295, v297, v298, v299, v30, v301, v302, v303, v305, v306, v307, v31, v313, v315, v316, v317, v319, v320, v321, v323, v324, v325, v327, v328, v329, v33, v331, v332, v333, v337, v339, v34, v340, v341, v343, v344, v345, v347, v348, v349, v35, v351, v352, v353, v355, v356, v357, v359, v360, v361, v363, v364, v365, v367, v368, v369, v371, v372, v373, v375, v376, v378, v379, v381, v382, v387, v389, v390, v391, v40, v42, v43, v44, v46, v47, v48, v5, v50, v51, v52, v54, v55, v56, v58, v59, v60, v66, v67, v68, v7, v70, v71, v72, v74, v75, v76, v80, v82, v83, v84, v86, v87, v88, v90, v91, v92, v94, v95, v96, v98, v99 b3FloatW4
	var v1, v225, v226, v227, v228, v383, v384, v385, v386, v4, v543, v9 float32
	var _ /* bA at bp+0 */ b3BodyStateW
	var _ /* bB at bp+208 */ b3BodyStateW
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = P, angularRelativeVelocity, anyRolling, c1, c2, cm, cp, deltaImpulse, deltaImpulse1, dp, dv, epsilonW, gA, gB, i, impulse, impulse1, invMass, inv_h, lengthSquared, lengthSquared1, linearNormalVelocity, linearTangentVelocity1, linearTangentVelocity2, manifoldCount, manifoldImpulse, manifoldIndex, manifolds, mask1, mask2, maxImpulse, maxImpulse1, maxLambda, negImpulse, newImpulse, newImpulse1, normalA, normalB, normalSeparation, normalize, normalize1, oldImpulse, oldImpulse1, oneW, pointIndex, r, r1, r2, r3, rA, rA1, rB, rB1, s3, scale, scale1, scaledImpulse, states, tangent1, tangent2, totalNormalImpulse, totalTwistLimit, twistSpeed, uA, uB, velocityBias, vn, vt, zeroW, v1, v10, v100, v102, v103, v104, v106, v107, v108, v110, v111, v112, v114, v115, v116, v118, v119, v120, v122, v123, v124, v126, v127, v128, v13, v130, v131, v132, v134, v135, v136, v138, v139, v14, v140, v142, v143, v144, v146, v147, v148, v15, v150, v151, v152, v154, v155, v156, v159, v160, v161, v163, v164, v165, v167, v168, v169, v17, v171, v172, v173, v175, v176, v177, v179, v18, v180, v181, v183, v184, v185, v187, v188, v189, v19, v191, v192, v193, v195, v196, v197, v199, v2, v200, v201, v203, v204, v205, v207, v208, v209, v21, v211, v212, v213, v215, v216, v217, v219, v22, v220, v221, v223, v224, v225, v226, v227, v228, v229, v23, v231, v232, v233, v234, v235, v237, v238, v239, v241, v242, v243, v245, v246, v247, v249, v25, v250, v251, v253, v254, v255, v257, v258, v259, v26, v261, v262, v263, v265, v266, v267, v269, v27, v271, v272, v273, v274, v275, v277, v278, v279, v281, v282, v283, v285, v286, v287, v289, v29, v290, v291, v293, v294, v295, v297, v298, v299, v30, v301, v302, v303, v305, v306, v307, v309, v31, v311, v312, v313, v315, v316, v317, v319, v320, v321, v323, v324, v325, v327, v328, v329, v33, v331, v332, v333, v335, v336, v337, v339, v34, v340, v341, v343, v344, v345, v347, v348, v349, v35, v351, v352, v353, v355, v356, v357, v359, v360, v361, v363, v364, v365, v367, v368, v369, v371, v372, v373, v375, v376, v378, v379, v38, v381, v382, v383, v384, v385, v386, v387, v389, v39, v390, v391, v4, v40, v42, v43, v44, v46, v47, v48, v5, v50, v51, v52, v54, v543, v55, v56, v58, v59, v60, v62, v63, v64, v66, v67, v68, v7, v70, v71, v72, v74, v75, v76, v78, v79, v80, v82, v83, v84, v86, v87, v88, v9, v90, v91, v92, v94, v95, v96, v98, v99
	states = (*b3StepContext)(unsafe.Pointer(context)).States
	v1 = (*b3StepContext)(unsafe.Pointer(context)).Inv_h
	v2 = b3FloatW4{
		X: v1,
		Y: v1,
		Z: v1,
		W: v1,
	}
	goto _3
_3:
	inv_h = v2
	v4 = float32FromFloat32(1)
	v5 = b3FloatW4{
		X: v4,
		Y: v4,
		Z: v4,
		W: v4,
	}
	goto _6
_6:
	oneW = v5
	v7 = b3FloatW4{}
	goto _8
_8:
	zeroW = v7
	v9 = float32FromFloat32(1.1920928955078125e-07)
	v10 = b3FloatW4{
		X: v9,
		Y: v9,
		Z: v9,
		W: v9,
	}
	goto _11
_11:
	epsilonW = v10
	c2 = b3GetConstraintW(tls, base, manifoldStarts, block.StartIndex, fixedManifoldCount)
	i = 0
	for {
		if !(i < int32FromUint16(block.Count)) {
			break
		}
		if fixedManifoldCount > 0 {
			v13 = fixedManifoldCount
		} else {
			v13 = (*b3ContactConstraintWide)(unsafe.Pointer(c2)).ManifoldCount
		}
		manifoldCount = v13
		*(*b3BodyStateW)(unsafe.Add(nil, bp)) = b3GatherBodies(tls, states, c2)
		*(*b3BodyStateW)(unsafe.Add(nil, bp+208)) = b3GatherBodies(tls, states, c2+16)
		v14 = (*b3ContactConstraintWide)(unsafe.Pointer(c2)).RollingResistance
		v15 = boolUint8(v14.X == float32FromFloat32(0) && v14.Y == float32FromFloat32(0) && v14.Z == float32FromFloat32(0) && v14.W == float32FromFloat32(0))
		goto _16
	_16:
		anyRolling = boolUint8(int32FromUint8(v15) == false1)
		v17 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).Dp
		v18 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).Dp
		v21 = v17.X
		v22 = v18.X
		v23 = b3FloatW4{
			X: v21.X - v22.X,
			Y: v21.Y - v22.Y,
			Z: v21.Z - v22.Z,
			W: v21.W - v22.W,
		}
		goto _24
	_24:
		v25 = v17.Y
		v26 = v18.Y
		v27 = b3FloatW4{
			X: v25.X - v26.X,
			Y: v25.Y - v26.Y,
			Z: v25.Z - v26.Z,
			W: v25.W - v26.W,
		}
		goto _28
	_28:
		v29 = v17.Z
		v30 = v18.Z
		v31 = b3FloatW4{
			X: v29.X - v30.X,
			Y: v29.Y - v30.Y,
			Z: v29.Z - v30.Z,
			W: v29.W - v30.W,
		}
		goto _32
	_32:
		v19 = b3Vec3W4{
			X: v23,
			Y: v27,
			Z: v31,
		}
		goto _20
	_20:
		dp = v19
		v33 = (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvMassA
		v34 = (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvMassB
		v35 = b3FloatW4{
			X: v33.X + v34.X,
			Y: v33.Y + v34.Y,
			Z: v33.Z + v34.Z,
			W: v33.W + v34.W,
		}
		goto _36
	_36:
		invMass = v35
		manifolds = b3GetManifoldsW(tls, c2)
		manifoldIndex = 0
		for {
			if !(manifoldIndex < manifoldCount) {
				break
			}
			cm = manifolds + uintptr(manifoldIndex)*1288
			v38 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
			v39 = dp
			v42 = v38.X
			v43 = v39.X
			v44 = b3FloatW4{
				X: float32(v42.X * v43.X),
				Y: float32(v42.Y * v43.Y),
				Z: float32(v42.Z * v43.Z),
				W: float32(v42.W * v43.W),
			}
			goto _45
		_45:
			v46 = v38.Y
			v47 = v39.Y
			v48 = b3FloatW4{
				X: float32(v46.X * v47.X),
				Y: float32(v46.Y * v47.Y),
				Z: float32(v46.Z * v47.Z),
				W: float32(v46.W * v47.W),
			}
			goto _49
		_49:
			v50 = v44
			v51 = v48
			v52 = b3FloatW4{
				X: v50.X + v51.X,
				Y: v50.Y + v51.Y,
				Z: v50.Z + v51.Z,
				W: v50.W + v51.W,
			}
			goto _53
		_53:
			v54 = v38.Z
			v55 = v39.Z
			v56 = b3FloatW4{
				X: float32(v54.X * v55.X),
				Y: float32(v54.Y * v55.Y),
				Z: float32(v54.Z * v55.Z),
				W: float32(v54.W * v55.W),
			}
			goto _57
		_57:
			v58 = v52
			v59 = v56
			v60 = b3FloatW4{
				X: v58.X + v59.X,
				Y: v58.Y + v59.Y,
				Z: v58.Z + v59.Z,
				W: v58.W + v59.W,
			}
			goto _61
		_61:
			v40 = v60
			goto _41
		_41:
			// Read the normal rather than put on stack to avoid register spills.
			normalSeparation = v40
			normalA = b3InvRotateVectorW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp))).Dq, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal)
			normalB = b3InvRotateVectorW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).Dq, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal)
			v62 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).V
			v63 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).V
			v66 = v62.X
			v67 = v63.X
			v68 = b3FloatW4{
				X: v66.X - v67.X,
				Y: v66.Y - v67.Y,
				Z: v66.Z - v67.Z,
				W: v66.W - v67.W,
			}
			goto _69
		_69:
			v70 = v62.Y
			v71 = v63.Y
			v72 = b3FloatW4{
				X: v70.X - v71.X,
				Y: v70.Y - v71.Y,
				Z: v70.Z - v71.Z,
				W: v70.W - v71.W,
			}
			goto _73
		_73:
			v74 = v62.Z
			v75 = v63.Z
			v76 = b3FloatW4{
				X: v74.X - v75.X,
				Y: v74.Y - v75.Y,
				Z: v74.Z - v75.Z,
				W: v74.W - v75.W,
			}
			goto _77
		_77:
			v64 = b3Vec3W4{
				X: v68,
				Y: v72,
				Z: v76,
			}
			goto _65
		_65:
			dv = v64
			v78 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
			v79 = dv
			v82 = v78.X
			v83 = v79.X
			v84 = b3FloatW4{
				X: float32(v82.X * v83.X),
				Y: float32(v82.Y * v83.Y),
				Z: float32(v82.Z * v83.Z),
				W: float32(v82.W * v83.W),
			}
			goto _85
		_85:
			v86 = v78.Y
			v87 = v79.Y
			v88 = b3FloatW4{
				X: float32(v86.X * v87.X),
				Y: float32(v86.Y * v87.Y),
				Z: float32(v86.Z * v87.Z),
				W: float32(v86.W * v87.W),
			}
			goto _89
		_89:
			v90 = v84
			v91 = v88
			v92 = b3FloatW4{
				X: v90.X + v91.X,
				Y: v90.Y + v91.Y,
				Z: v90.Z + v91.Z,
				W: v90.W + v91.W,
			}
			goto _93
		_93:
			v94 = v78.Z
			v95 = v79.Z
			v96 = b3FloatW4{
				X: float32(v94.X * v95.X),
				Y: float32(v94.Y * v95.Y),
				Z: float32(v94.Z * v95.Z),
				W: float32(v94.W * v95.W),
			}
			goto _97
		_97:
			v98 = v92
			v99 = v96
			v100 = b3FloatW4{
				X: v98.X + v99.X,
				Y: v98.Y + v99.Y,
				Z: v98.Z + v99.Z,
				W: v98.W + v99.W,
			}
			goto _101
		_101:
			v80 = v100
			goto _81
		_81:
			linearNormalVelocity = v80
			v102 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Tangent1
			v103 = dv
			v106 = v102.X
			v107 = v103.X
			v108 = b3FloatW4{
				X: float32(v106.X * v107.X),
				Y: float32(v106.Y * v107.Y),
				Z: float32(v106.Z * v107.Z),
				W: float32(v106.W * v107.W),
			}
			goto _109
		_109:
			v110 = v102.Y
			v111 = v103.Y
			v112 = b3FloatW4{
				X: float32(v110.X * v111.X),
				Y: float32(v110.Y * v111.Y),
				Z: float32(v110.Z * v111.Z),
				W: float32(v110.W * v111.W),
			}
			goto _113
		_113:
			v114 = v108
			v115 = v112
			v116 = b3FloatW4{
				X: v114.X + v115.X,
				Y: v114.Y + v115.Y,
				Z: v114.Z + v115.Z,
				W: v114.W + v115.W,
			}
			goto _117
		_117:
			v118 = v102.Z
			v119 = v103.Z
			v120 = b3FloatW4{
				X: float32(v118.X * v119.X),
				Y: float32(v118.Y * v119.Y),
				Z: float32(v118.Z * v119.Z),
				W: float32(v118.W * v119.W),
			}
			goto _121
		_121:
			v122 = v116
			v123 = v120
			v124 = b3FloatW4{
				X: v122.X + v123.X,
				Y: v122.Y + v123.Y,
				Z: v122.Z + v123.Z,
				W: v122.W + v123.W,
			}
			goto _125
		_125:
			v104 = v124
			goto _105
		_105:
			v126 = v104
			v127 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).TangentVelocity1
			v128 = b3FloatW4{
				X: v126.X - v127.X,
				Y: v126.Y - v127.Y,
				Z: v126.Z - v127.Z,
				W: v126.W - v127.W,
			}
			goto _129
		_129:
			linearTangentVelocity1 = v128
			v130 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Tangent2
			v131 = dv
			v134 = v130.X
			v135 = v131.X
			v136 = b3FloatW4{
				X: float32(v134.X * v135.X),
				Y: float32(v134.Y * v135.Y),
				Z: float32(v134.Z * v135.Z),
				W: float32(v134.W * v135.W),
			}
			goto _137
		_137:
			v138 = v130.Y
			v139 = v131.Y
			v140 = b3FloatW4{
				X: float32(v138.X * v139.X),
				Y: float32(v138.Y * v139.Y),
				Z: float32(v138.Z * v139.Z),
				W: float32(v138.W * v139.W),
			}
			goto _141
		_141:
			v142 = v136
			v143 = v140
			v144 = b3FloatW4{
				X: v142.X + v143.X,
				Y: v142.Y + v143.Y,
				Z: v142.Z + v143.Z,
				W: v142.W + v143.W,
			}
			goto _145
		_145:
			v146 = v130.Z
			v147 = v131.Z
			v148 = b3FloatW4{
				X: float32(v146.X * v147.X),
				Y: float32(v146.Y * v147.Y),
				Z: float32(v146.Z * v147.Z),
				W: float32(v146.W * v147.W),
			}
			goto _149
		_149:
			v150 = v144
			v151 = v148
			v152 = b3FloatW4{
				X: v150.X + v151.X,
				Y: v150.Y + v151.Y,
				Z: v150.Z + v151.Z,
				W: v150.W + v151.W,
			}
			goto _153
		_153:
			v132 = v152
			goto _133
		_133:
			v154 = v132
			v155 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).TangentVelocity2
			v156 = b3FloatW4{
				X: v154.X - v155.X,
				Y: v154.Y - v155.Y,
				Z: v154.Z - v155.Z,
				W: v154.W - v155.W,
			}
			goto _157
		_157:
			linearTangentVelocity2 = v156
			manifoldImpulse = zeroW
			totalNormalImpulse = zeroW
			totalTwistLimit = zeroW
			// Normal contraints
			pointIndex = 0
			for {
				if !(pointIndex < (*b3ContactManifoldWide)(unsafe.Pointer(cm)).PointCount) {
					break
				}
				cp = cm + 452 + uintptr(pointIndex)*208
				// Fixed anchor points for applying impulses
				rA = (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).AnchorAs
				rB = (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).AnchorBs
				v159 = normalB
				v160 = rB
				v163 = v159.X
				v164 = v160.X
				v165 = b3FloatW4{
					X: float32(v163.X * v164.X),
					Y: float32(v163.Y * v164.Y),
					Z: float32(v163.Z * v164.Z),
					W: float32(v163.W * v164.W),
				}
				goto _166
			_166:
				v167 = v159.Y
				v168 = v160.Y
				v169 = b3FloatW4{
					X: float32(v167.X * v168.X),
					Y: float32(v167.Y * v168.Y),
					Z: float32(v167.Z * v168.Z),
					W: float32(v167.W * v168.W),
				}
				goto _170
			_170:
				v171 = v165
				v172 = v169
				v173 = b3FloatW4{
					X: v171.X + v172.X,
					Y: v171.Y + v172.Y,
					Z: v171.Z + v172.Z,
					W: v171.W + v172.W,
				}
				goto _174
			_174:
				v175 = v159.Z
				v176 = v160.Z
				v177 = b3FloatW4{
					X: float32(v175.X * v176.X),
					Y: float32(v175.Y * v176.Y),
					Z: float32(v175.Z * v176.Z),
					W: float32(v175.W * v176.W),
				}
				goto _178
			_178:
				v179 = v173
				v180 = v177
				v181 = b3FloatW4{
					X: v179.X + v180.X,
					Y: v179.Y + v180.Y,
					Z: v179.Z + v180.Z,
					W: v179.W + v180.W,
				}
				goto _182
			_182:
				v161 = v181
				goto _162
			_162:
				v183 = normalA
				v184 = rA
				v187 = v183.X
				v188 = v184.X
				v189 = b3FloatW4{
					X: float32(v187.X * v188.X),
					Y: float32(v187.Y * v188.Y),
					Z: float32(v187.Z * v188.Z),
					W: float32(v187.W * v188.W),
				}
				goto _190
			_190:
				v191 = v183.Y
				v192 = v184.Y
				v193 = b3FloatW4{
					X: float32(v191.X * v192.X),
					Y: float32(v191.Y * v192.Y),
					Z: float32(v191.Z * v192.Z),
					W: float32(v191.W * v192.W),
				}
				goto _194
			_194:
				v195 = v189
				v196 = v193
				v197 = b3FloatW4{
					X: v195.X + v196.X,
					Y: v195.Y + v196.Y,
					Z: v195.Z + v196.Z,
					W: v195.W + v196.W,
				}
				goto _198
			_198:
				v199 = v183.Z
				v200 = v184.Z
				v201 = b3FloatW4{
					X: float32(v199.X * v200.X),
					Y: float32(v199.Y * v200.Y),
					Z: float32(v199.Z * v200.Z),
					W: float32(v199.W * v200.W),
				}
				goto _202
			_202:
				v203 = v197
				v204 = v201
				v205 = b3FloatW4{
					X: v203.X + v204.X,
					Y: v203.Y + v204.Y,
					Z: v203.Z + v204.Z,
					W: v203.W + v204.W,
				}
				goto _206
			_206:
				v185 = v205
				goto _186
			_186:
				v207 = v161
				v208 = v185
				v209 = b3FloatW4{
					X: v207.X - v208.X,
					Y: v207.Y - v208.Y,
					Z: v207.Z - v208.Z,
					W: v207.W - v208.W,
				}
				goto _210
			_210:
				v211 = normalSeparation
				v212 = v209
				v213 = b3FloatW4{
					X: v211.X + v212.X,
					Y: v211.Y + v212.Y,
					Z: v211.Z + v212.Z,
					W: v211.W + v212.W,
				}
				goto _214
			_214:
				v215 = v213
				v216 = (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).BaseSeparations
				v217 = b3FloatW4{
					X: v215.X + v216.X,
					Y: v215.Y + v216.Y,
					Z: v215.Z + v216.Z,
					W: v215.W + v216.W,
				}
				goto _218
			_218:
				s3 = v217
				v219 = s3
				v220 = inv_h
				v221 = b3FloatW4{
					X: float32(v219.X * v220.X),
					Y: float32(v219.Y * v220.Y),
					Z: float32(v219.Z * v220.Z),
					W: float32(v219.W * v220.W),
				}
				goto _222
			_222:
				v223 = zeroW
				v224 = v221
				if v223.X >= v224.X {
					v225 = v223.X
				} else {
					v225 = v224.X
				}
				r1.X = v225
				if v223.Y >= v224.Y {
					v226 = v223.Y
				} else {
					v226 = v224.Y
				}
				r1.Y = v226
				if v223.Z >= v224.Z {
					v227 = v223.Z
				} else {
					v227 = v224.Z
				}
				r1.Z = v227
				if v223.W >= v224.W {
					v228 = v223.W
				} else {
					v228 = v224.W
				}
				r1.W = v228
				v229 = r1
				goto _230
			_230:
				// Speculative bias, positive if separated and zero if overlapped
				velocityBias = v229
				v231 = rA
				v232 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
				v233 = v231.Y
				v234 = v232.Z
				v235 = b3FloatW4{
					X: float32(v233.X * v234.X),
					Y: float32(v233.Y * v234.Y),
					Z: float32(v233.Z * v234.Z),
					W: float32(v233.W * v234.W),
				}
				goto _236
			_236:
				v237 = v231.Z
				v238 = v232.Y
				v239 = b3FloatW4{
					X: float32(v237.X * v238.X),
					Y: float32(v237.Y * v238.Y),
					Z: float32(v237.Z * v238.Z),
					W: float32(v237.W * v238.W),
				}
				goto _240
			_240:
				v241 = v235
				v242 = v239
				v243 = b3FloatW4{
					X: v241.X - v242.X,
					Y: v241.Y - v242.Y,
					Z: v241.Z - v242.Z,
					W: v241.W - v242.W,
				}
				goto _244
			_244:
				c1.X = v243
				v245 = v231.Z
				v246 = v232.X
				v247 = b3FloatW4{
					X: float32(v245.X * v246.X),
					Y: float32(v245.Y * v246.Y),
					Z: float32(v245.Z * v246.Z),
					W: float32(v245.W * v246.W),
				}
				goto _248
			_248:
				v249 = v231.X
				v250 = v232.Z
				v251 = b3FloatW4{
					X: float32(v249.X * v250.X),
					Y: float32(v249.Y * v250.Y),
					Z: float32(v249.Z * v250.Z),
					W: float32(v249.W * v250.W),
				}
				goto _252
			_252:
				v253 = v247
				v254 = v251
				v255 = b3FloatW4{
					X: v253.X - v254.X,
					Y: v253.Y - v254.Y,
					Z: v253.Z - v254.Z,
					W: v253.W - v254.W,
				}
				goto _256
			_256:
				c1.Y = v255
				v257 = v231.X
				v258 = v232.Y
				v259 = b3FloatW4{
					X: float32(v257.X * v258.X),
					Y: float32(v257.Y * v258.Y),
					Z: float32(v257.Z * v258.Z),
					W: float32(v257.W * v258.W),
				}
				goto _260
			_260:
				v261 = v231.Y
				v262 = v232.X
				v263 = b3FloatW4{
					X: float32(v261.X * v262.X),
					Y: float32(v261.Y * v262.Y),
					Z: float32(v261.Z * v262.Z),
					W: float32(v261.W * v262.W),
				}
				goto _264
			_264:
				v265 = v259
				v266 = v263
				v267 = b3FloatW4{
					X: v265.X - v266.X,
					Y: v265.Y - v266.Y,
					Z: v265.Z - v266.Z,
					W: v265.W - v266.W,
				}
				goto _268
			_268:
				c1.Z = v267
				v269 = c1
				goto _270
			_270:
				uA = v269
				v271 = rB
				v272 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
				v273 = v271.Y
				v274 = v272.Z
				v275 = b3FloatW4{
					X: float32(v273.X * v274.X),
					Y: float32(v273.Y * v274.Y),
					Z: float32(v273.Z * v274.Z),
					W: float32(v273.W * v274.W),
				}
				goto _276
			_276:
				v277 = v271.Z
				v278 = v272.Y
				v279 = b3FloatW4{
					X: float32(v277.X * v278.X),
					Y: float32(v277.Y * v278.Y),
					Z: float32(v277.Z * v278.Z),
					W: float32(v277.W * v278.W),
				}
				goto _280
			_280:
				v281 = v275
				v282 = v279
				v283 = b3FloatW4{
					X: v281.X - v282.X,
					Y: v281.Y - v282.Y,
					Z: v281.Z - v282.Z,
					W: v281.W - v282.W,
				}
				goto _284
			_284:
				c1.X = v283
				v285 = v271.Z
				v286 = v272.X
				v287 = b3FloatW4{
					X: float32(v285.X * v286.X),
					Y: float32(v285.Y * v286.Y),
					Z: float32(v285.Z * v286.Z),
					W: float32(v285.W * v286.W),
				}
				goto _288
			_288:
				v289 = v271.X
				v290 = v272.Z
				v291 = b3FloatW4{
					X: float32(v289.X * v290.X),
					Y: float32(v289.Y * v290.Y),
					Z: float32(v289.Z * v290.Z),
					W: float32(v289.W * v290.W),
				}
				goto _292
			_292:
				v293 = v287
				v294 = v291
				v295 = b3FloatW4{
					X: v293.X - v294.X,
					Y: v293.Y - v294.Y,
					Z: v293.Z - v294.Z,
					W: v293.W - v294.W,
				}
				goto _296
			_296:
				c1.Y = v295
				v297 = v271.X
				v298 = v272.Y
				v299 = b3FloatW4{
					X: float32(v297.X * v298.X),
					Y: float32(v297.Y * v298.Y),
					Z: float32(v297.Z * v298.Z),
					W: float32(v297.W * v298.W),
				}
				goto _300
			_300:
				v301 = v271.Y
				v302 = v272.X
				v303 = b3FloatW4{
					X: float32(v301.X * v302.X),
					Y: float32(v301.Y * v302.Y),
					Z: float32(v301.Z * v302.Z),
					W: float32(v301.W * v302.W),
				}
				goto _304
			_304:
				v305 = v299
				v306 = v303
				v307 = b3FloatW4{
					X: v305.X - v306.X,
					Y: v305.Y - v306.Y,
					Z: v305.Z - v306.Z,
					W: v305.W - v306.W,
				}
				goto _308
			_308:
				c1.Z = v307
				v309 = c1
				goto _310
			_310:
				uB = v309
				gA = b3MulMVW(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvIA, uA)
				gB = b3MulMVW(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvIB, uB)
				v311 = uB
				v312 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
				v315 = v311.X
				v316 = v312.X
				v317 = b3FloatW4{
					X: float32(v315.X * v316.X),
					Y: float32(v315.Y * v316.Y),
					Z: float32(v315.Z * v316.Z),
					W: float32(v315.W * v316.W),
				}
				goto _318
			_318:
				v319 = v311.Y
				v320 = v312.Y
				v321 = b3FloatW4{
					X: float32(v319.X * v320.X),
					Y: float32(v319.Y * v320.Y),
					Z: float32(v319.Z * v320.Z),
					W: float32(v319.W * v320.W),
				}
				goto _322
			_322:
				v323 = v317
				v324 = v321
				v325 = b3FloatW4{
					X: v323.X + v324.X,
					Y: v323.Y + v324.Y,
					Z: v323.Z + v324.Z,
					W: v323.W + v324.W,
				}
				goto _326
			_326:
				v327 = v311.Z
				v328 = v312.Z
				v329 = b3FloatW4{
					X: float32(v327.X * v328.X),
					Y: float32(v327.Y * v328.Y),
					Z: float32(v327.Z * v328.Z),
					W: float32(v327.W * v328.W),
				}
				goto _330
			_330:
				v331 = v325
				v332 = v329
				v333 = b3FloatW4{
					X: v331.X + v332.X,
					Y: v331.Y + v332.Y,
					Z: v331.Z + v332.Z,
					W: v331.W + v332.W,
				}
				goto _334
			_334:
				v313 = v333
				goto _314
			_314:
				v335 = uA
				v336 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
				v339 = v335.X
				v340 = v336.X
				v341 = b3FloatW4{
					X: float32(v339.X * v340.X),
					Y: float32(v339.Y * v340.Y),
					Z: float32(v339.Z * v340.Z),
					W: float32(v339.W * v340.W),
				}
				goto _342
			_342:
				v343 = v335.Y
				v344 = v336.Y
				v345 = b3FloatW4{
					X: float32(v343.X * v344.X),
					Y: float32(v343.Y * v344.Y),
					Z: float32(v343.Z * v344.Z),
					W: float32(v343.W * v344.W),
				}
				goto _346
			_346:
				v347 = v341
				v348 = v345
				v349 = b3FloatW4{
					X: v347.X + v348.X,
					Y: v347.Y + v348.Y,
					Z: v347.Z + v348.Z,
					W: v347.W + v348.W,
				}
				goto _350
			_350:
				v351 = v335.Z
				v352 = v336.Z
				v353 = b3FloatW4{
					X: float32(v351.X * v352.X),
					Y: float32(v351.Y * v352.Y),
					Z: float32(v351.Z * v352.Z),
					W: float32(v351.W * v352.W),
				}
				goto _354
			_354:
				v355 = v349
				v356 = v353
				v357 = b3FloatW4{
					X: v355.X + v356.X,
					Y: v355.Y + v356.Y,
					Z: v355.Z + v356.Z,
					W: v355.W + v356.W,
				}
				goto _358
			_358:
				v337 = v357
				goto _338
			_338:
				v359 = v313
				v360 = v337
				v361 = b3FloatW4{
					X: v359.X - v360.X,
					Y: v359.Y - v360.Y,
					Z: v359.Z - v360.Z,
					W: v359.W - v360.W,
				}
				goto _362
			_362:
				v363 = linearNormalVelocity
				v364 = v361
				v365 = b3FloatW4{
					X: v363.X + v364.X,
					Y: v363.Y + v364.Y,
					Z: v363.Z + v364.Z,
					W: v363.W + v364.W,
				}
				goto _366
			_366:
				vn = v365
				v367 = vn
				v368 = velocityBias
				v369 = b3FloatW4{
					X: v367.X + v368.X,
					Y: v367.Y + v368.Y,
					Z: v367.Z + v368.Z,
					W: v367.W + v368.W,
				}
				goto _370
			_370:
				v371 = (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).NormalMasses
				v372 = v369
				v373 = b3FloatW4{
					X: float32(v371.X * v372.X),
					Y: float32(v371.Y * v372.Y),
					Z: float32(v371.Z * v372.Z),
					W: float32(v371.W * v372.W),
				}
				goto _374
			_374:
				negImpulse = v373
				v375 = negImpulse
				v376 = b3FloatW4{
					X: -v375.X,
					Y: -v375.Y,
					Z: -v375.Z,
					W: -v375.W,
				}
				goto _377
			_377:
				v378 = (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).NormalImpulses
				v379 = b3FloatW4{
					X: -v378.X,
					Y: -v378.Y,
					Z: -v378.Z,
					W: -v378.W,
				}
				goto _380
			_380:
				v381 = v376
				v382 = v379
				if v381.X >= v382.X {
					v383 = v381.X
				} else {
					v383 = v382.X
				}
				r1.X = v383
				if v381.Y >= v382.Y {
					v384 = v381.Y
				} else {
					v384 = v382.Y
				}
				r1.Y = v384
				if v381.Z >= v382.Z {
					v385 = v381.Z
				} else {
					v385 = v382.Z
				}
				r1.Z = v385
				if v381.W >= v382.W {
					v386 = v381.W
				} else {
					v386 = v382.W
				}
				r1.W = v386
				v387 = r1
				goto _388
			_388:
				impulse = v387
				v389 = (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).NormalImpulses
				v390 = impulse
				v391 = b3FloatW4{
					X: v389.X + v390.X,
					Y: v389.Y + v390.Y,
					Z: v389.Z + v390.Z,
					W: v389.W + v390.W,
				}
				goto _392
			_392:
				newImpulse = v391
				(*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).NormalImpulses = newImpulse
				v2 = (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).TotalNormalImpulses
				v5 = newImpulse
				v7 = b3FloatW4{
					X: v2.X + v5.X,
					Y: v2.Y + v5.Y,
					Z: v2.Z + v5.Z,
					W: v2.W + v5.W,
				}
				goto _396
			_396:
				(*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).TotalNormalImpulses = v7
				v2 = totalNormalImpulse
				v5 = newImpulse
				v7 = b3FloatW4{
					X: v2.X + v5.X,
					Y: v2.Y + v5.Y,
					Z: v2.Z + v5.Z,
					W: v2.W + v5.W,
				}
				goto _400
			_400:
				totalNormalImpulse = v7
				v2 = (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).LeverArms
				v5 = newImpulse
				v7 = b3FloatW4{
					X: float32(v2.X * v5.X),
					Y: float32(v2.Y * v5.Y),
					Z: float32(v2.Z * v5.Z),
					W: float32(v2.W * v5.W),
				}
				goto _404
			_404:
				v10 = totalTwistLimit
				v14 = v7
				v21 = b3FloatW4{
					X: v10.X + v14.X,
					Y: v10.Y + v14.Y,
					Z: v10.Z + v14.Z,
					W: v10.W + v14.W,
				}
				goto _408
			_408:
				totalTwistLimit = v21
				v2 = linearNormalVelocity
				v5 = invMass
				v7 = impulse
				v10 = b3FloatW4{
					X: v2.X + float32(v5.X*v7.X),
					Y: v2.Y + float32(v5.Y*v7.Y),
					Z: v2.Z + float32(v5.Z*v7.Z),
					W: v2.W + float32(v5.W*v7.W),
				}
				goto _413
			_413:
				linearNormalVelocity = v10
				v2 = manifoldImpulse
				v5 = impulse
				v7 = b3FloatW4{
					X: v2.X + v5.X,
					Y: v2.Y + v5.Y,
					Z: v2.Z + v5.Z,
					W: v2.W + v5.W,
				}
				goto _417
			_417:
				manifoldImpulse = v7
				v17 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
				v2 = impulse
				v18 = gA
				v5 = v2
				v7 = v18.X
				v10 = b3FloatW4{
					X: float32(v5.X * v7.X),
					Y: float32(v5.Y * v7.Y),
					Z: float32(v5.Z * v7.Z),
					W: float32(v5.W * v7.W),
				}
				goto _426
			_426:
				v14 = v17.X
				v21 = v10
				v22 = b3FloatW4{
					X: v14.X - v21.X,
					Y: v14.Y - v21.Y,
					Z: v14.Z - v21.Z,
					W: v14.W - v21.W,
				}
				goto _430
			_430:
				v23 = v2
				v25 = v18.Y
				v26 = b3FloatW4{
					X: float32(v23.X * v25.X),
					Y: float32(v23.Y * v25.Y),
					Z: float32(v23.Z * v25.Z),
					W: float32(v23.W * v25.W),
				}
				goto _434
			_434:
				v27 = v17.Y
				v29 = v26
				v30 = b3FloatW4{
					X: v27.X - v29.X,
					Y: v27.Y - v29.Y,
					Z: v27.Z - v29.Z,
					W: v27.W - v29.W,
				}
				goto _438
			_438:
				v31 = v2
				v33 = v18.Z
				v34 = b3FloatW4{
					X: float32(v31.X * v33.X),
					Y: float32(v31.Y * v33.Y),
					Z: float32(v31.Z * v33.Z),
					W: float32(v31.W * v33.W),
				}
				goto _442
			_442:
				v35 = v17.Z
				v40 = v34
				v42 = b3FloatW4{
					X: v35.X - v40.X,
					Y: v35.Y - v40.Y,
					Z: v35.Z - v40.Z,
					W: v35.W - v40.W,
				}
				goto _446
			_446:
				v19 = b3Vec3W4{
					X: v22,
					Y: v30,
					Z: v42,
				}
				goto _422
			_422:
				(*(*b3BodyStateW)(unsafe.Add(nil, bp))).W = v19
				v17 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
				v2 = impulse
				v18 = gB
				v5 = v2
				v7 = v18.X
				v10 = b3FloatW4{
					X: float32(v5.X * v7.X),
					Y: float32(v5.Y * v7.Y),
					Z: float32(v5.Z * v7.Z),
					W: float32(v5.W * v7.W),
				}
				goto _455
			_455:
				v14 = v17.X
				v21 = v10
				v22 = b3FloatW4{
					X: v14.X + v21.X,
					Y: v14.Y + v21.Y,
					Z: v14.Z + v21.Z,
					W: v14.W + v21.W,
				}
				goto _459
			_459:
				v23 = v2
				v25 = v18.Y
				v26 = b3FloatW4{
					X: float32(v23.X * v25.X),
					Y: float32(v23.Y * v25.Y),
					Z: float32(v23.Z * v25.Z),
					W: float32(v23.W * v25.W),
				}
				goto _463
			_463:
				v27 = v17.Y
				v29 = v26
				v30 = b3FloatW4{
					X: v27.X + v29.X,
					Y: v27.Y + v29.Y,
					Z: v27.Z + v29.Z,
					W: v27.W + v29.W,
				}
				goto _467
			_467:
				v31 = v2
				v33 = v18.Z
				v34 = b3FloatW4{
					X: float32(v31.X * v33.X),
					Y: float32(v31.Y * v33.Y),
					Z: float32(v31.Z * v33.Z),
					W: float32(v31.W * v33.W),
				}
				goto _471
			_471:
				v35 = v17.Z
				v40 = v34
				v42 = b3FloatW4{
					X: v35.X + v40.X,
					Y: v35.Y + v40.Y,
					Z: v35.Z + v40.Z,
					W: v35.W + v40.W,
				}
				goto _475
			_475:
				v19 = b3Vec3W4{
					X: v22,
					Y: v30,
					Z: v42,
				}
				goto _451
			_451:
				(*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W = v19
				goto _158
			_158:
				;
				pointIndex = pointIndex + 1
			}
			// Rolling resistance. After normal constraints because it uses the normal impulse.
			if anyRolling != 0 {
				v17 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
				v18 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
				v2 = v17.X
				v5 = v18.X
				v7 = b3FloatW4{
					X: v2.X - v5.X,
					Y: v2.Y - v5.Y,
					Z: v2.Z - v5.Z,
					W: v2.W - v5.W,
				}
				goto _483
			_483:
				v10 = v17.Y
				v14 = v18.Y
				v21 = b3FloatW4{
					X: v10.X - v14.X,
					Y: v10.Y - v14.Y,
					Z: v10.Z - v14.Z,
					W: v10.W - v14.W,
				}
				goto _487
			_487:
				v22 = v17.Z
				v23 = v18.Z
				v25 = b3FloatW4{
					X: v22.X - v23.X,
					Y: v22.Y - v23.Y,
					Z: v22.Z - v23.Z,
					W: v22.W - v23.W,
				}
				goto _491
			_491:
				v19 = b3Vec3W4{
					X: v7,
					Y: v21,
					Z: v25,
				}
				goto _479
			_479:
				// flip A/B order to negate
				deltaImpulse = b3MulMVW(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c2)).RollingMass, v19)
				oldImpulse = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse
				v38 = oldImpulse
				v39 = deltaImpulse
				v26 = v38.X
				v27 = v39.X
				v29 = b3FloatW4{
					X: v26.X + v27.X,
					Y: v26.Y + v27.Y,
					Z: v26.Z + v27.Z,
					W: v26.W + v27.W,
				}
				goto _499
			_499:
				v30 = v38.Y
				v31 = v39.Y
				v33 = b3FloatW4{
					X: v30.X + v31.X,
					Y: v30.Y + v31.Y,
					Z: v30.Z + v31.Z,
					W: v30.W + v31.W,
				}
				goto _503
			_503:
				v34 = v38.Z
				v35 = v39.Z
				v40 = b3FloatW4{
					X: v34.X + v35.X,
					Y: v34.Y + v35.Y,
					Z: v34.Z + v35.Z,
					W: v34.W + v35.W,
				}
				goto _507
			_507:
				v62 = b3Vec3W4{
					X: v29,
					Y: v33,
					Z: v40,
				}
				goto _495
			_495:
				(*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse = v62
				v2 = (*b3ContactConstraintWide)(unsafe.Pointer(c2)).RollingResistance
				v5 = totalNormalImpulse
				v7 = b3FloatW4{
					X: float32(v2.X * v5.X),
					Y: float32(v2.Y * v5.Y),
					Z: float32(v2.Z * v5.Z),
					W: float32(v2.W * v5.W),
				}
				goto _511
			_511:
				maxImpulse = v7
				v17 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse
				v18 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse
				v14 = v17.X
				v21 = v18.X
				v22 = b3FloatW4{
					X: float32(v14.X * v21.X),
					Y: float32(v14.Y * v21.Y),
					Z: float32(v14.Z * v21.Z),
					W: float32(v14.W * v21.W),
				}
				goto _519
			_519:
				v23 = v17.Y
				v25 = v18.Y
				v26 = b3FloatW4{
					X: float32(v23.X * v25.X),
					Y: float32(v23.Y * v25.Y),
					Z: float32(v23.Z * v25.Z),
					W: float32(v23.W * v25.W),
				}
				goto _523
			_523:
				v27 = v22
				v29 = v26
				v30 = b3FloatW4{
					X: v27.X + v29.X,
					Y: v27.Y + v29.Y,
					Z: v27.Z + v29.Z,
					W: v27.W + v29.W,
				}
				goto _527
			_527:
				v31 = v17.Z
				v33 = v18.Z
				v34 = b3FloatW4{
					X: float32(v31.X * v33.X),
					Y: float32(v31.Y * v33.Y),
					Z: float32(v31.Z * v33.Z),
					W: float32(v31.W * v33.W),
				}
				goto _531
			_531:
				v35 = v30
				v40 = v34
				v42 = b3FloatW4{
					X: v35.X + v40.X,
					Y: v35.Y + v40.Y,
					Z: v35.Z + v40.Z,
					W: v35.W + v40.W,
				}
				goto _535
			_535:
				v10 = v42
				goto _515
			_515:
				lengthSquared = v10
				v43 = epsilonW
				v44 = maxImpulse
				v46 = maxImpulse
				v47 = b3FloatW4{
					X: v43.X + float32(v44.X*v46.X),
					Y: v43.Y + float32(v44.Y*v46.Y),
					Z: v43.Z + float32(v44.Z*v46.Z),
					W: v43.W + float32(v44.W*v46.W),
				}
				goto _540
			_540:
				v48 = lengthSquared
				v50 = v47
				if v48.X > v50.X {
					v1 = float32FromFloat32(1)
				} else {
					v1 = float32FromFloat32(0)
				}
				r2.X = v1
				if v48.Y > v50.Y {
					v4 = float32FromFloat32(1)
				} else {
					v4 = float32FromFloat32(0)
				}
				r2.Y = v4
				if v48.Z > v50.Z {
					v9 = float32FromFloat32(1)
				} else {
					v9 = float32FromFloat32(0)
				}
				r2.Z = v9
				if v48.W > v50.W {
					v225 = float32FromFloat32(1)
				} else {
					v225 = float32FromFloat32(0)
				}
				r2.W = v225
				v51 = r2
				goto _548
			_548:
				// if ( magSqr > maxLambda * maxLambda + FLT_EPSILON )
				//{
				//	c->rollingImpulse *= maxLambda / sqrtf( magSqr );
				// }
				mask1 = v51
				v52 = lengthSquared
				v54 = b3FloatW4{
					X: sqrtf(tls, v52.X),
					Y: sqrtf(tls, v52.Y),
					Z: sqrtf(tls, v52.Z),
					W: sqrtf(tls, v52.W),
				}
				goto _551
			_551:
				v55 = v54
				v56 = epsilonW
				v58 = b3FloatW4{
					X: v55.X + v56.X,
					Y: v55.Y + v56.Y,
					Z: v55.Z + v56.Z,
					W: v55.W + v56.W,
				}
				goto _555
			_555:
				v59 = maxImpulse
				v60 = v58
				v66 = b3FloatW4{
					X: v59.X / v60.X,
					Y: v59.Y / v60.Y,
					Z: v59.Z / v60.Z,
					W: v59.W / v60.W,
				}
				goto _559
			_559:
				// No approximate _mm_rsqrt_ps here to maintain cross-platform determinism
				normalize = v66
				v67 = oneW
				v68 = normalize
				v70 = mask1
				if v70.X != float32FromFloat32(0) {
					v226 = v68.X
				} else {
					v226 = v67.X
				}
				r3.X = v226
				if v70.Y != float32FromFloat32(0) {
					v227 = v68.Y
				} else {
					v227 = v67.Y
				}
				r3.Y = v227
				if v70.Z != float32FromFloat32(0) {
					v228 = v68.Z
				} else {
					v228 = v67.Z
				}
				r3.Z = v228
				if v70.W != float32FromFloat32(0) {
					v383 = v68.W
				} else {
					v383 = v67.W
				}
				r3.W = v383
				v71 = r3
				goto _568
			_568:
				scale = v71
				v72 = scale
				v19 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse
				v74 = v72
				v75 = v19.X
				v76 = b3FloatW4{
					X: float32(v74.X * v75.X),
					Y: float32(v74.Y * v75.Y),
					Z: float32(v74.Z * v75.Z),
					W: float32(v74.W * v75.W),
				}
				goto _576
			_576:
				v80 = v72
				v82 = v19.Y
				v83 = b3FloatW4{
					X: float32(v80.X * v82.X),
					Y: float32(v80.Y * v82.Y),
					Z: float32(v80.Z * v82.Z),
					W: float32(v80.W * v82.W),
				}
				goto _580
			_580:
				v84 = v72
				v86 = v19.Z
				v87 = b3FloatW4{
					X: float32(v84.X * v86.X),
					Y: float32(v84.Y * v86.Y),
					Z: float32(v84.Z * v86.Z),
					W: float32(v84.W * v86.W),
				}
				goto _584
			_584:
				v38 = b3Vec3W4{
					X: v76,
					Y: v83,
					Z: v87,
				}
				goto _572
			_572:
				// Blend the result to ensure there are no +/-0 determinism problems, since 0 * -1 == -0.
				scaledImpulse = v38
				v88 = zeroW
				v90 = scaledImpulse.X
				v91 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingMask
				if v91.X != float32FromFloat32(0) {
					v384 = v90.X
				} else {
					v384 = v88.X
				}
				r3.X = v384
				if v91.Y != float32FromFloat32(0) {
					v385 = v90.Y
				} else {
					v385 = v88.Y
				}
				r3.Y = v385
				if v91.Z != float32FromFloat32(0) {
					v386 = v90.Z
				} else {
					v386 = v88.Z
				}
				r3.Z = v386
				if v91.W != float32FromFloat32(0) {
					v543 = v90.W
				} else {
					v543 = v88.W
				}
				r3.W = v543
				v92 = r3
				goto _593
			_593:
				(*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse.X = v92
				v2 = zeroW
				v5 = scaledImpulse.Y
				v7 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingMask
				if v7.X != float32FromFloat32(0) {
					v1 = v5.X
				} else {
					v1 = v2.X
				}
				r3.X = v1
				if v7.Y != float32FromFloat32(0) {
					v4 = v5.Y
				} else {
					v4 = v2.Y
				}
				r3.Y = v4
				if v7.Z != float32FromFloat32(0) {
					v9 = v5.Z
				} else {
					v9 = v2.Z
				}
				r3.Z = v9
				if v7.W != float32FromFloat32(0) {
					v225 = v5.W
				} else {
					v225 = v2.W
				}
				r3.W = v225
				v10 = r3
				goto _602
			_602:
				(*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse.Y = v10
				v2 = zeroW
				v5 = scaledImpulse.Z
				v7 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingMask
				if v7.X != float32FromFloat32(0) {
					v1 = v5.X
				} else {
					v1 = v2.X
				}
				r3.X = v1
				if v7.Y != float32FromFloat32(0) {
					v4 = v5.Y
				} else {
					v4 = v2.Y
				}
				r3.Y = v4
				if v7.Z != float32FromFloat32(0) {
					v9 = v5.Z
				} else {
					v9 = v2.Z
				}
				r3.Z = v9
				if v7.W != float32FromFloat32(0) {
					v225 = v5.W
				} else {
					v225 = v2.W
				}
				r3.W = v225
				v10 = r3
				goto _611
			_611:
				(*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse.Z = v10
				v17 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse
				v18 = oldImpulse
				v2 = v17.X
				v5 = v18.X
				v7 = b3FloatW4{
					X: v2.X - v5.X,
					Y: v2.Y - v5.Y,
					Z: v2.Z - v5.Z,
					W: v2.W - v5.W,
				}
				goto _619
			_619:
				v10 = v17.Y
				v14 = v18.Y
				v21 = b3FloatW4{
					X: v10.X - v14.X,
					Y: v10.Y - v14.Y,
					Z: v10.Z - v14.Z,
					W: v10.W - v14.W,
				}
				goto _623
			_623:
				v22 = v17.Z
				v23 = v18.Z
				v25 = b3FloatW4{
					X: v22.X - v23.X,
					Y: v22.Y - v23.Y,
					Z: v22.Z - v23.Z,
					W: v22.W - v23.W,
				}
				goto _627
			_627:
				v19 = b3Vec3W4{
					X: v7,
					Y: v21,
					Z: v25,
				}
				goto _615
			_615:
				deltaImpulse = v19
				(*(*b3BodyStateW)(unsafe.Add(nil, bp))).W = b3MulSubMVW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W, (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvIA, deltaImpulse)
				(*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W = b3MulAddMVW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W, (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvIB, deltaImpulse)
			}
			// Central twist friction. Friction goes last to improve stacking stability.
			v17 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
			v18 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
			v2 = v17.X
			v5 = v18.X
			v7 = b3FloatW4{
				X: v2.X - v5.X,
				Y: v2.Y - v5.Y,
				Z: v2.Z - v5.Z,
				W: v2.W - v5.W,
			}
			goto _635
		_635:
			v10 = v17.Y
			v14 = v18.Y
			v21 = b3FloatW4{
				X: v10.X - v14.X,
				Y: v10.Y - v14.Y,
				Z: v10.Z - v14.Z,
				W: v10.W - v14.W,
			}
			goto _639
		_639:
			v22 = v17.Z
			v23 = v18.Z
			v25 = b3FloatW4{
				X: v22.X - v23.X,
				Y: v22.Y - v23.Y,
				Z: v22.Z - v23.Z,
				W: v22.W - v23.W,
			}
			goto _643
		_643:
			v19 = b3Vec3W4{
				X: v7,
				Y: v21,
				Z: v25,
			}
			goto _631
		_631:
			v38 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
			v39 = v19
			v27 = v38.X
			v29 = v39.X
			v30 = b3FloatW4{
				X: float32(v27.X * v29.X),
				Y: float32(v27.Y * v29.Y),
				Z: float32(v27.Z * v29.Z),
				W: float32(v27.W * v29.W),
			}
			goto _651
		_651:
			v31 = v38.Y
			v33 = v39.Y
			v34 = b3FloatW4{
				X: float32(v31.X * v33.X),
				Y: float32(v31.Y * v33.Y),
				Z: float32(v31.Z * v33.Z),
				W: float32(v31.W * v33.W),
			}
			goto _655
		_655:
			v35 = v30
			v40 = v34
			v42 = b3FloatW4{
				X: v35.X + v40.X,
				Y: v35.Y + v40.Y,
				Z: v35.Z + v40.Z,
				W: v35.W + v40.W,
			}
			goto _659
		_659:
			v43 = v38.Z
			v44 = v39.Z
			v46 = b3FloatW4{
				X: float32(v43.X * v44.X),
				Y: float32(v43.Y * v44.Y),
				Z: float32(v43.Z * v44.Z),
				W: float32(v43.W * v44.W),
			}
			goto _663
		_663:
			v47 = v42
			v48 = v46
			v50 = b3FloatW4{
				X: v47.X + v48.X,
				Y: v47.Y + v48.Y,
				Z: v47.Z + v48.Z,
				W: v47.W + v48.W,
			}
			goto _667
		_667:
			v26 = v50
			goto _647
		_647:
			twistSpeed = v26
			v51 = (*b3ContactConstraintWide)(unsafe.Pointer(c2)).Friction
			v52 = totalTwistLimit
			v54 = b3FloatW4{
				X: float32(v51.X * v52.X),
				Y: float32(v51.Y * v52.Y),
				Z: float32(v51.Z * v52.Z),
				W: float32(v51.W * v52.W),
			}
			goto _671
		_671:
			maxLambda = v54
			oldImpulse1 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).TwistImpulse
			v55 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).TwistMass
			v56 = twistSpeed
			v58 = b3FloatW4{
				X: float32(v55.X * v56.X),
				Y: float32(v55.Y * v56.Y),
				Z: float32(v55.Z * v56.Z),
				W: float32(v55.W * v56.W),
			}
			goto _675
		_675:
			v59 = v58
			v60 = b3FloatW4{
				X: -v59.X,
				Y: -v59.Y,
				Z: -v59.Z,
				W: -v59.W,
			}
			goto _678
		_678:
			impulse1 = v60
			v66 = maxLambda
			v67 = b3FloatW4{
				X: -v66.X,
				Y: -v66.Y,
				Z: -v66.Z,
				W: -v66.W,
			}
			goto _681
		_681:
			v68 = v67
			v70 = oldImpulse1
			v71 = b3FloatW4{
				X: v68.X - v70.X,
				Y: v68.Y - v70.Y,
				Z: v68.Z - v70.Z,
				W: v68.W - v70.W,
			}
			goto _685
		_685:
			v72 = maxLambda
			v74 = oldImpulse1
			v75 = b3FloatW4{
				X: v72.X - v74.X,
				Y: v72.Y - v74.Y,
				Z: v72.Z - v74.Z,
				W: v72.W - v74.W,
			}
			goto _689
		_689:
			v76 = impulse1
			v80 = v75
			if v76.X <= v80.X {
				v1 = v76.X
			} else {
				v1 = v80.X
			}
			r.X = v1
			if v76.Y <= v80.Y {
				v4 = v76.Y
			} else {
				v4 = v80.Y
			}
			r.Y = v4
			if v76.Z <= v80.Z {
				v9 = v76.Z
			} else {
				v9 = v80.Z
			}
			r.Z = v9
			if v76.W <= v80.W {
				v225 = v76.W
			} else {
				v225 = v80.W
			}
			r.W = v225
			v82 = r
			goto _697
		_697:
			v83 = v71
			v84 = v82
			if v83.X >= v84.X {
				v226 = v83.X
			} else {
				v226 = v84.X
			}
			r1.X = v226
			if v83.Y >= v84.Y {
				v227 = v83.Y
			} else {
				v227 = v84.Y
			}
			r1.Y = v227
			if v83.Z >= v84.Z {
				v228 = v83.Z
			} else {
				v228 = v84.Z
			}
			r1.Z = v228
			if v83.W >= v84.W {
				v383 = v83.W
			} else {
				v383 = v84.W
			}
			r1.W = v383
			v86 = r1
			goto _705
		_705:
			impulse1 = v86
			v2 = oldImpulse1
			v5 = impulse1
			v7 = b3FloatW4{
				X: v2.X + v5.X,
				Y: v2.Y + v5.Y,
				Z: v2.Z + v5.Z,
				W: v2.W + v5.W,
			}
			goto _709
		_709:
			(*b3ContactManifoldWide)(unsafe.Pointer(cm)).TwistImpulse = v7
			v17 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
			v2 = impulse1
			v18 = b3MulMVW(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvIA, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal)
			v5 = v2
			v7 = v18.X
			v10 = b3FloatW4{
				X: float32(v5.X * v7.X),
				Y: float32(v5.Y * v7.Y),
				Z: float32(v5.Z * v7.Z),
				W: float32(v5.W * v7.W),
			}
			goto _718
		_718:
			v14 = v17.X
			v21 = v10
			v22 = b3FloatW4{
				X: v14.X - v21.X,
				Y: v14.Y - v21.Y,
				Z: v14.Z - v21.Z,
				W: v14.W - v21.W,
			}
			goto _722
		_722:
			v23 = v2
			v25 = v18.Y
			v26 = b3FloatW4{
				X: float32(v23.X * v25.X),
				Y: float32(v23.Y * v25.Y),
				Z: float32(v23.Z * v25.Z),
				W: float32(v23.W * v25.W),
			}
			goto _726
		_726:
			v27 = v17.Y
			v29 = v26
			v30 = b3FloatW4{
				X: v27.X - v29.X,
				Y: v27.Y - v29.Y,
				Z: v27.Z - v29.Z,
				W: v27.W - v29.W,
			}
			goto _730
		_730:
			v31 = v2
			v33 = v18.Z
			v34 = b3FloatW4{
				X: float32(v31.X * v33.X),
				Y: float32(v31.Y * v33.Y),
				Z: float32(v31.Z * v33.Z),
				W: float32(v31.W * v33.W),
			}
			goto _734
		_734:
			v35 = v17.Z
			v40 = v34
			v42 = b3FloatW4{
				X: v35.X - v40.X,
				Y: v35.Y - v40.Y,
				Z: v35.Z - v40.Z,
				W: v35.W - v40.W,
			}
			goto _738
		_738:
			v19 = b3Vec3W4{
				X: v22,
				Y: v30,
				Z: v42,
			}
			goto _714
		_714:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp))).W = v19
			v17 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
			v2 = impulse1
			v18 = b3MulMVW(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvIB, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal)
			v5 = v2
			v7 = v18.X
			v10 = b3FloatW4{
				X: float32(v5.X * v7.X),
				Y: float32(v5.Y * v7.Y),
				Z: float32(v5.Z * v7.Z),
				W: float32(v5.W * v7.W),
			}
			goto _747
		_747:
			v14 = v17.X
			v21 = v10
			v22 = b3FloatW4{
				X: v14.X + v21.X,
				Y: v14.Y + v21.Y,
				Z: v14.Z + v21.Z,
				W: v14.W + v21.W,
			}
			goto _751
		_751:
			v23 = v2
			v25 = v18.Y
			v26 = b3FloatW4{
				X: float32(v23.X * v25.X),
				Y: float32(v23.Y * v25.Y),
				Z: float32(v23.Z * v25.Z),
				W: float32(v23.W * v25.W),
			}
			goto _755
		_755:
			v27 = v17.Y
			v29 = v26
			v30 = b3FloatW4{
				X: v27.X + v29.X,
				Y: v27.Y + v29.Y,
				Z: v27.Z + v29.Z,
				W: v27.W + v29.W,
			}
			goto _759
		_759:
			v31 = v2
			v33 = v18.Z
			v34 = b3FloatW4{
				X: float32(v31.X * v33.X),
				Y: float32(v31.Y * v33.Y),
				Z: float32(v31.Z * v33.Z),
				W: float32(v31.W * v33.W),
			}
			goto _763
		_763:
			v35 = v17.Z
			v40 = v34
			v42 = b3FloatW4{
				X: v35.X + v40.X,
				Y: v35.Y + v40.Y,
				Z: v35.Z + v40.Z,
				W: v35.W + v40.W,
			}
			goto _767
		_767:
			v19 = b3Vec3W4{
				X: v22,
				Y: v30,
				Z: v42,
			}
			goto _743
		_743:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W = v19
			// Central friction. Friction goes last to improve stacking stability.
			tangent1 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Tangent1
			tangent2 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Tangent2
			// Fixed anchor points for applying impulses
			rA1 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).CenterA
			rB1 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).CenterB
			v17 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
			v18 = rB1
			v2 = v17.Y
			v5 = v18.Z
			v7 = b3FloatW4{
				X: float32(v2.X * v5.X),
				Y: float32(v2.Y * v5.Y),
				Z: float32(v2.Z * v5.Z),
				W: float32(v2.W * v5.W),
			}
			goto _773
		_773:
			v10 = v17.Z
			v14 = v18.Y
			v21 = b3FloatW4{
				X: float32(v10.X * v14.X),
				Y: float32(v10.Y * v14.Y),
				Z: float32(v10.Z * v14.Z),
				W: float32(v10.W * v14.W),
			}
			goto _777
		_777:
			v22 = v7
			v23 = v21
			v25 = b3FloatW4{
				X: v22.X - v23.X,
				Y: v22.Y - v23.Y,
				Z: v22.Z - v23.Z,
				W: v22.W - v23.W,
			}
			goto _781
		_781:
			c1.X = v25
			v26 = v17.Z
			v27 = v18.X
			v29 = b3FloatW4{
				X: float32(v26.X * v27.X),
				Y: float32(v26.Y * v27.Y),
				Z: float32(v26.Z * v27.Z),
				W: float32(v26.W * v27.W),
			}
			goto _785
		_785:
			v30 = v17.X
			v31 = v18.Z
			v33 = b3FloatW4{
				X: float32(v30.X * v31.X),
				Y: float32(v30.Y * v31.Y),
				Z: float32(v30.Z * v31.Z),
				W: float32(v30.W * v31.W),
			}
			goto _789
		_789:
			v34 = v29
			v35 = v33
			v40 = b3FloatW4{
				X: v34.X - v35.X,
				Y: v34.Y - v35.Y,
				Z: v34.Z - v35.Z,
				W: v34.W - v35.W,
			}
			goto _793
		_793:
			c1.Y = v40
			v42 = v17.X
			v43 = v18.Y
			v44 = b3FloatW4{
				X: float32(v42.X * v43.X),
				Y: float32(v42.Y * v43.Y),
				Z: float32(v42.Z * v43.Z),
				W: float32(v42.W * v43.W),
			}
			goto _797
		_797:
			v46 = v17.Y
			v47 = v18.X
			v48 = b3FloatW4{
				X: float32(v46.X * v47.X),
				Y: float32(v46.Y * v47.Y),
				Z: float32(v46.Z * v47.Z),
				W: float32(v46.W * v47.W),
			}
			goto _801
		_801:
			v50 = v44
			v51 = v48
			v52 = b3FloatW4{
				X: v50.X - v51.X,
				Y: v50.Y - v51.Y,
				Z: v50.Z - v51.Z,
				W: v50.W - v51.W,
			}
			goto _805
		_805:
			c1.Z = v52
			v19 = c1
			goto _807
		_807:
			v38 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
			v39 = rA1
			v54 = v38.Y
			v55 = v39.Z
			v56 = b3FloatW4{
				X: float32(v54.X * v55.X),
				Y: float32(v54.Y * v55.Y),
				Z: float32(v54.Z * v55.Z),
				W: float32(v54.W * v55.W),
			}
			goto _813
		_813:
			v58 = v38.Z
			v59 = v39.Y
			v60 = b3FloatW4{
				X: float32(v58.X * v59.X),
				Y: float32(v58.Y * v59.Y),
				Z: float32(v58.Z * v59.Z),
				W: float32(v58.W * v59.W),
			}
			goto _817
		_817:
			v66 = v56
			v67 = v60
			v68 = b3FloatW4{
				X: v66.X - v67.X,
				Y: v66.Y - v67.Y,
				Z: v66.Z - v67.Z,
				W: v66.W - v67.W,
			}
			goto _821
		_821:
			c1.X = v68
			v70 = v38.Z
			v71 = v39.X
			v72 = b3FloatW4{
				X: float32(v70.X * v71.X),
				Y: float32(v70.Y * v71.Y),
				Z: float32(v70.Z * v71.Z),
				W: float32(v70.W * v71.W),
			}
			goto _825
		_825:
			v74 = v38.X
			v75 = v39.Z
			v76 = b3FloatW4{
				X: float32(v74.X * v75.X),
				Y: float32(v74.Y * v75.Y),
				Z: float32(v74.Z * v75.Z),
				W: float32(v74.W * v75.W),
			}
			goto _829
		_829:
			v80 = v72
			v82 = v76
			v83 = b3FloatW4{
				X: v80.X - v82.X,
				Y: v80.Y - v82.Y,
				Z: v80.Z - v82.Z,
				W: v80.W - v82.W,
			}
			goto _833
		_833:
			c1.Y = v83
			v84 = v38.X
			v86 = v39.Y
			v87 = b3FloatW4{
				X: float32(v84.X * v86.X),
				Y: float32(v84.Y * v86.Y),
				Z: float32(v84.Z * v86.Z),
				W: float32(v84.W * v86.W),
			}
			goto _837
		_837:
			v88 = v38.Y
			v90 = v39.X
			v91 = b3FloatW4{
				X: float32(v88.X * v90.X),
				Y: float32(v88.Y * v90.Y),
				Z: float32(v88.Z * v90.Z),
				W: float32(v88.W * v90.W),
			}
			goto _841
		_841:
			v92 = v87
			v94 = v91
			v95 = b3FloatW4{
				X: v92.X - v94.X,
				Y: v92.Y - v94.Y,
				Z: v92.Z - v94.Z,
				W: v92.W - v94.W,
			}
			goto _845
		_845:
			c1.Z = v95
			v62 = c1
			goto _847
		_847:
			v63 = v19
			v64 = v62
			v96 = v63.X
			v98 = v64.X
			v99 = b3FloatW4{
				X: v96.X - v98.X,
				Y: v96.Y - v98.Y,
				Z: v96.Z - v98.Z,
				W: v96.W - v98.W,
			}
			goto _855
		_855:
			v100 = v63.Y
			v104 = v64.Y
			v106 = b3FloatW4{
				X: v100.X - v104.X,
				Y: v100.Y - v104.Y,
				Z: v100.Z - v104.Z,
				W: v100.W - v104.W,
			}
			goto _859
		_859:
			v107 = v63.Z
			v108 = v64.Z
			v110 = b3FloatW4{
				X: v107.X - v108.X,
				Y: v107.Y - v108.Y,
				Z: v107.Z - v108.Z,
				W: v107.W - v108.W,
			}
			goto _863
		_863:
			v78 = b3Vec3W4{
				X: v99,
				Y: v106,
				Z: v110,
			}
			goto _851
		_851:
			angularRelativeVelocity = v78
			v79 = angularRelativeVelocity
			v102 = tangent1
			v112 = v79.X
			v114 = v102.X
			v115 = b3FloatW4{
				X: float32(v112.X * v114.X),
				Y: float32(v112.Y * v114.Y),
				Z: float32(v112.Z * v114.Z),
				W: float32(v112.W * v114.W),
			}
			goto _871
		_871:
			v116 = v79.Y
			v118 = v102.Y
			v119 = b3FloatW4{
				X: float32(v116.X * v118.X),
				Y: float32(v116.Y * v118.Y),
				Z: float32(v116.Z * v118.Z),
				W: float32(v116.W * v118.W),
			}
			goto _875
		_875:
			v120 = v115
			v122 = v119
			v123 = b3FloatW4{
				X: v120.X + v122.X,
				Y: v120.Y + v122.Y,
				Z: v120.Z + v122.Z,
				W: v120.W + v122.W,
			}
			goto _879
		_879:
			v124 = v79.Z
			v126 = v102.Z
			v127 = b3FloatW4{
				X: float32(v124.X * v126.X),
				Y: float32(v124.Y * v126.Y),
				Z: float32(v124.Z * v126.Z),
				W: float32(v124.W * v126.W),
			}
			goto _883
		_883:
			v128 = v123
			v132 = v127
			v134 = b3FloatW4{
				X: v128.X + v132.X,
				Y: v128.Y + v132.Y,
				Z: v128.Z + v132.Z,
				W: v128.W + v132.W,
			}
			goto _887
		_887:
			v111 = v134
			goto _867
		_867:
			v135 = linearTangentVelocity1
			v136 = v111
			v138 = b3FloatW4{
				X: v135.X + v136.X,
				Y: v135.Y + v136.Y,
				Z: v135.Z + v136.Z,
				W: v135.W + v136.W,
			}
			goto _891
		_891:
			v103 = angularRelativeVelocity
			v130 = tangent2
			v140 = v103.X
			v142 = v130.X
			v143 = b3FloatW4{
				X: float32(v140.X * v142.X),
				Y: float32(v140.Y * v142.Y),
				Z: float32(v140.Z * v142.Z),
				W: float32(v140.W * v142.W),
			}
			goto _899
		_899:
			v144 = v103.Y
			v146 = v130.Y
			v147 = b3FloatW4{
				X: float32(v144.X * v146.X),
				Y: float32(v144.Y * v146.Y),
				Z: float32(v144.Z * v146.Z),
				W: float32(v144.W * v146.W),
			}
			goto _903
		_903:
			v148 = v143
			v150 = v147
			v151 = b3FloatW4{
				X: v148.X + v150.X,
				Y: v148.Y + v150.Y,
				Z: v148.Z + v150.Z,
				W: v148.W + v150.W,
			}
			goto _907
		_907:
			v152 = v103.Z
			v154 = v130.Z
			v155 = b3FloatW4{
				X: float32(v152.X * v154.X),
				Y: float32(v152.Y * v154.Y),
				Z: float32(v152.Z * v154.Z),
				W: float32(v152.W * v154.W),
			}
			goto _911
		_911:
			v156 = v151
			v161 = v155
			v163 = b3FloatW4{
				X: v156.X + v161.X,
				Y: v156.Y + v161.Y,
				Z: v156.Z + v161.Z,
				W: v156.W + v161.W,
			}
			goto _915
		_915:
			v139 = v163
			goto _895
		_895:
			v164 = linearTangentVelocity2
			v165 = v139
			v167 = b3FloatW4{
				X: v164.X + v165.X,
				Y: v164.Y + v165.Y,
				Z: v164.Z + v165.Z,
				W: v164.W + v165.W,
			}
			goto _919
		_919:
			// Relative tangent velocity at contact
			vt = b3Vec2W{
				X: v138,
				Y: v167,
			}
			// Incremental tangent impulse
			deltaImpulse1 = b3MulMV2W(tls, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).TangentMass, vt)
			v168 = deltaImpulse1.X
			v169 = b3FloatW4{
				X: -v168.X,
				Y: -v168.Y,
				Z: -v168.Z,
				W: -v168.W,
			}
			goto _922
		_922:
			v171 = deltaImpulse1.Y
			v172 = b3FloatW4{
				X: -v171.X,
				Y: -v171.Y,
				Z: -v171.Z,
				W: -v171.W,
			}
			goto _925
		_925:
			deltaImpulse1 = b3Vec2W{
				X: v169,
				Y: v172,
			}
			newImpulse1 = b3AddV2W(tls, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).FrictionImpulse, deltaImpulse1)
			v2 = (*b3ContactConstraintWide)(unsafe.Pointer(c2)).Friction
			v5 = totalNormalImpulse
			v7 = b3FloatW4{
				X: float32(v2.X * v5.X),
				Y: float32(v2.Y * v5.Y),
				Z: float32(v2.Z * v5.Z),
				W: float32(v2.W * v5.W),
			}
			goto _929
		_929:
			maxImpulse1 = v7
			v10 = newImpulse1.X
			v14 = newImpulse1.X
			v21 = b3FloatW4{
				X: float32(v10.X * v14.X),
				Y: float32(v10.Y * v14.Y),
				Z: float32(v10.Z * v14.Z),
				W: float32(v10.W * v14.W),
			}
			goto _933
		_933:
			v22 = newImpulse1.Y
			v23 = newImpulse1.Y
			v25 = b3FloatW4{
				X: float32(v22.X * v23.X),
				Y: float32(v22.Y * v23.Y),
				Z: float32(v22.Z * v23.Z),
				W: float32(v22.W * v23.W),
			}
			goto _937
		_937:
			v26 = v21
			v27 = v25
			v29 = b3FloatW4{
				X: v26.X + v27.X,
				Y: v26.Y + v27.Y,
				Z: v26.Z + v27.Z,
				W: v26.W + v27.W,
			}
			goto _941
		_941:
			// Clamp the accumulated impulse
			lengthSquared1 = v29
			v30 = maxImpulse1
			v31 = maxImpulse1
			v33 = b3FloatW4{
				X: float32(v30.X * v31.X),
				Y: float32(v30.Y * v31.Y),
				Z: float32(v30.Z * v31.Z),
				W: float32(v30.W * v31.W),
			}
			goto _945
		_945:
			v34 = lengthSquared1
			v35 = v33
			if v34.X > v35.X {
				v1 = float32FromFloat32(1)
			} else {
				v1 = float32FromFloat32(0)
			}
			r2.X = v1
			if v34.Y > v35.Y {
				v4 = float32FromFloat32(1)
			} else {
				v4 = float32FromFloat32(0)
			}
			r2.Y = v4
			if v34.Z > v35.Z {
				v9 = float32FromFloat32(1)
			} else {
				v9 = float32FromFloat32(0)
			}
			r2.Z = v9
			if v34.W > v35.W {
				v225 = float32FromFloat32(1)
			} else {
				v225 = float32FromFloat32(0)
			}
			r2.W = v225
			v40 = r2
			goto _953
		_953:
			// Max impulse can be zero
			mask2 = v40
			v42 = lengthSquared1
			v43 = b3FloatW4{
				X: sqrtf(tls, v42.X),
				Y: sqrtf(tls, v42.Y),
				Z: sqrtf(tls, v42.Z),
				W: sqrtf(tls, v42.W),
			}
			goto _956
		_956:
			v44 = v43
			v46 = epsilonW
			v47 = b3FloatW4{
				X: v44.X + v46.X,
				Y: v44.Y + v46.Y,
				Z: v44.Z + v46.Z,
				W: v44.W + v46.W,
			}
			goto _960
		_960:
			v48 = maxImpulse1
			v50 = v47
			v51 = b3FloatW4{
				X: v48.X / v50.X,
				Y: v48.Y / v50.Y,
				Z: v48.Z / v50.Z,
				W: v48.W / v50.W,
			}
			goto _964
		_964:
			// No approximate _mm_rsqrt_ps here to maintain cross-platform determinism. Add epsilon to avoid divide by
			// zero.
			normalize1 = v51
			v52 = oneW
			v54 = normalize1
			v55 = mask2
			if v55.X != float32FromFloat32(0) {
				v226 = v54.X
			} else {
				v226 = v52.X
			}
			r3.X = v226
			if v55.Y != float32FromFloat32(0) {
				v227 = v54.Y
			} else {
				v227 = v52.Y
			}
			r3.Y = v227
			if v55.Z != float32FromFloat32(0) {
				v228 = v54.Z
			} else {
				v228 = v52.Z
			}
			r3.Z = v228
			if v55.W != float32FromFloat32(0) {
				v383 = v54.W
			} else {
				v383 = v52.W
			}
			r3.W = v383
			v56 = r3
			goto _973
		_973:
			scale1 = v56
			v58 = scale1
			v59 = newImpulse1.X
			v60 = b3FloatW4{
				X: float32(v58.X * v59.X),
				Y: float32(v58.Y * v59.Y),
				Z: float32(v58.Z * v59.Z),
				W: float32(v58.W * v59.W),
			}
			goto _977
		_977:
			v66 = scale1
			v67 = newImpulse1.Y
			v68 = b3FloatW4{
				X: float32(v66.X * v67.X),
				Y: float32(v66.Y * v67.Y),
				Z: float32(v66.Z * v67.Z),
				W: float32(v66.W * v67.W),
			}
			goto _981
		_981:
			newImpulse1 = b3Vec2W{
				X: v60,
				Y: v68,
			}
			v2 = newImpulse1.X
			v5 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).FrictionImpulse.X
			v7 = b3FloatW4{
				X: v2.X - v5.X,
				Y: v2.Y - v5.Y,
				Z: v2.Z - v5.Z,
				W: v2.W - v5.W,
			}
			goto _985
		_985:
			v10 = newImpulse1.Y
			v14 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).FrictionImpulse.Y
			v21 = b3FloatW4{
				X: v10.X - v14.X,
				Y: v10.Y - v14.Y,
				Z: v10.Z - v14.Z,
				W: v10.W - v14.W,
			}
			goto _989
		_989:
			deltaImpulse1 = b3Vec2W{
				X: v7,
				Y: v21,
			}
			(*b3ContactManifoldWide)(unsafe.Pointer(cm)).FrictionImpulse = newImpulse1
			v2 = deltaImpulse1.X
			v17 = tangent1
			v5 = v2
			v7 = v17.X
			v10 = b3FloatW4{
				X: float32(v5.X * v7.X),
				Y: float32(v5.Y * v7.Y),
				Z: float32(v5.Z * v7.Z),
				W: float32(v5.W * v7.W),
			}
			goto _997
		_997:
			v14 = v2
			v21 = v17.Y
			v22 = b3FloatW4{
				X: float32(v14.X * v21.X),
				Y: float32(v14.Y * v21.Y),
				Z: float32(v14.Z * v21.Z),
				W: float32(v14.W * v21.W),
			}
			goto _1001
		_1001:
			v23 = v2
			v25 = v17.Z
			v26 = b3FloatW4{
				X: float32(v23.X * v25.X),
				Y: float32(v23.Y * v25.Y),
				Z: float32(v23.Z * v25.Z),
				W: float32(v23.W * v25.W),
			}
			goto _1005
		_1005:
			v18 = b3Vec3W4{
				X: v10,
				Y: v22,
				Z: v26,
			}
			goto _993
		_993:
			v27 = deltaImpulse1.Y
			v19 = tangent2
			v29 = v27
			v30 = v19.X
			v31 = b3FloatW4{
				X: float32(v29.X * v30.X),
				Y: float32(v29.Y * v30.Y),
				Z: float32(v29.Z * v30.Z),
				W: float32(v29.W * v30.W),
			}
			goto _1013
		_1013:
			v33 = v27
			v34 = v19.Y
			v35 = b3FloatW4{
				X: float32(v33.X * v34.X),
				Y: float32(v33.Y * v34.Y),
				Z: float32(v33.Z * v34.Z),
				W: float32(v33.W * v34.W),
			}
			goto _1017
		_1017:
			v40 = v27
			v42 = v19.Z
			v43 = b3FloatW4{
				X: float32(v40.X * v42.X),
				Y: float32(v40.Y * v42.Y),
				Z: float32(v40.Z * v42.Z),
				W: float32(v40.W * v42.W),
			}
			goto _1021
		_1021:
			v38 = b3Vec3W4{
				X: v31,
				Y: v35,
				Z: v43,
			}
			goto _1009
		_1009:
			v39 = v18
			v62 = v38
			v44 = v39.X
			v46 = v62.X
			v47 = b3FloatW4{
				X: v44.X + v46.X,
				Y: v44.Y + v46.Y,
				Z: v44.Z + v46.Z,
				W: v44.W + v46.W,
			}
			goto _1029
		_1029:
			v48 = v39.Y
			v50 = v62.Y
			v51 = b3FloatW4{
				X: v48.X + v50.X,
				Y: v48.Y + v50.Y,
				Z: v48.Z + v50.Z,
				W: v48.W + v50.W,
			}
			goto _1033
		_1033:
			v52 = v39.Z
			v54 = v62.Z
			v55 = b3FloatW4{
				X: v52.X + v54.X,
				Y: v52.Y + v54.Y,
				Z: v52.Z + v54.Z,
				W: v52.W + v54.W,
			}
			goto _1037
		_1037:
			v63 = b3Vec3W4{
				X: v47,
				Y: v51,
				Z: v55,
			}
			goto _1025
		_1025:
			// Apply delta impulse
			P = v63
			v64 = rA1
			v78 = P
			v56 = v64.Y
			v58 = v78.Z
			v59 = b3FloatW4{
				X: float32(v56.X * v58.X),
				Y: float32(v56.Y * v58.Y),
				Z: float32(v56.Z * v58.Z),
				W: float32(v56.W * v58.W),
			}
			goto _1043
		_1043:
			v60 = v64.Z
			v66 = v78.Y
			v67 = b3FloatW4{
				X: float32(v60.X * v66.X),
				Y: float32(v60.Y * v66.Y),
				Z: float32(v60.Z * v66.Z),
				W: float32(v60.W * v66.W),
			}
			goto _1047
		_1047:
			v68 = v59
			v70 = v67
			v71 = b3FloatW4{
				X: v68.X - v70.X,
				Y: v68.Y - v70.Y,
				Z: v68.Z - v70.Z,
				W: v68.W - v70.W,
			}
			goto _1051
		_1051:
			c1.X = v71
			v72 = v64.Z
			v74 = v78.X
			v75 = b3FloatW4{
				X: float32(v72.X * v74.X),
				Y: float32(v72.Y * v74.Y),
				Z: float32(v72.Z * v74.Z),
				W: float32(v72.W * v74.W),
			}
			goto _1055
		_1055:
			v76 = v64.X
			v80 = v78.Z
			v82 = b3FloatW4{
				X: float32(v76.X * v80.X),
				Y: float32(v76.Y * v80.Y),
				Z: float32(v76.Z * v80.Z),
				W: float32(v76.W * v80.W),
			}
			goto _1059
		_1059:
			v83 = v75
			v84 = v82
			v86 = b3FloatW4{
				X: v83.X - v84.X,
				Y: v83.Y - v84.Y,
				Z: v83.Z - v84.Z,
				W: v83.W - v84.W,
			}
			goto _1063
		_1063:
			c1.Y = v86
			v87 = v64.X
			v88 = v78.Y
			v90 = b3FloatW4{
				X: float32(v87.X * v88.X),
				Y: float32(v87.Y * v88.Y),
				Z: float32(v87.Z * v88.Z),
				W: float32(v87.W * v88.W),
			}
			goto _1067
		_1067:
			v91 = v64.Y
			v92 = v78.X
			v94 = b3FloatW4{
				X: float32(v91.X * v92.X),
				Y: float32(v91.Y * v92.Y),
				Z: float32(v91.Z * v92.Z),
				W: float32(v91.W * v92.W),
			}
			goto _1071
		_1071:
			v95 = v90
			v96 = v94
			v98 = b3FloatW4{
				X: v95.X - v96.X,
				Y: v95.Y - v96.Y,
				Z: v95.Z - v96.Z,
				W: v95.W - v96.W,
			}
			goto _1075
		_1075:
			c1.Z = v98
			v79 = c1
			goto _1077
		_1077:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp))).W = b3MulSubMVW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W, (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvIA, v79)
			v17 = rB1
			v18 = P
			v2 = v17.Y
			v5 = v18.Z
			v7 = b3FloatW4{
				X: float32(v2.X * v5.X),
				Y: float32(v2.Y * v5.Y),
				Z: float32(v2.Z * v5.Z),
				W: float32(v2.W * v5.W),
			}
			goto _1083
		_1083:
			v10 = v17.Z
			v14 = v18.Y
			v21 = b3FloatW4{
				X: float32(v10.X * v14.X),
				Y: float32(v10.Y * v14.Y),
				Z: float32(v10.Z * v14.Z),
				W: float32(v10.W * v14.W),
			}
			goto _1087
		_1087:
			v22 = v7
			v23 = v21
			v25 = b3FloatW4{
				X: v22.X - v23.X,
				Y: v22.Y - v23.Y,
				Z: v22.Z - v23.Z,
				W: v22.W - v23.W,
			}
			goto _1091
		_1091:
			c1.X = v25
			v26 = v17.Z
			v27 = v18.X
			v29 = b3FloatW4{
				X: float32(v26.X * v27.X),
				Y: float32(v26.Y * v27.Y),
				Z: float32(v26.Z * v27.Z),
				W: float32(v26.W * v27.W),
			}
			goto _1095
		_1095:
			v30 = v17.X
			v31 = v18.Z
			v33 = b3FloatW4{
				X: float32(v30.X * v31.X),
				Y: float32(v30.Y * v31.Y),
				Z: float32(v30.Z * v31.Z),
				W: float32(v30.W * v31.W),
			}
			goto _1099
		_1099:
			v34 = v29
			v35 = v33
			v40 = b3FloatW4{
				X: v34.X - v35.X,
				Y: v34.Y - v35.Y,
				Z: v34.Z - v35.Z,
				W: v34.W - v35.W,
			}
			goto _1103
		_1103:
			c1.Y = v40
			v42 = v17.X
			v43 = v18.Y
			v44 = b3FloatW4{
				X: float32(v42.X * v43.X),
				Y: float32(v42.Y * v43.Y),
				Z: float32(v42.Z * v43.Z),
				W: float32(v42.W * v43.W),
			}
			goto _1107
		_1107:
			v46 = v17.Y
			v47 = v18.X
			v48 = b3FloatW4{
				X: float32(v46.X * v47.X),
				Y: float32(v46.Y * v47.Y),
				Z: float32(v46.Z * v47.Z),
				W: float32(v46.W * v47.W),
			}
			goto _1111
		_1111:
			v50 = v44
			v51 = v48
			v52 = b3FloatW4{
				X: v50.X - v51.X,
				Y: v50.Y - v51.Y,
				Z: v50.Z - v51.Z,
				W: v50.W - v51.W,
			}
			goto _1115
		_1115:
			c1.Z = v52
			v19 = c1
			goto _1117
		_1117:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W = b3MulAddMVW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W, (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvIB, v19)
			v17 = P
			v2 = manifoldImpulse
			v18 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
			v5 = v2
			v7 = v18.X
			v10 = b3FloatW4{
				X: float32(v5.X * v7.X),
				Y: float32(v5.Y * v7.Y),
				Z: float32(v5.Z * v7.Z),
				W: float32(v5.W * v7.W),
			}
			goto _1126
		_1126:
			v14 = v17.X
			v21 = v10
			v22 = b3FloatW4{
				X: v14.X + v21.X,
				Y: v14.Y + v21.Y,
				Z: v14.Z + v21.Z,
				W: v14.W + v21.W,
			}
			goto _1130
		_1130:
			v23 = v2
			v25 = v18.Y
			v26 = b3FloatW4{
				X: float32(v23.X * v25.X),
				Y: float32(v23.Y * v25.Y),
				Z: float32(v23.Z * v25.Z),
				W: float32(v23.W * v25.W),
			}
			goto _1134
		_1134:
			v27 = v17.Y
			v29 = v26
			v30 = b3FloatW4{
				X: v27.X + v29.X,
				Y: v27.Y + v29.Y,
				Z: v27.Z + v29.Z,
				W: v27.W + v29.W,
			}
			goto _1138
		_1138:
			v31 = v2
			v33 = v18.Z
			v34 = b3FloatW4{
				X: float32(v31.X * v33.X),
				Y: float32(v31.Y * v33.Y),
				Z: float32(v31.Z * v33.Z),
				W: float32(v31.W * v33.W),
			}
			goto _1142
		_1142:
			v35 = v17.Z
			v40 = v34
			v42 = b3FloatW4{
				X: v35.X + v40.X,
				Y: v35.Y + v40.Y,
				Z: v35.Z + v40.Z,
				W: v35.W + v40.W,
			}
			goto _1146
		_1146:
			v19 = b3Vec3W4{
				X: v22,
				Y: v30,
				Z: v42,
			}
			goto _1122
		_1122:
			P = v19
			v17 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).V
			v2 = (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvMassA
			v18 = P
			v5 = v2
			v7 = v18.X
			v10 = b3FloatW4{
				X: float32(v5.X * v7.X),
				Y: float32(v5.Y * v7.Y),
				Z: float32(v5.Z * v7.Z),
				W: float32(v5.W * v7.W),
			}
			goto _1155
		_1155:
			v14 = v17.X
			v21 = v10
			v22 = b3FloatW4{
				X: v14.X - v21.X,
				Y: v14.Y - v21.Y,
				Z: v14.Z - v21.Z,
				W: v14.W - v21.W,
			}
			goto _1159
		_1159:
			v23 = v2
			v25 = v18.Y
			v26 = b3FloatW4{
				X: float32(v23.X * v25.X),
				Y: float32(v23.Y * v25.Y),
				Z: float32(v23.Z * v25.Z),
				W: float32(v23.W * v25.W),
			}
			goto _1163
		_1163:
			v27 = v17.Y
			v29 = v26
			v30 = b3FloatW4{
				X: v27.X - v29.X,
				Y: v27.Y - v29.Y,
				Z: v27.Z - v29.Z,
				W: v27.W - v29.W,
			}
			goto _1167
		_1167:
			v31 = v2
			v33 = v18.Z
			v34 = b3FloatW4{
				X: float32(v31.X * v33.X),
				Y: float32(v31.Y * v33.Y),
				Z: float32(v31.Z * v33.Z),
				W: float32(v31.W * v33.W),
			}
			goto _1171
		_1171:
			v35 = v17.Z
			v40 = v34
			v42 = b3FloatW4{
				X: v35.X - v40.X,
				Y: v35.Y - v40.Y,
				Z: v35.Z - v40.Z,
				W: v35.W - v40.W,
			}
			goto _1175
		_1175:
			v19 = b3Vec3W4{
				X: v22,
				Y: v30,
				Z: v42,
			}
			goto _1151
		_1151:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp))).V = v19
			v17 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).V
			v2 = (*b3ContactConstraintWide)(unsafe.Pointer(c2)).InvMassB
			v18 = P
			v5 = v2
			v7 = v18.X
			v10 = b3FloatW4{
				X: float32(v5.X * v7.X),
				Y: float32(v5.Y * v7.Y),
				Z: float32(v5.Z * v7.Z),
				W: float32(v5.W * v7.W),
			}
			goto _1184
		_1184:
			v14 = v17.X
			v21 = v10
			v22 = b3FloatW4{
				X: v14.X + v21.X,
				Y: v14.Y + v21.Y,
				Z: v14.Z + v21.Z,
				W: v14.W + v21.W,
			}
			goto _1188
		_1188:
			v23 = v2
			v25 = v18.Y
			v26 = b3FloatW4{
				X: float32(v23.X * v25.X),
				Y: float32(v23.Y * v25.Y),
				Z: float32(v23.Z * v25.Z),
				W: float32(v23.W * v25.W),
			}
			goto _1192
		_1192:
			v27 = v17.Y
			v29 = v26
			v30 = b3FloatW4{
				X: v27.X + v29.X,
				Y: v27.Y + v29.Y,
				Z: v27.Z + v29.Z,
				W: v27.W + v29.W,
			}
			goto _1196
		_1196:
			v31 = v2
			v33 = v18.Z
			v34 = b3FloatW4{
				X: float32(v31.X * v33.X),
				Y: float32(v31.Y * v33.Y),
				Z: float32(v31.Z * v33.Z),
				W: float32(v31.W * v33.W),
			}
			goto _1200
		_1200:
			v35 = v17.Z
			v40 = v34
			v42 = b3FloatW4{
				X: v35.X + v40.X,
				Y: v35.Y + v40.Y,
				Z: v35.Z + v40.Z,
				W: v35.W + v40.W,
			}
			goto _1204
		_1204:
			v19 = b3Vec3W4{
				X: v22,
				Y: v30,
				Z: v42,
			}
			goto _1180
		_1180:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).V = v19
			goto _37
		_37:
			;
			manifoldIndex = manifoldIndex + 1
		}
		b3ScatterBodies(tls, states, c2, bp)
		b3ScatterBodies(tls, states, c2+16, bp+208)
		c2 = b3NextConstraintW(tls, c2, manifoldCount)
		goto _12
	_12:
		;
		i = i + 1
	}
}

func b3SolveContacts_ConvexW4(tls *_Stack, block b3SolverBlock, context uintptr) {
	b3SolveContactsW(tls, block, context, (*(*b3GraphColor)(unsafe.Add(nil, (*b3StepContext)(unsafe.Pointer(context)).Graph+uintptr(block.ColorIndex)*120))).WideConstraints, uintptrFromInt32(0), int32(1))
}

func b3SolveContacts_MeshWideW4(tls *_Stack, block b3SolverBlock, context uintptr) {
	var color uintptr
	_ = color
	color = (*b3StepContext)(unsafe.Pointer(context)).Graph + uintptr(block.ColorIndex)*120
	b3SolveContactsW(tls, block, context, (*b3GraphColor)(unsafe.Pointer(color)).WideMeshConstraints, (*b3GraphColor)(unsafe.Pointer(color)).WideMeshManifoldStarts, 0)
}

// C documentation
//
//	// Solver set: setIndex + 4 arrays (note: contactIndices is int array, not contactSims)
func b3SerSolverSet(tls *_Stack, buf uintptr, set uintptr) {
	b3SnapW_I32(tls, buf, (*b3SolverSet)(unsafe.Pointer(set)).SetIndex)
	b3SnapW_I32(tls, buf, (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count)
	if (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count > 0 {
		b3SnapW_Bytes(tls, buf, (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data, (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count*int32FromInt64(252))
	}
	b3SnapW_I32(tls, buf, (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Count)
	if (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Count > 0 {
		b3SnapW_Bytes(tls, buf, (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Data, (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Count*int32FromInt64(64))
	}
	b3SnapW_I32(tls, buf, (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Count)
	if (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Count > 0 {
		b3SnapW_Bytes(tls, buf, (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Data, (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Count*int32FromInt64(444))
	}
	b3SnapW_I32(tls, buf, (*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Count)
	if (*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Count > 0 {
		b3SnapW_Bytes(tls, buf, (*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Data, (*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Count*int32FromInt64(4))
	}
	b3SnapW_I32(tls, buf, (*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Count)
	if (*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Count > 0 {
		b3SnapW_Bytes(tls, buf, (*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Data, (*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Count*int32FromInt64(4))
	}
}

func b3DesSolverSet(tls *_Stack, r uintptr, set uintptr) {
	var cnt, cnt1, cnt2, cnt3, cnt4, elemSize, elemSize1, elemSize2, elemSize3, elemSize4, newSize, newSize1, newSize2, newSize3, newSize4, oldSize, oldSize1, oldSize2, oldSize3, oldSize4 int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = cnt, cnt1, cnt2, cnt3, cnt4, elemSize, elemSize1, elemSize2, elemSize3, elemSize4, newSize, newSize1, newSize2, newSize3, newSize4, oldSize, oldSize1, oldSize2, oldSize3, oldSize4
	(*b3SolverSet)(unsafe.Pointer(set)).SetIndex = b3SnapR_I32(tls, r)
	cnt = b3SnapR_I32(tls, r)
	elemSize = int32FromInt64(252)
	if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 && int32FromUint8(b3SnapCheckCount(tls, r, cnt, elemSize, elemSize)) == false1 {
		(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
	if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 && cnt > 0 {
		if (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Capacity < cnt {
			oldSize = int32FromUint64(uint64FromInt32((*b3SolverSet)(unsafe.Pointer(set)).BodySims.Capacity) * uint64(252))
			newSize = int32FromUint64(uint64FromInt32(cnt) * uint64(252))
			(*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data = b3GrowAlloc(tls, (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data, oldSize, newSize)
			(*b3SolverSet)(unsafe.Pointer(set)).BodySims.Capacity = cnt
		}
		(*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count = cnt
		b3SnapR_Bytes(tls, r, (*b3SolverSet)(unsafe.Pointer(set)).BodySims.Data, cnt*elemSize)
	} else {
		if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 {
			(*b3SolverSet)(unsafe.Pointer(set)).BodySims.Count = 0
		}
	}
	cnt1 = b3SnapR_I32(tls, r)
	elemSize1 = int32FromInt64(64)
	if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 && int32FromUint8(b3SnapCheckCount(tls, r, cnt1, elemSize1, elemSize1)) == false1 {
		(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
	if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 && cnt1 > 0 {
		if (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity < cnt1 {
			oldSize1 = int32FromUint64(uint64FromInt32((*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity) * uint64(64))
			newSize1 = int32FromUint64(uint64FromInt32(cnt1) * uint64(64))
			(*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Data = b3GrowAlloc(tls, (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Data, oldSize1, newSize1)
			(*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Capacity = cnt1
		}
		(*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Count = cnt1
		b3SnapR_Bytes(tls, r, (*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Data, cnt1*elemSize1)
	} else {
		if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 {
			(*b3SolverSet)(unsafe.Pointer(set)).BodyStates.Count = 0
		}
	}
	cnt2 = b3SnapR_I32(tls, r)
	elemSize2 = int32FromInt64(444)
	if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 && int32FromUint8(b3SnapCheckCount(tls, r, cnt2, elemSize2, elemSize2)) == false1 {
		(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
	if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 && cnt2 > 0 {
		if (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Capacity < cnt2 {
			oldSize2 = int32FromUint64(uint64FromInt32((*b3SolverSet)(unsafe.Pointer(set)).JointSims.Capacity) * uint64(444))
			newSize2 = int32FromUint64(uint64FromInt32(cnt2) * uint64(444))
			(*b3SolverSet)(unsafe.Pointer(set)).JointSims.Data = b3GrowAlloc(tls, (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Data, oldSize2, newSize2)
			(*b3SolverSet)(unsafe.Pointer(set)).JointSims.Capacity = cnt2
		}
		(*b3SolverSet)(unsafe.Pointer(set)).JointSims.Count = cnt2
		b3SnapR_Bytes(tls, r, (*b3SolverSet)(unsafe.Pointer(set)).JointSims.Data, cnt2*elemSize2)
	} else {
		if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 {
			(*b3SolverSet)(unsafe.Pointer(set)).JointSims.Count = 0
		}
	}
	cnt3 = b3SnapR_I32(tls, r)
	elemSize3 = int32FromInt64(4)
	if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 && int32FromUint8(b3SnapCheckCount(tls, r, cnt3, elemSize3, elemSize3)) == false1 {
		(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
	if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 && cnt3 > 0 {
		if (*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Capacity < cnt3 {
			oldSize3 = int32FromUint64(uint64FromInt32((*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Capacity) * uint64(4))
			newSize3 = int32FromUint64(uint64FromInt32(cnt3) * uint64(4))
			(*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Data = b3GrowAlloc(tls, (*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Data, oldSize3, newSize3)
			(*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Capacity = cnt3
		}
		(*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Count = cnt3
		b3SnapR_Bytes(tls, r, (*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Data, cnt3*elemSize3)
	} else {
		if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 {
			(*b3SolverSet)(unsafe.Pointer(set)).ContactIndices.Count = 0
		}
	}
	cnt4 = b3SnapR_I32(tls, r)
	elemSize4 = int32FromInt64(4)
	if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 && int32FromUint8(b3SnapCheckCount(tls, r, cnt4, elemSize4, elemSize4)) == false1 {
		(*b3SnapReader)(unsafe.Pointer(r)).Ok = boolUint8(false1 != 0)
	}
	if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 && cnt4 > 0 {
		if (*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity < cnt4 {
			oldSize4 = int32FromUint64(uint64FromInt32((*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity) * uint64(4))
			newSize4 = int32FromUint64(uint64FromInt32(cnt4) * uint64(4))
			(*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Data = b3GrowAlloc(tls, (*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Data, oldSize4, newSize4)
			(*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Capacity = cnt4
		}
		(*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Count = cnt4
		b3SnapR_Bytes(tls, r, (*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Data, cnt4*elemSize4)
	} else {
		if (*b3SnapReader)(unsafe.Pointer(r)).Ok != 0 {
			(*b3SolverSet)(unsafe.Pointer(set)).IslandSims.Count = 0
		}
	}
}
