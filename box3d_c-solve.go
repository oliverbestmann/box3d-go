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
	var anyRolling uint8
	var c, v12, v13, v14, v16, v17, v20, v21, v24, v25, v29, v3, v30, v33, v34, v37, v38, v39, v4, v41, v42, v43, v45, v46, v49, v5, v50, v8, v9 b3Vec3W4
	var c1, cm, cp, manifolds, states uintptr
	var deltaImpulse1, newImpulse1, vt b3Vec2W
	var epsilonW, impulse, impulse1, invMass, inv_h, lengthSquared, lengthSquared1, linearNormalVelocity, linearTangentVelocity1, linearTangentVelocity2, manifoldImpulse, mask, mask1, maxImpulse, maxImpulse1, maxLambda, negImpulse, newImpulse, normalSeparation, normalize, normalize1, oldImpulse1, oneW, s3, scale, scale1, totalNormalImpulse, totalTwistLimit, twistSpeed, velocityBias, vn, zeroW b3FloatW
	var i, manifoldCount, manifoldIndex, pointIndex, v2 int32
	var v10, v18, v22, v26, v31, v35, v47, v51 b3FloatW4
	var _ /* bA at bp+0 */ b3BodyStateW
	var _ /* bB at bp+208 */ b3BodyStateW
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = P, angularRelativeVelocity, anyRolling, c, c1, cm, cp, deltaImpulse, deltaImpulse1, dp, dv, epsilonW, gA, gB, i, impulse, impulse1, invMass, inv_h, lengthSquared, lengthSquared1, linearNormalVelocity, linearTangentVelocity1, linearTangentVelocity2, manifoldCount, manifoldImpulse, manifoldIndex, manifolds, mask, mask1, maxImpulse, maxImpulse1, maxLambda, negImpulse, newImpulse, newImpulse1, normalA, normalB, normalSeparation, normalize, normalize1, oldImpulse, oldImpulse1, oneW, pointIndex, rA, rA1, rB, rB1, s3, scale, scale1, scaledImpulse, states, tangent1, tangent2, totalNormalImpulse, totalTwistLimit, twistSpeed, uA, uB, velocityBias, vn, vt, zeroW, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v22, v24, v25, v26, v29, v3, v30, v31, v33, v34, v35, v37, v38, v39, v4, v41, v42, v43, v45, v46, v47, v49, v5, v50, v51, v8, v9
	states = (*b3StepContext)(unsafe.Pointer(context)).States
	inv_h = b3SplatW4(tls, (*b3StepContext)(unsafe.Pointer(context)).Inv_h)
	oneW = b3SplatW4(tls, float32FromFloat32(1))
	zeroW = b3ZeroW4(tls)
	epsilonW = b3SplatW4(tls, float32FromFloat32(1.1920928955078125e-07))
	c1 = b3GetConstraintW(tls, base, manifoldStarts, block.StartIndex, fixedManifoldCount)
	i = 0
	for {
		if !(i < int32FromUint16(block.Count)) {
			break
		}
		if fixedManifoldCount > 0 {
			v2 = fixedManifoldCount
		} else {
			v2 = (*b3ContactConstraintWide)(unsafe.Pointer(c1)).ManifoldCount
		}
		manifoldCount = v2
		*(*b3BodyStateW)(unsafe.Add(nil, bp)) = b3GatherBodies(tls, states, c1)
		*(*b3BodyStateW)(unsafe.Add(nil, bp+208)) = b3GatherBodies(tls, states, c1+16)
		anyRolling = boolUint8(int32FromUint8(b3AllZeroW4(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).RollingResistance)) == false1)
		v3 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).Dp
		v4 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).Dp
		v5 = b3Vec3W4{
			X: b3SubW4(tls, v3.X, v4.X),
			Y: b3SubW4(tls, v3.Y, v4.Y),
			Z: b3SubW4(tls, v3.Z, v4.Z),
		}
		goto _6
	_6:
		dp = v5
		invMass = b3AddW4(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvMassA, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvMassB)
		manifolds = b3GetManifoldsW(tls, c1)
		manifoldIndex = 0
		for {
			if !(manifoldIndex < manifoldCount) {
				break
			}
			cm = manifolds + uintptr(manifoldIndex)*1288
			v8 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
			v9 = dp
			v10 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v8.X, v9.X), b3MulW4(tls, v8.Y, v9.Y)), b3MulW4(tls, v8.Z, v9.Z))
			goto _11
		_11:
			// Read the normal rather than put on stack to avoid register spills.
			normalSeparation = v10
			normalA = b3InvRotateVectorW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp))).Dq, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal)
			normalB = b3InvRotateVectorW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).Dq, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal)
			v12 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).V
			v13 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).V
			v14 = b3Vec3W4{
				X: b3SubW4(tls, v12.X, v13.X),
				Y: b3SubW4(tls, v12.Y, v13.Y),
				Z: b3SubW4(tls, v12.Z, v13.Z),
			}
			goto _15
		_15:
			dv = v14
			v16 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
			v17 = dv
			v18 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v16.X, v17.X), b3MulW4(tls, v16.Y, v17.Y)), b3MulW4(tls, v16.Z, v17.Z))
			goto _19
		_19:
			linearNormalVelocity = v18
			v20 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Tangent1
			v21 = dv
			v22 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v20.X, v21.X), b3MulW4(tls, v20.Y, v21.Y)), b3MulW4(tls, v20.Z, v21.Z))
			goto _23
		_23:
			linearTangentVelocity1 = b3SubW4(tls, v22, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).TangentVelocity1)
			v24 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Tangent2
			v25 = dv
			v26 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v24.X, v25.X), b3MulW4(tls, v24.Y, v25.Y)), b3MulW4(tls, v24.Z, v25.Z))
			goto _27
		_27:
			linearTangentVelocity2 = b3SubW4(tls, v26, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).TangentVelocity2)
			manifoldImpulse = zeroW
			totalNormalImpulse = zeroW
			totalTwistLimit = zeroW
			// Normal contraints
			pointIndex = 0
			for {
				if !(pointIndex < (*b3ContactManifoldWide)(unsafe.Pointer(cm)).PointCount) {
					break
				}
				cp = cm + 456 + uintptr(pointIndex)*208
				// Fixed anchor points for applying impulses
				rA = (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).AnchorAs
				rB = (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).AnchorBs
				v29 = normalB
				v30 = rB
				v31 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v29.X, v30.X), b3MulW4(tls, v29.Y, v30.Y)), b3MulW4(tls, v29.Z, v30.Z))
				goto _32
			_32:
				v33 = normalA
				v34 = rA
				v35 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v33.X, v34.X), b3MulW4(tls, v33.Y, v34.Y)), b3MulW4(tls, v33.Z, v34.Z))
				goto _36
			_36:
				s3 = b3AddW4(tls, b3AddW4(tls, normalSeparation, b3SubW4(tls, v31, v35)), (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).BaseSeparations)
				// Speculative bias, positive if separated and zero if overlapped
				velocityBias = b3MaxW4(tls, zeroW, b3MulW4(tls, s3, inv_h))
				v37 = rA
				v38 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
				c.X = b3SubW4(tls, b3MulW4(tls, v37.Y, v38.Z), b3MulW4(tls, v37.Z, v38.Y))
				c.Y = b3SubW4(tls, b3MulW4(tls, v37.Z, v38.X), b3MulW4(tls, v37.X, v38.Z))
				c.Z = b3SubW4(tls, b3MulW4(tls, v37.X, v38.Y), b3MulW4(tls, v37.Y, v38.X))
				v39 = c
				goto _40
			_40:
				uA = v39
				v41 = rB
				v42 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
				c.X = b3SubW4(tls, b3MulW4(tls, v41.Y, v42.Z), b3MulW4(tls, v41.Z, v42.Y))
				c.Y = b3SubW4(tls, b3MulW4(tls, v41.Z, v42.X), b3MulW4(tls, v41.X, v42.Z))
				c.Z = b3SubW4(tls, b3MulW4(tls, v41.X, v42.Y), b3MulW4(tls, v41.Y, v42.X))
				v43 = c
				goto _44
			_44:
				uB = v43
				gA = b3MulMVW(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvIA, uA)
				gB = b3MulMVW(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvIB, uB)
				v45 = uB
				v46 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
				v47 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v45.X, v46.X), b3MulW4(tls, v45.Y, v46.Y)), b3MulW4(tls, v45.Z, v46.Z))
				goto _48
			_48:
				v49 = uA
				v50 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
				v51 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v49.X, v50.X), b3MulW4(tls, v49.Y, v50.Y)), b3MulW4(tls, v49.Z, v50.Z))
				goto _52
			_52:
				vn = b3AddW4(tls, linearNormalVelocity, b3SubW4(tls, v47, v51))
				negImpulse = b3MulW4(tls, (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).NormalMasses, b3AddW4(tls, vn, velocityBias))
				impulse = b3MaxW4(tls, b3NegW4(tls, negImpulse), b3NegW4(tls, (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).NormalImpulses))
				newImpulse = b3AddW4(tls, (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).NormalImpulses, impulse)
				(*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).NormalImpulses = newImpulse
				(*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).TotalNormalImpulses = b3AddW4(tls, (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).TotalNormalImpulses, newImpulse)
				totalNormalImpulse = b3AddW4(tls, totalNormalImpulse, newImpulse)
				totalTwistLimit = b3AddW4(tls, totalTwistLimit, b3MulW4(tls, (*b3ContactConstraintPointWide)(unsafe.Pointer(cp)).LeverArms, newImpulse))
				linearNormalVelocity = b3MulAddW4(tls, linearNormalVelocity, invMass, impulse)
				manifoldImpulse = b3AddW4(tls, manifoldImpulse, impulse)
				v3 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
				v10 = impulse
				v4 = gA
				v5 = b3Vec3W4{
					X: b3SubW4(tls, v3.X, b3MulW4(tls, v10, v4.X)),
					Y: b3SubW4(tls, v3.Y, b3MulW4(tls, v10, v4.Y)),
					Z: b3SubW4(tls, v3.Z, b3MulW4(tls, v10, v4.Z)),
				}
				goto _57
			_57:
				(*(*b3BodyStateW)(unsafe.Add(nil, bp))).W = v5
				v3 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
				v10 = impulse
				v4 = gB
				v5 = b3Vec3W4{
					X: b3AddW4(tls, v3.X, b3MulW4(tls, v10, v4.X)),
					Y: b3AddW4(tls, v3.Y, b3MulW4(tls, v10, v4.Y)),
					Z: b3AddW4(tls, v3.Z, b3MulW4(tls, v10, v4.Z)),
				}
				goto _62
			_62:
				(*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W = v5
				goto _28
			_28:
				;
				pointIndex = pointIndex + 1
			}
			// Rolling resistance. After normal constraints because it uses the normal impulse.
			if anyRolling != 0 {
				v3 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
				v4 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
				v5 = b3Vec3W4{
					X: b3SubW4(tls, v3.X, v4.X),
					Y: b3SubW4(tls, v3.Y, v4.Y),
					Z: b3SubW4(tls, v3.Z, v4.Z),
				}
				goto _66
			_66:
				// flip A/B order to negate
				deltaImpulse = b3MulMVW(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).RollingMass, v5)
				oldImpulse = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse
				v8 = oldImpulse
				v9 = deltaImpulse
				v12 = b3Vec3W4{
					X: b3AddW4(tls, v8.X, v9.X),
					Y: b3AddW4(tls, v8.Y, v9.Y),
					Z: b3AddW4(tls, v8.Z, v9.Z),
				}
				goto _70
			_70:
				(*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse = v12
				maxImpulse = b3MulW4(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).RollingResistance, totalNormalImpulse)
				v3 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse
				v4 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse
				v10 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v3.X, v4.X), b3MulW4(tls, v3.Y, v4.Y)), b3MulW4(tls, v3.Z, v4.Z))
				goto _74
			_74:
				lengthSquared = v10
				// if ( magSqr > maxLambda * maxLambda + FLT_EPSILON )
				//{
				//	c->rollingImpulse *= maxLambda / sqrtf( magSqr );
				// }
				mask = b3GreaterThanW4(tls, lengthSquared, b3MulAddW4(tls, epsilonW, maxImpulse, maxImpulse))
				// No approximate _mm_rsqrt_ps here to maintain cross-platform determinism
				normalize = b3DivW4(tls, maxImpulse, b3AddW4(tls, b3SqrtW4(tls, lengthSquared), epsilonW))
				scale = b3BlendW4(tls, oneW, normalize, mask)
				v18 = scale
				v5 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse
				v8 = b3Vec3W4{
					X: b3MulW4(tls, v18, v5.X),
					Y: b3MulW4(tls, v18, v5.Y),
					Z: b3MulW4(tls, v18, v5.Z),
				}
				goto _78
			_78:
				// Blend the result to ensure there are no +/-0 determinism problems, since 0 * -1 == -0.
				scaledImpulse = v8
				(*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse.X = b3BlendW4(tls, zeroW, scaledImpulse.X, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingMask)
				(*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse.Y = b3BlendW4(tls, zeroW, scaledImpulse.Y, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingMask)
				(*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse.Z = b3BlendW4(tls, zeroW, scaledImpulse.Z, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingMask)
				v3 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).RollingImpulse
				v4 = oldImpulse
				v5 = b3Vec3W4{
					X: b3SubW4(tls, v3.X, v4.X),
					Y: b3SubW4(tls, v3.Y, v4.Y),
					Z: b3SubW4(tls, v3.Z, v4.Z),
				}
				goto _82
			_82:
				deltaImpulse = v5
				(*(*b3BodyStateW)(unsafe.Add(nil, bp))).W = b3MulSubMVW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvIA, deltaImpulse)
				(*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W = b3MulAddMVW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvIB, deltaImpulse)
			}
			// Central twist friction. Friction goes last to improve stacking stability.
			v3 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
			v4 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
			v5 = b3Vec3W4{
				X: b3SubW4(tls, v3.X, v4.X),
				Y: b3SubW4(tls, v3.Y, v4.Y),
				Z: b3SubW4(tls, v3.Z, v4.Z),
			}
			goto _86
		_86:
			v8 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
			v9 = v5
			v10 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v8.X, v9.X), b3MulW4(tls, v8.Y, v9.Y)), b3MulW4(tls, v8.Z, v9.Z))
			goto _90
		_90:
			twistSpeed = v10
			maxLambda = b3MulW4(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).Friction, totalTwistLimit)
			oldImpulse1 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).TwistImpulse
			impulse1 = b3NegW4(tls, b3MulW4(tls, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).TwistMass, twistSpeed))
			impulse1 = b3MaxW4(tls, b3SubW4(tls, b3NegW4(tls, maxLambda), oldImpulse1), b3MinW4(tls, impulse1, b3SubW4(tls, maxLambda, oldImpulse1)))
			(*b3ContactManifoldWide)(unsafe.Pointer(cm)).TwistImpulse = b3AddW4(tls, oldImpulse1, impulse1)
			v3 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
			v10 = impulse1
			v4 = b3MulMVW(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvIA, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal)
			v5 = b3Vec3W4{
				X: b3SubW4(tls, v3.X, b3MulW4(tls, v10, v4.X)),
				Y: b3SubW4(tls, v3.Y, b3MulW4(tls, v10, v4.Y)),
				Z: b3SubW4(tls, v3.Z, b3MulW4(tls, v10, v4.Z)),
			}
			goto _95
		_95:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp))).W = v5
			v3 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
			v10 = impulse1
			v4 = b3MulMVW(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvIB, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal)
			v5 = b3Vec3W4{
				X: b3AddW4(tls, v3.X, b3MulW4(tls, v10, v4.X)),
				Y: b3AddW4(tls, v3.Y, b3MulW4(tls, v10, v4.Y)),
				Z: b3AddW4(tls, v3.Z, b3MulW4(tls, v10, v4.Z)),
			}
			goto _100
		_100:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W = v5
			// Central friction. Friction goes last to improve stacking stability.
			tangent1 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Tangent1
			tangent2 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Tangent2
			// Fixed anchor points for applying impulses
			rA1 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).CenterA
			rB1 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).CenterB
			v3 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W
			v4 = rB1
			c.X = b3SubW4(tls, b3MulW4(tls, v3.Y, v4.Z), b3MulW4(tls, v3.Z, v4.Y))
			c.Y = b3SubW4(tls, b3MulW4(tls, v3.Z, v4.X), b3MulW4(tls, v3.X, v4.Z))
			c.Z = b3SubW4(tls, b3MulW4(tls, v3.X, v4.Y), b3MulW4(tls, v3.Y, v4.X))
			v5 = c
			goto _104
		_104:
			v8 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W
			v9 = rA1
			c.X = b3SubW4(tls, b3MulW4(tls, v8.Y, v9.Z), b3MulW4(tls, v8.Z, v9.Y))
			c.Y = b3SubW4(tls, b3MulW4(tls, v8.Z, v9.X), b3MulW4(tls, v8.X, v9.Z))
			c.Z = b3SubW4(tls, b3MulW4(tls, v8.X, v9.Y), b3MulW4(tls, v8.Y, v9.X))
			v12 = c
			goto _108
		_108:
			v13 = v5
			v14 = v12
			v16 = b3Vec3W4{
				X: b3SubW4(tls, v13.X, v14.X),
				Y: b3SubW4(tls, v13.Y, v14.Y),
				Z: b3SubW4(tls, v13.Z, v14.Z),
			}
			goto _112
		_112:
			angularRelativeVelocity = v16
			v17 = angularRelativeVelocity
			v20 = tangent1
			v10 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v17.X, v20.X), b3MulW4(tls, v17.Y, v20.Y)), b3MulW4(tls, v17.Z, v20.Z))
			goto _116
		_116:
			v21 = angularRelativeVelocity
			v24 = tangent2
			v18 = b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, v21.X, v24.X), b3MulW4(tls, v21.Y, v24.Y)), b3MulW4(tls, v21.Z, v24.Z))
			goto _120
		_120:
			// Relative tangent velocity at contact
			vt = b3Vec2W{
				X: b3AddW4(tls, linearTangentVelocity1, v10),
				Y: b3AddW4(tls, linearTangentVelocity2, v18),
			}
			// Incremental tangent impulse
			deltaImpulse1 = b3MulMV2W(tls, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).TangentMass, vt)
			deltaImpulse1 = b3Vec2W{
				X: b3NegW4(tls, deltaImpulse1.X),
				Y: b3NegW4(tls, deltaImpulse1.Y),
			}
			newImpulse1 = b3AddV2W(tls, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).FrictionImpulse, deltaImpulse1)
			maxImpulse1 = b3MulW4(tls, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).Friction, totalNormalImpulse)
			// Clamp the accumulated impulse
			lengthSquared1 = b3AddW4(tls, b3MulW4(tls, newImpulse1.X, newImpulse1.X), b3MulW4(tls, newImpulse1.Y, newImpulse1.Y))
			// Max impulse can be zero
			mask1 = b3GreaterThanW4(tls, lengthSquared1, b3MulW4(tls, maxImpulse1, maxImpulse1))
			// No approximate _mm_rsqrt_ps here to maintain cross-platform determinism. Add epsilon to avoid divide by
			// zero.
			normalize1 = b3DivW4(tls, maxImpulse1, b3AddW4(tls, b3SqrtW4(tls, lengthSquared1), epsilonW))
			scale1 = b3BlendW4(tls, oneW, normalize1, mask1)
			newImpulse1 = b3Vec2W{
				X: b3MulW4(tls, scale1, newImpulse1.X),
				Y: b3MulW4(tls, scale1, newImpulse1.Y),
			}
			deltaImpulse1 = b3Vec2W{
				X: b3SubW4(tls, newImpulse1.X, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).FrictionImpulse.X),
				Y: b3SubW4(tls, newImpulse1.Y, (*b3ContactManifoldWide)(unsafe.Pointer(cm)).FrictionImpulse.Y),
			}
			(*b3ContactManifoldWide)(unsafe.Pointer(cm)).FrictionImpulse = newImpulse1
			v10 = deltaImpulse1.X
			v3 = tangent1
			v4 = b3Vec3W4{
				X: b3MulW4(tls, v10, v3.X),
				Y: b3MulW4(tls, v10, v3.Y),
				Z: b3MulW4(tls, v10, v3.Z),
			}
			goto _124
		_124:
			v18 = deltaImpulse1.Y
			v5 = tangent2
			v8 = b3Vec3W4{
				X: b3MulW4(tls, v18, v5.X),
				Y: b3MulW4(tls, v18, v5.Y),
				Z: b3MulW4(tls, v18, v5.Z),
			}
			goto _128
		_128:
			v9 = v4
			v12 = v8
			v13 = b3Vec3W4{
				X: b3AddW4(tls, v9.X, v12.X),
				Y: b3AddW4(tls, v9.Y, v12.Y),
				Z: b3AddW4(tls, v9.Z, v12.Z),
			}
			goto _132
		_132:
			// Apply delta impulse
			P = v13
			v14 = rA1
			v16 = P
			c.X = b3SubW4(tls, b3MulW4(tls, v14.Y, v16.Z), b3MulW4(tls, v14.Z, v16.Y))
			c.Y = b3SubW4(tls, b3MulW4(tls, v14.Z, v16.X), b3MulW4(tls, v14.X, v16.Z))
			c.Z = b3SubW4(tls, b3MulW4(tls, v14.X, v16.Y), b3MulW4(tls, v14.Y, v16.X))
			v17 = c
			goto _136
		_136:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp))).W = b3MulSubMVW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp))).W, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvIA, v17)
			v3 = rB1
			v4 = P
			c.X = b3SubW4(tls, b3MulW4(tls, v3.Y, v4.Z), b3MulW4(tls, v3.Z, v4.Y))
			c.Y = b3SubW4(tls, b3MulW4(tls, v3.Z, v4.X), b3MulW4(tls, v3.X, v4.Z))
			c.Z = b3SubW4(tls, b3MulW4(tls, v3.X, v4.Y), b3MulW4(tls, v3.Y, v4.X))
			v5 = c
			goto _140
		_140:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W = b3MulAddMVW(tls, (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).W, (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvIB, v5)
			v3 = P
			v10 = manifoldImpulse
			v4 = (*b3ContactManifoldWide)(unsafe.Pointer(cm)).Normal
			v5 = b3Vec3W4{
				X: b3AddW4(tls, v3.X, b3MulW4(tls, v10, v4.X)),
				Y: b3AddW4(tls, v3.Y, b3MulW4(tls, v10, v4.Y)),
				Z: b3AddW4(tls, v3.Z, b3MulW4(tls, v10, v4.Z)),
			}
			goto _145
		_145:
			P = v5
			v3 = (*(*b3BodyStateW)(unsafe.Add(nil, bp))).V
			v10 = (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvMassA
			v4 = P
			v5 = b3Vec3W4{
				X: b3SubW4(tls, v3.X, b3MulW4(tls, v10, v4.X)),
				Y: b3SubW4(tls, v3.Y, b3MulW4(tls, v10, v4.Y)),
				Z: b3SubW4(tls, v3.Z, b3MulW4(tls, v10, v4.Z)),
			}
			goto _150
		_150:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp))).V = v5
			v3 = (*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).V
			v10 = (*b3ContactConstraintWide)(unsafe.Pointer(c1)).InvMassB
			v4 = P
			v5 = b3Vec3W4{
				X: b3AddW4(tls, v3.X, b3MulW4(tls, v10, v4.X)),
				Y: b3AddW4(tls, v3.Y, b3MulW4(tls, v10, v4.Y)),
				Z: b3AddW4(tls, v3.Z, b3MulW4(tls, v10, v4.Z)),
			}
			goto _155
		_155:
			(*(*b3BodyStateW)(unsafe.Add(nil, bp+208))).V = v5
			goto _7
		_7:
			;
			manifoldIndex = manifoldIndex + 1
		}
		b3ScatterBodies(tls, states, c1, bp)
		b3ScatterBodies(tls, states, c1+16, bp+208)
		c1 = b3NextConstraintW(tls, c1, manifoldCount)
		goto _1
	_1:
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
