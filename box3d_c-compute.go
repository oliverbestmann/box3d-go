package b3

import (
	"unsafe"
	"reflect"
)

var _ unsafe.Pointer
var _ reflect.Type

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

func b3ComputeCapsuleMass(tls *_Stack, shape uintptr, density float32) (r1 MassData) {
	var b4, c1, c2, center, direction, dv, u, v1, v16, v17, v18, v2, v20, v21, v26, v27, v5, v8, v9 Vec3
	var cylinderHeight, cylinderMass, cylinderVolume, lengthSquared, mass, r, s, sphereMass, sphereVolume, steiner, xw, xx, xy, xz, yw, yy, yz, zw, zz, v10, v3, v6 float32
	var inertia, out, out1, rotation, v12, v13, v14, v24, v34, v35, v37, v38, v39, v43, v47, v51, v53, v54 Matrix3
	var out2 MassData
	var q1, v23 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b4, c1, c2, center, cylinderHeight, cylinderMass, cylinderVolume, direction, dv, inertia, lengthSquared, mass, out, out1, out2, q1, r, rotation, s, sphereMass, sphereVolume, steiner, u, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v12, v13, v14, v16, v17, v18, v2, v20, v21, v23, v24, v26, v27, v3, v34, v35, v37, v38, v39, v43, v47, v5, v51, v53, v54, v6, v8, v9
	c1 = (*Capsule)(unsafe.Pointer(shape)).Center1
	c2 = (*Capsule)(unsafe.Pointer(shape)).Center2
	r = (*Capsule)(unsafe.Pointer(shape)).Radius
	v1 = c1
	v2 = c2
	dv = Vec3{
		X: v2.X - v1.X,
		Y: v2.Y - v1.Y,
		Z: v2.Z - v1.Z,
	}
	v5 = dv
	v8 = v5
	v9 = v5
	v10 = float32(v8.X*v9.X) + float32(v8.Y*v9.Y) + float32(v8.Z*v9.Z)
	goto _11
_11:
	v6 = sqrtf(tls, v10)
	goto _7
_7:
	v3 = v6
	goto _4
_4:
	// Cylinder
	cylinderHeight = v3
	cylinderVolume = float32(float32(float32(float32FromFloat32(3.14159265359)*r)*r) * cylinderHeight)
	cylinderMass = float32(cylinderVolume * density)
	// Sphere
	sphereVolume = float32(float32(float32(float32(float32FromFloat32(4)/float32FromFloat32(3)*float32FromFloat32(3.14159265359))*r)*r) * r)
	sphereMass = float32(sphereVolume * density)
	v12 = b3CylinderInertia(tls, cylinderMass, r, cylinderHeight)
	v13 = b3SphereInertia(tls, sphereMass, r)
	v14 = Matrix3{
		Cx: Vec3{
			X: v12.Cx.X + v13.Cx.X,
			Y: v12.Cx.Y + v13.Cx.Y,
			Z: v12.Cx.Z + v13.Cx.Z,
		},
		Cy: Vec3{
			X: v12.Cy.X + v13.Cy.X,
			Y: v12.Cy.Y + v13.Cy.Y,
			Z: v12.Cy.Z + v13.Cy.Z,
		},
		Cz: Vec3{
			X: v12.Cz.X + v13.Cz.X,
			Y: v12.Cz.Y + v13.Cz.Y,
			Z: v12.Cz.Z + v13.Cz.Z,
		},
	}
	goto _15
_15:
	// Local accumulated inertia
	inertia = v14
	steiner = float32(float32(float32(float32FromFloat32(0.125)*sphereMass)*(float32(float32FromFloat32(3)*r)+float32(float32FromFloat32(2)*cylinderHeight))) * cylinderHeight)
	inertia.Cx.X += steiner
	inertia.Cz.Z += steiner
	// Align capsule axis with chosen up-axis
	rotation = b3Mat3_identity
	if float32(cylinderHeight*cylinderHeight) > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		v1 = c2
		v2 = c1
		v5 = Vec3{
			X: v1.X - v2.X,
			Y: v1.Y - v2.Y,
			Z: v1.Z - v2.Z,
		}
		goto _19
	_19:
		v8 = v5
		lengthSquared = float32(v8.X*v8.X) + float32(v8.Y*v8.Y) + float32(v8.Z*v8.Z)
		if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
			s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
			u = Vec3{
				X: float32(s * v8.X),
				Y: float32(s * v8.Y),
				Z: float32(s * v8.Z),
			}
			v9 = u
			goto _22
		}
		v9 = Vec3{}
		goto _22
	_22:
		direction = v9
		q1 = b3ComputeQuatBetweenUnitVectors(tls, b3Vec3_axisY, direction)
		v23 = q1
		xx = float32(v23.V.X * v23.V.X)
		yy = float32(v23.V.Y * v23.V.Y)
		zz = float32(v23.V.Z * v23.V.Z)
		xy = float32(v23.V.X * v23.V.Y)
		xz = float32(v23.V.X * v23.V.Z)
		xw = float32(v23.V.X * v23.S)
		yz = float32(v23.V.Y * v23.V.Z)
		yw = float32(v23.V.Y * v23.S)
		zw = float32(v23.V.Z * v23.S)
		v12 = Matrix3{
			Cx: Vec3{
				X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
				Y: float32(float32FromFloat32(2) * (xy + zw)),
				Z: float32(float32FromFloat32(2) * (xz - yw)),
			},
			Cy: Vec3{
				X: float32(float32FromFloat32(2) * (xy - zw)),
				Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
				Z: float32(float32FromFloat32(2) * (yz + xw)),
			},
			Cz: Vec3{
				X: float32(float32FromFloat32(2) * (xz + yw)),
				Y: float32(float32FromFloat32(2) * (yz - xw)),
				Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
			},
		}
		goto _25
	_25:
		rotation = v12
	}
	mass = sphereMass + cylinderMass
	v1 = c1
	v2 = c2
	v5 = Vec3{
		X: v1.X + v2.X,
		Y: v1.Y + v2.Y,
		Z: v1.Z + v2.Z,
	}
	goto _29
_29:
	v3 = float32FromFloat32(0.5)
	v8 = v5
	v9 = Vec3{
		X: float32(v3 * v8.X),
		Y: float32(v3 * v8.Y),
		Z: float32(v3 * v8.Z),
	}
	goto _33
_33:
	center = v9
	out2.Mass = mass
	out2.Center = center
	// Rotate the central inertia into the shape frame
	v12 = rotation
	out1.Cx = Vec3{
		X: v12.Cx.X,
		Y: v12.Cy.X,
		Z: v12.Cz.X,
	}
	out1.Cy = Vec3{
		X: v12.Cx.Y,
		Y: v12.Cy.Y,
		Z: v12.Cz.Y,
	}
	out1.Cz = Vec3{
		X: v12.Cx.Z,
		Y: v12.Cy.Z,
		Z: v12.Cz.Z,
	}
	v13 = out1
	goto _36
_36:
	v14 = inertia
	v24 = v13
	v34 = v14
	v1 = v24.Cx
	b4 = Vec3{
		X: float32(v34.Cx.X*v1.X) + float32(v34.Cy.X*v1.Y) + float32(v34.Cz.X*v1.Z),
		Y: float32(v34.Cx.Y*v1.X) + float32(v34.Cy.Y*v1.Y) + float32(v34.Cz.Y*v1.Z),
		Z: float32(v34.Cx.Z*v1.X) + float32(v34.Cy.Z*v1.Y) + float32(v34.Cz.Z*v1.Z),
	}
	v2 = b4
	goto _42
_42:
	out.Cx = v2
	v35 = v14
	v5 = v24.Cy
	b4 = Vec3{
		X: float32(v35.Cx.X*v5.X) + float32(v35.Cy.X*v5.Y) + float32(v35.Cz.X*v5.Z),
		Y: float32(v35.Cx.Y*v5.X) + float32(v35.Cy.Y*v5.Y) + float32(v35.Cz.Y*v5.Z),
		Z: float32(v35.Cx.Z*v5.X) + float32(v35.Cy.Z*v5.Y) + float32(v35.Cz.Z*v5.Z),
	}
	v8 = b4
	goto _46
_46:
	out.Cy = v8
	v37 = v14
	v9 = v24.Cz
	b4 = Vec3{
		X: float32(v37.Cx.X*v9.X) + float32(v37.Cy.X*v9.Y) + float32(v37.Cz.X*v9.Z),
		Y: float32(v37.Cx.Y*v9.X) + float32(v37.Cy.Y*v9.Y) + float32(v37.Cz.Y*v9.Z),
		Z: float32(v37.Cx.Z*v9.X) + float32(v37.Cy.Z*v9.Y) + float32(v37.Cz.Z*v9.Z),
	}
	v16 = b4
	goto _50
_50:
	out.Cz = v16
	v38 = out
	goto _52
_52:
	v39 = rotation
	v43 = v38
	v47 = v39
	v17 = v43.Cx
	b4 = Vec3{
		X: float32(v47.Cx.X*v17.X) + float32(v47.Cy.X*v17.Y) + float32(v47.Cz.X*v17.Z),
		Y: float32(v47.Cx.Y*v17.X) + float32(v47.Cy.Y*v17.Y) + float32(v47.Cz.Y*v17.Z),
		Z: float32(v47.Cx.Z*v17.X) + float32(v47.Cy.Z*v17.Y) + float32(v47.Cz.Z*v17.Z),
	}
	v18 = b4
	goto _58
_58:
	out.Cx = v18
	v51 = v39
	v20 = v43.Cy
	b4 = Vec3{
		X: float32(v51.Cx.X*v20.X) + float32(v51.Cy.X*v20.Y) + float32(v51.Cz.X*v20.Z),
		Y: float32(v51.Cx.Y*v20.X) + float32(v51.Cy.Y*v20.Y) + float32(v51.Cz.Y*v20.Z),
		Z: float32(v51.Cx.Z*v20.X) + float32(v51.Cy.Z*v20.Y) + float32(v51.Cz.Z*v20.Z),
	}
	v21 = b4
	goto _62
_62:
	out.Cy = v21
	v53 = v39
	v26 = v43.Cz
	b4 = Vec3{
		X: float32(v53.Cx.X*v26.X) + float32(v53.Cy.X*v26.Y) + float32(v53.Cz.X*v26.Z),
		Y: float32(v53.Cx.Y*v26.X) + float32(v53.Cy.Y*v26.Y) + float32(v53.Cz.Y*v26.Z),
		Z: float32(v53.Cx.Z*v26.X) + float32(v53.Cy.Z*v26.Y) + float32(v53.Cz.Z*v26.Z),
	}
	v27 = b4
	goto _66
_66:
	out.Cz = v27
	v54 = out
	goto _68
_68:
	out2.Inertia = v54
	return out2
}

func b3ComputeCapsuleAABB(tls *_Stack, shape uintptr, transform Transform) (r1 AABB) {
	var aabb AABB
	var c, center1, center2, extent, rv, t1, t2, t3, v10, v11, v13, v14, v15, v17, v19, v21, v22, v24, v26, v27, v28, v3, v32, v33, v34, v35, v37, v39, v4, v40, v42, v43, v44, v46, v48, v5, v50, v51, v53, v55, v56, v57, v59, v6, v60, v61, v78, v79, v8, v80 Vec3
	var r, v20, v38, v49, v63, v64, v65, v67, v68, v69, v70, v72, v73, v74, v75, v77, v9 float32
	var v1, v30 Transform
	var v2, v31 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, c, center1, center2, extent, r, rv, t1, t2, t3, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v3, v30, v31, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v5, v50, v51, v53, v55, v56, v57, v59, v6, v60, v61, v63, v64, v65, v67, v68, v69, v70, v72, v73, v74, v75, v77, v78, v79, v8, v80, v9
	r = (*Capsule)(unsafe.Pointer(shape)).Radius
	v1 = transform
	v2 = v1.Q
	v3 = (*Capsule)(unsafe.Pointer(shape)).Center1
	v4 = v2.V
	v5 = v3
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _7
_7:
	t1 = v6
	v8 = t1
	v9 = v2.S
	v10 = v3
	v11 = Vec3{
		X: v8.X + float32(v9*v10.X),
		Y: v8.Y + float32(v9*v10.Y),
		Z: v8.Z + float32(v9*v10.Z),
	}
	goto _12
_12:
	t2 = v11
	v13 = v2.V
	v14 = t2
	c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
	c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
	c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	v15 = c
	goto _16
_16:
	t3 = v15
	v19 = v3
	v20 = float32FromFloat32(2)
	v21 = t3
	v22 = Vec3{
		X: v19.X + float32(v20*v21.X),
		Y: v19.Y + float32(v20*v21.Y),
		Z: v19.Z + float32(v20*v21.Z),
	}
	goto _23
_23:
	v17 = v22
	goto _18
_18:
	rv = v17
	v26 = rv
	v27 = v1.P
	v28 = Vec3{
		X: v26.X + v27.X,
		Y: v26.Y + v27.Y,
		Z: v26.Z + v27.Z,
	}
	goto _29
_29:
	v24 = v28
	goto _25
_25:
	center1 = v24
	v30 = transform
	v31 = v30.Q
	v32 = (*Capsule)(unsafe.Pointer(shape)).Center2
	v33 = v31.V
	v34 = v32
	c.X = float32(v33.Y*v34.Z) - float32(v33.Z*v34.Y)
	c.Y = float32(v33.Z*v34.X) - float32(v33.X*v34.Z)
	c.Z = float32(v33.X*v34.Y) - float32(v33.Y*v34.X)
	v35 = c
	goto _36
_36:
	t1 = v35
	v37 = t1
	v38 = v31.S
	v39 = v32
	v40 = Vec3{
		X: v37.X + float32(v38*v39.X),
		Y: v37.Y + float32(v38*v39.Y),
		Z: v37.Z + float32(v38*v39.Z),
	}
	goto _41
_41:
	t2 = v40
	v42 = v31.V
	v43 = t2
	c.X = float32(v42.Y*v43.Z) - float32(v42.Z*v43.Y)
	c.Y = float32(v42.Z*v43.X) - float32(v42.X*v43.Z)
	c.Z = float32(v42.X*v43.Y) - float32(v42.Y*v43.X)
	v44 = c
	goto _45
_45:
	t3 = v44
	v48 = v32
	v49 = float32FromFloat32(2)
	v50 = t3
	v51 = Vec3{
		X: v48.X + float32(v49*v50.X),
		Y: v48.Y + float32(v49*v50.Y),
		Z: v48.Z + float32(v49*v50.Z),
	}
	goto _52
_52:
	v46 = v51
	goto _47
_47:
	rv = v46
	v55 = rv
	v56 = v30.P
	v57 = Vec3{
		X: v55.X + v56.X,
		Y: v55.Y + v56.Y,
		Z: v55.Z + v56.Z,
	}
	goto _58
_58:
	v53 = v57
	goto _54
_54:
	center2 = v53
	extent = Vec3{
		X: r,
		Y: r,
		Z: r,
	}
	v59 = center1
	v60 = center2
	v63 = v59.X
	v64 = v60.X
	if v63 < v64 {
		v67 = v63
	} else {
		v67 = v64
	}
	v65 = v67
	goto _66
_66:
	v68 = v59.Y
	v69 = v60.Y
	if v68 < v69 {
		v72 = v68
	} else {
		v72 = v69
	}
	v70 = v72
	goto _71
_71:
	v73 = v59.Z
	v74 = v60.Z
	if v73 < v74 {
		v77 = v73
	} else {
		v77 = v74
	}
	v75 = v77
	goto _76
_76:
	v61 = Vec3{
		X: v65,
		Y: v70,
		Z: v75,
	}
	goto _62
_62:
	v78 = v61
	v79 = extent
	v80 = Vec3{
		X: v78.X - v79.X,
		Y: v78.Y - v79.Y,
		Z: v78.Z - v79.Z,
	}
	goto _81
_81:
	aabb.LowerBound = v80
	v3 = center1
	v4 = center2
	v9 = v3.X
	v20 = v4.X
	if v9 > v20 {
		v49 = v9
	} else {
		v49 = v20
	}
	v38 = v49
	goto _89
_89:
	v63 = v3.Y
	v64 = v4.Y
	if v63 > v64 {
		v67 = v63
	} else {
		v67 = v64
	}
	v65 = v67
	goto _94
_94:
	v68 = v3.Z
	v69 = v4.Z
	if v68 > v69 {
		v72 = v68
	} else {
		v72 = v69
	}
	v70 = v72
	goto _99
_99:
	v5 = Vec3{
		X: v38,
		Y: v65,
		Z: v70,
	}
	goto _85
_85:
	v6 = v5
	v8 = extent
	v10 = Vec3{
		X: v6.X + v8.X,
		Y: v6.Y + v8.Y,
		Z: v6.Z + v8.Z,
	}
	goto _104
_104:
	aabb.UpperBound = v10
	return aabb
}

func b3ComputeSweptCapsuleAABB(tls *_Stack, shape uintptr, xf1 Transform, xf2 Transform) (r1 AABB) {
	var a8, b8, c, c1, d, r, rv, t1, t2, t3, v10, v100, v101, v102, v104, v106, v108, v109, v11, v111, v113, v114, v115, v117, v118, v119, v13, v136, v137, v138, v14, v15, v155, v156, v157, v17, v174, v175, v176, v178, v179, v180, v19, v197, v198, v199, v21, v216, v217, v218, v22, v235, v236, v237, v24, v26, v27, v28, v3, v32, v33, v34, v35, v37, v39, v4, v40, v42, v43, v44, v46, v48, v5, v50, v51, v53, v55, v56, v57, v6, v61, v62, v63, v64, v66, v68, v69, v71, v72, v73, v75, v77, v79, v8, v80, v82, v84, v85, v86, v90, v91, v92, v93, v95, v97, v98 Vec3
	var aabb AABB
	var v1, v30, v59, v88 Transform
	var v107, v121, v122, v123, v125, v126, v127, v128, v130, v131, v132, v133, v135, v140, v141, v142, v144, v145, v146, v147, v149, v150, v151, v152, v154, v159, v160, v161, v163, v164, v165, v166, v168, v169, v170, v171, v173, v182, v183, v184, v186, v187, v188, v189, v191, v192, v193, v194, v196, v20, v201, v202, v203, v205, v206, v207, v208, v210, v211, v212, v213, v215, v220, v221, v222, v224, v225, v226, v227, v229, v230, v231, v232, v234, v38, v49, v67, v78, v9, v96 float32
	var v2, v31, v60, v89 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a8, aabb, b8, c, c1, d, r, rv, t1, t2, t3, v1, v10, v100, v101, v102, v104, v106, v107, v108, v109, v11, v111, v113, v114, v115, v117, v118, v119, v121, v122, v123, v125, v126, v127, v128, v13, v130, v131, v132, v133, v135, v136, v137, v138, v14, v140, v141, v142, v144, v145, v146, v147, v149, v15, v150, v151, v152, v154, v155, v156, v157, v159, v160, v161, v163, v164, v165, v166, v168, v169, v17, v170, v171, v173, v174, v175, v176, v178, v179, v180, v182, v183, v184, v186, v187, v188, v189, v19, v191, v192, v193, v194, v196, v197, v198, v199, v2, v20, v201, v202, v203, v205, v206, v207, v208, v21, v210, v211, v212, v213, v215, v216, v217, v218, v22, v220, v221, v222, v224, v225, v226, v227, v229, v230, v231, v232, v234, v235, v236, v237, v24, v26, v27, v28, v3, v30, v31, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v5, v50, v51, v53, v55, v56, v57, v59, v6, v60, v61, v62, v63, v64, v66, v67, v68, v69, v71, v72, v73, v75, v77, v78, v79, v8, v80, v82, v84, v85, v86, v88, v89, v9, v90, v91, v92, v93, v95, v96, v97, v98
	r = Vec3{
		X: (*Capsule)(unsafe.Pointer(shape)).Radius,
		Y: (*Capsule)(unsafe.Pointer(shape)).Radius,
		Z: (*Capsule)(unsafe.Pointer(shape)).Radius,
	}
	v1 = xf1
	v2 = v1.Q
	v3 = (*Capsule)(unsafe.Pointer(shape)).Center1
	v4 = v2.V
	v5 = v3
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _7
_7:
	t1 = v6
	v8 = t1
	v9 = v2.S
	v10 = v3
	v11 = Vec3{
		X: v8.X + float32(v9*v10.X),
		Y: v8.Y + float32(v9*v10.Y),
		Z: v8.Z + float32(v9*v10.Z),
	}
	goto _12
_12:
	t2 = v11
	v13 = v2.V
	v14 = t2
	c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
	c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
	c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	v15 = c
	goto _16
_16:
	t3 = v15
	v19 = v3
	v20 = float32FromFloat32(2)
	v21 = t3
	v22 = Vec3{
		X: v19.X + float32(v20*v21.X),
		Y: v19.Y + float32(v20*v21.Y),
		Z: v19.Z + float32(v20*v21.Z),
	}
	goto _23
_23:
	v17 = v22
	goto _18
_18:
	rv = v17
	v26 = rv
	v27 = v1.P
	v28 = Vec3{
		X: v26.X + v27.X,
		Y: v26.Y + v27.Y,
		Z: v26.Z + v27.Z,
	}
	goto _29
_29:
	v24 = v28
	goto _25
_25:
	a8 = v24
	v30 = xf1
	v31 = v30.Q
	v32 = (*Capsule)(unsafe.Pointer(shape)).Center2
	v33 = v31.V
	v34 = v32
	c.X = float32(v33.Y*v34.Z) - float32(v33.Z*v34.Y)
	c.Y = float32(v33.Z*v34.X) - float32(v33.X*v34.Z)
	c.Z = float32(v33.X*v34.Y) - float32(v33.Y*v34.X)
	v35 = c
	goto _36
_36:
	t1 = v35
	v37 = t1
	v38 = v31.S
	v39 = v32
	v40 = Vec3{
		X: v37.X + float32(v38*v39.X),
		Y: v37.Y + float32(v38*v39.Y),
		Z: v37.Z + float32(v38*v39.Z),
	}
	goto _41
_41:
	t2 = v40
	v42 = v31.V
	v43 = t2
	c.X = float32(v42.Y*v43.Z) - float32(v42.Z*v43.Y)
	c.Y = float32(v42.Z*v43.X) - float32(v42.X*v43.Z)
	c.Z = float32(v42.X*v43.Y) - float32(v42.Y*v43.X)
	v44 = c
	goto _45
_45:
	t3 = v44
	v48 = v32
	v49 = float32FromFloat32(2)
	v50 = t3
	v51 = Vec3{
		X: v48.X + float32(v49*v50.X),
		Y: v48.Y + float32(v49*v50.Y),
		Z: v48.Z + float32(v49*v50.Z),
	}
	goto _52
_52:
	v46 = v51
	goto _47
_47:
	rv = v46
	v55 = rv
	v56 = v30.P
	v57 = Vec3{
		X: v55.X + v56.X,
		Y: v55.Y + v56.Y,
		Z: v55.Z + v56.Z,
	}
	goto _58
_58:
	v53 = v57
	goto _54
_54:
	b8 = v53
	v59 = xf2
	v60 = v59.Q
	v61 = (*Capsule)(unsafe.Pointer(shape)).Center1
	v62 = v60.V
	v63 = v61
	c.X = float32(v62.Y*v63.Z) - float32(v62.Z*v63.Y)
	c.Y = float32(v62.Z*v63.X) - float32(v62.X*v63.Z)
	c.Z = float32(v62.X*v63.Y) - float32(v62.Y*v63.X)
	v64 = c
	goto _65
_65:
	t1 = v64
	v66 = t1
	v67 = v60.S
	v68 = v61
	v69 = Vec3{
		X: v66.X + float32(v67*v68.X),
		Y: v66.Y + float32(v67*v68.Y),
		Z: v66.Z + float32(v67*v68.Z),
	}
	goto _70
_70:
	t2 = v69
	v71 = v60.V
	v72 = t2
	c.X = float32(v71.Y*v72.Z) - float32(v71.Z*v72.Y)
	c.Y = float32(v71.Z*v72.X) - float32(v71.X*v72.Z)
	c.Z = float32(v71.X*v72.Y) - float32(v71.Y*v72.X)
	v73 = c
	goto _74
_74:
	t3 = v73
	v77 = v61
	v78 = float32FromFloat32(2)
	v79 = t3
	v80 = Vec3{
		X: v77.X + float32(v78*v79.X),
		Y: v77.Y + float32(v78*v79.Y),
		Z: v77.Z + float32(v78*v79.Z),
	}
	goto _81
_81:
	v75 = v80
	goto _76
_76:
	rv = v75
	v84 = rv
	v85 = v59.P
	v86 = Vec3{
		X: v84.X + v85.X,
		Y: v84.Y + v85.Y,
		Z: v84.Z + v85.Z,
	}
	goto _87
_87:
	v82 = v86
	goto _83
_83:
	c1 = v82
	v88 = xf2
	v89 = v88.Q
	v90 = (*Capsule)(unsafe.Pointer(shape)).Center2
	v91 = v89.V
	v92 = v90
	c.X = float32(v91.Y*v92.Z) - float32(v91.Z*v92.Y)
	c.Y = float32(v91.Z*v92.X) - float32(v91.X*v92.Z)
	c.Z = float32(v91.X*v92.Y) - float32(v91.Y*v92.X)
	v93 = c
	goto _94
_94:
	t1 = v93
	v95 = t1
	v96 = v89.S
	v97 = v90
	v98 = Vec3{
		X: v95.X + float32(v96*v97.X),
		Y: v95.Y + float32(v96*v97.Y),
		Z: v95.Z + float32(v96*v97.Z),
	}
	goto _99
_99:
	t2 = v98
	v100 = v89.V
	v101 = t2
	c.X = float32(v100.Y*v101.Z) - float32(v100.Z*v101.Y)
	c.Y = float32(v100.Z*v101.X) - float32(v100.X*v101.Z)
	c.Z = float32(v100.X*v101.Y) - float32(v100.Y*v101.X)
	v102 = c
	goto _103
_103:
	t3 = v102
	v106 = v90
	v107 = float32FromFloat32(2)
	v108 = t3
	v109 = Vec3{
		X: v106.X + float32(v107*v108.X),
		Y: v106.Y + float32(v107*v108.Y),
		Z: v106.Z + float32(v107*v108.Z),
	}
	goto _110
_110:
	v104 = v109
	goto _105
_105:
	rv = v104
	v113 = rv
	v114 = v88.P
	v115 = Vec3{
		X: v113.X + v114.X,
		Y: v113.Y + v114.Y,
		Z: v113.Z + v114.Z,
	}
	goto _116
_116:
	v111 = v115
	goto _112
_112:
	d = v111
	v117 = a8
	v118 = b8
	v121 = v117.X
	v122 = v118.X
	if v121 < v122 {
		v125 = v121
	} else {
		v125 = v122
	}
	v123 = v125
	goto _124
_124:
	v126 = v117.Y
	v127 = v118.Y
	if v126 < v127 {
		v130 = v126
	} else {
		v130 = v127
	}
	v128 = v130
	goto _129
_129:
	v131 = v117.Z
	v132 = v118.Z
	if v131 < v132 {
		v135 = v131
	} else {
		v135 = v132
	}
	v133 = v135
	goto _134
_134:
	v119 = Vec3{
		X: v123,
		Y: v128,
		Z: v133,
	}
	goto _120
_120:
	v136 = c1
	v137 = d
	v140 = v136.X
	v141 = v137.X
	if v140 < v141 {
		v144 = v140
	} else {
		v144 = v141
	}
	v142 = v144
	goto _143
_143:
	v145 = v136.Y
	v146 = v137.Y
	if v145 < v146 {
		v149 = v145
	} else {
		v149 = v146
	}
	v147 = v149
	goto _148
_148:
	v150 = v136.Z
	v151 = v137.Z
	if v150 < v151 {
		v154 = v150
	} else {
		v154 = v151
	}
	v152 = v154
	goto _153
_153:
	v138 = Vec3{
		X: v142,
		Y: v147,
		Z: v152,
	}
	goto _139
_139:
	v155 = v119
	v156 = v138
	v159 = v155.X
	v160 = v156.X
	if v159 < v160 {
		v163 = v159
	} else {
		v163 = v160
	}
	v161 = v163
	goto _162
_162:
	v164 = v155.Y
	v165 = v156.Y
	if v164 < v165 {
		v168 = v164
	} else {
		v168 = v165
	}
	v166 = v168
	goto _167
_167:
	v169 = v155.Z
	v170 = v156.Z
	if v169 < v170 {
		v173 = v169
	} else {
		v173 = v170
	}
	v171 = v173
	goto _172
_172:
	v157 = Vec3{
		X: v161,
		Y: v166,
		Z: v171,
	}
	goto _158
_158:
	v174 = v157
	v175 = r
	v176 = Vec3{
		X: v174.X - v175.X,
		Y: v174.Y - v175.Y,
		Z: v174.Z - v175.Z,
	}
	goto _177
_177:
	v178 = a8
	v179 = b8
	v182 = v178.X
	v183 = v179.X
	if v182 > v183 {
		v186 = v182
	} else {
		v186 = v183
	}
	v184 = v186
	goto _185
_185:
	v187 = v178.Y
	v188 = v179.Y
	if v187 > v188 {
		v191 = v187
	} else {
		v191 = v188
	}
	v189 = v191
	goto _190
_190:
	v192 = v178.Z
	v193 = v179.Z
	if v192 > v193 {
		v196 = v192
	} else {
		v196 = v193
	}
	v194 = v196
	goto _195
_195:
	v180 = Vec3{
		X: v184,
		Y: v189,
		Z: v194,
	}
	goto _181
_181:
	v197 = c1
	v198 = d
	v201 = v197.X
	v202 = v198.X
	if v201 > v202 {
		v205 = v201
	} else {
		v205 = v202
	}
	v203 = v205
	goto _204
_204:
	v206 = v197.Y
	v207 = v198.Y
	if v206 > v207 {
		v210 = v206
	} else {
		v210 = v207
	}
	v208 = v210
	goto _209
_209:
	v211 = v197.Z
	v212 = v198.Z
	if v211 > v212 {
		v215 = v211
	} else {
		v215 = v212
	}
	v213 = v215
	goto _214
_214:
	v199 = Vec3{
		X: v203,
		Y: v208,
		Z: v213,
	}
	goto _200
_200:
	v216 = v180
	v217 = v199
	v220 = v216.X
	v221 = v217.X
	if v220 > v221 {
		v224 = v220
	} else {
		v224 = v221
	}
	v222 = v224
	goto _223
_223:
	v225 = v216.Y
	v226 = v217.Y
	if v225 > v226 {
		v229 = v225
	} else {
		v229 = v226
	}
	v227 = v229
	goto _228
_228:
	v230 = v216.Z
	v231 = v217.Z
	if v230 > v231 {
		v234 = v230
	} else {
		v234 = v231
	}
	v232 = v234
	goto _233
_233:
	v218 = Vec3{
		X: v222,
		Y: v227,
		Z: v232,
	}
	goto _219
_219:
	v235 = v218
	v236 = r
	v237 = Vec3{
		X: v235.X + v236.X,
		Y: v235.Y + v236.Y,
		Z: v235.Z + v236.Z,
	}
	goto _238
_238:
	aabb = AABB{
		LowerBound: v176,
		UpperBound: v237,
	}
	return aabb
}

func b3ComputeCompoundAABB(tls *_Stack, shape uintptr, transform1 Transform) (r AABB) {
	var b4, c, center, extent, rv, t1, t2, t3, v102, v103, v106, v107, v109, v11, v110, v111, v113, v114, v115, v12, v16, v17, v18, v19, v21, v23, v24, v26, v27, v28, v30, v32, v34, v35, v37, v39, v4, v40, v41, v47, v48, v6, v62, v63, v7, v77, v78, v8, v95, v97, v98, v99 Vec3
	var m2, out, v105, v44, v46, v92 Matrix3
	var out1, v117, v2, v3, v94 AABB
	var xw, xx, xy, xz, yw, yy, yz, zw, zz, v10, v101, v22, v33, v50, v51, v53, v54, v55, v57, v58, v59, v61, v65, v66, v68, v69, v70, v72, v73, v74, v76, v80, v81, v83, v84, v85, v87, v88, v89, v91 float32
	var v1, v14 Transform
	var v15, v43 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b4, c, center, extent, m2, out, out1, rv, t1, t2, t3, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v101, v102, v103, v105, v106, v107, v109, v11, v110, v111, v113, v114, v115, v117, v12, v14, v15, v16, v17, v18, v19, v2, v21, v22, v23, v24, v26, v27, v28, v3, v30, v32, v33, v34, v35, v37, v39, v4, v40, v41, v43, v44, v46, v47, v48, v50, v51, v53, v54, v55, v57, v58, v59, v6, v61, v62, v63, v65, v66, v68, v69, v7, v70, v72, v73, v74, v76, v77, v78, v8, v80, v81, v83, v84, v85, v87, v88, v89, v91, v92, v94, v95, v97, v98, v99
	v1 = transform1
	v2 = (*CompoundData)(unsafe.Pointer(shape)).Bounds
	v3 = v2
	v6 = v3.UpperBound
	v7 = v3.LowerBound
	v8 = Vec3{
		X: v6.X + v7.X,
		Y: v6.Y + v7.Y,
		Z: v6.Z + v7.Z,
	}
	goto _9
_9:
	v10 = float32FromFloat32(0.5)
	v11 = v8
	v12 = Vec3{
		X: float32(v10 * v11.X),
		Y: float32(v10 * v11.Y),
		Z: float32(v10 * v11.Z),
	}
	goto _13
_13:
	v4 = v12
	goto _5
_5:
	v14 = v1
	v15 = v14.Q
	v16 = v4
	v17 = v15.V
	v18 = v16
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v19 = c
	goto _20
_20:
	t1 = v19
	v21 = t1
	v22 = v15.S
	v23 = v16
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	t2 = v24
	v26 = v15.V
	v27 = t2
	c.X = float32(v26.Y*v27.Z) - float32(v26.Z*v27.Y)
	c.Y = float32(v26.Z*v27.X) - float32(v26.X*v27.Z)
	c.Z = float32(v26.X*v27.Y) - float32(v26.Y*v27.X)
	v28 = c
	goto _29
_29:
	t3 = v28
	v32 = v16
	v33 = float32FromFloat32(2)
	v34 = t3
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	v30 = v35
	goto _31
_31:
	rv = v30
	v39 = rv
	v40 = v14.P
	v41 = Vec3{
		X: v39.X + v40.X,
		Y: v39.Y + v40.Y,
		Z: v39.Z + v40.Z,
	}
	goto _42
_42:
	v37 = v41
	goto _38
_38:
	center = v37
	v43 = v1.Q
	xx = float32(v43.V.X * v43.V.X)
	yy = float32(v43.V.Y * v43.V.Y)
	zz = float32(v43.V.Z * v43.V.Z)
	xy = float32(v43.V.X * v43.V.Y)
	xz = float32(v43.V.X * v43.V.Z)
	xw = float32(v43.V.X * v43.S)
	yz = float32(v43.V.Y * v43.V.Z)
	yw = float32(v43.V.Y * v43.S)
	zw = float32(v43.V.Z * v43.S)
	v44 = Matrix3{
		Cx: Vec3{
			X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
			Y: float32(float32FromFloat32(2) * (xy + zw)),
			Z: float32(float32FromFloat32(2) * (xz - yw)),
		},
		Cy: Vec3{
			X: float32(float32FromFloat32(2) * (xy - zw)),
			Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
			Z: float32(float32FromFloat32(2) * (yz + xw)),
		},
		Cz: Vec3{
			X: float32(float32FromFloat32(2) * (xz + yw)),
			Y: float32(float32FromFloat32(2) * (yz - xw)),
			Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
		},
	}
	goto _45
_45:
	m2 = v44
	v46 = m2
	v47 = v46.Cx
	v50 = v47.X
	if v50 < float32FromInt32(0) {
		v53 = -v50
	} else {
		v53 = v50
	}
	v51 = v53
	goto _52
_52:
	v54 = v47.Y
	if v54 < float32FromInt32(0) {
		v57 = -v54
	} else {
		v57 = v54
	}
	v55 = v57
	goto _56
_56:
	v58 = v47.Z
	if v58 < float32FromInt32(0) {
		v61 = -v58
	} else {
		v61 = v58
	}
	v59 = v61
	goto _60
_60:
	v48 = Vec3{
		X: v51,
		Y: v55,
		Z: v59,
	}
	goto _49
_49:
	out.Cx = v48
	v62 = v46.Cy
	v65 = v62.X
	if v65 < float32FromInt32(0) {
		v68 = -v65
	} else {
		v68 = v65
	}
	v66 = v68
	goto _67
_67:
	v69 = v62.Y
	if v69 < float32FromInt32(0) {
		v72 = -v69
	} else {
		v72 = v69
	}
	v70 = v72
	goto _71
_71:
	v73 = v62.Z
	if v73 < float32FromInt32(0) {
		v76 = -v73
	} else {
		v76 = v73
	}
	v74 = v76
	goto _75
_75:
	v63 = Vec3{
		X: v66,
		Y: v70,
		Z: v74,
	}
	goto _64
_64:
	out.Cy = v63
	v77 = v46.Cz
	v80 = v77.X
	if v80 < float32FromInt32(0) {
		v83 = -v80
	} else {
		v83 = v80
	}
	v81 = v83
	goto _82
_82:
	v84 = v77.Y
	if v84 < float32FromInt32(0) {
		v87 = -v84
	} else {
		v87 = v84
	}
	v85 = v87
	goto _86
_86:
	v88 = v77.Z
	if v88 < float32FromInt32(0) {
		v91 = -v88
	} else {
		v91 = v88
	}
	v89 = v91
	goto _90
_90:
	v78 = Vec3{
		X: v81,
		Y: v85,
		Z: v89,
	}
	goto _79
_79:
	out.Cz = v78
	v92 = out
	goto _93
_93:
	v94 = v2
	v97 = v94.UpperBound
	v98 = v94.LowerBound
	v99 = Vec3{
		X: v97.X - v98.X,
		Y: v97.Y - v98.Y,
		Z: v97.Z - v98.Z,
	}
	goto _100
_100:
	v101 = float32FromFloat32(0.5)
	v102 = v99
	v103 = Vec3{
		X: float32(v101 * v102.X),
		Y: float32(v101 * v102.Y),
		Z: float32(v101 * v102.Z),
	}
	goto _104
_104:
	v95 = v103
	goto _96
_96:
	v105 = v92
	v106 = v95
	b4 = Vec3{
		X: float32(v105.Cx.X*v106.X) + float32(v105.Cy.X*v106.Y) + float32(v105.Cz.X*v106.Z),
		Y: float32(v105.Cx.Y*v106.X) + float32(v105.Cy.Y*v106.Y) + float32(v105.Cz.Y*v106.Z),
		Z: float32(v105.Cx.Z*v106.X) + float32(v105.Cy.Z*v106.Y) + float32(v105.Cz.Z*v106.Z),
	}
	v107 = b4
	goto _108
_108:
	extent = v107
	v109 = center
	v110 = extent
	v111 = Vec3{
		X: v109.X - v110.X,
		Y: v109.Y - v110.Y,
		Z: v109.Z - v110.Z,
	}
	goto _112
_112:
	v113 = center
	v114 = extent
	v115 = Vec3{
		X: v113.X + v114.X,
		Y: v113.Y + v114.Y,
		Z: v113.Z + v114.Z,
	}
	goto _116
_116:
	out1 = AABB{
		LowerBound: v111,
		UpperBound: v115,
	}
	v117 = out1
	goto _118
_118:
	return v117
}

func b3ComputeConvexManifold(tls *_Stack, world2 uintptr, workerIndex int32, contact uintptr, shapeA uintptr, xfA WorldTransform, shapeB uintptr, xfB WorldTransform, _arena b3Arena) (r uint8) {
	bp := tls.Alloc(320)
	defer tls.Free(320)
	*(*b3Arena)(unsafe.Pointer(bp)) = _arena
	var C, transformBtoA, v51 Transform
	var aligned, currentCount, i, i1, i2, index, index1, j, newCapacity, newSize, oldCount, oldSize, pointCapacity, v2, v58, v62 int32
	var allocator b3BlockAllocator
	var allocator1, allocator2, cache, manifold, manifolds, pointBuffer, pt1, pt2, source, target, v1, v3, v57 uintptr
	var b6, c, d, t1, t11, t2, t21, t3, t31, v10, v11, v13, v15, v16, v18, v20, v21, v23, v24, v30, v31, v32, v33, v35, v37, v38, v40, v41, v42, v44, v46, v48, v49, v9 Vec3
	var matrixA, v72, v74 Matrix3
	var q1, v27, v29, v7, v8 Quat
	var typeA, typeB b3ShapeType1
	var xw, xx, xy, xz, yw, yy, yz, zw, zz, v14, v19, v25, v36, v47 float32
	var v5, v6 WorldTransform
	var v53 bool
	var v83, v84 Pos
	var v91 FeaturePair
	var v92 uint32_t
	var _ /* geomManifold at bp+24 */ LocalManifold
	var _ /* oldPoints at bp+88 */ [4]ManifoldPoint
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, aligned, allocator, allocator1, allocator2, b6, c, cache, currentCount, d, i, i1, i2, index, index1, j, manifold, manifolds, matrixA, newCapacity, newSize, oldCount, oldSize, pointBuffer, pointCapacity, pt1, pt2, q1, source, t1, t11, t2, t21, t3, t31, target, transformBtoA, typeA, typeB, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v11, v13, v14, v15, v16, v18, v19, v2, v20, v21, v23, v24, v25, v27, v29, v3, v30, v31, v32, v33, v35, v36, v37, v38, v40, v41, v42, v44, v46, v47, v48, v49, v5, v51, v53, v57, v58, v6, v62, v7, v72, v74, v8, v83, v84, v9, v91, v92
	typeA = (*b3Shape)(unsafe.Pointer(shapeA)).Type1
	typeB = (*b3Shape)(unsafe.Pointer(shapeB)).Type1
	cache = contact + 152
	pointCapacity = int32(_B3_MAX_MANIFOLD_POINTS)
	v1 = bp
	v2 = int32FromUint64(uint64FromInt32(pointCapacity) * uint64(24))
	if v2 == int32FromInt32(0) {
		v3 = uintptrFromInt32(0)
		goto _4
	}
	aligned = ((*b3Arena)(unsafe.Pointer(v1)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
	if aligned+v2 > (*b3Arena)(unsafe.Pointer(v1)).Capacity {
		v3 = b3ArenaOverflowAlloc(tls, v1, v2)
		goto _4
	}
	(*b3Arena)(unsafe.Pointer(v1)).Index = aligned + v2
	if (*b3Arena)(unsafe.Pointer(v1)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v1)).Shared)).MaxIndex {
		(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v1)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v1)).Index
	}
	v3 = (*b3Arena)(unsafe.Pointer(v1)).Memory + uintptr(aligned)
	goto _4
_4:
	pointBuffer = v3
	*(*LocalManifold)(unsafe.Add(nil, bp+24)) = LocalManifold{}
	(*(*LocalManifold)(unsafe.Add(nil, bp+24))).points = pointBuffer
	v5 = xfA
	v6 = xfB
	v7 = v5.Q
	v8 = v6.Q
	v9 = v8.V
	v10 = v7.V
	c.X = float32(v9.Y*v10.Z) - float32(v9.Z*v10.Y)
	c.Y = float32(v9.Z*v10.X) - float32(v9.X*v10.Z)
	c.Z = float32(v9.X*v10.Y) - float32(v9.Y*v10.X)
	v11 = c
	goto _12
_12:
	t11 = v11
	v13 = t11
	v14 = v7.S
	v15 = v8.V
	v16 = Vec3{
		X: v13.X + float32(v14*v15.X),
		Y: v13.Y + float32(v14*v15.Y),
		Z: v13.Z + float32(v14*v15.Z),
	}
	goto _17
_17:
	t21 = v16
	v18 = t21
	v19 = v8.S
	v20 = v7.V
	v21 = Vec3{
		X: v18.X - float32(v19*v20.X),
		Y: v18.Y - float32(v19*v20.Y),
		Z: v18.Z - float32(v19*v20.Z),
	}
	goto _22
_22:
	t31 = v21
	v23 = v7.V
	v24 = v8.V
	v25 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y) + float32(v23.Z*v24.Z)
	goto _26
_26:
	q1 = Quat{
		V: t31,
		S: float32(v7.S*v8.S) + v25,
	}
	v27 = q1
	goto _28
_28:
	C.Q = v27
	d = Vec3{
		X: v6.P.X - v5.P.X,
		Y: v6.P.Y - v5.P.Y,
		Z: v6.P.Z - v5.P.Z,
	}
	v29 = v5.Q
	v30 = d
	v31 = v29.V
	v32 = v30
	c.X = float32(v31.Y*v32.Z) - float32(v31.Z*v32.Y)
	c.Y = float32(v31.Z*v32.X) - float32(v31.X*v32.Z)
	c.Z = float32(v31.X*v32.Y) - float32(v31.Y*v32.X)
	v33 = c
	goto _34
_34:
	t1 = v33
	v35 = t1
	v36 = v29.S
	v37 = v30
	v38 = Vec3{
		X: v35.X - float32(v36*v37.X),
		Y: v35.Y - float32(v36*v37.Y),
		Z: v35.Z - float32(v36*v37.Z),
	}
	goto _39
_39:
	t2 = v38
	v40 = v29.V
	v41 = t2
	c.X = float32(v40.Y*v41.Z) - float32(v40.Z*v41.Y)
	c.Y = float32(v40.Z*v41.X) - float32(v40.X*v41.Z)
	c.Z = float32(v40.X*v41.Y) - float32(v40.Y*v41.X)
	v42 = c
	goto _43
_43:
	t3 = v42
	v46 = v30
	v47 = float32FromFloat32(2)
	v48 = t3
	v49 = Vec3{
		X: v46.X + float32(v47*v48.X),
		Y: v46.Y + float32(v47*v48.Y),
		Z: v46.Z + float32(v47*v48.Z),
	}
	goto _50
_50:
	v44 = v49
	goto _45
_45:
	C.P = v44
	v51 = C
	goto _52
_52:
	transformBtoA = v51
	if typeA == int32(b3_sphereShape) {
		if v53 = !!(typeB == int32(b3_sphereShape)); !v53 {
			b3InternalAssert(tls, __ccgo_ts+8858, __ccgo_ts+7283, int32FromInt32(500))
		}
		_ = v53 || bool(int32FromInt32(0) != 0)
		b3CollideSpheres(tls, bp+24, pointCapacity, shapeA+176, shapeB+176, transformBtoA)
	} else {
		if typeA == int32(b3_capsuleShape) {
			if typeB == int32(b3_sphereShape) {
				b3CollideCapsuleAndSphere(tls, bp+24, pointCapacity, shapeA+176, shapeB+176, transformBtoA)
			} else {
				if v53 = !!(typeB == int32(b3_capsuleShape)); !v53 {
					b3InternalAssert(tls, __ccgo_ts+8882, __ccgo_ts+7283, int32FromInt32(511))
				}
				_ = v53 || bool(int32FromInt32(0) != 0)
				b3CollideCapsules(tls, bp+24, pointCapacity, shapeA+176, shapeB+176, transformBtoA)
			}
		} else {
			if v53 = !!(typeA == int32(b3_hullShape)); !v53 {
				b3InternalAssert(tls, __ccgo_ts+8907, __ccgo_ts+7283, int32FromInt32(517))
			}
			_ = v53 || bool(int32FromInt32(0) != 0)
			if typeB == int32(b3_sphereShape) {
				b3CollideHullAndSphere(tls, bp+24, pointCapacity, *(*uintptr)(unsafe.Add(unsafe.Pointer(shapeA), 176)), shapeB+176, transformBtoA, cache)
			} else {
				if typeB == int32(b3_capsuleShape) {
					b3CollideHullAndCapsule(tls, bp+24, pointCapacity, *(*uintptr)(unsafe.Add(unsafe.Pointer(shapeA), 176)), shapeB+176, transformBtoA, cache)
				} else {
					if v53 = !!(typeB == int32(b3_hullShape)); !v53 {
						b3InternalAssert(tls, __ccgo_ts+8929, __ccgo_ts+7283, int32FromInt32(531))
					}
					_ = v53 || bool(int32FromInt32(0) != 0)
					b3CollideHullsAtWidth(tls, bp+24, pointCapacity, *(*uintptr)(unsafe.Add(unsafe.Pointer(shapeA), 176)), *(*uintptr)(unsafe.Add(unsafe.Pointer(shapeB), 176)), transformBtoA, cache, (*b3World)(unsafe.Pointer(world2)).SimdWidth)
					(*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world2)).TaskContexts.Data+uintptr(workerIndex)*3848))).SatCallCount += int32(1)
					(*(*b3TaskContext)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(world2)).TaskContexts.Data+uintptr(workerIndex)*3848))).SatCacheHitCount += int32FromUint8((*b3ContactCache)(unsafe.Pointer(cache)).SatCache.Hit)
				}
			}
		}
	}
	if (*(*LocalManifold)(unsafe.Add(nil, bp+24))).PointCount == 0 {
		if (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount > 0 {
			v1 = world2
			v2 = (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount
			if v2 == int32FromInt32(0) {
				goto _59
			}
			index1 = v2 - int32(1)
			__builtin_b3LockMutex(tls, (*b3World)(unsafe.Pointer(v1)).ManifoldAllocatorMutex)
			if v53 = !!(0 <= index1 && index1 < (*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Count); !v53 {
				b3InternalAssert(tls, __ccgo_ts+8145, __ccgo_ts+8211, int32FromInt32(359))
			}
			_ = v53 || bool(int32FromInt32(0) != 0)
			allocator2 = (*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Data + uintptr(index1)*40
			b3FreeElement(tls, allocator2, (*b3Contact)(unsafe.Pointer(contact)).Manifolds)
			__builtin_b3UnlockMutex(tls, (*b3World)(unsafe.Pointer(v1)).ManifoldAllocatorMutex)
		_59:
			;
			(*b3Contact)(unsafe.Pointer(contact)).Manifolds = uintptrFromInt32(0)
			(*b3Contact)(unsafe.Pointer(contact)).ManifoldCount = 0
		}
		return boolUint8(false1 != 0)
	}
	oldCount = 0
	if (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount == 0 {
		v1 = world2
		v2 = int32(1)
		if v2 == int32FromInt32(0) {
			v3 = uintptrFromInt32(0)
			goto _64
		}
		index = v2 - int32(1)
		__builtin_b3LockMutex(tls, (*b3World)(unsafe.Pointer(v1)).ManifoldAllocatorMutex)
		currentCount = (*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Count
		i = currentCount
		for {
			if !(i < v2) {
				break
			}
			allocator = b3CreateBlockAllocator(tls, int32FromUint64(uint64FromInt32(i+int32FromInt32(1))*uint64(268)), int32FromInt32(2)*(int32FromInt32(1)<<int32FromInt32(_B3_BLOCK_EXPONENT)))
			if (*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Count >= (*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Capacity {
				oldSize = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Capacity) * uint64(40))
				if (*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Capacity == 0 {
					v58 = int32(8)
				} else {
					v58 = int32(2) * (*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Capacity
				}
				newCapacity = v58
				newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(40))
				(*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Data, oldSize, newSize)
				(*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Capacity = newCapacity
			}
			v57 = v1 + 4040 + 8
			v62 = *(*int32)(unsafe.Pointer(v57))
			*(*int32)(unsafe.Pointer(v57)) = *(*int32)(unsafe.Pointer(v57)) + 1
			*(*b3BlockAllocator)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Data+uintptr(v62)*40)) = allocator
			goto _65
		_65:
			;
			i = i + 1
		}
		if v53 = !!(0 <= index && index < (*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Count); !v53 {
			b3InternalAssert(tls, __ccgo_ts+8145, __ccgo_ts+8211, int32FromInt32(343))
		}
		_ = v53 || bool(int32FromInt32(0) != 0)
		allocator1 = (*b3World)(unsafe.Pointer(v1)).ManifoldAllocators.Data + uintptr(index)*40
		manifolds = b3AllocateElement(tls, allocator1)
		__builtin_b3UnlockMutex(tls, (*b3World)(unsafe.Pointer(v1)).ManifoldAllocatorMutex)
		memset(tls, manifolds, 0, uint64FromInt32(v2)*uint64(268))
		v3 = manifolds
		goto _64
	_64:
		(*b3Contact)(unsafe.Pointer(contact)).Manifolds = v3
		(*b3Contact)(unsafe.Pointer(contact)).ManifoldCount = int32(1)
	} else {
		oldCount = (*(*Manifold)(unsafe.Add(nil, (*b3Contact)(unsafe.Pointer(contact)).Manifolds))).PointCount
		memcpy(tls, bp+88, (*b3Contact)(unsafe.Pointer(contact)).Manifolds, uint64FromInt32(oldCount)*uint64(56))
	}
	manifold = (*b3Contact)(unsafe.Pointer(contact)).Manifolds
	if v53 = !!(0 < (*(*LocalManifold)(unsafe.Add(nil, bp+24))).PointCount && (*(*LocalManifold)(unsafe.Add(nil, bp+24))).PointCount <= int32(_B3_MAX_MANIFOLD_POINTS)); !v53 {
		b3InternalAssert(tls, __ccgo_ts+8951, __ccgo_ts+7283, int32FromInt32(566))
	}
	_ = v53 || bool(int32FromInt32(0) != 0)
	(*Manifold)(unsafe.Pointer(manifold)).PointCount = (*(*LocalManifold)(unsafe.Add(nil, bp+24))).PointCount
	v7 = xfA.Q
	xx = float32(v7.V.X * v7.V.X)
	yy = float32(v7.V.Y * v7.V.Y)
	zz = float32(v7.V.Z * v7.V.Z)
	xy = float32(v7.V.X * v7.V.Y)
	xz = float32(v7.V.X * v7.V.Z)
	xw = float32(v7.V.X * v7.S)
	yz = float32(v7.V.Y * v7.V.Z)
	yw = float32(v7.V.Y * v7.S)
	zw = float32(v7.V.Z * v7.S)
	v72 = Matrix3{
		Cx: Vec3{
			X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
			Y: float32(float32FromFloat32(2) * (xy + zw)),
			Z: float32(float32FromFloat32(2) * (xz - yw)),
		},
		Cy: Vec3{
			X: float32(float32FromFloat32(2) * (xy - zw)),
			Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
			Z: float32(float32FromFloat32(2) * (yz + xw)),
		},
		Cz: Vec3{
			X: float32(float32FromFloat32(2) * (xz + yw)),
			Y: float32(float32FromFloat32(2) * (yz - xw)),
			Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
		},
	}
	goto _73
_73:
	matrixA = v72
	v74 = matrixA
	v9 = (*(*LocalManifold)(unsafe.Add(nil, bp+24))).Normal
	b6 = Vec3{
		X: float32(v74.Cx.X*v9.X) + float32(v74.Cy.X*v9.Y) + float32(v74.Cz.X*v9.Z),
		Y: float32(v74.Cx.Y*v9.X) + float32(v74.Cy.Y*v9.Y) + float32(v74.Cz.Y*v9.Z),
		Z: float32(v74.Cx.Z*v9.X) + float32(v74.Cy.Z*v9.Y) + float32(v74.Cz.Z*v9.Z),
	}
	v10 = b6
	goto _77
_77:
	(*Manifold)(unsafe.Pointer(manifold)).Normal = v10
	// Store point data in contact
	i1 = 0
	for {
		if !(i1 < (*(*LocalManifold)(unsafe.Add(nil, bp+24))).PointCount) {
			break
		}
		source = (*(*LocalManifold)(unsafe.Add(nil, bp+24))).points + uintptr(i1)*24
		target = manifold + uintptr(i1)*56
		// Contact points are computed in frame A
		v72 = matrixA
		v9 = (*LocalManifoldPoint)(unsafe.Pointer(source)).Point
		b6 = Vec3{
			X: float32(v72.Cx.X*v9.X) + float32(v72.Cy.X*v9.Y) + float32(v72.Cz.X*v9.Z),
			Y: float32(v72.Cx.Y*v9.X) + float32(v72.Cy.Y*v9.Y) + float32(v72.Cz.Y*v9.Z),
			Z: float32(v72.Cx.Z*v9.X) + float32(v72.Cy.Z*v9.Y) + float32(v72.Cz.Z*v9.Z),
		}
		v10 = b6
		goto _82
	_82:
		(*ManifoldPoint)(unsafe.Pointer(target)).AnchorA = v10
		v83 = xfA.P
		v84 = xfB.P
		v9 = Vec3{
			X: v83.X - v84.X,
			Y: v83.Y - v84.Y,
			Z: v83.Z - v84.Z,
		}
		goto _86
	_86:
		v10 = (*ManifoldPoint)(unsafe.Pointer(target)).AnchorA
		v11 = v9
		v13 = Vec3{
			X: v10.X + v11.X,
			Y: v10.Y + v11.Y,
			Z: v10.Z + v11.Z,
		}
		goto _90
	_90:
		(*ManifoldPoint)(unsafe.Pointer(target)).AnchorB = v13
		(*ManifoldPoint)(unsafe.Pointer(target)).Separation = (*LocalManifoldPoint)(unsafe.Pointer(source)).Separation
		v91 = (*LocalManifoldPoint)(unsafe.Pointer(source)).Pair
		v92 = uint32(v91.Owner1)<<int32FromInt32(24) | uint32(v91.Index1)<<int32FromInt32(16) | uint32(v91.Owner2)<<int32FromInt32(8) | uint32(v91.Index2)
		goto _93
	_93:
		(*ManifoldPoint)(unsafe.Pointer(target)).FeatureId = v92
		(*ManifoldPoint)(unsafe.Pointer(target)).TriangleIndex = -int32(1)
		(*ManifoldPoint)(unsafe.Pointer(target)).NormalVelocity = float32FromFloat32(0)
		goto _78
	_78:
		;
		i1 = i1 + 1
	}
	// Copy impulses from old points
	i2 = 0
	for {
		if !(i2 < (*(*LocalManifold)(unsafe.Add(nil, bp+24))).PointCount) {
			break
		}
		pt2 = manifold + uintptr(i2)*56
		(*ManifoldPoint)(unsafe.Pointer(pt2)).TotalNormalImpulse = float32FromFloat32(0)
		(*ManifoldPoint)(unsafe.Pointer(pt2)).Persisted = boolUint8(false1 != 0)
		j = 0
		for {
			if !(j < oldCount) {
				break
			}
			pt1 = bp + 88 + uintptr(j)*56
			if (*ManifoldPoint)(unsafe.Pointer(pt2)).FeatureId == (*ManifoldPoint)(unsafe.Pointer(pt1)).FeatureId {
				(*ManifoldPoint)(unsafe.Pointer(pt2)).NormalImpulse = (*ManifoldPoint)(unsafe.Pointer(pt1)).NormalImpulse
				(*ManifoldPoint)(unsafe.Pointer(pt2)).Persisted = boolUint8(true1 != 0)
				// claimed
				(*ManifoldPoint)(unsafe.Pointer(pt1)).FeatureId = uint32FromUint32(0xffffffff)
				break
			}
			goto _95
		_95:
			;
			j = j + 1
		}
		if int32FromUint8((*ManifoldPoint)(unsafe.Pointer(pt2)).Persisted) == false1 {
			(*ManifoldPoint)(unsafe.Pointer(pt2)).NormalImpulse = float32FromFloat32(0)
		}
		goto _94
	_94:
		;
		i2 = i2 + 1
	}
	return boolUint8(true1 != 0)
}

func b3ComputeSeparatingAxisAtWidth(tls *_Stack, hullA uintptr, hullB uintptr, xfB Transform, earlyReturn uint8, simdWidth int32) (r b3AxisQuery) {
	_ = uint64FromInt64(4)
	return b3ComputeSeparatingAxisW4(tls, hullA, hullB, xfB, earlyReturn)
}

func b3ComputeSeparatingAxis(tls *_Stack, hullA uintptr, hullB uintptr, xfB Transform, earlyReturn uint8) (r b3AxisQuery) {
	return b3ComputeSeparatingAxisAtWidth(tls, hullA, hullB, xfB, earlyReturn, b3GetSIMDWidth(tls))
}

func b3ComputeWitnessPoints(tls *_Stack, simplex uintptr, vertexA uintptr, vertexB uintptr) {
	var count int32
	var d, d1, sum, v11, v12, v15, v17, v19, v3, v5, v6, v9 Vec3
	var vs uintptr
	var v1 bool
	var v10, v2, v4, v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = count, d, d1, sum, vs, v1, v10, v11, v12, v15, v17, v19, v2, v3, v4, v5, v6, v8, v9
	vs = simplex
	count = (*Simplex)(unsafe.Pointer(simplex)).Count
	if v1 = !!(int32(1) <= count && count <= int32(4)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+10573, __ccgo_ts+10532, int32FromInt32(678))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	switch count {
	case int32(1):
		*(*Vec3)(unsafe.Add(nil, vertexA)) = (*(*SimplexVertex)(unsafe.Add(nil, vs))).WA
		*(*Vec3)(unsafe.Add(nil, vertexB)) = (*(*SimplexVertex)(unsafe.Add(nil, vs))).WB
	case int32(2):
		v2 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).A
		v3 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).WA
		v4 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A
		v5 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).WA
		d = Vec3{
			X: float32(v2*v3.X) + float32(v4*v5.X),
			Y: float32(v2*v3.Y) + float32(v4*v5.Y),
			Z: float32(v2*v3.Z) + float32(v4*v5.Z),
		}
		v6 = d
		goto _7
	_7:
		*(*Vec3)(unsafe.Add(nil, vertexA)) = v6
		v2 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).A
		v3 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).WB
		v4 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A
		v5 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).WB
		d = Vec3{
			X: float32(v2*v3.X) + float32(v4*v5.X),
			Y: float32(v2*v3.Y) + float32(v4*v5.Y),
			Z: float32(v2*v3.Z) + float32(v4*v5.Z),
		}
		v6 = d
		goto _13
	_13:
		*(*Vec3)(unsafe.Add(nil, vertexB)) = v6
	case int32(3):
		v2 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).A
		v3 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).WA
		v4 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A
		v5 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).WA
		v8 = (*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).A
		v6 = (*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).WA
		d1 = Vec3{
			X: float32(v2*v3.X) + float32(v4*v5.X) + float32(v8*v6.X),
			Y: float32(v2*v3.Y) + float32(v4*v5.Y) + float32(v8*v6.Y),
			Z: float32(v2*v3.Z) + float32(v4*v5.Z) + float32(v8*v6.Z),
		}
		v9 = d1
		goto _21
	_21:
		*(*Vec3)(unsafe.Add(nil, vertexA)) = v9
		v2 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).A
		v3 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).WB
		v4 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A
		v5 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).WB
		v8 = (*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).A
		v6 = (*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).WB
		d1 = Vec3{
			X: float32(v2*v3.X) + float32(v4*v5.X) + float32(v8*v6.X),
			Y: float32(v2*v3.Y) + float32(v4*v5.Y) + float32(v8*v6.Y),
			Z: float32(v2*v3.Z) + float32(v4*v5.Z) + float32(v8*v6.Z),
		}
		v9 = d1
		goto _29
	_29:
		*(*Vec3)(unsafe.Add(nil, vertexB)) = v9
	case int32(4):
		v2 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).A
		v3 = (*(*SimplexVertex)(unsafe.Add(nil, vs))).WA
		v4 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).A
		v5 = (*(*SimplexVertex)(unsafe.Add(nil, vs+1*48))).WA
		d = Vec3{
			X: float32(v2*v3.X) + float32(v4*v5.X),
			Y: float32(v2*v3.Y) + float32(v4*v5.Y),
			Z: float32(v2*v3.Z) + float32(v4*v5.Z),
		}
		v6 = d
		goto _35
	_35:
		v8 = (*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).A
		v9 = (*(*SimplexVertex)(unsafe.Add(nil, vs+2*48))).WA
		v10 = (*(*SimplexVertex)(unsafe.Add(nil, vs+3*48))).A
		v11 = (*(*SimplexVertex)(unsafe.Add(nil, vs+3*48))).WA
		d = Vec3{
			X: float32(v8*v9.X) + float32(v10*v11.X),
			Y: float32(v8*v9.Y) + float32(v10*v11.Y),
			Z: float32(v8*v9.Z) + float32(v10*v11.Z),
		}
		v12 = d
		goto _41
	_41:
		v15 = v6
		v17 = v12
		v19 = Vec3{
			X: v15.X + v17.X,
			Y: v15.Y + v17.Y,
			Z: v15.Z + v17.Z,
		}
		goto _45
	_45:
		// Force identical points and *zero* distance
		sum = v19
		*(*Vec3)(unsafe.Add(nil, vertexA)) = sum
		*(*Vec3)(unsafe.Add(nil, vertexB)) = sum
	default:
		if v1 = bool(!!!(__ccgo_ts+10598 != 0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+10621, __ccgo_ts+10532, int32FromInt32(708))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		break
	}
}

func b3ComputeHeightFieldAABB(tls *_Stack, shape uintptr, transform1 Transform) (r AABB) {
	var b4, c, center, extent, rv, t1, t2, t3, v102, v103, v106, v107, v109, v11, v110, v111, v113, v114, v115, v12, v16, v17, v18, v19, v21, v23, v24, v26, v27, v28, v30, v32, v34, v35, v37, v39, v4, v40, v41, v47, v48, v6, v62, v63, v7, v77, v78, v8, v95, v97, v98, v99 Vec3
	var m2, out, v105, v44, v46, v92 Matrix3
	var out1, v117, v2, v3, v94 AABB
	var xw, xx, xy, xz, yw, yy, yz, zw, zz, v10, v101, v22, v33, v50, v51, v53, v54, v55, v57, v58, v59, v61, v65, v66, v68, v69, v70, v72, v73, v74, v76, v80, v81, v83, v84, v85, v87, v88, v89, v91 float32
	var v1, v14 Transform
	var v15, v43 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b4, c, center, extent, m2, out, out1, rv, t1, t2, t3, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v101, v102, v103, v105, v106, v107, v109, v11, v110, v111, v113, v114, v115, v117, v12, v14, v15, v16, v17, v18, v19, v2, v21, v22, v23, v24, v26, v27, v28, v3, v30, v32, v33, v34, v35, v37, v39, v4, v40, v41, v43, v44, v46, v47, v48, v50, v51, v53, v54, v55, v57, v58, v59, v6, v61, v62, v63, v65, v66, v68, v69, v7, v70, v72, v73, v74, v76, v77, v78, v8, v80, v81, v83, v84, v85, v87, v88, v89, v91, v92, v94, v95, v97, v98, v99
	v1 = transform1
	v2 = (*HeightFieldData)(unsafe.Pointer(shape)).Aabb
	v3 = v2
	v6 = v3.UpperBound
	v7 = v3.LowerBound
	v8 = Vec3{
		X: v6.X + v7.X,
		Y: v6.Y + v7.Y,
		Z: v6.Z + v7.Z,
	}
	goto _9
_9:
	v10 = float32FromFloat32(0.5)
	v11 = v8
	v12 = Vec3{
		X: float32(v10 * v11.X),
		Y: float32(v10 * v11.Y),
		Z: float32(v10 * v11.Z),
	}
	goto _13
_13:
	v4 = v12
	goto _5
_5:
	v14 = v1
	v15 = v14.Q
	v16 = v4
	v17 = v15.V
	v18 = v16
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v19 = c
	goto _20
_20:
	t1 = v19
	v21 = t1
	v22 = v15.S
	v23 = v16
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	t2 = v24
	v26 = v15.V
	v27 = t2
	c.X = float32(v26.Y*v27.Z) - float32(v26.Z*v27.Y)
	c.Y = float32(v26.Z*v27.X) - float32(v26.X*v27.Z)
	c.Z = float32(v26.X*v27.Y) - float32(v26.Y*v27.X)
	v28 = c
	goto _29
_29:
	t3 = v28
	v32 = v16
	v33 = float32FromFloat32(2)
	v34 = t3
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	v30 = v35
	goto _31
_31:
	rv = v30
	v39 = rv
	v40 = v14.P
	v41 = Vec3{
		X: v39.X + v40.X,
		Y: v39.Y + v40.Y,
		Z: v39.Z + v40.Z,
	}
	goto _42
_42:
	v37 = v41
	goto _38
_38:
	center = v37
	v43 = v1.Q
	xx = float32(v43.V.X * v43.V.X)
	yy = float32(v43.V.Y * v43.V.Y)
	zz = float32(v43.V.Z * v43.V.Z)
	xy = float32(v43.V.X * v43.V.Y)
	xz = float32(v43.V.X * v43.V.Z)
	xw = float32(v43.V.X * v43.S)
	yz = float32(v43.V.Y * v43.V.Z)
	yw = float32(v43.V.Y * v43.S)
	zw = float32(v43.V.Z * v43.S)
	v44 = Matrix3{
		Cx: Vec3{
			X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
			Y: float32(float32FromFloat32(2) * (xy + zw)),
			Z: float32(float32FromFloat32(2) * (xz - yw)),
		},
		Cy: Vec3{
			X: float32(float32FromFloat32(2) * (xy - zw)),
			Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
			Z: float32(float32FromFloat32(2) * (yz + xw)),
		},
		Cz: Vec3{
			X: float32(float32FromFloat32(2) * (xz + yw)),
			Y: float32(float32FromFloat32(2) * (yz - xw)),
			Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
		},
	}
	goto _45
_45:
	m2 = v44
	v46 = m2
	v47 = v46.Cx
	v50 = v47.X
	if v50 < float32FromInt32(0) {
		v53 = -v50
	} else {
		v53 = v50
	}
	v51 = v53
	goto _52
_52:
	v54 = v47.Y
	if v54 < float32FromInt32(0) {
		v57 = -v54
	} else {
		v57 = v54
	}
	v55 = v57
	goto _56
_56:
	v58 = v47.Z
	if v58 < float32FromInt32(0) {
		v61 = -v58
	} else {
		v61 = v58
	}
	v59 = v61
	goto _60
_60:
	v48 = Vec3{
		X: v51,
		Y: v55,
		Z: v59,
	}
	goto _49
_49:
	out.Cx = v48
	v62 = v46.Cy
	v65 = v62.X
	if v65 < float32FromInt32(0) {
		v68 = -v65
	} else {
		v68 = v65
	}
	v66 = v68
	goto _67
_67:
	v69 = v62.Y
	if v69 < float32FromInt32(0) {
		v72 = -v69
	} else {
		v72 = v69
	}
	v70 = v72
	goto _71
_71:
	v73 = v62.Z
	if v73 < float32FromInt32(0) {
		v76 = -v73
	} else {
		v76 = v73
	}
	v74 = v76
	goto _75
_75:
	v63 = Vec3{
		X: v66,
		Y: v70,
		Z: v74,
	}
	goto _64
_64:
	out.Cy = v63
	v77 = v46.Cz
	v80 = v77.X
	if v80 < float32FromInt32(0) {
		v83 = -v80
	} else {
		v83 = v80
	}
	v81 = v83
	goto _82
_82:
	v84 = v77.Y
	if v84 < float32FromInt32(0) {
		v87 = -v84
	} else {
		v87 = v84
	}
	v85 = v87
	goto _86
_86:
	v88 = v77.Z
	if v88 < float32FromInt32(0) {
		v91 = -v88
	} else {
		v91 = v88
	}
	v89 = v91
	goto _90
_90:
	v78 = Vec3{
		X: v81,
		Y: v85,
		Z: v89,
	}
	goto _79
_79:
	out.Cz = v78
	v92 = out
	goto _93
_93:
	v94 = v2
	v97 = v94.UpperBound
	v98 = v94.LowerBound
	v99 = Vec3{
		X: v97.X - v98.X,
		Y: v97.Y - v98.Y,
		Z: v97.Z - v98.Z,
	}
	goto _100
_100:
	v101 = float32FromFloat32(0.5)
	v102 = v99
	v103 = Vec3{
		X: float32(v101 * v102.X),
		Y: float32(v101 * v102.Y),
		Z: float32(v101 * v102.Z),
	}
	goto _104
_104:
	v95 = v103
	goto _96
_96:
	v105 = v92
	v106 = v95
	b4 = Vec3{
		X: float32(v105.Cx.X*v106.X) + float32(v105.Cy.X*v106.Y) + float32(v105.Cz.X*v106.Z),
		Y: float32(v105.Cx.Y*v106.X) + float32(v105.Cy.Y*v106.Y) + float32(v105.Cz.Y*v106.Z),
		Z: float32(v105.Cx.Z*v106.X) + float32(v105.Cy.Z*v106.Y) + float32(v105.Cz.Z*v106.Z),
	}
	v107 = b4
	goto _108
_108:
	extent = v107
	v109 = center
	v110 = extent
	v111 = Vec3{
		X: v109.X - v110.X,
		Y: v109.Y - v110.Y,
		Z: v109.Z - v110.Z,
	}
	goto _112
_112:
	v113 = center
	v114 = extent
	v115 = Vec3{
		X: v113.X + v114.X,
		Y: v113.Y + v114.Y,
		Z: v113.Z + v114.Z,
	}
	goto _116
_116:
	out1 = AABB{
		LowerBound: v111,
		UpperBound: v115,
	}
	v117 = out1
	goto _118
_118:
	return v117
}

func b3ComputeHullWorkSizes(tls *_Stack, pointCount int32, clampedMaxCount int32) (r b3HullWorkSizes) {
	var offset, v26 size_t
	var s b3HullWorkSizes
	var v1, v2, v3, v5 int32
	_, _, _, _, _, _, _ = offset, s, v1, v2, v26, v3, v5
	s.N = pointCount
	s.M = clampedMaxCount
	// Vertices: 4 initial hull vertices + at most one per remaining input point. No free list.
	s.VertexCapacity = pointCount + int32(4)
	// Edges and faces use free-list recycling; capacity is proportional to live hull size.
	// edgeCapacity: peak is ~twice live edges plus cone edges. Minimum 48.
	v1 = int32(48)
	v2 = int32(24)*s.M - int32(48)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _4
_4:
	s.EdgeCapacity = v3
	// faceCapacity: peak intermediate state live faces (<=2*M-4) plus full cone (<=3*M-6). Minimum 16.
	v1 = int32(16)
	v2 = int32(5)*s.M - int32(10)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _9
_9:
	s.FaceCapacity = v3
	// Horizon/cone bounded by current half-edge count. Merged faces by face count.
	v1 = int32(6)
	v2 = int32(3)*s.M - int32(6)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _14
_14:
	s.HorizonCapacity = v3
	s.ConeCapacity = s.HorizonCapacity
	v1 = int32(4)
	v2 = int32(2)*s.M - int32(4)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _19
_19:
	s.MergedFacesCapacity = v3
	// Horizon DFS depth is bounded by the number of live faces (Euler: <=2*M-4).
	v1 = int32(4)
	v2 = int32(2)*s.M - int32(4)
	if v1 > v2 {
		v5 = v1
	} else {
		v5 = v2
	}
	v3 = v5
	goto _24
_24:
	s.HorizonStackCapacity = v3
	offset = uint64(0)
	s.OffsetVertex = offset
	v26 = (offset + uint64FromInt32(s.VertexCapacity)*uint64(48) + uint64(7)) & ^uint64FromUint32(7)
	goto _27
_27:
	offset = v26
	s.OffsetEdge = offset
	v26 = (offset + uint64FromInt32(s.EdgeCapacity)*uint64(48) + uint64(7)) & ^uint64FromUint32(7)
	goto _29
_29:
	offset = v26
	s.OffsetFace = offset
	v26 = (offset + uint64FromInt32(s.FaceCapacity)*uint64(128) + uint64(7)) & ^uint64FromUint32(7)
	goto _31
_31:
	offset = v26
	s.OffsetHorizon = offset
	v26 = (offset + uint64FromInt32(s.HorizonCapacity)*uint64(8) + uint64(7)) & ^uint64FromUint32(7)
	goto _33
_33:
	offset = v26
	s.OffsetCone = offset
	v26 = (offset + uint64FromInt32(s.ConeCapacity)*uint64(8) + uint64(7)) & ^uint64FromUint32(7)
	goto _35
_35:
	offset = v26
	s.OffsetMergedFaces = offset
	v26 = (offset + uint64FromInt32(s.MergedFacesCapacity)*uint64(8) + uint64(7)) & ^uint64FromUint32(7)
	goto _37
_37:
	offset = v26
	s.OffsetHorizonStack = offset
	v26 = (offset + uint64FromInt32(s.HorizonStackCapacity)*uint64(32) + uint64(7)) & ^uint64FromUint32(7)
	goto _39
_39:
	offset = v26
	s.OffsetShiftedPoints = offset
	offset = offset + uint64FromInt32(pointCount)*uint64(12)
	s.TotalBytes = offset
	return s
}

func b3ComputeHullMass(tls *_Stack, shape uintptr, density float32) (r MassData) {
	var out MassData
	var v1 float32
	var v2, v3 Matrix3
	_, _, _, _ = out, v1, v2, v3
	out.Mass = float32(density * (*HullData)(unsafe.Pointer(shape)).Volume)
	out.Center = (*HullData)(unsafe.Pointer(shape)).Center
	// Inertia about the center of mass
	v1 = density
	v2 = (*HullData)(unsafe.Pointer(shape)).CentralInertia
	v3 = Matrix3{
		Cx: Vec3{
			X: float32(v1 * v2.Cx.X),
			Y: float32(v1 * v2.Cx.Y),
			Z: float32(v1 * v2.Cx.Z),
		},
		Cy: Vec3{
			X: float32(v1 * v2.Cy.X),
			Y: float32(v1 * v2.Cy.Y),
			Z: float32(v1 * v2.Cy.Z),
		},
		Cz: Vec3{
			X: float32(v1 * v2.Cz.X),
			Y: float32(v1 * v2.Cz.Y),
			Z: float32(v1 * v2.Cz.Z),
		},
	}
	goto _4
_4:
	out.Inertia = v3
	return out
}

func b3ComputeHullAABB(tls *_Stack, shape uintptr, transform1 Transform) (r AABB) {
	var b4, c, center, extent, rv, t1, t2, t3, v102, v103, v106, v107, v109, v11, v110, v111, v113, v114, v115, v12, v16, v17, v18, v19, v21, v23, v24, v26, v27, v28, v30, v32, v34, v35, v37, v39, v4, v40, v41, v47, v48, v6, v62, v63, v7, v77, v78, v8, v95, v97, v98, v99 Vec3
	var m2, out, v105, v44, v46, v92 Matrix3
	var out1, v117, v2, v3, v94 AABB
	var xw, xx, xy, xz, yw, yy, yz, zw, zz, v10, v101, v22, v33, v50, v51, v53, v54, v55, v57, v58, v59, v61, v65, v66, v68, v69, v70, v72, v73, v74, v76, v80, v81, v83, v84, v85, v87, v88, v89, v91 float32
	var v1, v14 Transform
	var v15, v43 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b4, c, center, extent, m2, out, out1, rv, t1, t2, t3, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v101, v102, v103, v105, v106, v107, v109, v11, v110, v111, v113, v114, v115, v117, v12, v14, v15, v16, v17, v18, v19, v2, v21, v22, v23, v24, v26, v27, v28, v3, v30, v32, v33, v34, v35, v37, v39, v4, v40, v41, v43, v44, v46, v47, v48, v50, v51, v53, v54, v55, v57, v58, v59, v6, v61, v62, v63, v65, v66, v68, v69, v7, v70, v72, v73, v74, v76, v77, v78, v8, v80, v81, v83, v84, v85, v87, v88, v89, v91, v92, v94, v95, v97, v98, v99
	v1 = transform1
	v2 = (*HullData)(unsafe.Pointer(shape)).Aabb
	v3 = v2
	v6 = v3.UpperBound
	v7 = v3.LowerBound
	v8 = Vec3{
		X: v6.X + v7.X,
		Y: v6.Y + v7.Y,
		Z: v6.Z + v7.Z,
	}
	goto _9
_9:
	v10 = float32FromFloat32(0.5)
	v11 = v8
	v12 = Vec3{
		X: float32(v10 * v11.X),
		Y: float32(v10 * v11.Y),
		Z: float32(v10 * v11.Z),
	}
	goto _13
_13:
	v4 = v12
	goto _5
_5:
	v14 = v1
	v15 = v14.Q
	v16 = v4
	v17 = v15.V
	v18 = v16
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v19 = c
	goto _20
_20:
	t1 = v19
	v21 = t1
	v22 = v15.S
	v23 = v16
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	t2 = v24
	v26 = v15.V
	v27 = t2
	c.X = float32(v26.Y*v27.Z) - float32(v26.Z*v27.Y)
	c.Y = float32(v26.Z*v27.X) - float32(v26.X*v27.Z)
	c.Z = float32(v26.X*v27.Y) - float32(v26.Y*v27.X)
	v28 = c
	goto _29
_29:
	t3 = v28
	v32 = v16
	v33 = float32FromFloat32(2)
	v34 = t3
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	v30 = v35
	goto _31
_31:
	rv = v30
	v39 = rv
	v40 = v14.P
	v41 = Vec3{
		X: v39.X + v40.X,
		Y: v39.Y + v40.Y,
		Z: v39.Z + v40.Z,
	}
	goto _42
_42:
	v37 = v41
	goto _38
_38:
	center = v37
	v43 = v1.Q
	xx = float32(v43.V.X * v43.V.X)
	yy = float32(v43.V.Y * v43.V.Y)
	zz = float32(v43.V.Z * v43.V.Z)
	xy = float32(v43.V.X * v43.V.Y)
	xz = float32(v43.V.X * v43.V.Z)
	xw = float32(v43.V.X * v43.S)
	yz = float32(v43.V.Y * v43.V.Z)
	yw = float32(v43.V.Y * v43.S)
	zw = float32(v43.V.Z * v43.S)
	v44 = Matrix3{
		Cx: Vec3{
			X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
			Y: float32(float32FromFloat32(2) * (xy + zw)),
			Z: float32(float32FromFloat32(2) * (xz - yw)),
		},
		Cy: Vec3{
			X: float32(float32FromFloat32(2) * (xy - zw)),
			Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
			Z: float32(float32FromFloat32(2) * (yz + xw)),
		},
		Cz: Vec3{
			X: float32(float32FromFloat32(2) * (xz + yw)),
			Y: float32(float32FromFloat32(2) * (yz - xw)),
			Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
		},
	}
	goto _45
_45:
	m2 = v44
	v46 = m2
	v47 = v46.Cx
	v50 = v47.X
	if v50 < float32FromInt32(0) {
		v53 = -v50
	} else {
		v53 = v50
	}
	v51 = v53
	goto _52
_52:
	v54 = v47.Y
	if v54 < float32FromInt32(0) {
		v57 = -v54
	} else {
		v57 = v54
	}
	v55 = v57
	goto _56
_56:
	v58 = v47.Z
	if v58 < float32FromInt32(0) {
		v61 = -v58
	} else {
		v61 = v58
	}
	v59 = v61
	goto _60
_60:
	v48 = Vec3{
		X: v51,
		Y: v55,
		Z: v59,
	}
	goto _49
_49:
	out.Cx = v48
	v62 = v46.Cy
	v65 = v62.X
	if v65 < float32FromInt32(0) {
		v68 = -v65
	} else {
		v68 = v65
	}
	v66 = v68
	goto _67
_67:
	v69 = v62.Y
	if v69 < float32FromInt32(0) {
		v72 = -v69
	} else {
		v72 = v69
	}
	v70 = v72
	goto _71
_71:
	v73 = v62.Z
	if v73 < float32FromInt32(0) {
		v76 = -v73
	} else {
		v76 = v73
	}
	v74 = v76
	goto _75
_75:
	v63 = Vec3{
		X: v66,
		Y: v70,
		Z: v74,
	}
	goto _64
_64:
	out.Cy = v63
	v77 = v46.Cz
	v80 = v77.X
	if v80 < float32FromInt32(0) {
		v83 = -v80
	} else {
		v83 = v80
	}
	v81 = v83
	goto _82
_82:
	v84 = v77.Y
	if v84 < float32FromInt32(0) {
		v87 = -v84
	} else {
		v87 = v84
	}
	v85 = v87
	goto _86
_86:
	v88 = v77.Z
	if v88 < float32FromInt32(0) {
		v91 = -v88
	} else {
		v91 = v88
	}
	v89 = v91
	goto _90
_90:
	v78 = Vec3{
		X: v81,
		Y: v85,
		Z: v89,
	}
	goto _79
_79:
	out.Cz = v78
	v92 = out
	goto _93
_93:
	v94 = v2
	v97 = v94.UpperBound
	v98 = v94.LowerBound
	v99 = Vec3{
		X: v97.X - v98.X,
		Y: v97.Y - v98.Y,
		Z: v97.Z - v98.Z,
	}
	goto _100
_100:
	v101 = float32FromFloat32(0.5)
	v102 = v99
	v103 = Vec3{
		X: float32(v101 * v102.X),
		Y: float32(v101 * v102.Y),
		Z: float32(v101 * v102.Z),
	}
	goto _104
_104:
	v95 = v103
	goto _96
_96:
	v105 = v92
	v106 = v95
	b4 = Vec3{
		X: float32(v105.Cx.X*v106.X) + float32(v105.Cy.X*v106.Y) + float32(v105.Cz.X*v106.Z),
		Y: float32(v105.Cx.Y*v106.X) + float32(v105.Cy.Y*v106.Y) + float32(v105.Cz.Y*v106.Z),
		Z: float32(v105.Cx.Z*v106.X) + float32(v105.Cy.Z*v106.Y) + float32(v105.Cz.Z*v106.Z),
	}
	v107 = b4
	goto _108
_108:
	extent = v107
	v109 = center
	v110 = extent
	v111 = Vec3{
		X: v109.X - v110.X,
		Y: v109.Y - v110.Y,
		Z: v109.Z - v110.Z,
	}
	goto _112
_112:
	v113 = center
	v114 = extent
	v115 = Vec3{
		X: v113.X + v114.X,
		Y: v113.Y + v114.Y,
		Z: v113.Z + v114.Z,
	}
	goto _116
_116:
	out1 = AABB{
		LowerBound: v111,
		UpperBound: v115,
	}
	v117 = out1
	goto _118
_118:
	return v117
}

func b3ComputeSweptHullAABB(tls *_Stack, shape uintptr, xf1 Transform, xf2 Transform) (r AABB) {
	var aabb1, aabb2, out1, out2, v117, v120, v121, v2, v212, v235, v237, v238, v277, v3, v94 AABB
	var b8, c, center, extent, rv, t1, t2, t3, v102, v103, v106, v107, v109, v11, v110, v111, v113, v114, v115, v12, v122, v124, v125, v126, v129, v130, v134, v135, v136, v137, v139, v141, v142, v144, v145, v146, v148, v150, v152, v153, v155, v157, v158, v159, v16, v165, v166, v17, v18, v180, v181, v19, v195, v196, v21, v213, v215, v216, v217, v220, v221, v224, v225, v227, v228, v229, v23, v231, v232, v233, v239, v24, v240, v241, v258, v259, v26, v260, v27, v28, v30, v32, v34, v35, v37, v39, v4, v40, v41, v47, v48, v6, v62, v63, v7, v77, v78, v8, v95, v97, v98, v99 Vec3
	var m2, out, v105, v162, v164, v210, v223, v44, v46, v92 Matrix3
	var xw, xx, xy, xz, yw, yy, yz, zw, zz, v10, v101, v128, v140, v151, v168, v169, v171, v172, v173, v175, v176, v177, v179, v183, v184, v186, v187, v188, v190, v191, v192, v194, v198, v199, v201, v202, v203, v205, v206, v207, v209, v219, v22, v243, v244, v245, v247, v248, v249, v250, v252, v253, v254, v255, v257, v262, v263, v264, v266, v267, v268, v269, v271, v272, v273, v274, v276, v33, v50, v51, v53, v54, v55, v57, v58, v59, v61, v65, v66, v68, v69, v70, v72, v73, v74, v76, v80, v81, v83, v84, v85, v87, v88, v89, v91 float32
	var v1, v119, v132, v14 Transform
	var v133, v15, v161, v43 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb1, aabb2, b8, c, center, extent, m2, out, out1, out2, rv, t1, t2, t3, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v101, v102, v103, v105, v106, v107, v109, v11, v110, v111, v113, v114, v115, v117, v119, v12, v120, v121, v122, v124, v125, v126, v128, v129, v130, v132, v133, v134, v135, v136, v137, v139, v14, v140, v141, v142, v144, v145, v146, v148, v15, v150, v151, v152, v153, v155, v157, v158, v159, v16, v161, v162, v164, v165, v166, v168, v169, v17, v171, v172, v173, v175, v176, v177, v179, v18, v180, v181, v183, v184, v186, v187, v188, v19, v190, v191, v192, v194, v195, v196, v198, v199, v2, v201, v202, v203, v205, v206, v207, v209, v21, v210, v212, v213, v215, v216, v217, v219, v22, v220, v221, v223, v224, v225, v227, v228, v229, v23, v231, v232, v233, v235, v237, v238, v239, v24, v240, v241, v243, v244, v245, v247, v248, v249, v250, v252, v253, v254, v255, v257, v258, v259, v26, v260, v262, v263, v264, v266, v267, v268, v269, v27, v271, v272, v273, v274, v276, v277, v28, v3, v30, v32, v33, v34, v35, v37, v39, v4, v40, v41, v43, v44, v46, v47, v48, v50, v51, v53, v54, v55, v57, v58, v59, v6, v61, v62, v63, v65, v66, v68, v69, v7, v70, v72, v73, v74, v76, v77, v78, v8, v80, v81, v83, v84, v85, v87, v88, v89, v91, v92, v94, v95, v97, v98, v99
	v1 = xf1
	v2 = (*HullData)(unsafe.Pointer(shape)).Aabb
	v3 = v2
	v6 = v3.UpperBound
	v7 = v3.LowerBound
	v8 = Vec3{
		X: v6.X + v7.X,
		Y: v6.Y + v7.Y,
		Z: v6.Z + v7.Z,
	}
	goto _9
_9:
	v10 = float32FromFloat32(0.5)
	v11 = v8
	v12 = Vec3{
		X: float32(v10 * v11.X),
		Y: float32(v10 * v11.Y),
		Z: float32(v10 * v11.Z),
	}
	goto _13
_13:
	v4 = v12
	goto _5
_5:
	v14 = v1
	v15 = v14.Q
	v16 = v4
	v17 = v15.V
	v18 = v16
	c.X = float32(v17.Y*v18.Z) - float32(v17.Z*v18.Y)
	c.Y = float32(v17.Z*v18.X) - float32(v17.X*v18.Z)
	c.Z = float32(v17.X*v18.Y) - float32(v17.Y*v18.X)
	v19 = c
	goto _20
_20:
	t1 = v19
	v21 = t1
	v22 = v15.S
	v23 = v16
	v24 = Vec3{
		X: v21.X + float32(v22*v23.X),
		Y: v21.Y + float32(v22*v23.Y),
		Z: v21.Z + float32(v22*v23.Z),
	}
	goto _25
_25:
	t2 = v24
	v26 = v15.V
	v27 = t2
	c.X = float32(v26.Y*v27.Z) - float32(v26.Z*v27.Y)
	c.Y = float32(v26.Z*v27.X) - float32(v26.X*v27.Z)
	c.Z = float32(v26.X*v27.Y) - float32(v26.Y*v27.X)
	v28 = c
	goto _29
_29:
	t3 = v28
	v32 = v16
	v33 = float32FromFloat32(2)
	v34 = t3
	v35 = Vec3{
		X: v32.X + float32(v33*v34.X),
		Y: v32.Y + float32(v33*v34.Y),
		Z: v32.Z + float32(v33*v34.Z),
	}
	goto _36
_36:
	v30 = v35
	goto _31
_31:
	rv = v30
	v39 = rv
	v40 = v14.P
	v41 = Vec3{
		X: v39.X + v40.X,
		Y: v39.Y + v40.Y,
		Z: v39.Z + v40.Z,
	}
	goto _42
_42:
	v37 = v41
	goto _38
_38:
	center = v37
	v43 = v1.Q
	xx = float32(v43.V.X * v43.V.X)
	yy = float32(v43.V.Y * v43.V.Y)
	zz = float32(v43.V.Z * v43.V.Z)
	xy = float32(v43.V.X * v43.V.Y)
	xz = float32(v43.V.X * v43.V.Z)
	xw = float32(v43.V.X * v43.S)
	yz = float32(v43.V.Y * v43.V.Z)
	yw = float32(v43.V.Y * v43.S)
	zw = float32(v43.V.Z * v43.S)
	v44 = Matrix3{
		Cx: Vec3{
			X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
			Y: float32(float32FromFloat32(2) * (xy + zw)),
			Z: float32(float32FromFloat32(2) * (xz - yw)),
		},
		Cy: Vec3{
			X: float32(float32FromFloat32(2) * (xy - zw)),
			Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
			Z: float32(float32FromFloat32(2) * (yz + xw)),
		},
		Cz: Vec3{
			X: float32(float32FromFloat32(2) * (xz + yw)),
			Y: float32(float32FromFloat32(2) * (yz - xw)),
			Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
		},
	}
	goto _45
_45:
	m2 = v44
	v46 = m2
	v47 = v46.Cx
	v50 = v47.X
	if v50 < float32FromInt32(0) {
		v53 = -v50
	} else {
		v53 = v50
	}
	v51 = v53
	goto _52
_52:
	v54 = v47.Y
	if v54 < float32FromInt32(0) {
		v57 = -v54
	} else {
		v57 = v54
	}
	v55 = v57
	goto _56
_56:
	v58 = v47.Z
	if v58 < float32FromInt32(0) {
		v61 = -v58
	} else {
		v61 = v58
	}
	v59 = v61
	goto _60
_60:
	v48 = Vec3{
		X: v51,
		Y: v55,
		Z: v59,
	}
	goto _49
_49:
	out.Cx = v48
	v62 = v46.Cy
	v65 = v62.X
	if v65 < float32FromInt32(0) {
		v68 = -v65
	} else {
		v68 = v65
	}
	v66 = v68
	goto _67
_67:
	v69 = v62.Y
	if v69 < float32FromInt32(0) {
		v72 = -v69
	} else {
		v72 = v69
	}
	v70 = v72
	goto _71
_71:
	v73 = v62.Z
	if v73 < float32FromInt32(0) {
		v76 = -v73
	} else {
		v76 = v73
	}
	v74 = v76
	goto _75
_75:
	v63 = Vec3{
		X: v66,
		Y: v70,
		Z: v74,
	}
	goto _64
_64:
	out.Cy = v63
	v77 = v46.Cz
	v80 = v77.X
	if v80 < float32FromInt32(0) {
		v83 = -v80
	} else {
		v83 = v80
	}
	v81 = v83
	goto _82
_82:
	v84 = v77.Y
	if v84 < float32FromInt32(0) {
		v87 = -v84
	} else {
		v87 = v84
	}
	v85 = v87
	goto _86
_86:
	v88 = v77.Z
	if v88 < float32FromInt32(0) {
		v91 = -v88
	} else {
		v91 = v88
	}
	v89 = v91
	goto _90
_90:
	v78 = Vec3{
		X: v81,
		Y: v85,
		Z: v89,
	}
	goto _79
_79:
	out.Cz = v78
	v92 = out
	goto _93
_93:
	v94 = v2
	v97 = v94.UpperBound
	v98 = v94.LowerBound
	v99 = Vec3{
		X: v97.X - v98.X,
		Y: v97.Y - v98.Y,
		Z: v97.Z - v98.Z,
	}
	goto _100
_100:
	v101 = float32FromFloat32(0.5)
	v102 = v99
	v103 = Vec3{
		X: float32(v101 * v102.X),
		Y: float32(v101 * v102.Y),
		Z: float32(v101 * v102.Z),
	}
	goto _104
_104:
	v95 = v103
	goto _96
_96:
	v105 = v92
	v106 = v95
	b8 = Vec3{
		X: float32(v105.Cx.X*v106.X) + float32(v105.Cy.X*v106.Y) + float32(v105.Cz.X*v106.Z),
		Y: float32(v105.Cx.Y*v106.X) + float32(v105.Cy.Y*v106.Y) + float32(v105.Cz.Y*v106.Z),
		Z: float32(v105.Cx.Z*v106.X) + float32(v105.Cy.Z*v106.Y) + float32(v105.Cz.Z*v106.Z),
	}
	v107 = b8
	goto _108
_108:
	extent = v107
	v109 = center
	v110 = extent
	v111 = Vec3{
		X: v109.X - v110.X,
		Y: v109.Y - v110.Y,
		Z: v109.Z - v110.Z,
	}
	goto _112
_112:
	v113 = center
	v114 = extent
	v115 = Vec3{
		X: v113.X + v114.X,
		Y: v113.Y + v114.Y,
		Z: v113.Z + v114.Z,
	}
	goto _116
_116:
	out2 = AABB{
		LowerBound: v111,
		UpperBound: v115,
	}
	v117 = out2
	goto _118
_118:
	aabb1 = v117
	v119 = xf2
	v120 = (*HullData)(unsafe.Pointer(shape)).Aabb
	v121 = v120
	v124 = v121.UpperBound
	v125 = v121.LowerBound
	v126 = Vec3{
		X: v124.X + v125.X,
		Y: v124.Y + v125.Y,
		Z: v124.Z + v125.Z,
	}
	goto _127
_127:
	v128 = float32FromFloat32(0.5)
	v129 = v126
	v130 = Vec3{
		X: float32(v128 * v129.X),
		Y: float32(v128 * v129.Y),
		Z: float32(v128 * v129.Z),
	}
	goto _131
_131:
	v122 = v130
	goto _123
_123:
	v132 = v119
	v133 = v132.Q
	v134 = v122
	v135 = v133.V
	v136 = v134
	c.X = float32(v135.Y*v136.Z) - float32(v135.Z*v136.Y)
	c.Y = float32(v135.Z*v136.X) - float32(v135.X*v136.Z)
	c.Z = float32(v135.X*v136.Y) - float32(v135.Y*v136.X)
	v137 = c
	goto _138
_138:
	t1 = v137
	v139 = t1
	v140 = v133.S
	v141 = v134
	v142 = Vec3{
		X: v139.X + float32(v140*v141.X),
		Y: v139.Y + float32(v140*v141.Y),
		Z: v139.Z + float32(v140*v141.Z),
	}
	goto _143
_143:
	t2 = v142
	v144 = v133.V
	v145 = t2
	c.X = float32(v144.Y*v145.Z) - float32(v144.Z*v145.Y)
	c.Y = float32(v144.Z*v145.X) - float32(v144.X*v145.Z)
	c.Z = float32(v144.X*v145.Y) - float32(v144.Y*v145.X)
	v146 = c
	goto _147
_147:
	t3 = v146
	v150 = v134
	v151 = float32FromFloat32(2)
	v152 = t3
	v153 = Vec3{
		X: v150.X + float32(v151*v152.X),
		Y: v150.Y + float32(v151*v152.Y),
		Z: v150.Z + float32(v151*v152.Z),
	}
	goto _154
_154:
	v148 = v153
	goto _149
_149:
	rv = v148
	v157 = rv
	v158 = v132.P
	v159 = Vec3{
		X: v157.X + v158.X,
		Y: v157.Y + v158.Y,
		Z: v157.Z + v158.Z,
	}
	goto _160
_160:
	v155 = v159
	goto _156
_156:
	center = v155
	v161 = v119.Q
	xx = float32(v161.V.X * v161.V.X)
	yy = float32(v161.V.Y * v161.V.Y)
	zz = float32(v161.V.Z * v161.V.Z)
	xy = float32(v161.V.X * v161.V.Y)
	xz = float32(v161.V.X * v161.V.Z)
	xw = float32(v161.V.X * v161.S)
	yz = float32(v161.V.Y * v161.V.Z)
	yw = float32(v161.V.Y * v161.S)
	zw = float32(v161.V.Z * v161.S)
	v162 = Matrix3{
		Cx: Vec3{
			X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
			Y: float32(float32FromFloat32(2) * (xy + zw)),
			Z: float32(float32FromFloat32(2) * (xz - yw)),
		},
		Cy: Vec3{
			X: float32(float32FromFloat32(2) * (xy - zw)),
			Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
			Z: float32(float32FromFloat32(2) * (yz + xw)),
		},
		Cz: Vec3{
			X: float32(float32FromFloat32(2) * (xz + yw)),
			Y: float32(float32FromFloat32(2) * (yz - xw)),
			Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
		},
	}
	goto _163
_163:
	m2 = v162
	v164 = m2
	v165 = v164.Cx
	v168 = v165.X
	if v168 < float32FromInt32(0) {
		v171 = -v168
	} else {
		v171 = v168
	}
	v169 = v171
	goto _170
_170:
	v172 = v165.Y
	if v172 < float32FromInt32(0) {
		v175 = -v172
	} else {
		v175 = v172
	}
	v173 = v175
	goto _174
_174:
	v176 = v165.Z
	if v176 < float32FromInt32(0) {
		v179 = -v176
	} else {
		v179 = v176
	}
	v177 = v179
	goto _178
_178:
	v166 = Vec3{
		X: v169,
		Y: v173,
		Z: v177,
	}
	goto _167
_167:
	out.Cx = v166
	v180 = v164.Cy
	v183 = v180.X
	if v183 < float32FromInt32(0) {
		v186 = -v183
	} else {
		v186 = v183
	}
	v184 = v186
	goto _185
_185:
	v187 = v180.Y
	if v187 < float32FromInt32(0) {
		v190 = -v187
	} else {
		v190 = v187
	}
	v188 = v190
	goto _189
_189:
	v191 = v180.Z
	if v191 < float32FromInt32(0) {
		v194 = -v191
	} else {
		v194 = v191
	}
	v192 = v194
	goto _193
_193:
	v181 = Vec3{
		X: v184,
		Y: v188,
		Z: v192,
	}
	goto _182
_182:
	out.Cy = v181
	v195 = v164.Cz
	v198 = v195.X
	if v198 < float32FromInt32(0) {
		v201 = -v198
	} else {
		v201 = v198
	}
	v199 = v201
	goto _200
_200:
	v202 = v195.Y
	if v202 < float32FromInt32(0) {
		v205 = -v202
	} else {
		v205 = v202
	}
	v203 = v205
	goto _204
_204:
	v206 = v195.Z
	if v206 < float32FromInt32(0) {
		v209 = -v206
	} else {
		v209 = v206
	}
	v207 = v209
	goto _208
_208:
	v196 = Vec3{
		X: v199,
		Y: v203,
		Z: v207,
	}
	goto _197
_197:
	out.Cz = v196
	v210 = out
	goto _211
_211:
	v212 = v120
	v215 = v212.UpperBound
	v216 = v212.LowerBound
	v217 = Vec3{
		X: v215.X - v216.X,
		Y: v215.Y - v216.Y,
		Z: v215.Z - v216.Z,
	}
	goto _218
_218:
	v219 = float32FromFloat32(0.5)
	v220 = v217
	v221 = Vec3{
		X: float32(v219 * v220.X),
		Y: float32(v219 * v220.Y),
		Z: float32(v219 * v220.Z),
	}
	goto _222
_222:
	v213 = v221
	goto _214
_214:
	v223 = v210
	v224 = v213
	b8 = Vec3{
		X: float32(v223.Cx.X*v224.X) + float32(v223.Cy.X*v224.Y) + float32(v223.Cz.X*v224.Z),
		Y: float32(v223.Cx.Y*v224.X) + float32(v223.Cy.Y*v224.Y) + float32(v223.Cz.Y*v224.Z),
		Z: float32(v223.Cx.Z*v224.X) + float32(v223.Cy.Z*v224.Y) + float32(v223.Cz.Z*v224.Z),
	}
	v225 = b8
	goto _226
_226:
	extent = v225
	v227 = center
	v228 = extent
	v229 = Vec3{
		X: v227.X - v228.X,
		Y: v227.Y - v228.Y,
		Z: v227.Z - v228.Z,
	}
	goto _230
_230:
	v231 = center
	v232 = extent
	v233 = Vec3{
		X: v231.X + v232.X,
		Y: v231.Y + v232.Y,
		Z: v231.Z + v232.Z,
	}
	goto _234
_234:
	out2 = AABB{
		LowerBound: v229,
		UpperBound: v233,
	}
	v235 = out2
	goto _236
_236:
	aabb2 = v235
	v237 = aabb1
	v238 = aabb2
	v239 = v237.LowerBound
	v240 = v238.LowerBound
	v243 = v239.X
	v244 = v240.X
	if v243 < v244 {
		v247 = v243
	} else {
		v247 = v244
	}
	v245 = v247
	goto _246
_246:
	v248 = v239.Y
	v249 = v240.Y
	if v248 < v249 {
		v252 = v248
	} else {
		v252 = v249
	}
	v250 = v252
	goto _251
_251:
	v253 = v239.Z
	v254 = v240.Z
	if v253 < v254 {
		v257 = v253
	} else {
		v257 = v254
	}
	v255 = v257
	goto _256
_256:
	v241 = Vec3{
		X: v245,
		Y: v250,
		Z: v255,
	}
	goto _242
_242:
	out1.LowerBound = v241
	v258 = v237.UpperBound
	v259 = v238.UpperBound
	v262 = v258.X
	v263 = v259.X
	if v262 > v263 {
		v266 = v262
	} else {
		v266 = v263
	}
	v264 = v266
	goto _265
_265:
	v267 = v258.Y
	v268 = v259.Y
	if v267 > v268 {
		v271 = v267
	} else {
		v271 = v268
	}
	v269 = v271
	goto _270
_270:
	v272 = v258.Z
	v273 = v259.Z
	if v272 > v273 {
		v276 = v272
	} else {
		v276 = v273
	}
	v274 = v276
	goto _275
_275:
	v260 = Vec3{
		X: v264,
		Y: v269,
		Z: v274,
	}
	goto _261
_261:
	out1.UpperBound = v260
	v277 = out1
	goto _278
_278:
	return v277
}

func b3ComputeHullExtent(tls *_Stack, hull1 uintptr, origin Vec3) (r b3ShapeExtent) {
	var extent b3ShapeExtent
	var index int32
	var point, v10, v24, v25, v26, v5, v6, v7, v9 Vec3
	var points, v1, v2 uintptr
	var v12, v13, v15, v16, v17, v19, v20, v21, v23, v28, v29, v30, v32, v33, v34, v35, v37, v38, v39, v40, v42 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = extent, index, point, points, v1, v10, v12, v13, v15, v16, v17, v19, v2, v20, v21, v23, v24, v25, v26, v28, v29, v30, v32, v33, v34, v35, v37, v38, v39, v40, v42, v5, v6, v7, v9
	v1 = hull1
	if (*HullData)(unsafe.Pointer(v1)).PointOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).PointOffset))
	goto _3
_3:
	points = v2
	extent.MinExtent = (*HullData)(unsafe.Pointer(hull1)).InnerRadius
	extent.MaxExtent = b3Vec3_zero
	index = 0
	for {
		if !(index < (*HullData)(unsafe.Pointer(hull1)).VertexCount) {
			break
		}
		point = *(*Vec3)(unsafe.Add(nil, points+uintptr(index)*12))
		v5 = point
		v6 = origin
		v7 = Vec3{
			X: v5.X - v6.X,
			Y: v5.Y - v6.Y,
			Z: v5.Z - v6.Z,
		}
		goto _8
	_8:
		v9 = v7
		v12 = v9.X
		if v12 < float32FromInt32(0) {
			v15 = -v12
		} else {
			v15 = v12
		}
		v13 = v15
		goto _14
	_14:
		v16 = v9.Y
		if v16 < float32FromInt32(0) {
			v19 = -v16
		} else {
			v19 = v16
		}
		v17 = v19
		goto _18
	_18:
		v20 = v9.Z
		if v20 < float32FromInt32(0) {
			v23 = -v20
		} else {
			v23 = v20
		}
		v21 = v23
		goto _22
	_22:
		v10 = Vec3{
			X: v13,
			Y: v17,
			Z: v21,
		}
		goto _11
	_11:
		v24 = extent.MaxExtent
		v25 = v10
		v28 = v24.X
		v29 = v25.X
		if v28 > v29 {
			v32 = v28
		} else {
			v32 = v29
		}
		v30 = v32
		goto _31
	_31:
		v33 = v24.Y
		v34 = v25.Y
		if v33 > v34 {
			v37 = v33
		} else {
			v37 = v34
		}
		v35 = v37
		goto _36
	_36:
		v38 = v24.Z
		v39 = v25.Z
		if v38 > v39 {
			v42 = v38
		} else {
			v42 = v39
		}
		v40 = v42
		goto _41
	_41:
		v26 = Vec3{
			X: v30,
			Y: v35,
			Z: v40,
		}
		goto _27
	_27:
		extent.MaxExtent = v26
		goto _4
	_4:
		;
		index = index + 1
	}
	return extent
}

func b3ComputeHullProjectedArea(tls *_Stack, hull3 uintptr, direction Vec3) (r float32) {
	var a4, area, v25, v27, v28, v29, v31 float32
	var baseEdge, edgeIndex, faceCount, i int32
	var c, e1, e2, n, p1, p2, p3, v11, v12, v13, v15, v16, v17, v19, v20, v21, v23, v24 Vec3
	var edge, face, hullEdges, hullFaces, hullPoints, v1, v2, v4, v5, v7, v8 uintptr
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a4, area, baseEdge, c, e1, e2, edge, edgeIndex, face, faceCount, hullEdges, hullFaces, hullPoints, i, n, p1, p2, p3, v1, v11, v12, v13, v15, v16, v17, v19, v2, v20, v21, v23, v24, v25, v27, v28, v29, v31, v4, v5, v7, v8
	area = float32FromFloat32(0)
	faceCount = (*HullData)(unsafe.Pointer(hull3)).FaceCount
	v1 = hull3
	if (*HullData)(unsafe.Pointer(v1)).FaceOffset == int32FromInt32(0) {
		v2 = uintptrFromInt32(0)
		goto _3
	}
	v2 = uintptr(int64(v1) + int64((*HullData)(unsafe.Pointer(v1)).FaceOffset))
	goto _3
_3:
	hullFaces = v2
	v4 = hull3
	if (*HullData)(unsafe.Pointer(v4)).EdgeOffset == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _6
	}
	v5 = uintptr(int64(v4) + int64((*HullData)(unsafe.Pointer(v4)).EdgeOffset))
	goto _6
_6:
	hullEdges = v5
	v7 = hull3
	if (*HullData)(unsafe.Pointer(v7)).PointOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _9
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).PointOffset))
	goto _9
_9:
	hullPoints = v8
	i = 0
	for {
		if !(i < faceCount) {
			break
		}
		face = hullFaces + uintptr(i)
		baseEdge = int32FromUint8((*HullFace)(unsafe.Pointer(face)).Edge)
		edge = hullEdges + uintptr(baseEdge)*4
		p1 = *(*Vec3)(unsafe.Add(nil, hullPoints+uintptr((*HullHalfEdge)(unsafe.Pointer(edge)).Origin)*12))
		edgeIndex = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge)).Next)
		edge = hullEdges + uintptr(edgeIndex)*4
		p2 = *(*Vec3)(unsafe.Add(nil, hullPoints+uintptr((*HullHalfEdge)(unsafe.Pointer(edge)).Origin)*12))
		edgeIndex = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge)).Next)
		for cond := true; cond; cond = edgeIndex != baseEdge {
			edge = hullEdges + uintptr(edgeIndex)*4
			p3 = *(*Vec3)(unsafe.Add(nil, hullPoints+uintptr((*HullHalfEdge)(unsafe.Pointer(edge)).Origin)*12))
			v11 = p2
			v12 = p1
			v13 = Vec3{
				X: v11.X - v12.X,
				Y: v11.Y - v12.Y,
				Z: v11.Z - v12.Z,
			}
			goto _14
		_14:
			e1 = v13
			v15 = p3
			v16 = p1
			v17 = Vec3{
				X: v15.X - v16.X,
				Y: v15.Y - v16.Y,
				Z: v15.Z - v16.Z,
			}
			goto _18
		_18:
			e2 = v17
			v19 = e1
			v20 = e2
			c.X = float32(v19.Y*v20.Z) - float32(v19.Z*v20.Y)
			c.Y = float32(v19.Z*v20.X) - float32(v19.X*v20.Z)
			c.Z = float32(v19.X*v20.Y) - float32(v19.Y*v20.X)
			v21 = c
			goto _22
		_22:
			n = v21
			v23 = n
			v24 = direction
			v25 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y) + float32(v23.Z*v24.Z)
			goto _26
		_26:
			a4 = v25
			v27 = a4
			v28 = float32FromFloat32(0)
			if v27 > v28 {
				v31 = v27
			} else {
				v31 = v28
			}
			v29 = v31
			goto _30
		_30:
			area = area + v29
			p2 = p3
			edgeIndex = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge)).Next)
		}
		goto _10
	_10:
		;
		i = i + 1
	}
	return float32(float32FromFloat32(0.5) * area)
}

// C documentation
//
//	// Approximate cosine and sine for determinism. In my testing cosf and sinf produced
//	// the same results on x64 and ARM using MSVC, GCC, and Clang. However, I don't trust
//	// this result.
//	// https://en.wikipedia.org/wiki/Bh%C4%81skara_I%27s_sine_approximation_formula
func b3ComputeCosSin(tls *_Stack, radians1 float32) (r CosSin) {
	var c, invMag, mag, pi2, s, x, y, y1, y2, y21, y22, y23, v1 float32
	var cs CosSin
	_, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, cs, invMag, mag, pi2, s, x, y, y1, y2, y21, y22, y23, v1
	v1 = remainderf(tls, radians1, float32(float32FromFloat32(2)*float32FromFloat32(3.14159265359)))
	goto _2
_2:
	x = v1
	pi2 = float32(float32FromFloat32(3.14159265359) * float32FromFloat32(3.14159265359))
	if x < float32(-float32FromFloat32(0.5)*float32FromFloat32(3.14159265359)) {
		y = x + float32FromFloat32(3.14159265359)
		y21 = float32(y * y)
		c = -(pi2 - float32(float32FromFloat32(4)*y21)) / (pi2 + y21)
	} else {
		if x > float32(float32FromFloat32(0.5)*float32FromFloat32(3.14159265359)) {
			y1 = x - float32FromFloat32(3.14159265359)
			y22 = float32(y1 * y1)
			c = -(pi2 - float32(float32FromFloat32(4)*y22)) / (pi2 + y22)
		} else {
			y23 = float32(x * x)
			c = (pi2 - float32(float32FromFloat32(4)*y23)) / (pi2 + y23)
		}
	}
	if x < float32FromFloat32(0) {
		y2 = x + float32FromFloat32(3.14159265359)
		s = float32(float32(-float32FromFloat32(16)*y2)*(float32FromFloat32(3.14159265359)-y2)) / (float32(float32FromFloat32(5)*pi2) - float32(float32(float32FromFloat32(4)*y2)*(float32FromFloat32(3.14159265359)-y2)))
	} else {
		s = float32(float32(float32FromFloat32(16)*x)*(float32FromFloat32(3.14159265359)-x)) / (float32(float32FromFloat32(5)*pi2) - float32(float32(float32FromFloat32(4)*x)*(float32FromFloat32(3.14159265359)-x)))
	}
	mag = sqrtf(tls, float32(s*s)+float32(c*c))
	if mag > float32FromFloat32(0) {
		v1 = float32FromFloat32(1) / mag
	} else {
		v1 = float32FromFloat32(0)
	}
	invMag = v1
	cs = CosSin{
		Cosine: float32(c * invMag),
		Sine:   float32(s * invMag),
	}
	return cs
}

func b3ComputeQuatBetweenUnitVectors(tls *_Stack, v1 Vec3, v2 Vec3) (r Quat) {
	var aa, lengthSq, s, tolerance, v111, v4, v8, v9 float32
	var c, c1, m, v11, v13, v14, v15, v21, v25, v3 Vec3
	var out, qn, v46, v47, v48, v51 Quat
	var v12 bool
	var v6 uint8
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aa, c, c1, lengthSq, m, out, qn, s, tolerance, v11, v111, v12, v13, v14, v15, v21, v25, v3, v4, v46, v47, v48, v51, v6, v8, v9
	v11 = v1
	v21 = v11
	v3 = v11
	v4 = float32(v21.X*v3.X) + float32(v21.Y*v3.Y) + float32(v21.Z*v3.Z)
	goto _5
_5:
	aa = v4
	v8 = float32FromFloat32(1) - aa
	if v8 < float32FromInt32(0) {
		v111 = -v8
	} else {
		v111 = v8
	}
	v9 = v111
	goto _10
_10:
	v6 = boolUint8(v9 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
	goto _7
_7:
	;
	if v12 = !!(v6 != 0); !v12 {
		b3InternalAssert(tls, __ccgo_ts+21073, __ccgo_ts+21094, int32FromInt32(313))
	}
	_ = v12 || bool(int32FromInt32(0) != 0)
	v11 = v2
	v21 = v11
	v3 = v11
	v4 = float32(v21.X*v3.X) + float32(v21.Y*v3.Y) + float32(v21.Z*v3.Z)
	goto _17
_17:
	aa = v4
	v8 = float32FromFloat32(1) - aa
	if v8 < float32FromInt32(0) {
		v111 = -v8
	} else {
		v111 = v8
	}
	v9 = v111
	goto _22
_22:
	v6 = boolUint8(v9 < float32(float32FromFloat32(100)*float32FromFloat32(1.1920928955078125e-07)))
	goto _19
_19:
	;
	if v12 = !!(v6 != 0); !v12 {
		b3InternalAssert(tls, __ccgo_ts+21126, __ccgo_ts+21094, int32FromInt32(314))
	}
	_ = v12 || bool(int32FromInt32(0) != 0)
	v11 = v1
	v21 = v2
	v4 = float32FromFloat32(0.5)
	if v12 = !!(float32FromFloat32(0) <= v4 && v4 <= float32FromFloat32(1)); !v12 {
		b3InternalAssert(tls, __ccgo_ts+9783, __ccgo_ts+6355, int32FromInt32(354))
	}
	_ = v12 || bool(int32FromInt32(0) != 0)
	c1 = Vec3{
		X: float32((float32FromFloat32(1)-v4)*v11.X) + float32(v4*v21.X),
		Y: float32((float32FromFloat32(1)-v4)*v11.Y) + float32(v4*v21.Y),
		Z: float32((float32FromFloat32(1)-v4)*v11.Z) + float32(v4*v21.Z),
	}
	v3 = c1
	goto _30
_30:
	m = v3
	tolerance = float32(float32FromFloat32(100) * float32FromFloat32(1.1920928955078125e-07))
	v13 = m
	v8 = float32(v13.X*v13.X) + float32(v13.Y*v13.Y) + float32(v13.Z*v13.Z)
	goto _33
_33:
	if v8 > float32(tolerance*tolerance) {
		v14 = v1
		v15 = m
		c.X = float32(v14.Y*v15.Z) - float32(v14.Z*v15.Y)
		c.Y = float32(v14.Z*v15.X) - float32(v14.X*v15.Z)
		c.Z = float32(v14.X*v15.Y) - float32(v14.Y*v15.X)
		v25 = c
		goto _37
	_37:
		out.V = v25
		v11 = v1
		v21 = m
		v4 = float32(v11.X*v21.X) + float32(v11.Y*v21.Y) + float32(v11.Z*v21.Z)
		goto _41
	_41:
		out.S = v4
	} else {
		// Anti-parallel: Use a perpendicular vector
		v4 = v1.X
		if v4 < float32FromInt32(0) {
			v9 = -v4
		} else {
			v9 = v4
		}
		v8 = v9
		goto _44
	_44:
		if v8 > float32FromFloat32(0.5) {
			out.V.X = v1.Y
			out.V.Y = -v1.X
			out.V.Z = float32FromFloat32(0)
		} else {
			out.V.X = float32FromFloat32(0)
			out.V.Y = v1.Z
			out.V.Z = -v1.Y
		}
		out.S = float32FromFloat32(0)
	}
	// The algorithm is simplified and made more accurate by normalizing at the end
	v46 = out
	v47 = v46
	v48 = v46
	v4 = float32(v47.V.X*v48.V.X) + float32(v47.V.Y*v48.V.Y) + float32(v47.V.Z*v48.V.Z) + float32(v47.S*v48.S)
	goto _50
_50:
	lengthSq = v4
	if lengthSq > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
		s = float32FromFloat32(1) / sqrtf(tls, lengthSq)
		qn = Quat{
			V: Vec3{
				X: float32(s * v46.V.X),
				Y: float32(s * v46.V.Y),
				Z: float32(s * v46.V.Z),
			},
			S: float32(s * v46.S),
		}
		v51 = qn
		goto _52
	}
	v51 = b3Quat_identity
	goto _52
_52:
	return v51
}

func b3ComputeMeshAABB(tls *_Stack, shape uintptr, transform1 Transform, scale Vec3) (r AABB) {
	var b9, c, center, extent, rv, scaledLower, scaledUpper, t1, t2, t3, v1, v10, v108, v109, v11, v123, v124, v141, v143, v144, v145, v148, v149, v152, v153, v155, v156, v157, v159, v160, v161, v2, v28, v29, v3, v30, v5, v50, v52, v53, v54, v57, v58, v6, v62, v63, v64, v65, v67, v69, v7, v70, v72, v73, v74, v76, v78, v80, v81, v83, v85, v86, v87, v9, v93, v94 Vec3
	var bounds, out1, v140, v163, v48, v49 AABB
	var m2, out, v138, v151, v90, v92 Matrix3
	var xw, xx, xy, xz, yw, yy, yz, zw, zz, v100, v101, v103, v104, v105, v107, v111, v112, v114, v115, v116, v118, v119, v120, v122, v126, v127, v129, v13, v130, v131, v133, v134, v135, v137, v14, v147, v15, v17, v18, v19, v20, v22, v23, v24, v25, v27, v32, v33, v34, v36, v37, v38, v39, v41, v42, v43, v44, v46, v56, v68, v79, v96, v97, v99 float32
	var v47, v60 Transform
	var v61, v89 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = b9, bounds, c, center, extent, m2, out, out1, rv, scaledLower, scaledUpper, t1, t2, t3, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v10, v100, v101, v103, v104, v105, v107, v108, v109, v11, v111, v112, v114, v115, v116, v118, v119, v120, v122, v123, v124, v126, v127, v129, v13, v130, v131, v133, v134, v135, v137, v138, v14, v140, v141, v143, v144, v145, v147, v148, v149, v15, v151, v152, v153, v155, v156, v157, v159, v160, v161, v163, v17, v18, v19, v2, v20, v22, v23, v24, v25, v27, v28, v29, v3, v30, v32, v33, v34, v36, v37, v38, v39, v41, v42, v43, v44, v46, v47, v48, v49, v5, v50, v52, v53, v54, v56, v57, v58, v6, v60, v61, v62, v63, v64, v65, v67, v68, v69, v7, v70, v72, v73, v74, v76, v78, v79, v80, v81, v83, v85, v86, v87, v89, v9, v90, v92, v93, v94, v96, v97, v99
	v1 = scale
	v2 = (*MeshData)(unsafe.Pointer(shape)).Bounds.LowerBound
	v3 = Vec3{
		X: float32(v1.X * v2.X),
		Y: float32(v1.Y * v2.Y),
		Z: float32(v1.Z * v2.Z),
	}
	goto _4
_4:
	scaledLower = v3
	v5 = scale
	v6 = (*MeshData)(unsafe.Pointer(shape)).Bounds.UpperBound
	v7 = Vec3{
		X: float32(v5.X * v6.X),
		Y: float32(v5.Y * v6.Y),
		Z: float32(v5.Z * v6.Z),
	}
	goto _8
_8:
	scaledUpper = v7
	v9 = scaledLower
	v10 = scaledUpper
	v13 = v9.X
	v14 = v10.X
	if v13 < v14 {
		v17 = v13
	} else {
		v17 = v14
	}
	v15 = v17
	goto _16
_16:
	v18 = v9.Y
	v19 = v10.Y
	if v18 < v19 {
		v22 = v18
	} else {
		v22 = v19
	}
	v20 = v22
	goto _21
_21:
	v23 = v9.Z
	v24 = v10.Z
	if v23 < v24 {
		v27 = v23
	} else {
		v27 = v24
	}
	v25 = v27
	goto _26
_26:
	v11 = Vec3{
		X: v15,
		Y: v20,
		Z: v25,
	}
	goto _12
_12:
	v28 = scaledLower
	v29 = scaledUpper
	v32 = v28.X
	v33 = v29.X
	if v32 > v33 {
		v36 = v32
	} else {
		v36 = v33
	}
	v34 = v36
	goto _35
_35:
	v37 = v28.Y
	v38 = v29.Y
	if v37 > v38 {
		v41 = v37
	} else {
		v41 = v38
	}
	v39 = v41
	goto _40
_40:
	v42 = v28.Z
	v43 = v29.Z
	if v42 > v43 {
		v46 = v42
	} else {
		v46 = v43
	}
	v44 = v46
	goto _45
_45:
	v30 = Vec3{
		X: v34,
		Y: v39,
		Z: v44,
	}
	goto _31
_31:
	bounds = AABB{
		LowerBound: v11,
		UpperBound: v30,
	}
	v47 = transform1
	v48 = bounds
	v49 = v48
	v52 = v49.UpperBound
	v53 = v49.LowerBound
	v54 = Vec3{
		X: v52.X + v53.X,
		Y: v52.Y + v53.Y,
		Z: v52.Z + v53.Z,
	}
	goto _55
_55:
	v56 = float32FromFloat32(0.5)
	v57 = v54
	v58 = Vec3{
		X: float32(v56 * v57.X),
		Y: float32(v56 * v57.Y),
		Z: float32(v56 * v57.Z),
	}
	goto _59
_59:
	v50 = v58
	goto _51
_51:
	v60 = v47
	v61 = v60.Q
	v62 = v50
	v63 = v61.V
	v64 = v62
	c.X = float32(v63.Y*v64.Z) - float32(v63.Z*v64.Y)
	c.Y = float32(v63.Z*v64.X) - float32(v63.X*v64.Z)
	c.Z = float32(v63.X*v64.Y) - float32(v63.Y*v64.X)
	v65 = c
	goto _66
_66:
	t1 = v65
	v67 = t1
	v68 = v61.S
	v69 = v62
	v70 = Vec3{
		X: v67.X + float32(v68*v69.X),
		Y: v67.Y + float32(v68*v69.Y),
		Z: v67.Z + float32(v68*v69.Z),
	}
	goto _71
_71:
	t2 = v70
	v72 = v61.V
	v73 = t2
	c.X = float32(v72.Y*v73.Z) - float32(v72.Z*v73.Y)
	c.Y = float32(v72.Z*v73.X) - float32(v72.X*v73.Z)
	c.Z = float32(v72.X*v73.Y) - float32(v72.Y*v73.X)
	v74 = c
	goto _75
_75:
	t3 = v74
	v78 = v62
	v79 = float32FromFloat32(2)
	v80 = t3
	v81 = Vec3{
		X: v78.X + float32(v79*v80.X),
		Y: v78.Y + float32(v79*v80.Y),
		Z: v78.Z + float32(v79*v80.Z),
	}
	goto _82
_82:
	v76 = v81
	goto _77
_77:
	rv = v76
	v85 = rv
	v86 = v60.P
	v87 = Vec3{
		X: v85.X + v86.X,
		Y: v85.Y + v86.Y,
		Z: v85.Z + v86.Z,
	}
	goto _88
_88:
	v83 = v87
	goto _84
_84:
	center = v83
	v89 = v47.Q
	xx = float32(v89.V.X * v89.V.X)
	yy = float32(v89.V.Y * v89.V.Y)
	zz = float32(v89.V.Z * v89.V.Z)
	xy = float32(v89.V.X * v89.V.Y)
	xz = float32(v89.V.X * v89.V.Z)
	xw = float32(v89.V.X * v89.S)
	yz = float32(v89.V.Y * v89.V.Z)
	yw = float32(v89.V.Y * v89.S)
	zw = float32(v89.V.Z * v89.S)
	v90 = Matrix3{
		Cx: Vec3{
			X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
			Y: float32(float32FromFloat32(2) * (xy + zw)),
			Z: float32(float32FromFloat32(2) * (xz - yw)),
		},
		Cy: Vec3{
			X: float32(float32FromFloat32(2) * (xy - zw)),
			Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
			Z: float32(float32FromFloat32(2) * (yz + xw)),
		},
		Cz: Vec3{
			X: float32(float32FromFloat32(2) * (xz + yw)),
			Y: float32(float32FromFloat32(2) * (yz - xw)),
			Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
		},
	}
	goto _91
_91:
	m2 = v90
	v92 = m2
	v93 = v92.Cx
	v96 = v93.X
	if v96 < float32FromInt32(0) {
		v99 = -v96
	} else {
		v99 = v96
	}
	v97 = v99
	goto _98
_98:
	v100 = v93.Y
	if v100 < float32FromInt32(0) {
		v103 = -v100
	} else {
		v103 = v100
	}
	v101 = v103
	goto _102
_102:
	v104 = v93.Z
	if v104 < float32FromInt32(0) {
		v107 = -v104
	} else {
		v107 = v104
	}
	v105 = v107
	goto _106
_106:
	v94 = Vec3{
		X: v97,
		Y: v101,
		Z: v105,
	}
	goto _95
_95:
	out.Cx = v94
	v108 = v92.Cy
	v111 = v108.X
	if v111 < float32FromInt32(0) {
		v114 = -v111
	} else {
		v114 = v111
	}
	v112 = v114
	goto _113
_113:
	v115 = v108.Y
	if v115 < float32FromInt32(0) {
		v118 = -v115
	} else {
		v118 = v115
	}
	v116 = v118
	goto _117
_117:
	v119 = v108.Z
	if v119 < float32FromInt32(0) {
		v122 = -v119
	} else {
		v122 = v119
	}
	v120 = v122
	goto _121
_121:
	v109 = Vec3{
		X: v112,
		Y: v116,
		Z: v120,
	}
	goto _110
_110:
	out.Cy = v109
	v123 = v92.Cz
	v126 = v123.X
	if v126 < float32FromInt32(0) {
		v129 = -v126
	} else {
		v129 = v126
	}
	v127 = v129
	goto _128
_128:
	v130 = v123.Y
	if v130 < float32FromInt32(0) {
		v133 = -v130
	} else {
		v133 = v130
	}
	v131 = v133
	goto _132
_132:
	v134 = v123.Z
	if v134 < float32FromInt32(0) {
		v137 = -v134
	} else {
		v137 = v134
	}
	v135 = v137
	goto _136
_136:
	v124 = Vec3{
		X: v127,
		Y: v131,
		Z: v135,
	}
	goto _125
_125:
	out.Cz = v124
	v138 = out
	goto _139
_139:
	v140 = v48
	v143 = v140.UpperBound
	v144 = v140.LowerBound
	v145 = Vec3{
		X: v143.X - v144.X,
		Y: v143.Y - v144.Y,
		Z: v143.Z - v144.Z,
	}
	goto _146
_146:
	v147 = float32FromFloat32(0.5)
	v148 = v145
	v149 = Vec3{
		X: float32(v147 * v148.X),
		Y: float32(v147 * v148.Y),
		Z: float32(v147 * v148.Z),
	}
	goto _150
_150:
	v141 = v149
	goto _142
_142:
	v151 = v138
	v152 = v141
	b9 = Vec3{
		X: float32(v151.Cx.X*v152.X) + float32(v151.Cy.X*v152.Y) + float32(v151.Cz.X*v152.Z),
		Y: float32(v151.Cx.Y*v152.X) + float32(v151.Cy.Y*v152.Y) + float32(v151.Cz.Y*v152.Z),
		Z: float32(v151.Cx.Z*v152.X) + float32(v151.Cy.Z*v152.Y) + float32(v151.Cz.Z*v152.Z),
	}
	v153 = b9
	goto _154
_154:
	extent = v153
	v155 = center
	v156 = extent
	v157 = Vec3{
		X: v155.X - v156.X,
		Y: v155.Y - v156.Y,
		Z: v155.Z - v156.Z,
	}
	goto _158
_158:
	v159 = center
	v160 = extent
	v161 = Vec3{
		X: v159.X + v160.X,
		Y: v159.Y + v160.Y,
		Z: v159.Z + v160.Z,
	}
	goto _162
_162:
	out1 = AABB{
		LowerBound: v157,
		UpperBound: v161,
	}
	v163 = out1
	goto _164
_164:
	return v163
}

func b3ComputeMeshManifolds(tls *_Stack, world2 uintptr, workerIndex int32, contact uintptr, shapeA uintptr, materialMap uintptr, xfA WorldTransform, shapeB uintptr, xfB WorldTransform, isFast uint8, _arena b3Arena) (r uint8) {
	bp := tls.Alloc(1200)
	defer tls.Free(1200)
	*(*b3Arena)(unsafe.Pointer(bp + 56)) = _arena
	var C, transformAtoB, v61 Transform
	var acceptedManifoldCount, aligned, baseIndex, bestIndex, clusterCount, clusterIndex, clusterIndex1, clusterPointCount, currentCount, i, i1, i2, i3, i4, i5, i6, i7, i8, i9, index, index1, index2, index3, j, j1, j2, j3, j4, k, k1, k2, lane, laneCount, manifoldCount, manifoldPointCount, materialIndex, newCapacity, newSize, oldManifoldCount, oldPointCount, oldSize, pointBufferCapacity, pointCapacity, pointCount, pointCount1, pointOffset, reducedCount, survivorCount, survivorIndex, tentativeManifoldCount, tentativeTriangleCount, totalPointCount, triangleCount, triangleFlags1, triangleIndex1, triangleIndex2, triangleIndex3, v12, v4, v67, v71, v77, v8, v81, v92 int32
	var acceptedManifolds, allocator1, allocator2, am, ap, cache, cluster, clusterMemberships, clusterPoints, clusters, cm, cm1, cm2, consumed, context, cp, flags, hullB, m1, m2, manifold, manifold1, manifold2, manifold3, manifoldBuffer, manifolds, matchedManifold, materialB, materialIndices, materialsA, meshContact, meshTriangles, meshVertices, oldManifolds, oldPt, pointBuffer, source, survivor, survivorIndices, survivorVertices, target, tentativeManifolds, tentativeTriangles, triangleCaches, triangles, vertices, vertices1, v11, v13, v3, v5, v66, v68, v7, v70, v72, v74, v9 uintptr
	var addedEdge1, addedEdge2, addedEdge3, addedVertex1, addedVertex2, addedVertex3, allowClustering, enableSpeculative, meshFlipped, shouldCollide uint8
	var allocator b3BlockAllocator
	var ax, ay, az, cxx, cxy, cxz, cyx, cyy, cyz, czx, czy, czz, hullLowerX, hullLowerY, hullLowerZ, hullUpperX, hullUpperY, hullUpperZ, lowerX, lowerY, lowerZ, marginW, px, py, pz, separated, upperX, upperY, upperZ b3FloatW4
	var b10, c, clusterNormal, d, hullMargin, manifoldNormal, meshScale, normal, offsetA, scale, t1, t11, t12, t2, t21, t22, t3, t31, t32, tangentVelocityA, tangentVelocityB, triangleNormal, u, v19, v20, v21, v23, v25, v26, v28, v30, v31, v33, v34, v40, v41, v42, v43, v45, v47, v48, v50, v51, v52, v54, v56, v58, v59 Vec3
	var bestDot, clusterThreshold, cosManifoldAngle, cosNormalAngle, cosTriangleAngle, dot, friction, hullSpeculativeDistance, invCount, lengthSquared, linearSlop, minSeparation, normalMatchTolerance, radiusB, restOffset, restitution, s, sampleCount, xw, xx, xy, xz, yw, yy, yz, zw, zz, v24, v29, v35, v46, v57, v75 float32
	var cullMask, keepMask, lane1, lane2, v119 uint32_t
	var feature TriangleFeature
	var hullBounds, v90 AABB
	var material SurfaceMaterial
	var matrixB, relativeMatrix, v64 Matrix3
	var meshTriangle, triangle MeshTriangle
	var q2, v17, v18, v37, v39, v63 Quat
	var q_i, q_i1, q_j, q_j1, q_l, q_m, q_r, q_sp uint64
	var q_st [32]struct {
		Q_l uint64
		Q_r uint64
	}
	var tentativeTriangle, tentativeTriangle1, tmp, tmp1, tmp2, tmp3, tmp4, tmp5, tmp6, tmp7, tmp8 b3TentativeTriangle
	var triangle2, v160 b3Triangle
	var triangleFlags uint8_t
	var vx, vy, vz [3]b3FloatW4
	var v1 bool
	var v15, v16 WorldTransform
	var v269, v270 Pos
	var v297 FeaturePair
	var _ /* foundEdges at bp+80 */ b3FoundEdges
	var _ /* foundVertices at bp+600 */ b3FoundVertices
	var _ /* lanes at bp+860 */ [3][4]Vec3
	var _ /* result at bp+0 */ b3Triangle
	var _ /* triangle at bp+1004 */ b3Triangle
	var _ /* x at bp+1056 */ [3][4]float32
	var _ /* y at bp+1104 */ [3][4]float32
	var _ /* z at bp+1152 */ [3][4]float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = C, acceptedManifoldCount, acceptedManifolds, addedEdge1, addedEdge2, addedEdge3, addedVertex1, addedVertex2, addedVertex3, aligned, allocator, allocator1, allocator2, allowClustering, am, ap, ax, ay, az, b10, baseIndex, bestDot, bestIndex, c, cache, cluster, clusterCount, clusterIndex, clusterIndex1, clusterMemberships, clusterNormal, clusterPointCount, clusterPoints, clusterThreshold, clusters, cm, cm1, cm2, consumed, context, cosManifoldAngle, cosNormalAngle, cosTriangleAngle, cp, cullMask, currentCount, cxx, cxy, cxz, cyx, cyy, cyz, czx, czy, czz, d, dot, enableSpeculative, feature, flags, friction, hullB, hullBounds, hullLowerX, hullLowerY, hullLowerZ, hullMargin, hullSpeculativeDistance, hullUpperX, hullUpperY, hullUpperZ, i, i1, i2, i3, i4, i5, i6, i7, i8, i9, index, index1, index2, index3, invCount, j, j1, j2, j3, j4, k, k1, k2, keepMask, lane, lane1, lane2, laneCount, lengthSquared, linearSlop, lowerX, lowerY, lowerZ, m1, m2, manifold, manifold1, manifold2, manifold3, manifoldBuffer, manifoldCount, manifoldNormal, manifoldPointCount, manifolds, marginW, matchedManifold, material, materialB, materialIndex, materialIndices, materialsA, matrixB, meshContact, meshFlipped, meshScale, meshTriangle, meshTriangles, meshVertices, minSeparation, newCapacity, newSize, normal, normalMatchTolerance, offsetA, oldManifoldCount, oldManifolds, oldPointCount, oldPt, oldSize, pointBuffer, pointBufferCapacity, pointCapacity, pointCount, pointCount1, pointOffset, px, py, pz, q2, q_i, q_i1, q_j, q_j1, q_l, q_m, q_r, q_sp, q_st, radiusB, reducedCount, relativeMatrix, restOffset, restitution, s, sampleCount, scale, separated, shouldCollide, source, survivor, survivorCount, survivorIndex, survivorIndices, survivorVertices, t1, t11, t12, t2, t21, t22, t3, t31, t32, tangentVelocityA, tangentVelocityB, target, tentativeManifoldCount, tentativeManifolds, tentativeTriangle, tentativeTriangle1, tentativeTriangleCount, tentativeTriangles, tmp, tmp1, tmp2, tmp3, tmp4, tmp5, tmp6, tmp7, tmp8, totalPointCount, transformAtoB, triangle, triangle2, triangleCaches, triangleCount, triangleFlags, triangleFlags1, triangleIndex1, triangleIndex2, triangleIndex3, triangleNormal, triangles, u, upperX, upperY, upperZ, vertices, vertices1, vx, vy, vz, xw, xx, xy, xz, yw, yy, yz, zw, zz, v1, v11, v119, v12, v13, v15, v16, v160, v17, v18, v19, v20, v21, v23, v24, v25, v26, v269, v270, v28, v29, v297, v3, v30, v31, v33, v34, v35, v37, v39, v4, v40, v41, v42, v43, v45, v46, v47, v48, v5, v50, v51, v52, v54, v56, v57, v58, v59, v61, v63, v64, v66, v67, v68, v7, v70, v71, v72, v74, v75, v77, v8, v81, v9, v90, v92
	if v1 = !!((*b3Shape)(unsafe.Pointer(shapeA)).Type1 == int32(b3_meshShape) || (*b3Shape)(unsafe.Pointer(shapeA)).Type1 == int32(b3_heightShape)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+22122, __ccgo_ts+22185, int32FromInt32(532))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	_ = uint64FromInt64(4)
	_ = uint64FromInt64(4)
	_ = uint64FromInt64(4)
	if v1 = !!(0 <= workerIndex && workerIndex < (*b3World)(unsafe.Pointer(world2)).TaskContexts.Count); !v1 {
		b3InternalAssert(tls, __ccgo_ts+22536, __ccgo_ts+22185, int32FromInt32(537))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	context = (*b3World)(unsafe.Pointer(world2)).TaskContexts.Data + uintptr(workerIndex)*3848
	b3RefreshCache(tls, contact, shapeA, xfA, shapeB+40, *(*b3Arena)(unsafe.Add(nil, bp+56)))
	// Collide with triangles and build manifolds
	meshContact = contact + 152
	triangleCount = (*b3MeshContact)(unsafe.Pointer(meshContact)).TriangleCache.Count
	v3 = bp + 56
	v4 = int32FromUint64(uint64FromInt32(triangleCount) * uint64(8))
	if v4 == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _6
	}
	aligned = ((*b3Arena)(unsafe.Pointer(v3)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
	if aligned+v4 > (*b3Arena)(unsafe.Pointer(v3)).Capacity {
		v5 = b3ArenaOverflowAlloc(tls, v3, v4)
		goto _6
	}
	(*b3Arena)(unsafe.Pointer(v3)).Index = aligned + v4
	if (*b3Arena)(unsafe.Pointer(v3)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex {
		(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v3)).Index
	}
	v5 = (*b3Arena)(unsafe.Pointer(v3)).Memory + uintptr(aligned)
	goto _6
_6:
	acceptedManifolds = v5
	acceptedManifoldCount = 0
	v7 = bp + 56
	v8 = int32FromUint64(uint64FromInt32(triangleCount) * uint64(8))
	if v8 == int32FromInt32(0) {
		v9 = uintptrFromInt32(0)
		goto _10
	}
	aligned = ((*b3Arena)(unsafe.Pointer(v7)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
	if aligned+v8 > (*b3Arena)(unsafe.Pointer(v7)).Capacity {
		v9 = b3ArenaOverflowAlloc(tls, v7, v8)
		goto _10
	}
	(*b3Arena)(unsafe.Pointer(v7)).Index = aligned + v8
	if (*b3Arena)(unsafe.Pointer(v7)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v7)).Shared)).MaxIndex {
		(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v7)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v7)).Index
	}
	v9 = (*b3Arena)(unsafe.Pointer(v7)).Memory + uintptr(aligned)
	goto _10
_10:
	tentativeManifolds = v9
	tentativeManifoldCount = 0
	v11 = bp + 56
	v12 = int32FromUint64(uint64FromInt32(triangleCount) * uint64(8))
	if v12 == int32FromInt32(0) {
		v13 = uintptrFromInt32(0)
		goto _14
	}
	aligned = ((*b3Arena)(unsafe.Pointer(v11)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
	if aligned+v12 > (*b3Arena)(unsafe.Pointer(v11)).Capacity {
		v13 = b3ArenaOverflowAlloc(tls, v11, v12)
		goto _14
	}
	(*b3Arena)(unsafe.Pointer(v11)).Index = aligned + v12
	if (*b3Arena)(unsafe.Pointer(v11)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v11)).Shared)).MaxIndex {
		(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v11)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v11)).Index
	}
	v13 = (*b3Arena)(unsafe.Pointer(v11)).Memory + uintptr(aligned)
	goto _14
_14:
	tentativeTriangles = v13
	tentativeTriangleCount = 0
	(*(*b3FoundEdges)(unsafe.Add(nil, bp+80))).Count = 0
	(*(*b3FoundVertices)(unsafe.Add(nil, bp+600))).Count = 0
	v15 = xfB
	v16 = xfA
	v17 = v15.Q
	v18 = v16.Q
	v19 = v18.V
	v20 = v17.V
	c.X = float32(v19.Y*v20.Z) - float32(v19.Z*v20.Y)
	c.Y = float32(v19.Z*v20.X) - float32(v19.X*v20.Z)
	c.Z = float32(v19.X*v20.Y) - float32(v19.Y*v20.X)
	v21 = c
	goto _22
_22:
	t12 = v21
	v23 = t12
	v24 = v17.S
	v25 = v18.V
	v26 = Vec3{
		X: v23.X + float32(v24*v25.X),
		Y: v23.Y + float32(v24*v25.Y),
		Z: v23.Z + float32(v24*v25.Z),
	}
	goto _27
_27:
	t22 = v26
	v28 = t22
	v29 = v18.S
	v30 = v17.V
	v31 = Vec3{
		X: v28.X - float32(v29*v30.X),
		Y: v28.Y - float32(v29*v30.Y),
		Z: v28.Z - float32(v29*v30.Z),
	}
	goto _32
_32:
	t32 = v31
	v33 = v17.V
	v34 = v18.V
	v35 = float32(v33.X*v34.X) + float32(v33.Y*v34.Y) + float32(v33.Z*v34.Z)
	goto _36
_36:
	q2 = Quat{
		V: t32,
		S: float32(v17.S*v18.S) + v35,
	}
	v37 = q2
	goto _38
_38:
	C.Q = v37
	d = Vec3{
		X: v16.P.X - v15.P.X,
		Y: v16.P.Y - v15.P.Y,
		Z: v16.P.Z - v15.P.Z,
	}
	v39 = v15.Q
	v40 = d
	v41 = v39.V
	v42 = v40
	c.X = float32(v41.Y*v42.Z) - float32(v41.Z*v42.Y)
	c.Y = float32(v41.Z*v42.X) - float32(v41.X*v42.Z)
	c.Z = float32(v41.X*v42.Y) - float32(v41.Y*v42.X)
	v43 = c
	goto _44
_44:
	t11 = v43
	v45 = t11
	v46 = v39.S
	v47 = v40
	v48 = Vec3{
		X: v45.X - float32(v46*v47.X),
		Y: v45.Y - float32(v46*v47.Y),
		Z: v45.Z - float32(v46*v47.Z),
	}
	goto _49
_49:
	t21 = v48
	v50 = v39.V
	v51 = t21
	c.X = float32(v50.Y*v51.Z) - float32(v50.Z*v51.Y)
	c.Y = float32(v50.Z*v51.X) - float32(v50.X*v51.Z)
	c.Z = float32(v50.X*v51.Y) - float32(v50.Y*v51.X)
	v52 = c
	goto _53
_53:
	t31 = v52
	v56 = v40
	v57 = float32FromFloat32(2)
	v58 = t31
	v59 = Vec3{
		X: v56.X + float32(v57*v58.X),
		Y: v56.Y + float32(v57*v58.Y),
		Z: v56.Z + float32(v57*v58.Z),
	}
	goto _60
_60:
	v54 = v59
	goto _55
_55:
	C.P = v54
	v61 = C
	goto _62
_62:
	// This transform converts from mesh frame into the shapeB frame
	transformAtoB = v61
	v63 = transformAtoB.Q
	xx = float32(v63.V.X * v63.V.X)
	yy = float32(v63.V.Y * v63.V.Y)
	zz = float32(v63.V.Z * v63.V.Z)
	xy = float32(v63.V.X * v63.V.Y)
	xz = float32(v63.V.X * v63.V.Z)
	xw = float32(v63.V.X * v63.S)
	yz = float32(v63.V.Y * v63.V.Z)
	yw = float32(v63.V.Y * v63.S)
	zw = float32(v63.V.Z * v63.S)
	v64 = Matrix3{
		Cx: Vec3{
			X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
			Y: float32(float32FromFloat32(2) * (xy + zw)),
			Z: float32(float32FromFloat32(2) * (xz - yw)),
		},
		Cy: Vec3{
			X: float32(float32FromFloat32(2) * (xy - zw)),
			Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
			Z: float32(float32FromFloat32(2) * (yz + xw)),
		},
		Cz: Vec3{
			X: float32(float32FromFloat32(2) * (xz + yw)),
			Y: float32(float32FromFloat32(2) * (yz - xw)),
			Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
		},
	}
	goto _65
_65:
	relativeMatrix = v64
	linearSlop = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	// This should push apart shapes after a time of impact event.
	// In the past I've called this `polygon skin`, but PhysX and Unreal
	// call it `rest offset` which seems appropriate in this case.
	// It leads to a small visual gap but seems to improve the quality of mesh
	// collision, especially for hull versus mesh.
	restOffset = float32(float32FromFloat32(1) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	enableSpeculative = uint8(boolUint32((*b3Contact)(unsafe.Pointer(contact)).Flags&uint32(b3_enableSpeculativePoints) != 0))
	// Make room for clip points
	pointBufferCapacity = int32(_B3_MAX_POINTS_PER_TRIANGLE) * triangleCount
	v3 = bp + 56
	v4 = int32FromUint64(uint64FromInt32(pointBufferCapacity) * uint64(24))
	if v4 == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _69
	}
	aligned = ((*b3Arena)(unsafe.Pointer(v3)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
	if aligned+v4 > (*b3Arena)(unsafe.Pointer(v3)).Capacity {
		v5 = b3ArenaOverflowAlloc(tls, v3, v4)
		goto _69
	}
	(*b3Arena)(unsafe.Pointer(v3)).Index = aligned + v4
	if (*b3Arena)(unsafe.Pointer(v3)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex {
		(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v3)).Index
	}
	v5 = (*b3Arena)(unsafe.Pointer(v3)).Memory + uintptr(aligned)
	goto _69
_69:
	pointBuffer = v5
	totalPointCount = 0
	v7 = bp + 56
	v8 = int32FromUint64(uint64FromInt32(triangleCount) * uint64(64))
	if v8 == int32FromInt32(0) {
		v9 = uintptrFromInt32(0)
		goto _73
	}
	aligned = ((*b3Arena)(unsafe.Pointer(v7)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
	if aligned+v8 > (*b3Arena)(unsafe.Pointer(v7)).Capacity {
		v9 = b3ArenaOverflowAlloc(tls, v7, v8)
		goto _73
	}
	(*b3Arena)(unsafe.Pointer(v7)).Index = aligned + v8
	if (*b3Arena)(unsafe.Pointer(v7)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v7)).Shared)).MaxIndex {
		(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v7)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v7)).Index
	}
	v9 = (*b3Arena)(unsafe.Pointer(v7)).Memory + uintptr(aligned)
	goto _73
_73:
	manifoldBuffer = v9
	manifoldCount = 0
	triangleCaches = (*b3MeshContact)(unsafe.Pointer(meshContact)).TriangleCache.Data
	if (*b3Shape)(unsafe.Pointer(shapeB)).Type1 == int32(b3_hullShape) {
		v11 = *(*uintptr)(unsafe.Add(unsafe.Pointer(shapeB), 176))
	} else {
		v11 = uintptrFromInt32(0)
	}
	hullB = v11
	if enableSpeculative != 0 {
		v75 = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	} else {
		v75 = float32FromFloat32(0)
	}
	hullSpeculativeDistance = v75
	hullMargin = Vec3{
		X: hullSpeculativeDistance,
		Y: hullSpeculativeDistance,
		Z: hullSpeculativeDistance,
	}
	v13 = bp + 56
	v12 = int32FromUint64(uint64FromInt32(triangleCount) * uint64(4))
	if v12 == int32FromInt32(0) {
		v66 = uintptrFromInt32(0)
		goto _79
	}
	aligned = ((*b3Arena)(unsafe.Pointer(v13)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
	if aligned+v12 > (*b3Arena)(unsafe.Pointer(v13)).Capacity {
		v66 = b3ArenaOverflowAlloc(tls, v13, v12)
		goto _79
	}
	(*b3Arena)(unsafe.Pointer(v13)).Index = aligned + v12
	if (*b3Arena)(unsafe.Pointer(v13)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v13)).Shared)).MaxIndex {
		(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v13)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v13)).Index
	}
	v66 = (*b3Arena)(unsafe.Pointer(v13)).Memory + uintptr(aligned)
	goto _79
_79:
	survivorIndices = v66
	v68 = bp + 56
	v67 = int32FromUint64(uint64FromInt32(int32(3)*triangleCount) * uint64(12))
	if v67 == int32FromInt32(0) {
		v70 = uintptrFromInt32(0)
		goto _83
	}
	aligned = ((*b3Arena)(unsafe.Pointer(v68)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
	if aligned+v67 > (*b3Arena)(unsafe.Pointer(v68)).Capacity {
		v70 = b3ArenaOverflowAlloc(tls, v68, v67)
		goto _83
	}
	(*b3Arena)(unsafe.Pointer(v68)).Index = aligned + v67
	if (*b3Arena)(unsafe.Pointer(v68)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v68)).Shared)).MaxIndex {
		(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v68)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v68)).Index
	}
	v70 = (*b3Arena)(unsafe.Pointer(v68)).Memory + uintptr(aligned)
	goto _83
_83:
	survivorVertices = v70
	survivorCount = 0
	meshTriangles = uintptrFromInt32(0)
	meshVertices = uintptrFromInt32(0)
	meshScale = b3Vec3_one
	meshFlipped = boolUint8(false1 != 0)
	if (*b3Shape)(unsafe.Pointer(shapeA)).Type1 == int32(b3_meshShape) {
		v72 = uintptr(unsafe.Pointer((*(*Mesh)(unsafe.Add(unsafe.Pointer(shapeA), 176))).Data))
		if (*MeshData)(unsafe.Pointer(v72)).TriangleOffset == int32FromInt32(0) {
			v74 = uintptrFromInt32(0)
			goto _86
		}
		v74 = uintptr(int64(v72) + int64((*MeshData)(unsafe.Pointer(v72)).TriangleOffset))
		goto _86
	_86:
		meshTriangles = v74
		v3 = uintptr(unsafe.Pointer((*(*Mesh)(unsafe.Add(unsafe.Pointer(shapeA), 176))).Data))
		if (*MeshData)(unsafe.Pointer(v3)).VertexOffset == int32FromInt32(0) {
			v5 = uintptrFromInt32(0)
			goto _89
		}
		v5 = uintptr(int64(v3) + int64((*MeshData)(unsafe.Pointer(v3)).VertexOffset))
		goto _89
	_89:
		meshVertices = v5
		meshScale = (*(*Mesh)(unsafe.Add(unsafe.Pointer(shapeA), 176))).Scale
		meshFlipped = boolUint8(float32(float32(meshScale.X*meshScale.Y)*meshScale.Z) < float32FromFloat32(0))
	}
	cxx = b3SplatW4(tls, relativeMatrix.Cx.X)
	cxy = b3SplatW4(tls, relativeMatrix.Cx.Y)
	cxz = b3SplatW4(tls, relativeMatrix.Cx.Z)
	cyx = b3SplatW4(tls, relativeMatrix.Cy.X)
	cyy = b3SplatW4(tls, relativeMatrix.Cy.Y)
	cyz = b3SplatW4(tls, relativeMatrix.Cy.Z)
	czx = b3SplatW4(tls, relativeMatrix.Cz.X)
	czy = b3SplatW4(tls, relativeMatrix.Cz.Y)
	czz = b3SplatW4(tls, relativeMatrix.Cz.Z)
	px = b3SplatW4(tls, transformAtoB.P.X)
	py = b3SplatW4(tls, transformAtoB.P.Y)
	pz = b3SplatW4(tls, transformAtoB.P.Z)
	marginW = b3SplatW4(tls, hullMargin.X)
	if hullB != uintptrFromInt32(0) {
		v90 = (*HullData)(unsafe.Pointer(hullB)).Aabb
	} else {
		v90 = AABB{}
	}
	hullBounds = v90
	hullLowerX = b3SplatW4(tls, hullBounds.LowerBound.X)
	hullLowerY = b3SplatW4(tls, hullBounds.LowerBound.Y)
	hullLowerZ = b3SplatW4(tls, hullBounds.LowerBound.Z)
	hullUpperX = b3SplatW4(tls, hullBounds.UpperBound.X)
	hullUpperY = b3SplatW4(tls, hullBounds.UpperBound.Y)
	hullUpperZ = b3SplatW4(tls, hullBounds.UpperBound.Z)
	baseIndex = 0
	for {
		if !(baseIndex < triangleCount) {
			break
		}
		v4 = int32(4)
		v8 = triangleCount - baseIndex
		if v4 < v8 {
			v67 = v4
		} else {
			v67 = v8
		}
		v12 = v67
		goto _95
	_95:
		laneCount = v12
		lane = 0
		for {
			if !(lane < int32(4)) {
				break
			}
			v71 = lane
			v77 = laneCount - int32(1)
			if v71 < v77 {
				v92 = v71
			} else {
				v92 = v77
			}
			v81 = v92
			goto _101
		_101:
			index2 = baseIndex + v81
			triangleIndex1 = (*(*b3TriangleCache)(unsafe.Add(nil, triangleCaches+uintptr(index2)*20))).TriangleIndex
			if meshVertices != uintptrFromInt32(0) {
				meshTriangle = *(*MeshTriangle)(unsafe.Add(nil, meshTriangles+uintptr(triangleIndex1)*12))
				v19 = meshScale
				v20 = *(*Vec3)(unsafe.Add(nil, meshVertices+uintptr(meshTriangle.Index1)*12))
				v21 = Vec3{
					X: float32(v19.X * v20.X),
					Y: float32(v19.Y * v20.Y),
					Z: float32(v19.Z * v20.Z),
				}
				goto _106
			_106:
				*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(lane)*12)) = v21
				if meshFlipped != 0 {
					v4 = meshTriangle.Index3
				} else {
					v4 = meshTriangle.Index2
				}
				v19 = meshScale
				v20 = *(*Vec3)(unsafe.Add(nil, meshVertices+uintptr(v4)*12))
				v21 = Vec3{
					X: float32(v19.X * v20.X),
					Y: float32(v19.Y * v20.Y),
					Z: float32(v19.Z * v20.Z),
				}
				goto _111
			_111:
				*(*Vec3)(unsafe.Add(nil, bp+860+1*48+uintptr(lane)*12)) = v21
				if meshFlipped != 0 {
					v4 = meshTriangle.Index2
				} else {
					v4 = meshTriangle.Index3
				}
				v19 = meshScale
				v20 = *(*Vec3)(unsafe.Add(nil, meshVertices+uintptr(v4)*12))
				v21 = Vec3{
					X: float32(v19.X * v20.X),
					Y: float32(v19.Y * v20.Y),
					Z: float32(v19.Z * v20.Z),
				}
				goto _116
			_116:
				*(*Vec3)(unsafe.Add(nil, bp+860+2*48+uintptr(lane)*12)) = v21
			} else {
				if v1 = !!((*b3Shape)(unsafe.Pointer(shapeA)).Type1 == int32(b3_heightShape)); !v1 {
					b3InternalAssert(tls, __ccgo_ts+22372, __ccgo_ts+22185, int32FromInt32(642))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				*(*b3Triangle)(unsafe.Add(nil, bp+1004)) = b3GetHeightFieldTriangle(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shapeA), 176)), triangleIndex1)
				*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(lane)*12)) = *(*Vec3)(unsafe.Add(nil, bp+1004))
				*(*Vec3)(unsafe.Add(nil, bp+860+1*48+uintptr(lane)*12)) = *(*Vec3)(unsafe.Add(nil, bp+1004+1*12))
				*(*Vec3)(unsafe.Add(nil, bp+860+2*48+uintptr(lane)*12)) = *(*Vec3)(unsafe.Add(nil, bp+1004+2*12))
			}
			goto _97
		_97:
			;
			lane = lane + 1
		}
		k = 0
		for {
			if !(k < int32(3)) {
				break
			}
			ax = b3SetW4(tls, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48))).X, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48+1*12))).X, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48+2*12))).X, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48+3*12))).X)
			ay = b3SetW4(tls, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48))).Y, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48+1*12))).Y, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48+2*12))).Y, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48+3*12))).Y)
			az = b3SetW4(tls, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48))).Z, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48+1*12))).Z, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48+2*12))).Z, (*(*Vec3)(unsafe.Add(nil, bp+860+uintptr(k)*48+3*12))).Z)
			vx[k] = b3AddW4(tls, b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, cxx, ax), b3MulW4(tls, cyx, ay)), b3MulW4(tls, czx, az)), px)
			vy[k] = b3AddW4(tls, b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, cxy, ax), b3MulW4(tls, cyy, ay)), b3MulW4(tls, czy, az)), py)
			vz[k] = b3AddW4(tls, b3AddW4(tls, b3AddW4(tls, b3MulW4(tls, cxz, ax), b3MulW4(tls, cyz, ay)), b3MulW4(tls, czz, az)), pz)
			goto _118
		_118:
			;
			k = k + 1
		}
		keepMask = uint32(1)<<laneCount - uint32(1)
		if hullB != uintptrFromInt32(0) {
			lowerX = b3SubW4(tls, b3MinW4(tls, vx[0], b3MinW4(tls, vx[int32(1)], vx[int32(2)])), marginW)
			lowerY = b3SubW4(tls, b3MinW4(tls, vy[0], b3MinW4(tls, vy[int32(1)], vy[int32(2)])), marginW)
			lowerZ = b3SubW4(tls, b3MinW4(tls, vz[0], b3MinW4(tls, vz[int32(1)], vz[int32(2)])), marginW)
			upperX = b3AddW4(tls, b3MaxW4(tls, vx[0], b3MaxW4(tls, vx[int32(1)], vx[int32(2)])), marginW)
			upperY = b3AddW4(tls, b3MaxW4(tls, vy[0], b3MaxW4(tls, vy[int32(1)], vy[int32(2)])), marginW)
			upperZ = b3AddW4(tls, b3MaxW4(tls, vz[0], b3MaxW4(tls, vz[int32(1)], vz[int32(2)])), marginW)
			separated = b3OrW4(tls, b3LessThanW4(tls, upperX, hullLowerX), b3GreaterThanW4(tls, lowerX, hullUpperX))
			separated = b3OrW4(tls, separated, b3OrW4(tls, b3LessThanW4(tls, upperY, hullLowerY), b3GreaterThanW4(tls, lowerY, hullUpperY)))
			separated = b3OrW4(tls, separated, b3OrW4(tls, b3LessThanW4(tls, upperZ, hullLowerZ), b3GreaterThanW4(tls, lowerZ, hullUpperZ)))
			cullMask = keepMask & uint32FromInt32(b3MoveMaskW4(tls, separated))
			keepMask = keepMask & ^cullMask
			for cullMask != uint32(0) {
				v119 = uint32FromInt32(__builtin_ctz(tls, cullMask))
				goto _120
			_120:
				lane1 = v119
				cullMask = cullMask & (cullMask - uint32(1))
				*(*b3SATCache)(unsafe.Pointer(triangleCaches + uintptr(baseIndex+int32FromUint32(lane1))*20 + 4)) = b3SATCache{}
			}
		}
		if keepMask == uint32(0) {
			goto _91
		}
		k1 = 0
		for {
			if !(k1 < int32(3)) {
				break
			}
			b3StoreW4(tls, bp+1056+uintptr(k1)*16, vx[k1])
			b3StoreW4(tls, bp+1104+uintptr(k1)*16, vy[k1])
			b3StoreW4(tls, bp+1152+uintptr(k1)*16, vz[k1])
			goto _121
		_121:
			;
			k1 = k1 + 1
		}
		for keepMask != uint32(0) {
			v119 = uint32FromInt32(__builtin_ctz(tls, keepMask))
			goto _123
		_123:
			lane2 = v119
			keepMask = keepMask & (keepMask - uint32(1))
			survivor = survivorVertices + uintptr(int32(3)*survivorCount)*12
			*(*Vec3)(unsafe.Add(nil, survivor)) = Vec3{
				X: *(*float32)(unsafe.Add(nil, bp+1056+uintptr(lane2)*4)),
				Y: *(*float32)(unsafe.Add(nil, bp+1104+uintptr(lane2)*4)),
				Z: *(*float32)(unsafe.Add(nil, bp+1152+uintptr(lane2)*4)),
			}
			*(*Vec3)(unsafe.Add(nil, survivor+1*12)) = Vec3{
				X: *(*float32)(unsafe.Add(nil, bp+1056+1*16+uintptr(lane2)*4)),
				Y: *(*float32)(unsafe.Add(nil, bp+1104+1*16+uintptr(lane2)*4)),
				Z: *(*float32)(unsafe.Add(nil, bp+1152+1*16+uintptr(lane2)*4)),
			}
			*(*Vec3)(unsafe.Add(nil, survivor+2*12)) = Vec3{
				X: *(*float32)(unsafe.Add(nil, bp+1056+2*16+uintptr(lane2)*4)),
				Y: *(*float32)(unsafe.Add(nil, bp+1104+2*16+uintptr(lane2)*4)),
				Z: *(*float32)(unsafe.Add(nil, bp+1152+2*16+uintptr(lane2)*4)),
			}
			*(*int32)(unsafe.Add(nil, survivorIndices+uintptr(survivorCount)*4)) = baseIndex + int32FromUint32(lane2)
			survivorCount = survivorCount + int32(1)
		}
		goto _91
	_91:
		;
		baseIndex = baseIndex + int32(4)
	}
	survivorIndex = 0
	for {
		if !(survivorIndex < survivorCount && totalPointCount+int32(3) < pointBufferCapacity) {
			break
		}
		index3 = *(*int32)(unsafe.Add(nil, survivorIndices+uintptr(survivorIndex)*4))
		triangleIndex2 = (*(*b3TriangleCache)(unsafe.Add(nil, triangleCaches+uintptr(index3)*20))).TriangleIndex
		vertices1 = survivorVertices + uintptr(int32(3)*survivorIndex)*12
		if (*b3Shape)(unsafe.Pointer(shapeA)).Type1 == int32(b3_meshShape) {
			v3 = shapeA + 176
			v4 = triangleIndex2
			if v1 = !!(0 <= v4 && v4 < (*MeshData)(unsafe.Pointer(uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(v3)).Data)))).TriangleCount); !v1 {
				b3InternalAssert(tls, __ccgo_ts+22608, __ccgo_ts+22672, int32FromInt32(143))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			v5 = uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(v3)).Data))
			if (*MeshData)(unsafe.Pointer(v5)).TriangleOffset == int32FromInt32(0) {
				v7 = uintptrFromInt32(0)
				goto _130
			}
			v7 = uintptr(int64(v5) + int64((*MeshData)(unsafe.Pointer(v5)).TriangleOffset))
			goto _130
		_130:
			triangles = v7
			v9 = uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(v3)).Data))
			if (*MeshData)(unsafe.Pointer(v9)).FlagsOffset == int32FromInt32(0) {
				v11 = uintptrFromInt32(0)
				goto _133
			}
			v11 = uintptr(int64(v9) + int64((*MeshData)(unsafe.Pointer(v9)).FlagsOffset))
			goto _133
		_133:
			flags = v11
			v13 = uintptr(unsafe.Pointer((*Mesh)(unsafe.Pointer(v3)).Data))
			if (*MeshData)(unsafe.Pointer(v13)).VertexOffset == int32FromInt32(0) {
				v66 = uintptrFromInt32(0)
				goto _136
			}
			v66 = uintptr(int64(v13) + int64((*MeshData)(unsafe.Pointer(v13)).VertexOffset))
			goto _136
		_136:
			vertices = v66
			triangle = *(*MeshTriangle)(unsafe.Add(nil, triangles+uintptr(v4)*12))
			triangleFlags = *(*uint8_t)(unsafe.Add(nil, flags+uintptr(v4)))
			scale = (*Mesh)(unsafe.Pointer(v3)).Scale
			v19 = scale
			v20 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index1)*12))
			v21 = Vec3{
				X: float32(v19.X * v20.X),
				Y: float32(v19.Y * v20.Y),
				Z: float32(v19.Z * v20.Z),
			}
			goto _140
		_140:
			*(*Vec3)(unsafe.Add(nil, bp)) = v21
			(*(*b3Triangle)(unsafe.Add(nil, bp))).I1 = triangle.Index1
			if float32(float32(scale.X*scale.Y)*scale.Z) < float32FromFloat32(0) {
				v23 = scale
				v25 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index3)*12))
				v26 = Vec3{
					X: float32(v23.X * v25.X),
					Y: float32(v23.Y * v25.Y),
					Z: float32(v23.Z * v25.Z),
				}
				goto _144
			_144:
				*(*Vec3)(unsafe.Add(nil, bp+1*12)) = v26
				v28 = scale
				v30 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index2)*12))
				v31 = Vec3{
					X: float32(v28.X * v30.X),
					Y: float32(v28.Y * v30.Y),
					Z: float32(v28.Z * v30.Z),
				}
				goto _148
			_148:
				*(*Vec3)(unsafe.Add(nil, bp+2*12)) = v31
				(*(*b3Triangle)(unsafe.Add(nil, bp))).I2 = triangle.Index3
				(*(*b3Triangle)(unsafe.Add(nil, bp))).I3 = triangle.Index2
				(*(*b3Triangle)(unsafe.Add(nil, bp))).Flags = 0
				if int32FromUint8(triangleFlags)&int32(b3_inverseConcaveEdge1) != 0 {
					v8 = int32(b3_concaveEdge1)
				} else {
					v8 = 0
				}
				(*(*b3Triangle)(unsafe.Add(nil, bp))).Flags |= v8
				if int32FromUint8(triangleFlags)&int32(b3_inverseConcaveEdge2) != 0 {
					v12 = int32(b3_concaveEdge2)
				} else {
					v12 = 0
				}
				(*(*b3Triangle)(unsafe.Add(nil, bp))).Flags |= v12
				if int32FromUint8(triangleFlags)&int32(b3_inverseConcaveEdge3) != 0 {
					v67 = int32(b3_concaveEdge3)
				} else {
					v67 = 0
				}
				(*(*b3Triangle)(unsafe.Add(nil, bp))).Flags |= v67
			} else {
				v33 = scale
				v34 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index2)*12))
				v40 = Vec3{
					X: float32(v33.X * v34.X),
					Y: float32(v33.Y * v34.Y),
					Z: float32(v33.Z * v34.Z),
				}
				goto _155
			_155:
				*(*Vec3)(unsafe.Add(nil, bp+1*12)) = v40
				v41 = scale
				v42 = *(*Vec3)(unsafe.Add(nil, vertices+uintptr(triangle.Index3)*12))
				v43 = Vec3{
					X: float32(v41.X * v42.X),
					Y: float32(v41.Y * v42.Y),
					Z: float32(v41.Z * v42.Z),
				}
				goto _159
			_159:
				*(*Vec3)(unsafe.Add(nil, bp+2*12)) = v43
				(*(*b3Triangle)(unsafe.Add(nil, bp))).I2 = triangle.Index2
				(*(*b3Triangle)(unsafe.Add(nil, bp))).I3 = triangle.Index3
				(*(*b3Triangle)(unsafe.Add(nil, bp))).Flags = int32FromUint8(triangleFlags)
			}
			v160 = *(*b3Triangle)(unsafe.Add(nil, bp))
			goto _161
		_161:
			triangle2 = v160
		} else {
			if v1 = !!((*b3Shape)(unsafe.Pointer(shapeA)).Type1 == int32(b3_heightShape)); !v1 {
				b3InternalAssert(tls, __ccgo_ts+22372, __ccgo_ts+22185, int32FromInt32(728))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			triangle2 = b3GetHeightFieldTriangle(tls, *(*uintptr)(unsafe.Add(unsafe.Pointer(shapeA), 176)), triangleIndex2)
		}
		cache = triangleCaches + uintptr(index3)*20 + 4
		pointCapacity = pointBufferCapacity - totalPointCount
		manifold = manifoldBuffer + uintptr(manifoldCount)*64
		(*LocalManifold)(unsafe.Pointer(manifold)).points = pointBuffer + uintptr(totalPointCount)*24
		(*LocalManifold)(unsafe.Pointer(manifold)).PointCount = 0
		(*LocalManifold)(unsafe.Pointer(manifold)).TriangleFlags = triangle2.Flags
		(*LocalManifold)(unsafe.Pointer(manifold)).Feature = int32(b3_featureNone)
		switch (*b3Shape)(unsafe.Pointer(shapeB)).Type1 {
		case int32(b3_capsuleShape):
			b3CollideTriangleAndCapsule(tls, manifold, pointCapacity, vertices1, shapeB+176, cache)
		case int32(b3_hullShape):
			// Cached edge contact is dangerous at high speed because the hull can rotate around the edge and tunnel
			// through the triangle.
			if isFast != 0 && int32FromUint8((*b3ContactCache)(unsafe.Pointer(cache)).SatCache.Type1) == int32(b3_edgePairAxis) {
				(*b3ContactCache)(unsafe.Pointer(cache)).SatCache = b3SATCache{}
			}
			b3CollideTriangleAndHull(tls, manifold, pointCapacity, *(*Vec3)(unsafe.Add(nil, vertices1)), *(*Vec3)(unsafe.Add(nil, vertices1+1*12)), *(*Vec3)(unsafe.Add(nil, vertices1+2*12)), triangle2.Flags, hullB, cache, enableSpeculative)
			*(*int32)(unsafe.Add(nil, context+136)) += int32(1)
			*(*int32)(unsafe.Add(nil, context+140)) += int32FromUint8((*b3ContactCache)(unsafe.Pointer(cache)).SatCache.Hit)
		case int32(b3_sphereShape):
			b3CollideTriangleAndSphere(tls, manifold, pointCapacity, vertices1, shapeB+176)
		default:
			if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
				b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+22185, int32FromInt32(765))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			return boolUint8(false1 != 0)
		}
		manifoldPointCount = (*LocalManifold)(unsafe.Pointer(manifold)).PointCount
		if manifoldPointCount > 0 {
			if v1 = !!((*LocalManifold)(unsafe.Pointer(manifold)).Feature != int32(b3_featureNone)); !v1 {
				b3InternalAssert(tls, __ccgo_ts+22695, __ccgo_ts+22185, int32FromInt32(773))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			manifoldCount = manifoldCount + int32(1)
			totalPointCount = totalPointCount + manifoldPointCount
			(*LocalManifold)(unsafe.Pointer(manifold)).TriangleIndex = triangleIndex2
			v19 = *(*Vec3)(unsafe.Add(nil, vertices1))
			v20 = *(*Vec3)(unsafe.Add(nil, vertices1+1*12))
			v21 = v19
			v23 = Vec3{
				X: v20.X - v21.X,
				Y: v20.Y - v21.Y,
				Z: v20.Z - v21.Z,
			}
			goto _169
		_169:
			v25 = *(*Vec3)(unsafe.Add(nil, vertices1+2*12))
			v26 = v19
			v28 = Vec3{
				X: v25.X - v26.X,
				Y: v25.Y - v26.Y,
				Z: v25.Z - v26.Z,
			}
			goto _173
		_173:
			v30 = v23
			v31 = v28
			c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
			c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
			c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
			v33 = c
			goto _177
		_177:
			normal = v33
			v40 = normal
			lengthSquared = float32(v40.X*v40.X) + float32(v40.Y*v40.Y) + float32(v40.Z*v40.Z)
			if lengthSquared > float32(float32FromFloat32(1000)*float32FromFloat32(1.1754943508222875e-38)) {
				s = float32FromFloat32(1) / sqrtf(tls, lengthSquared)
				u = Vec3{
					X: float32(s * v40.X),
					Y: float32(s * v40.Y),
					Z: float32(s * v40.Z),
				}
				v41 = u
				goto _182
			}
			v41 = Vec3{}
			goto _182
		_182:
			v34 = v41
			goto _179
		_179:
			(*LocalManifold)(unsafe.Pointer(manifold)).TriangleNormal = v34
			(*LocalManifold)(unsafe.Pointer(manifold)).I1 = triangle2.I1
			(*LocalManifold)(unsafe.Pointer(manifold)).I2 = triangle2.I2
			(*LocalManifold)(unsafe.Pointer(manifold)).I3 = triangle2.I3
			if (*LocalManifold)(unsafe.Pointer(manifold)).Feature == int32(b3_featureTriangleFace) || bool(_B3_FORCE_GHOST_COLLISIONS != 0) {
				b3AddEdge(tls, bp+80, triangle2.I1, triangle2.I2)
				b3AddEdge(tls, bp+80, triangle2.I2, triangle2.I3)
				b3AddEdge(tls, bp+80, triangle2.I3, triangle2.I1)
				b3AddVertex(tls, bp+600, triangle2.I1)
				b3AddVertex(tls, bp+600, triangle2.I2)
				b3AddVertex(tls, bp+600, triangle2.I3)
				v4 = acceptedManifoldCount
				acceptedManifoldCount = acceptedManifoldCount + 1
				*(*uintptr)(unsafe.Add(nil, acceptedManifolds+uintptr(v4)*8)) = manifold
			} else {
				if (*LocalManifold)(unsafe.Pointer(manifold)).Feature == int32(b3_featureHullFace) {
					v19 = (*LocalManifold)(unsafe.Pointer(manifold)).TriangleNormal
					v20 = (*LocalManifold)(unsafe.Pointer(manifold)).Normal
					v24 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
					goto _187
				_187:
					cosNormalAngle = v24
					if cosNormalAngle > float32FromFloat32(0.5) {
						b3AddEdge(tls, bp+80, triangle2.I1, triangle2.I2)
						b3AddEdge(tls, bp+80, triangle2.I2, triangle2.I3)
						b3AddEdge(tls, bp+80, triangle2.I3, triangle2.I1)
						b3AddVertex(tls, bp+600, triangle2.I1)
						b3AddVertex(tls, bp+600, triangle2.I2)
						b3AddVertex(tls, bp+600, triangle2.I3)
						v4 = acceptedManifoldCount
						acceptedManifoldCount = acceptedManifoldCount + 1
						*(*uintptr)(unsafe.Add(nil, acceptedManifolds+uintptr(v4)*8)) = manifold
					} else {
						minSeparation = (*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points))).Separation
						i1 = int32(1)
						for {
							if !(i1 < manifoldPointCount) {
								break
							}
							v24 = minSeparation
							v29 = (*(*LocalManifoldPoint)(unsafe.Add(nil, (*LocalManifold)(unsafe.Pointer(manifold)).points+uintptr(i1)*24))).Separation
							if v24 < v29 {
								v46 = v24
							} else {
								v46 = v29
							}
							v35 = v46
							goto _193
						_193:
							minSeparation = v35
							goto _189
						_189:
							;
							i1 = i1 + 1
						}
						if minSeparation < float32(-float32FromFloat32(2)*linearSlop) {
							// Deep overlap
							b3AddEdge(tls, bp+80, triangle2.I1, triangle2.I2)
							b3AddEdge(tls, bp+80, triangle2.I2, triangle2.I3)
							b3AddEdge(tls, bp+80, triangle2.I3, triangle2.I1)
							b3AddVertex(tls, bp+600, triangle2.I1)
							b3AddVertex(tls, bp+600, triangle2.I2)
							b3AddVertex(tls, bp+600, triangle2.I3)
							v4 = acceptedManifoldCount
							acceptedManifoldCount = acceptedManifoldCount + 1
							*(*uintptr)(unsafe.Add(nil, acceptedManifolds+uintptr(v4)*8)) = manifold
						} else {
							tentativeTriangle = b3TentativeTriangle{
								SquaredDistance: (*LocalManifold)(unsafe.Pointer(manifold)).SquaredDistance,
								Index:           tentativeManifoldCount,
							}
							v4 = tentativeTriangleCount
							tentativeTriangleCount = tentativeTriangleCount + 1
							*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(v4)*8)) = tentativeTriangle
							v4 = tentativeManifoldCount
							tentativeManifoldCount = tentativeManifoldCount + 1
							*(*uintptr)(unsafe.Add(nil, tentativeManifolds+uintptr(v4)*8)) = manifold
						}
					}
				} else {
					tentativeTriangle1 = b3TentativeTriangle{
						SquaredDistance: (*LocalManifold)(unsafe.Pointer(manifold)).SquaredDistance,
						Index:           tentativeManifoldCount,
					}
					v4 = tentativeTriangleCount
					tentativeTriangleCount = tentativeTriangleCount + 1
					*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(v4)*8)) = tentativeTriangle1
					v4 = tentativeManifoldCount
					tentativeManifoldCount = tentativeManifoldCount + 1
					*(*uintptr)(unsafe.Add(nil, tentativeManifolds+uintptr(v4)*8)) = manifold
				}
			}
		}
		goto _124
	_124:
		;
		survivorIndex = survivorIndex + 1
	}
	if v1 = !!(acceptedManifoldCount <= triangleCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+22731, __ccgo_ts+22185, int32FromInt32(846))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(tentativeManifoldCount <= triangleCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+22770, __ccgo_ts+22185, int32FromInt32(847))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(tentativeTriangleCount <= triangleCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+22810, __ccgo_ts+22185, int32FromInt32(848))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if (*b3Shape)(unsafe.Pointer(shapeB)).Type1 == int32(b3_sphereShape) {
		// Sort triangles so the closest triangles are processed first
		if tentativeTriangleCount > int32(1) {
			q_l = uint64(0)
			q_r = uint64FromInt32(tentativeTriangleCount - int32(1))
			q_sp = uint64(0)
			for int32(1) != 0 {
				if q_r-q_l+uint64(1) >= uint64(_Q_THRESH) {
					q_m = q_l + (q_r-q_l)>>int32(1)
					if (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8))).SquaredDistance < (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l)+int32(1))*8))).SquaredDistance {
						if (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_r))*8))).SquaredDistance < (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8))).SquaredDistance {
							tmp = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l)+int32(1))*8))
							*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l)+int32(1))*8)) = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_r))*8))
							*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_r))*8)) = tmp
						} else {
							tmp1 = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l)+int32(1))*8))
							*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l)+int32(1))*8)) = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8))
							*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8)) = tmp1
							if (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_r))*8))).SquaredDistance < (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8))).SquaredDistance {
								tmp2 = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8))
								*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8)) = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_r))*8))
								*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_r))*8)) = tmp2
							}
						}
					} else {
						if (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_r))*8))).SquaredDistance < (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8))).SquaredDistance {
							tmp3 = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8))
							*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8)) = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_r))*8))
							*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_r))*8)) = tmp3
							if (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8))).SquaredDistance < (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l)+int32(1))*8))).SquaredDistance {
								tmp4 = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l)+int32(1))*8))
								*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l)+int32(1))*8)) = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8))
								*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8)) = tmp4
							}
						}
					}
					tmp5 = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l))*8))
					*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l))*8)) = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8))
					*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_m))*8)) = tmp5
					q_i = q_l + uint64(1)
					q_j = q_r
					for int32(1) != 0 {
						for cond := true; cond; cond = (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_i))*8))).SquaredDistance < (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l))*8))).SquaredDistance {
							q_i = q_i + 1
						}
						for cond := true; cond; cond = (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l))*8))).SquaredDistance < (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_j))*8))).SquaredDistance {
							q_j = q_j - 1
						}
						if q_i >= q_j {
							break
						}
						tmp6 = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_i))*8))
						*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_i))*8)) = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_j))*8))
						*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_j))*8)) = tmp6
					}
					q_i = q_j + uint64(1)
					tmp7 = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l))*8))
					*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_l))*8)) = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_j))*8))
					*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_j))*8)) = tmp7
					q_j = q_j - 1           /* Now have two subfiles: [q_l,q_j] and [q_i,q_r].			     * Dealing with them depends on which one is bigger. */
					if q_j-q_l >= q_r-q_i { /* If the second subfile is only a single element, it needs		     * no further processing.  The first subfile will be processed	     * on the next iteration (both subfiles cannot be only a single	     * element, due to Q_THRESH). */
						if q_i == q_r {
							q_r = q_j
						} else { /* Otherwise, both subfiles need processing.				 * Push the larger subfile onto the stack. */
							q_st[q_sp].Q_l = q_l
							q_st[q_sp].Q_r = q_j
							q_sp = q_sp + 1 /* Process the smaller subfile on the next iteration. */
							q_l = q_i
						}
					} else { /* If the second subfile is only a single element, it needs		     * no further processing.  The first subfile will be processed	     * on the next iteration (both subfiles cannot be only a single	     * element, due to Q_THRESH). */
						if q_l == q_j {
							q_l = q_i
						} else { /* Otherwise, both subfiles need processing.				 * Push the larger subfile onto the stack. */
							q_st[q_sp].Q_l = q_i
							q_st[q_sp].Q_r = q_r
							q_sp = q_sp + 1 /* Process the smaller subfile on the next iteration. */
							q_r = q_j
						}
					}
				} else { /* For each item starting with the second... */
					q_i1 = q_l + uint64(1)
					for {
						if !(q_i1 <= q_r) {
							break
						} /* move it down the array so that the first part is sorted. */
						q_j1 = q_i1
						for {
							if !(q_j1 > q_l && (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_j1))*8))).SquaredDistance < (*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_j1)-int32(1))*8))).SquaredDistance) {
								break
							}
							tmp8 = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_j1))*8))
							*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_j1))*8)) = *(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_j1)-int32(1))*8))
							*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(int32FromUint64(q_j1)-int32(1))*8)) = tmp8
							goto _204
						_204:
							;
							q_j1 = q_j1 - 1
						}
						goto _203
					_203:
						;
						q_i1 = q_i1 + 1
					} /* Pop subfiles from the stack, until it gets empty. */
					if q_sp == uint64(0) {
						break
					}
					q_sp = q_sp - 1
					q_l = q_st[q_sp].Q_l
					q_r = q_st[q_sp].Q_r
				}
			}
		}
		// Add tentative manifolds in sorted order. Avoid adding manifolds that generate ghost collisions.
		i2 = 0
		for {
			if !(i2 < tentativeTriangleCount) {
				break
			}
			m1 = *(*uintptr)(unsafe.Add(nil, tentativeManifolds+uintptr((*(*b3TentativeTriangle)(unsafe.Add(nil, tentativeTriangles+uintptr(i2)*8))).Index)*8))
			addedEdge1 = b3AddEdge(tls, bp+80, (*LocalManifold)(unsafe.Pointer(m1)).I1, (*LocalManifold)(unsafe.Pointer(m1)).I2)
			addedEdge2 = b3AddEdge(tls, bp+80, (*LocalManifold)(unsafe.Pointer(m1)).I2, (*LocalManifold)(unsafe.Pointer(m1)).I3)
			addedEdge3 = b3AddEdge(tls, bp+80, (*LocalManifold)(unsafe.Pointer(m1)).I3, (*LocalManifold)(unsafe.Pointer(m1)).I1)
			addedVertex1 = b3AddVertex(tls, bp+600, (*LocalManifold)(unsafe.Pointer(m1)).I1)
			addedVertex2 = b3AddVertex(tls, bp+600, (*LocalManifold)(unsafe.Pointer(m1)).I2)
			addedVertex3 = b3AddVertex(tls, bp+600, (*LocalManifold)(unsafe.Pointer(m1)).I3)
			feature = (*LocalManifold)(unsafe.Pointer(m1)).Feature
			shouldCollide = boolUint8(false1 != 0)
			switch feature {
			case int32(b3_featureNone):
				fallthrough
			case int32(b3_featureTriangleFace):
				if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
					b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+22185, int32FromInt32(886))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
			case int32(b3_featureEdge1):
				shouldCollide = addedEdge1
			case int32(b3_featureEdge2):
				shouldCollide = addedEdge2
			case int32(b3_featureEdge3):
				shouldCollide = addedEdge3
			case int32(b3_featureVertex1):
				shouldCollide = addedVertex1
			case int32(b3_featureVertex2):
				shouldCollide = addedVertex2
			case int32(b3_featureVertex3):
				shouldCollide = addedVertex3
			default:
				if v1 = bool(!!(int32FromInt32(false1) != 0)); !v1 {
					b3InternalAssert(tls, __ccgo_ts+47, __ccgo_ts+22185, int32FromInt32(914))
				}
				_ = v1 || bool(int32FromInt32(0) != 0)
				break
			}
			if int32FromUint8(shouldCollide) == int32(true1) {
				v4 = acceptedManifoldCount
				acceptedManifoldCount = acceptedManifoldCount + 1
				*(*uintptr)(unsafe.Add(nil, acceptedManifolds+uintptr(v4)*8)) = m1
			}
			goto _205
		_205:
			;
			i2 = i2 + 1
		}
	} else {
		// Problem: hull can tunnel if time of impact is at concave edge
		// Example: flat box sliding down a ramp to a flat bottom
		// Solution: only ignore flat edges
		i3 = 0
		for {
			if !(i3 < tentativeManifoldCount) {
				break
			}
			m2 = *(*uintptr)(unsafe.Add(nil, tentativeManifolds+uintptr(i3)*8))
			triangleFlags1 = (*LocalManifold)(unsafe.Pointer(m2)).TriangleFlags
			if triangleFlags1&int32(b3_allFlatEdges) == int32(b3_allFlatEdges) {
				goto _209
			}
			if triangleFlags1&int32(b3_flatEdge1) == int32(b3_flatEdge1) {
				if b3FindEdge(tls, bp+80, (*LocalManifold)(unsafe.Pointer(m2)).I1, (*LocalManifold)(unsafe.Pointer(m2)).I2) != 0 {
					goto _209
				}
			}
			if triangleFlags1&int32(b3_flatEdge2) == int32(b3_flatEdge2) {
				if b3FindEdge(tls, bp+80, (*LocalManifold)(unsafe.Pointer(m2)).I2, (*LocalManifold)(unsafe.Pointer(m2)).I3) != 0 {
					goto _209
				}
			}
			if triangleFlags1&int32(b3_flatEdge3) == int32(b3_flatEdge3) {
				if b3FindEdge(tls, bp+80, (*LocalManifold)(unsafe.Pointer(m2)).I3, (*LocalManifold)(unsafe.Pointer(m2)).I1) != 0 {
					goto _209
				}
			}
			v4 = acceptedManifoldCount
			acceptedManifoldCount = acceptedManifoldCount + 1
			*(*uintptr)(unsafe.Add(nil, acceptedManifolds+uintptr(v4)*8)) = m2
			goto _209
		_209:
			;
			i3 = i3 + 1
		}
	}
	if v1 = !!(acceptedManifoldCount <= triangleCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+22731, __ccgo_ts+22185, int32FromInt32(967))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if acceptedManifoldCount == 0 {
		if (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount > 0 {
			v3 = world2
			v4 = (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount
			if v4 == int32FromInt32(0) {
				goto _214
			}
			index1 = v4 - int32(1)
			__builtin_b3LockMutex(tls, (*b3World)(unsafe.Pointer(v3)).ManifoldAllocatorMutex)
			if v1 = !!(0 <= index1 && index1 < (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Count); !v1 {
				b3InternalAssert(tls, __ccgo_ts+8145, __ccgo_ts+8211, int32FromInt32(359))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			allocator2 = (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Data + uintptr(index1)*40
			b3FreeElement(tls, allocator2, (*b3Contact)(unsafe.Pointer(contact)).Manifolds)
			__builtin_b3UnlockMutex(tls, (*b3World)(unsafe.Pointer(v3)).ManifoldAllocatorMutex)
		_214:
			;
			(*b3Contact)(unsafe.Pointer(contact)).Manifolds = uintptrFromInt32(0)
			(*b3Contact)(unsafe.Pointer(contact)).ManifoldCount = 0
		}
		return boolUint8(false1 != 0)
	}
	v3 = bp + 56
	v4 = int32FromUint64(uint64FromInt32(acceptedManifoldCount) * uint64(40))
	if v4 == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _219
	}
	aligned = ((*b3Arena)(unsafe.Pointer(v3)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
	if aligned+v4 > (*b3Arena)(unsafe.Pointer(v3)).Capacity {
		v5 = b3ArenaOverflowAlloc(tls, v3, v4)
		goto _219
	}
	(*b3Arena)(unsafe.Pointer(v3)).Index = aligned + v4
	if (*b3Arena)(unsafe.Pointer(v3)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex {
		(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v3)).Index
	}
	v5 = (*b3Arena)(unsafe.Pointer(v3)).Memory + uintptr(aligned)
	goto _219
_219:
	clusters = v5
	v7 = bp + 56
	v8 = int32FromUint64(uint64FromInt32(acceptedManifoldCount) * uint64(4))
	if v8 == int32FromInt32(0) {
		v9 = uintptrFromInt32(0)
		goto _223
	}
	aligned = ((*b3Arena)(unsafe.Pointer(v7)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
	if aligned+v8 > (*b3Arena)(unsafe.Pointer(v7)).Capacity {
		v9 = b3ArenaOverflowAlloc(tls, v7, v8)
		goto _223
	}
	(*b3Arena)(unsafe.Pointer(v7)).Index = aligned + v8
	if (*b3Arena)(unsafe.Pointer(v7)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v7)).Shared)).MaxIndex {
		(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v7)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v7)).Index
	}
	v9 = (*b3Arena)(unsafe.Pointer(v7)).Memory + uintptr(aligned)
	goto _223
_223:
	clusterMemberships = v9
	// Cluster tolerance is tighter than the warm starting manifold matching tolerance. These
	// serve different purposes.
	clusterThreshold = float32FromFloat32(0.996)
	clusterCount = 0
	clusterPointCount = 0
	i4 = 0
	for {
		if !(i4 < acceptedManifoldCount) {
			break
		}
		*(*int32)(unsafe.Add(nil, clusterMemberships+uintptr(i4)*4)) = -int32(1)
		manifold1 = *(*uintptr)(unsafe.Add(nil, acceptedManifolds+uintptr(i4)*8))
		clusterPointCount = clusterPointCount + (*LocalManifold)(unsafe.Pointer(manifold1)).PointCount
		// Cluster based on the triangle normal and contact normal.
		// The first cluster found is accepted because the tolerance is tight.
		// todo consider requiring the triangles to be connect by an edge.
		// todo consider looking for the best cluster instead of the first one within tolerance
		// This bool is here to allow quick testing with and without clustering.
		allowClustering = boolUint8(true1 != 0)
		manifoldNormal = (*LocalManifold)(unsafe.Pointer(manifold1)).Normal
		triangleNormal = (*LocalManifold)(unsafe.Pointer(manifold1)).TriangleNormal
		clusterIndex = -int32(1)
		j = 0
		for {
			if !(j < clusterCount && allowClustering != 0) {
				break
			}
			v19 = (*(*b3Cluster)(unsafe.Add(nil, clusters+uintptr(j)*40))).ManifoldNormal
			v20 = manifoldNormal
			v24 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
			goto _229
		_229:
			cosManifoldAngle = v24
			v21 = (*(*b3Cluster)(unsafe.Add(nil, clusters+uintptr(j)*40))).TriangleNormal
			v23 = triangleNormal
			v29 = float32(v21.X*v23.X) + float32(v21.Y*v23.Y) + float32(v21.Z*v23.Z)
			goto _233
		_233:
			cosTriangleAngle = v29
			if cosManifoldAngle <= clusterThreshold || cosTriangleAngle <= clusterThreshold {
				goto _225
			}
			// Found a cluster
			clusterIndex = j
			break
		_225:
			;
			j = j + 1
		}
		if clusterIndex != -int32(1) {
			*(*int32)(unsafe.Add(nil, clusterMemberships+uintptr(i4)*4)) = clusterIndex
			(*(*b3Cluster)(unsafe.Add(nil, clusters+uintptr(clusterIndex)*40))).PointCapacity += (*LocalManifold)(unsafe.Pointer(manifold1)).PointCount
		} else {
			(*(*b3Cluster)(unsafe.Add(nil, clusters+uintptr(clusterCount)*40))).ManifoldNormal = manifoldNormal
			(*(*b3Cluster)(unsafe.Add(nil, clusters+uintptr(clusterCount)*40))).TriangleNormal = triangleNormal
			(*(*b3Cluster)(unsafe.Add(nil, clusters+uintptr(clusterCount)*40))).PointCapacity = (*LocalManifold)(unsafe.Pointer(manifold1)).PointCount
			*(*int32)(unsafe.Add(nil, clusterMemberships+uintptr(i4)*4)) = clusterCount
			clusterCount = clusterCount + int32(1)
		}
		goto _224
	_224:
		;
		i4 = i4 + 1
	}
	if clusterPointCount == 0 {
		return boolUint8(false1 != 0)
	}
	v3 = bp + 56
	v4 = int32FromUint64(uint64FromInt32(clusterPointCount) * uint64(24))
	if v4 == int32FromInt32(0) {
		v5 = uintptrFromInt32(0)
		goto _237
	}
	aligned = ((*b3Arena)(unsafe.Pointer(v3)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
	if aligned+v4 > (*b3Arena)(unsafe.Pointer(v3)).Capacity {
		v5 = b3ArenaOverflowAlloc(tls, v3, v4)
		goto _237
	}
	(*b3Arena)(unsafe.Pointer(v3)).Index = aligned + v4
	if (*b3Arena)(unsafe.Pointer(v3)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex {
		(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v3)).Index
	}
	v5 = (*b3Arena)(unsafe.Pointer(v3)).Memory + uintptr(aligned)
	goto _237
_237:
	// Setup clusters
	clusterPoints = v5
	pointOffset = 0
	i5 = 0
	for {
		if !(i5 < clusterCount) {
			break
		}
		cluster = clusters + uintptr(i5)*40
		(*b3Cluster)(unsafe.Pointer(cluster)).Points = clusterPoints + uintptr(pointOffset)*24
		(*b3Cluster)(unsafe.Pointer(cluster)).PointCount = 0
		pointOffset = pointOffset + (*b3Cluster)(unsafe.Pointer(cluster)).PointCapacity
		goto _238
	_238:
		;
		i5 = i5 + 1
	}
	// Populate clusters
	i6 = 0
	for {
		if !(i6 < acceptedManifoldCount) {
			break
		}
		clusterIndex1 = *(*int32)(unsafe.Add(nil, clusterMemberships+uintptr(i6)*4))
		if clusterIndex1 == -int32(1) {
			goto _239
		}
		if v1 = !!(0 <= clusterIndex1 && clusterIndex1 < clusterCount); !v1 {
			b3InternalAssert(tls, __ccgo_ts+22850, __ccgo_ts+22185, int32FromInt32(1087))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		am = *(*uintptr)(unsafe.Add(nil, acceptedManifolds+uintptr(i6)*8))
		cm = clusters + uintptr(clusterIndex1)*40
		j1 = 0
		for {
			if !(j1 < (*LocalManifold)(unsafe.Pointer(am)).PointCount) {
				break
			}
			if v1 = !!((*b3Cluster)(unsafe.Pointer(cm)).PointCount < (*b3Cluster)(unsafe.Pointer(cm)).PointCapacity); !v1 {
				b3InternalAssert(tls, __ccgo_ts+22899, __ccgo_ts+22185, int32FromInt32(1093))
			}
			_ = v1 || bool(int32FromInt32(0) != 0)
			ap = (*LocalManifold)(unsafe.Pointer(am)).points + uintptr(j1)*24
			cp = (*b3Cluster)(unsafe.Pointer(cm)).Points + uintptr((*b3Cluster)(unsafe.Pointer(cm)).PointCount)*24
			(*LocalManifoldPoint)(unsafe.Pointer(cp)).TriangleIndex = (*LocalManifold)(unsafe.Pointer(am)).TriangleIndex
			(*LocalManifoldPoint)(unsafe.Pointer(cp)).Point = (*LocalManifoldPoint)(unsafe.Pointer(ap)).Point
			(*LocalManifoldPoint)(unsafe.Pointer(cp)).Separation = (*LocalManifoldPoint)(unsafe.Pointer(ap)).Separation
			(*LocalManifoldPoint)(unsafe.Pointer(cp)).Pair = (*LocalManifoldPoint)(unsafe.Pointer(ap)).Pair
			*(*int32)(unsafe.Add(nil, cm+36)) += int32(1)
			goto _241
		_241:
			;
			j1 = j1 + 1
		}
		goto _239
	_239:
		;
		i6 = i6 + 1
	}
	// Simplify clusters
	i7 = 0
	for {
		if !(i7 < clusterCount) {
			break
		}
		cm1 = clusters + uintptr(i7)*40
		if v1 = !!((*b3Cluster)(unsafe.Pointer(cm1)).PointCount == (*b3Cluster)(unsafe.Pointer(cm1)).PointCapacity); !v1 {
			b3InternalAssert(tls, __ccgo_ts+22934, __ccgo_ts+22185, int32FromInt32(1109))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		reducedCount = b3ReduceCluster(tls, (*b3Cluster)(unsafe.Pointer(cm1)).Points, (*b3Cluster)(unsafe.Pointer(cm1)).PointCount, (*b3Cluster)(unsafe.Pointer(cm1)).TriangleNormal, *(*b3Arena)(unsafe.Add(nil, bp+56)))
		(*b3Cluster)(unsafe.Pointer(cm1)).PointCount = reducedCount
		goto _243
	_243:
		;
		i7 = i7 + 1
	}
	// Make a temporary copy of previous manifolds
	oldManifoldCount = (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount
	oldManifolds = uintptrFromInt32(0)
	if oldManifoldCount > 0 {
		v3 = bp + 56
		v4 = int32FromUint64(uint64FromInt32(oldManifoldCount) * uint64(268))
		if v4 == int32FromInt32(0) {
			v5 = uintptrFromInt32(0)
			goto _248
		}
		aligned = ((*b3Arena)(unsafe.Pointer(v3)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
		if aligned+v4 > (*b3Arena)(unsafe.Pointer(v3)).Capacity {
			v5 = b3ArenaOverflowAlloc(tls, v3, v4)
			goto _248
		}
		(*b3Arena)(unsafe.Pointer(v3)).Index = aligned + v4
		if (*b3Arena)(unsafe.Pointer(v3)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex {
			(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v3)).Index
		}
		v5 = (*b3Arena)(unsafe.Pointer(v3)).Memory + uintptr(aligned)
		goto _248
	_248:
		oldManifolds = v5
		memcpy(tls, oldManifolds, (*b3Contact)(unsafe.Pointer(contact)).Manifolds, uint64FromInt32(oldManifoldCount)*uint64(268))
	}
	// Resize manifolds if needed
	if oldManifoldCount != clusterCount {
		v3 = world2
		v4 = (*b3Contact)(unsafe.Pointer(contact)).ManifoldCount
		if v4 == int32FromInt32(0) {
			goto _251
		}
		index1 = v4 - int32(1)
		__builtin_b3LockMutex(tls, (*b3World)(unsafe.Pointer(v3)).ManifoldAllocatorMutex)
		if v1 = !!(0 <= index1 && index1 < (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+8145, __ccgo_ts+8211, int32FromInt32(359))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		allocator2 = (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Data + uintptr(index1)*40
		b3FreeElement(tls, allocator2, (*b3Contact)(unsafe.Pointer(contact)).Manifolds)
		__builtin_b3UnlockMutex(tls, (*b3World)(unsafe.Pointer(v3)).ManifoldAllocatorMutex)
	_251:
		;
		v3 = world2
		v4 = clusterCount
		if v4 == int32FromInt32(0) {
			v5 = uintptrFromInt32(0)
			goto _256
		}
		index = v4 - int32(1)
		__builtin_b3LockMutex(tls, (*b3World)(unsafe.Pointer(v3)).ManifoldAllocatorMutex)
		currentCount = (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Count
		i = currentCount
		for {
			if !(i < v4) {
				break
			}
			allocator = b3CreateBlockAllocator(tls, int32FromUint64(uint64FromInt32(i+int32FromInt32(1))*uint64(268)), int32FromInt32(2)*(int32FromInt32(1)<<int32FromInt32(_B3_BLOCK_EXPONENT)))
			if (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Count >= (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Capacity {
				oldSize = int32FromUint64(uint64FromInt32((*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Capacity) * uint64(40))
				if (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Capacity == 0 {
					v8 = int32(8)
				} else {
					v8 = int32(2) * (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Capacity
				}
				newCapacity = v8
				newSize = int32FromUint64(uint64FromInt32(newCapacity) * uint64(40))
				(*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Data = b3GrowAlloc(tls, (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Data, oldSize, newSize)
				(*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Capacity = newCapacity
			}
			v7 = v3 + 4040 + 8
			v12 = *(*int32)(unsafe.Pointer(v7))
			*(*int32)(unsafe.Pointer(v7)) = *(*int32)(unsafe.Pointer(v7)) + 1
			*(*b3BlockAllocator)(unsafe.Add(nil, (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Data+uintptr(v12)*40)) = allocator
			goto _257
		_257:
			;
			i = i + 1
		}
		if v1 = !!(0 <= index && index < (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Count); !v1 {
			b3InternalAssert(tls, __ccgo_ts+8145, __ccgo_ts+8211, int32FromInt32(343))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		allocator1 = (*b3World)(unsafe.Pointer(v3)).ManifoldAllocators.Data + uintptr(index)*40
		manifolds = b3AllocateElement(tls, allocator1)
		__builtin_b3UnlockMutex(tls, (*b3World)(unsafe.Pointer(v3)).ManifoldAllocatorMutex)
		memset(tls, manifolds, 0, uint64FromInt32(v4)*uint64(268))
		v5 = manifolds
		goto _256
	_256:
		(*b3Contact)(unsafe.Pointer(contact)).Manifolds = v5
		(*b3Contact)(unsafe.Pointer(contact)).ManifoldCount = int32FromUint16(uint16FromInt32(clusterCount))
	} else {
		// Mem zero manifolds
		memset(tls, (*b3Contact)(unsafe.Pointer(contact)).Manifolds, 0, uint64FromInt32((*b3Contact)(unsafe.Pointer(contact)).ManifoldCount)*uint64(268))
	}
	consumed = uintptrFromInt32(0)
	if oldManifoldCount > 0 {
		v3 = bp + 56
		v4 = int32FromUint64(uint64FromInt32(oldManifoldCount) * uint64(1))
		if v4 == int32FromInt32(0) {
			v5 = uintptrFromInt32(0)
			goto _265
		}
		aligned = ((*b3Arena)(unsafe.Pointer(v3)).Index + (int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))) & ^(int32FromInt32(_B3_ARENA_ALIGNMENT) - int32FromInt32(1))
		if aligned+v4 > (*b3Arena)(unsafe.Pointer(v3)).Capacity {
			v5 = b3ArenaOverflowAlloc(tls, v3, v4)
			goto _265
		}
		(*b3Arena)(unsafe.Pointer(v3)).Index = aligned + v4
		if (*b3Arena)(unsafe.Pointer(v3)).Index > (*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex {
			(*b3ArenaSharedState)(unsafe.Pointer((*b3Arena)(unsafe.Pointer(v3)).Shared)).MaxIndex = (*b3Arena)(unsafe.Pointer(v3)).Index
		}
		v5 = (*b3Arena)(unsafe.Pointer(v3)).Memory + uintptr(aligned)
		goto _265
	_265:
		consumed = v5
		memset(tls, consumed, 0, uint64FromInt32(oldManifoldCount)*uint64(1))
	}
	v17 = xfB.Q
	xx = float32(v17.V.X * v17.V.X)
	yy = float32(v17.V.Y * v17.V.Y)
	zz = float32(v17.V.Z * v17.V.Z)
	xy = float32(v17.V.X * v17.V.Y)
	xz = float32(v17.V.X * v17.V.Z)
	xw = float32(v17.V.X * v17.S)
	yz = float32(v17.V.Y * v17.V.Z)
	yw = float32(v17.V.Y * v17.S)
	zw = float32(v17.V.Z * v17.S)
	v64 = Matrix3{
		Cx: Vec3{
			X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
			Y: float32(float32FromFloat32(2) * (xy + zw)),
			Z: float32(float32FromFloat32(2) * (xz - yw)),
		},
		Cy: Vec3{
			X: float32(float32FromFloat32(2) * (xy - zw)),
			Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
			Z: float32(float32FromFloat32(2) * (yz + xw)),
		},
		Cz: Vec3{
			X: float32(float32FromFloat32(2) * (xz + yw)),
			Y: float32(float32FromFloat32(2) * (yz - xw)),
			Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
		},
	}
	goto _268
_268:
	matrixB = v64
	v269 = xfB.P
	v270 = xfA.P
	v19 = Vec3{
		X: v269.X - v270.X,
		Y: v269.Y - v270.Y,
		Z: v269.Z - v270.Z,
	}
	goto _272
_272:
	offsetA = v19
	normalMatchTolerance = float32FromFloat32(0.995)
	i8 = 0
	for {
		if !(i8 < clusterCount) {
			break
		}
		cm2 = clusters + uintptr(i8)*40
		pointCount = (*b3Cluster)(unsafe.Pointer(cm2)).PointCount
		if v1 = !!(0 < pointCount && pointCount <= int32(_B3_MAX_MANIFOLD_POINTS)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+22970, __ccgo_ts+22185, int32FromInt32(1151))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		manifold2 = (*b3Contact)(unsafe.Pointer(contact)).Manifolds + uintptr(i8)*268
		(*Manifold)(unsafe.Pointer(manifold2)).PointCount = pointCount
		v64 = matrixB
		v19 = (*b3Cluster)(unsafe.Pointer(cm2)).ManifoldNormal
		b10 = Vec3{
			X: float32(v64.Cx.X*v19.X) + float32(v64.Cy.X*v19.Y) + float32(v64.Cz.X*v19.Z),
			Y: float32(v64.Cx.Y*v19.X) + float32(v64.Cy.Y*v19.Y) + float32(v64.Cz.Y*v19.Z),
			Z: float32(v64.Cx.Z*v19.X) + float32(v64.Cy.Z*v19.Y) + float32(v64.Cz.Z*v19.Z),
		}
		v20 = b10
		goto _278
	_278:
		(*Manifold)(unsafe.Pointer(manifold2)).Normal = v20
		v64 = matrixB
		v19 = (*b3Cluster)(unsafe.Pointer(cm2)).ManifoldNormal
		b10 = Vec3{
			X: float32(v64.Cx.X*v19.X) + float32(v64.Cy.X*v19.Y) + float32(v64.Cz.X*v19.Z),
			Y: float32(v64.Cx.Y*v19.X) + float32(v64.Cy.Y*v19.Y) + float32(v64.Cz.Y*v19.Z),
			Z: float32(v64.Cx.Z*v19.X) + float32(v64.Cy.Z*v19.Y) + float32(v64.Cz.Z*v19.Z),
		}
		v20 = b10
		goto _282
	_282:
		clusterNormal = v20
		bestDot = normalMatchTolerance
		bestIndex = -int32(1)
		j2 = 0
		for {
			if !(j2 < oldManifoldCount) {
				break
			}
			if int32FromUint8(*(*uint8)(unsafe.Add(nil, consumed+uintptr(j2)))) == int32(true1) {
				goto _283
			}
			v19 = (*(*Manifold)(unsafe.Add(nil, oldManifolds+uintptr(j2)*268))).Normal
			v20 = clusterNormal
			v24 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
			goto _287
		_287:
			dot = v24
			if dot > bestDot {
				bestIndex = j2
				bestDot = dot
			}
			goto _283
		_283:
			;
			j2 = j2 + 1
		}
		matchedManifold = uintptrFromInt32(0)
		if bestIndex != -int32(1) {
			matchedManifold = oldManifolds + uintptr(bestIndex)*268
			(*Manifold)(unsafe.Pointer(manifold2)).FrictionImpulse = (*Manifold)(unsafe.Pointer(matchedManifold)).FrictionImpulse
			(*Manifold)(unsafe.Pointer(manifold2)).RollingImpulse = (*Manifold)(unsafe.Pointer(matchedManifold)).RollingImpulse
			(*Manifold)(unsafe.Pointer(manifold2)).TwistImpulse = (*Manifold)(unsafe.Pointer(matchedManifold)).TwistImpulse
			*(*uint8)(unsafe.Add(nil, consumed+uintptr(bestIndex))) = boolUint8(true1 != 0)
		}
		j3 = 0
		for {
			if !(j3 < pointCount) {
				break
			}
			source = (*b3Cluster)(unsafe.Pointer(cm2)).Points + uintptr(j3)*24
			target = manifold2 + uintptr(j3)*56
			// Contact points are computed in frame B
			v64 = matrixB
			v19 = (*LocalManifoldPoint)(unsafe.Pointer(source)).Point
			b10 = Vec3{
				X: float32(v64.Cx.X*v19.X) + float32(v64.Cy.X*v19.Y) + float32(v64.Cz.X*v19.Z),
				Y: float32(v64.Cx.Y*v19.X) + float32(v64.Cy.Y*v19.Y) + float32(v64.Cz.Y*v19.Z),
				Z: float32(v64.Cx.Z*v19.X) + float32(v64.Cy.Z*v19.Y) + float32(v64.Cz.Z*v19.Z),
			}
			v20 = b10
			goto _292
		_292:
			(*ManifoldPoint)(unsafe.Pointer(target)).AnchorB = v20
			v19 = (*ManifoldPoint)(unsafe.Pointer(target)).AnchorB
			v20 = offsetA
			v21 = Vec3{
				X: v19.X + v20.X,
				Y: v19.Y + v20.Y,
				Z: v19.Z + v20.Z,
			}
			goto _296
		_296:
			(*ManifoldPoint)(unsafe.Pointer(target)).AnchorA = v21
			(*ManifoldPoint)(unsafe.Pointer(target)).Separation = (*LocalManifoldPoint)(unsafe.Pointer(source)).Separation - restOffset
			v297 = (*LocalManifoldPoint)(unsafe.Pointer(source)).Pair
			v119 = uint32(v297.Owner1)<<int32FromInt32(24) | uint32(v297.Index1)<<int32FromInt32(16) | uint32(v297.Owner2)<<int32FromInt32(8) | uint32(v297.Index2)
			goto _299
		_299:
			(*ManifoldPoint)(unsafe.Pointer(target)).FeatureId = v119
			(*ManifoldPoint)(unsafe.Pointer(target)).TriangleIndex = (*LocalManifoldPoint)(unsafe.Pointer(source)).TriangleIndex
			// Preserve normal impulse if possible
			if matchedManifold != uintptrFromInt32(0) {
				oldPointCount = (*Manifold)(unsafe.Pointer(matchedManifold)).PointCount
				k2 = 0
				for {
					if !(k2 < oldPointCount) {
						break
					}
					oldPt = matchedManifold + uintptr(k2)*56
					if (*ManifoldPoint)(unsafe.Pointer(target)).FeatureId == (*ManifoldPoint)(unsafe.Pointer(oldPt)).FeatureId && (*ManifoldPoint)(unsafe.Pointer(target)).TriangleIndex == (*ManifoldPoint)(unsafe.Pointer(oldPt)).TriangleIndex {
						(*ManifoldPoint)(unsafe.Pointer(target)).NormalImpulse = (*ManifoldPoint)(unsafe.Pointer(oldPt)).NormalImpulse
						(*ManifoldPoint)(unsafe.Pointer(target)).Persisted = boolUint8(true1 != 0)
						// claimed
						(*ManifoldPoint)(unsafe.Pointer(oldPt)).TriangleIndex = -int32(1)
						break
					}
					goto _300
				_300:
					;
					k2 = k2 + 1
				}
			}
			goto _288
		_288:
			;
			j3 = j3 + 1
		}
		goto _273
	_273:
		;
		i8 = i8 + 1
	}
	v3 = shapeA
	if (*b3Shape)(unsafe.Pointer(v3)).Materials != uintptrFromInt32(0) {
		v7 = (*b3Shape)(unsafe.Pointer(v3)).Materials
	} else {
		v7 = v3 + 80
	}
	v5 = v7
	goto _303
_303:
	materialsA = v5
	v9 = shapeB
	if (*b3Shape)(unsafe.Pointer(v9)).Materials != uintptrFromInt32(0) {
		v13 = (*b3Shape)(unsafe.Pointer(v9)).Materials
	} else {
		v13 = v9 + 80
	}
	v11 = v13
	goto _307
_307:
	materialB = v11
	tangentVelocityA = b3Vec3_zero
	// Update friction and restitution if the mesh has per triangle material
	if (*b3Shape)(unsafe.Pointer(shapeA)).MaterialCount > 0 {
		friction = float32FromFloat32(0)
		restitution = float32FromFloat32(0)
		sampleCount = float32FromFloat32(0)
		if (*b3Shape)(unsafe.Pointer(shapeA)).Type1 == int32(b3_meshShape) {
			v66 = uintptr(unsafe.Pointer((*(*Mesh)(unsafe.Add(unsafe.Pointer(shapeA), 176))).Data))
			if (*MeshData)(unsafe.Pointer(v66)).MaterialOffset == int32FromInt32(0) {
				v68 = uintptrFromInt32(0)
				goto _311
			}
			v68 = uintptr(int64(v66) + int64((*MeshData)(unsafe.Pointer(v66)).MaterialOffset))
			goto _311
		_311:
			materialIndices = v68
		} else {
			v3 = *(*uintptr)(unsafe.Add(unsafe.Pointer(shapeA), 176))
			if (*HeightFieldData)(unsafe.Pointer(v3)).MaterialOffset == int32FromInt32(0) {
				v5 = uintptrFromInt32(0)
				goto _314
			}
			v5 = uintptr(int64(v3) + int64((*HeightFieldData)(unsafe.Pointer(v3)).MaterialOffset))
			goto _314
		_314:
			materialIndices = v5
		}
		i9 = 0
		for {
			if !(i9 < clusterCount) {
				break
			}
			manifold3 = (*b3Contact)(unsafe.Pointer(contact)).Manifolds + uintptr(i9)*268
			pointCount1 = (*Manifold)(unsafe.Pointer(manifold3)).PointCount
			j4 = 0
			for {
				if !(j4 < pointCount1) {
					break
				}
				triangleIndex3 = (*(*ManifoldPoint)(unsafe.Add(nil, manifold3+uintptr(j4)*56))).TriangleIndex
				if (*b3Shape)(unsafe.Pointer(shapeA)).Type1 == int32(b3_meshShape) {
					materialIndex = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(triangleIndex3))))
					if materialMap != uintptrFromInt32(0) {
						materialIndex = int32FromUint16(*(*uint16_t)(unsafe.Add(nil, materialMap+uintptr(materialIndex)*2)))
					}
				} else {
					materialIndex = int32FromUint8(*(*uint8_t)(unsafe.Add(nil, materialIndices+uintptr(triangleIndex3>>int32(1)))))
				}
				v4 = materialIndex
				v8 = 0
				v12 = (*b3Shape)(unsafe.Pointer(shapeA)).MaterialCount - int32(1)
				if v4 < v8 {
					v71 = v8
				} else {
					if v12 < v4 {
						v77 = v12
					} else {
						v77 = v4
					}
					v71 = v77
				}
				v67 = v71
				goto _321
			_321:
				materialIndex = v67
				material = *(*SurfaceMaterial)(unsafe.Add(nil, materialsA+uintptr(materialIndex)*40))
				friction = friction + (*(*func(*_Stack, float32, uint64, float32, uint64) float32)(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world2)).FrictionCallback})))(tls, material.Friction, material.UserMaterialId, (*SurfaceMaterial)(unsafe.Pointer(materialB)).Friction, (*SurfaceMaterial)(unsafe.Pointer(materialB)).UserMaterialId)
				restitution = restitution + (*(*func(*_Stack, float32, uint64, float32, uint64) float32)(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world2)).RestitutionCallback})))(tls, material.Restitution, material.UserMaterialId, (*SurfaceMaterial)(unsafe.Pointer(materialB)).Restitution, (*SurfaceMaterial)(unsafe.Pointer(materialB)).UserMaterialId)
				v19 = tangentVelocityA
				v20 = material.TangentVelocity
				v21 = Vec3{
					X: v19.X + v20.X,
					Y: v19.Y + v20.Y,
					Z: v19.Z + v20.Z,
				}
				goto _327
			_327:
				tangentVelocityA = v21
				sampleCount = sampleCount + float32FromFloat32(1)
				goto _316
			_316:
				;
				j4 = j4 + 1
			}
			goto _315
		_315:
			;
			i9 = i9 + 1
		}
		if sampleCount > float32FromFloat32(0) {
			invCount = float32FromFloat32(1) / sampleCount
			(*b3Contact)(unsafe.Pointer(contact)).Friction = float32(invCount * friction)
			(*b3Contact)(unsafe.Pointer(contact)).Restitution = float32(invCount * restitution)
			v24 = invCount
			v19 = tangentVelocityA
			v20 = Vec3{
				X: float32(v24 * v19.X),
				Y: float32(v24 * v19.Y),
				Z: float32(v24 * v19.Z),
			}
			goto _331
		_331:
			tangentVelocityA = v20
		}
		if v1 = !!(b3IsValidFloat(tls, (*b3Contact)(unsafe.Pointer(contact)).Friction) != 0 && (*b3Contact)(unsafe.Pointer(contact)).Friction >= float32FromFloat32(0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+23025, __ccgo_ts+22185, int32FromInt32(1284))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
		if v1 = !!(b3IsValidFloat(tls, (*b3Contact)(unsafe.Pointer(contact)).Restitution) != 0 && (*b3Contact)(unsafe.Pointer(contact)).Restitution >= float32FromFloat32(0)); !v1 {
			b3InternalAssert(tls, __ccgo_ts+23090, __ccgo_ts+22185, int32FromInt32(1285))
		}
		_ = v1 || bool(int32FromInt32(0) != 0)
	} else {
		// Keep these updated in case the values on the shapes are modified
		(*b3Contact)(unsafe.Pointer(contact)).Friction = (*(*func(*_Stack, float32, uint64, float32, uint64) float32)(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world2)).FrictionCallback})))(tls, (*(*SurfaceMaterial)(unsafe.Add(nil, materialsA))).Friction, (*(*SurfaceMaterial)(unsafe.Add(nil, materialsA))).UserMaterialId, (*SurfaceMaterial)(unsafe.Pointer(materialB)).Friction, (*SurfaceMaterial)(unsafe.Pointer(materialB)).UserMaterialId)
		(*b3Contact)(unsafe.Pointer(contact)).Restitution = (*(*func(*_Stack, float32, uint64, float32, uint64) float32)(unsafe.Pointer(&struct{ uintptr }{(*b3World)(unsafe.Pointer(world2)).RestitutionCallback})))(tls, (*(*SurfaceMaterial)(unsafe.Add(nil, materialsA))).Restitution, (*(*SurfaceMaterial)(unsafe.Add(nil, materialsA))).UserMaterialId, (*SurfaceMaterial)(unsafe.Pointer(materialB)).Restitution, (*SurfaceMaterial)(unsafe.Pointer(materialB)).UserMaterialId)
		tangentVelocityA = (*(*SurfaceMaterial)(unsafe.Add(nil, materialsA))).TangentVelocity
	}
	v17 = xfA.Q
	v19 = tangentVelocityA
	v20 = v17.V
	v21 = v19
	c.X = float32(v20.Y*v21.Z) - float32(v20.Z*v21.Y)
	c.Y = float32(v20.Z*v21.X) - float32(v20.X*v21.Z)
	c.Z = float32(v20.X*v21.Y) - float32(v20.Y*v21.X)
	v23 = c
	goto _339
_339:
	t1 = v23
	v25 = t1
	v24 = v17.S
	v26 = v19
	v28 = Vec3{
		X: v25.X + float32(v24*v26.X),
		Y: v25.Y + float32(v24*v26.Y),
		Z: v25.Z + float32(v24*v26.Z),
	}
	goto _344
_344:
	t2 = v28
	v30 = v17.V
	v31 = t2
	c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
	c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
	c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
	v33 = c
	goto _348
_348:
	t3 = v33
	v40 = v19
	v29 = float32FromFloat32(2)
	v41 = t3
	v42 = Vec3{
		X: v40.X + float32(v29*v41.X),
		Y: v40.Y + float32(v29*v41.Y),
		Z: v40.Z + float32(v29*v41.Z),
	}
	goto _355
_355:
	v34 = v42
	goto _350
_350:
	tangentVelocityA = v34
	radiusB = float32FromFloat32(0)
	if (*b3Shape)(unsafe.Pointer(shapeB)).Type1 == int32(b3_sphereShape) {
		radiusB = (*(*Sphere)(unsafe.Add(unsafe.Pointer(shapeB), 176))).Radius
	} else {
		if (*b3Shape)(unsafe.Pointer(shapeB)).Type1 == int32(b3_capsuleShape) {
			radiusB = (*b3Shape)(unsafe.Pointer(shapeB)).ｆ__ccgo21_176.Capsule.Radius
		} else {
			if (*b3Shape)(unsafe.Pointer(shapeB)).Type1 == int32(b3_hullShape) {
				radiusB = (*HullData)(unsafe.Pointer(*(*uintptr)(unsafe.Add(unsafe.Pointer(shapeB), 176)))).InnerRadius
			}
		}
	}
	(*b3Contact)(unsafe.Pointer(contact)).RollingResistance = float32((*SurfaceMaterial)(unsafe.Pointer(materialB)).RollingResistance * radiusB)
	v17 = xfB.Q
	v19 = (*SurfaceMaterial)(unsafe.Pointer(materialB)).TangentVelocity
	v20 = v17.V
	v21 = v19
	c.X = float32(v20.Y*v21.Z) - float32(v20.Z*v21.Y)
	c.Y = float32(v20.Z*v21.X) - float32(v20.X*v21.Z)
	c.Z = float32(v20.X*v21.Y) - float32(v20.Y*v21.X)
	v23 = c
	goto _361
_361:
	t1 = v23
	v25 = t1
	v24 = v17.S
	v26 = v19
	v28 = Vec3{
		X: v25.X + float32(v24*v26.X),
		Y: v25.Y + float32(v24*v26.Y),
		Z: v25.Z + float32(v24*v26.Z),
	}
	goto _366
_366:
	t2 = v28
	v30 = v17.V
	v31 = t2
	c.X = float32(v30.Y*v31.Z) - float32(v30.Z*v31.Y)
	c.Y = float32(v30.Z*v31.X) - float32(v30.X*v31.Z)
	c.Z = float32(v30.X*v31.Y) - float32(v30.Y*v31.X)
	v33 = c
	goto _370
_370:
	t3 = v33
	v40 = v19
	v29 = float32FromFloat32(2)
	v41 = t3
	v42 = Vec3{
		X: v40.X + float32(v29*v41.X),
		Y: v40.Y + float32(v29*v41.Y),
		Z: v40.Z + float32(v29*v41.Z),
	}
	goto _377
_377:
	v34 = v42
	goto _372
_372:
	tangentVelocityB = v34
	v43 = tangentVelocityA
	v45 = tangentVelocityB
	v47 = Vec3{
		X: v43.X - v45.X,
		Y: v43.Y - v45.Y,
		Z: v43.Z - v45.Z,
	}
	goto _381
_381:
	(*b3Contact)(unsafe.Pointer(contact)).TangentVelocity = v47
	return boolUint8(true1 != 0)
}

// C documentation
//
//	// Tight world-space bounds of a query's swept geometry, so the viewer can frame any query and not
//	// just the overlap AABB. Mover and proxy points are origin relative. A cast sweeps the shape from the
//	// origin to origin plus translation. The overlap AABB is already a world-space box.
func b3RecComputeQueryBounds(tls *_Stack, q uintptr) {
	bp := tls.Alloc(3072)
	defer tls.Free(3072)
	var a6, v74 AABB
	var count1, i, i1, i2, n, v7 int32
	var end, v10, v12, v2, v4, v8 Pos
	var local [128]Vec3
	var r, v13, v17, v21, v28, v29, v3, v30, v47, v48, v49, v66, v9 Vec3
	var radius1, v25, v32, v33, v34, v36, v37, v38, v39, v41, v42, v43, v44, v46, v51, v52, v53, v55, v56, v57, v58, v60, v61, v62, v63, v65 float32
	var v23 uintptr
	var v26 bool
	var _ /* world at bp+0 */ [256]Vec3
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = a6, count1, end, i, i1, i2, local, n, r, radius1, v10, v12, v13, v17, v2, v21, v23, v25, v26, v28, v29, v3, v30, v32, v33, v34, v36, v37, v38, v39, v4, v41, v42, v43, v44, v46, v47, v48, v49, v51, v52, v53, v55, v56, v57, v58, v60, v61, v62, v63, v65, v66, v7, v74, v8, v9
	if (*b3RecDrawQuery)(unsafe.Pointer(q)).Kind == int32(_B3_RECQ_OVERLAP_AABB) {
		return
	}
	count1 = 0
	radius1 = float32FromFloat32(0)
	switch (*b3RecDrawQuery)(unsafe.Pointer(q)).Kind {
	case int32(_B3_RECQ_CAST_MOVER):
		fallthrough
	case int32(_B3_RECQ_COLLIDE_MOVER):
		local[0] = (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Center1
		local[int32(1)] = (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Center2
		count1 = int32(2)
		radius1 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Mover.Radius
	case int32(_B3_RECQ_OVERLAP_SHAPE):
		fallthrough
	case int32(_B3_RECQ_CAST_SHAPE):
		count1 = (*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyCount
		i1 = 0
		for {
			if !(i1 < count1) {
				break
			}
			local[i1] = *(*Vec3)(unsafe.Add(nil, q+72+uintptr(i1)*12))
			goto _1
		_1:
			;
			i1 = i1 + 1
		}
		radius1 = (*b3RecDrawQuery)(unsafe.Pointer(q)).ProxyRadius
	default:
		break
	}
	if count1 == 0 {
		local[0] = b3Vec3_zero
		count1 = int32(1)
	}
	v2 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
	v3 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Translation
	v4 = Pos{
		X: v2.X + v3.X,
		Y: v2.Y + v3.Y,
		Z: v2.Z + v3.Z,
	}
	goto _5
_5:
	// Sweep each point across the translation. A non-cast query has zero translation, so both ends
	// coincide and the duplicates fold away.
	end = v4
	*(*[256]Vec3)(unsafe.Add(nil, bp)) = [256]Vec3{}
	n = 0
	i2 = 0
	for {
		if !(i2 < count1) {
			break
		}
		v7 = n
		n = n + 1
		v8 = (*b3RecDrawQuery)(unsafe.Pointer(q)).Origin
		v9 = local[i2]
		v10 = Pos{
			X: v8.X + v9.X,
			Y: v8.Y + v9.Y,
			Z: v8.Z + v9.Z,
		}
		goto _11
	_11:
		v12 = v10
		v13 = Vec3{
			X: v12.X,
			Y: v12.Y,
			Z: v12.Z,
		}
		goto _14
	_14:
		(*(*[256]Vec3)(unsafe.Add(nil, bp)))[v7] = v13
		v7 = n
		n = n + 1
		v2 = end
		v3 = local[i2]
		v4 = Pos{
			X: v2.X + v3.X,
			Y: v2.Y + v3.Y,
			Z: v2.Z + v3.Z,
		}
		goto _19
	_19:
		v8 = v4
		v9 = Vec3{
			X: v8.X,
			Y: v8.Y,
			Z: v8.Z,
		}
		goto _22
	_22:
		(*(*[256]Vec3)(unsafe.Add(nil, bp)))[v7] = v9
		goto _6
	_6:
		;
		i2 = i2 + 1
	}
	v23 = bp
	v7 = n
	v25 = radius1
	if v26 = !!(v7 > int32FromInt32(0)); !v26 {
		b3InternalAssert(tls, __ccgo_ts+6345, __ccgo_ts+6355, int32FromInt32(948))
	}
	_ = v26 || bool(int32FromInt32(0) != 0)
	a6 = AABB{
		LowerBound: *(*Vec3)(unsafe.Add(nil, v23)),
		UpperBound: *(*Vec3)(unsafe.Add(nil, v23)),
	}
	i = int32(1)
	for {
		if !(i < v7) {
			break
		}
		v3 = a6.LowerBound
		v9 = *(*Vec3)(unsafe.Add(nil, v23+uintptr(i)*12))
		v32 = v3.X
		v33 = v9.X
		if v32 < v33 {
			v36 = v32
		} else {
			v36 = v33
		}
		v34 = v36
		goto _35
	_35:
		v37 = v3.Y
		v38 = v9.Y
		if v37 < v38 {
			v41 = v37
		} else {
			v41 = v38
		}
		v39 = v41
		goto _40
	_40:
		v42 = v3.Z
		v43 = v9.Z
		if v42 < v43 {
			v46 = v42
		} else {
			v46 = v43
		}
		v44 = v46
		goto _45
	_45:
		v13 = Vec3{
			X: v34,
			Y: v39,
			Z: v44,
		}
		goto _31
	_31:
		a6.LowerBound = v13
		v17 = a6.UpperBound
		v21 = *(*Vec3)(unsafe.Add(nil, v23+uintptr(i)*12))
		v51 = v17.X
		v52 = v21.X
		if v51 > v52 {
			v55 = v51
		} else {
			v55 = v52
		}
		v53 = v55
		goto _54
	_54:
		v56 = v17.Y
		v57 = v21.Y
		if v56 > v57 {
			v60 = v56
		} else {
			v60 = v57
		}
		v58 = v60
		goto _59
	_59:
		v61 = v17.Z
		v62 = v21.Z
		if v61 > v62 {
			v65 = v61
		} else {
			v65 = v62
		}
		v63 = v65
		goto _64
	_64:
		v28 = Vec3{
			X: v53,
			Y: v58,
			Z: v63,
		}
		goto _50
	_50:
		a6.UpperBound = v28
		goto _27
	_27:
		;
		i = i + 1
	}
	r = Vec3{
		X: v25,
		Y: v25,
		Z: v25,
	}
	v29 = a6.LowerBound
	v30 = r
	v47 = Vec3{
		X: v29.X - v30.X,
		Y: v29.Y - v30.Y,
		Z: v29.Z - v30.Z,
	}
	goto _69
_69:
	a6.LowerBound = v47
	v48 = a6.UpperBound
	v49 = r
	v66 = Vec3{
		X: v48.X + v49.X,
		Y: v48.Y + v49.Y,
		Z: v48.Z + v49.Z,
	}
	goto _73
_73:
	a6.UpperBound = v66
	v74 = a6
	goto _75
_75:
	(*b3RecDrawQuery)(unsafe.Pointer(q)).Aabb = v74
}

func b3ComputeProxyAABB(tls *_Stack, proxy uintptr) (r1 AABB) {
	var aabb AABB
	var i int32
	var points uintptr
	var r, v2, v3, v4 Vec3
	var v10, v11, v12, v13, v15, v16, v17, v18, v20, v6, v7, v8 float32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = aabb, i, points, r, v10, v11, v12, v13, v15, v16, v17, v18, v2, v20, v3, v4, v6, v7, v8
	points = uintptr(unsafe.Pointer((*ShapeProxy)(unsafe.Pointer(proxy)).Points))
	aabb = AABB{
		LowerBound: *(*Vec3)(unsafe.Add(nil, points)),
		UpperBound: *(*Vec3)(unsafe.Add(nil, points)),
	}
	i = int32(1)
	for {
		if !(i < (*ShapeProxy)(unsafe.Pointer(proxy)).Count) {
			break
		}
		v2 = aabb.LowerBound
		v3 = *(*Vec3)(unsafe.Add(nil, points+uintptr(i)*12))
		v6 = v2.X
		v7 = v3.X
		if v6 < v7 {
			v10 = v6
		} else {
			v10 = v7
		}
		v8 = v10
		goto _9
	_9:
		v11 = v2.Y
		v12 = v3.Y
		if v11 < v12 {
			v15 = v11
		} else {
			v15 = v12
		}
		v13 = v15
		goto _14
	_14:
		v16 = v2.Z
		v17 = v3.Z
		if v16 < v17 {
			v20 = v16
		} else {
			v20 = v17
		}
		v18 = v20
		goto _19
	_19:
		v4 = Vec3{
			X: v8,
			Y: v13,
			Z: v18,
		}
		goto _5
	_5:
		aabb.LowerBound = v4
		v2 = aabb.UpperBound
		v3 = *(*Vec3)(unsafe.Add(nil, points+uintptr(i)*12))
		v6 = v2.X
		v7 = v3.X
		if v6 > v7 {
			v10 = v6
		} else {
			v10 = v7
		}
		v8 = v10
		goto _28
	_28:
		v11 = v2.Y
		v12 = v3.Y
		if v11 > v12 {
			v15 = v11
		} else {
			v15 = v12
		}
		v13 = v15
		goto _33
	_33:
		v16 = v2.Z
		v17 = v3.Z
		if v16 > v17 {
			v20 = v16
		} else {
			v20 = v17
		}
		v18 = v20
		goto _38
	_38:
		v4 = Vec3{
			X: v8,
			Y: v13,
			Z: v18,
		}
		goto _24
	_24:
		aabb.UpperBound = v4
		goto _1
	_1:
		;
		i = i + 1
	}
	r = Vec3{
		X: (*ShapeProxy)(unsafe.Pointer(proxy)).Radius,
		Y: (*ShapeProxy)(unsafe.Pointer(proxy)).Radius,
		Z: (*ShapeProxy)(unsafe.Pointer(proxy)).Radius,
	}
	v2 = aabb.LowerBound
	v3 = r
	v4 = Vec3{
		X: v2.X - v3.X,
		Y: v2.Y - v3.Y,
		Z: v2.Z - v3.Z,
	}
	goto _43
_43:
	aabb.LowerBound = v4
	v2 = aabb.UpperBound
	v3 = r
	v4 = Vec3{
		X: v2.X + v3.X,
		Y: v2.Y + v3.Y,
		Z: v2.Z + v3.Z,
	}
	goto _47
_47:
	aabb.UpperBound = v4
	return aabb
}

// C documentation
//
//	// A block is a range of tasks, a start index and count as a sub-array. Each worker receives at
//	// most M blocks of work. The workers may receive less blocks if there is not sufficient work.
//	// Each block of work has a minimum number of elements (block size). This in turn may limit the
//	// number of blocks. If there are many elements then the block size is increased so there are
//	// still at most M blocks of work per worker. M is a tunable number that has two goals:
//	// 1. keep M small to reduce overhead
//	// 2. keep M large enough for other workers to be able to steal work
//	// The block size is a power of two to make math efficient.
func b3ComputeBlockCount(tls *_Stack, itemCount int32, minSize int32, maxBlockCount int32) (r b3BlockDim) {
	var dim b3BlockDim
	var v1 bool
	_, _ = dim, v1
	dim = b3BlockDim{}
	if itemCount == 0 {
		return dim
	}
	if itemCount <= minSize*maxBlockCount {
		dim.Size = minSize
	} else {
		dim.Size = (itemCount + maxBlockCount - int32(1)) / maxBlockCount
	}
	dim.Count = (itemCount + dim.Size - int32(1)) / dim.Size
	if v1 = !!(dim.Count >= int32FromInt32(1)); !v1 {
		b3InternalAssert(tls, __ccgo_ts+31743, __ccgo_ts+31091, int32FromInt32(927))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	if v1 = !!(dim.Size*dim.Count >= itemCount); !v1 {
		b3InternalAssert(tls, __ccgo_ts+31758, __ccgo_ts+31091, int32FromInt32(928))
	}
	_ = v1 || bool(int32FromInt32(0) != 0)
	return dim
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

func b3ComputeSphereMass(tls *_Stack, shape uintptr, density float32) (r MassData) {
	var center Vec3
	var ixx, mass, radius, volume float32
	var out MassData
	var v1 Matrix3
	_, _, _, _, _, _, _ = center, ixx, mass, out, radius, volume, v1
	center = (*Sphere)(unsafe.Pointer(shape)).Center
	radius = (*Sphere)(unsafe.Pointer(shape)).Radius
	volume = float32(float32(float32(float32(float32FromFloat32(4)/float32FromFloat32(3)*float32FromFloat32(3.14159265359))*radius)*radius) * radius)
	mass = float32(volume * density)
	ixx = float32(float32(float32(float32FromFloat32(0.4)*mass)*radius) * radius)
	out.Mass = mass
	out.Center = center
	// Inertia about the center of mass
	v1 = Matrix3{
		Cx: Vec3{
			X: ixx,
		},
		Cy: Vec3{
			Y: ixx,
		},
		Cz: Vec3{
			Z: ixx,
		},
	}
	goto _2
_2:
	out.Inertia = v1
	return out
}

func b3ComputeSphereAABB(tls *_Stack, shape uintptr, transform Transform) (r AABB) {
	var c, center, extent, rv, t1, t2, t3, v10, v11, v13, v14, v15, v17, v19, v21, v22, v24, v26, v27, v28, v3, v30, v31, v32, v34, v35, v36, v4, v5, v6, v8 Vec3
	var radius, v20, v9 float32
	var v1 Transform
	var v2 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, center, extent, radius, rv, t1, t2, t3, v1, v10, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v3, v30, v31, v32, v34, v35, v36, v4, v5, v6, v8, v9
	v1 = transform
	v2 = v1.Q
	v3 = (*Sphere)(unsafe.Pointer(shape)).Center
	v4 = v2.V
	v5 = v3
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _7
_7:
	t1 = v6
	v8 = t1
	v9 = v2.S
	v10 = v3
	v11 = Vec3{
		X: v8.X + float32(v9*v10.X),
		Y: v8.Y + float32(v9*v10.Y),
		Z: v8.Z + float32(v9*v10.Z),
	}
	goto _12
_12:
	t2 = v11
	v13 = v2.V
	v14 = t2
	c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
	c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
	c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	v15 = c
	goto _16
_16:
	t3 = v15
	v19 = v3
	v20 = float32FromFloat32(2)
	v21 = t3
	v22 = Vec3{
		X: v19.X + float32(v20*v21.X),
		Y: v19.Y + float32(v20*v21.Y),
		Z: v19.Z + float32(v20*v21.Z),
	}
	goto _23
_23:
	v17 = v22
	goto _18
_18:
	rv = v17
	v26 = rv
	v27 = v1.P
	v28 = Vec3{
		X: v26.X + v27.X,
		Y: v26.Y + v27.Y,
		Z: v26.Z + v27.Z,
	}
	goto _29
_29:
	v24 = v28
	goto _25
_25:
	center = v24
	radius = (*Sphere)(unsafe.Pointer(shape)).Radius
	extent = Vec3{
		X: radius,
		Y: radius,
		Z: radius,
	}
	v30 = center
	v31 = extent
	v32 = Vec3{
		X: v30.X - v31.X,
		Y: v30.Y - v31.Y,
		Z: v30.Z - v31.Z,
	}
	goto _33
_33:
	v34 = center
	v35 = extent
	v36 = Vec3{
		X: v34.X + v35.X,
		Y: v34.Y + v35.Y,
		Z: v34.Z + v35.Z,
	}
	goto _37
_37:
	return AABB{
		LowerBound: v32,
		UpperBound: v36,
	}
}

func b3ComputeSweptSphereAABB(tls *_Stack, shape uintptr, xf1 Transform, xf2 Transform) (r1 AABB) {
	var c, center1, center2, r, rv, t1, t2, t3, v10, v101, v102, v103, v11, v13, v14, v15, v17, v19, v21, v22, v24, v26, v27, v28, v3, v32, v33, v34, v35, v37, v39, v4, v40, v42, v43, v44, v46, v48, v5, v50, v51, v53, v55, v56, v57, v59, v6, v60, v61, v78, v79, v8, v80, v82, v83, v84 Vec3
	var v1, v30 Transform
	var v100, v20, v38, v49, v63, v64, v65, v67, v68, v69, v70, v72, v73, v74, v75, v77, v86, v87, v88, v9, v90, v91, v92, v93, v95, v96, v97, v98 float32
	var v2, v31 Quat
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = c, center1, center2, r, rv, t1, t2, t3, v1, v10, v100, v101, v102, v103, v11, v13, v14, v15, v17, v19, v2, v20, v21, v22, v24, v26, v27, v28, v3, v30, v31, v32, v33, v34, v35, v37, v38, v39, v4, v40, v42, v43, v44, v46, v48, v49, v5, v50, v51, v53, v55, v56, v57, v59, v6, v60, v61, v63, v64, v65, v67, v68, v69, v70, v72, v73, v74, v75, v77, v78, v79, v8, v80, v82, v83, v84, v86, v87, v88, v9, v90, v91, v92, v93, v95, v96, v97, v98
	r = Vec3{
		X: (*Sphere)(unsafe.Pointer(shape)).Radius,
		Y: (*Sphere)(unsafe.Pointer(shape)).Radius,
		Z: (*Sphere)(unsafe.Pointer(shape)).Radius,
	}
	v1 = xf1
	v2 = v1.Q
	v3 = (*Sphere)(unsafe.Pointer(shape)).Center
	v4 = v2.V
	v5 = v3
	c.X = float32(v4.Y*v5.Z) - float32(v4.Z*v5.Y)
	c.Y = float32(v4.Z*v5.X) - float32(v4.X*v5.Z)
	c.Z = float32(v4.X*v5.Y) - float32(v4.Y*v5.X)
	v6 = c
	goto _7
_7:
	t1 = v6
	v8 = t1
	v9 = v2.S
	v10 = v3
	v11 = Vec3{
		X: v8.X + float32(v9*v10.X),
		Y: v8.Y + float32(v9*v10.Y),
		Z: v8.Z + float32(v9*v10.Z),
	}
	goto _12
_12:
	t2 = v11
	v13 = v2.V
	v14 = t2
	c.X = float32(v13.Y*v14.Z) - float32(v13.Z*v14.Y)
	c.Y = float32(v13.Z*v14.X) - float32(v13.X*v14.Z)
	c.Z = float32(v13.X*v14.Y) - float32(v13.Y*v14.X)
	v15 = c
	goto _16
_16:
	t3 = v15
	v19 = v3
	v20 = float32FromFloat32(2)
	v21 = t3
	v22 = Vec3{
		X: v19.X + float32(v20*v21.X),
		Y: v19.Y + float32(v20*v21.Y),
		Z: v19.Z + float32(v20*v21.Z),
	}
	goto _23
_23:
	v17 = v22
	goto _18
_18:
	rv = v17
	v26 = rv
	v27 = v1.P
	v28 = Vec3{
		X: v26.X + v27.X,
		Y: v26.Y + v27.Y,
		Z: v26.Z + v27.Z,
	}
	goto _29
_29:
	v24 = v28
	goto _25
_25:
	center1 = v24
	v30 = xf2
	v31 = v30.Q
	v32 = (*Sphere)(unsafe.Pointer(shape)).Center
	v33 = v31.V
	v34 = v32
	c.X = float32(v33.Y*v34.Z) - float32(v33.Z*v34.Y)
	c.Y = float32(v33.Z*v34.X) - float32(v33.X*v34.Z)
	c.Z = float32(v33.X*v34.Y) - float32(v33.Y*v34.X)
	v35 = c
	goto _36
_36:
	t1 = v35
	v37 = t1
	v38 = v31.S
	v39 = v32
	v40 = Vec3{
		X: v37.X + float32(v38*v39.X),
		Y: v37.Y + float32(v38*v39.Y),
		Z: v37.Z + float32(v38*v39.Z),
	}
	goto _41
_41:
	t2 = v40
	v42 = v31.V
	v43 = t2
	c.X = float32(v42.Y*v43.Z) - float32(v42.Z*v43.Y)
	c.Y = float32(v42.Z*v43.X) - float32(v42.X*v43.Z)
	c.Z = float32(v42.X*v43.Y) - float32(v42.Y*v43.X)
	v44 = c
	goto _45
_45:
	t3 = v44
	v48 = v32
	v49 = float32FromFloat32(2)
	v50 = t3
	v51 = Vec3{
		X: v48.X + float32(v49*v50.X),
		Y: v48.Y + float32(v49*v50.Y),
		Z: v48.Z + float32(v49*v50.Z),
	}
	goto _52
_52:
	v46 = v51
	goto _47
_47:
	rv = v46
	v55 = rv
	v56 = v30.P
	v57 = Vec3{
		X: v55.X + v56.X,
		Y: v55.Y + v56.Y,
		Z: v55.Z + v56.Z,
	}
	goto _58
_58:
	v53 = v57
	goto _54
_54:
	center2 = v53
	v59 = center1
	v60 = center2
	v63 = v59.X
	v64 = v60.X
	if v63 < v64 {
		v67 = v63
	} else {
		v67 = v64
	}
	v65 = v67
	goto _66
_66:
	v68 = v59.Y
	v69 = v60.Y
	if v68 < v69 {
		v72 = v68
	} else {
		v72 = v69
	}
	v70 = v72
	goto _71
_71:
	v73 = v59.Z
	v74 = v60.Z
	if v73 < v74 {
		v77 = v73
	} else {
		v77 = v74
	}
	v75 = v77
	goto _76
_76:
	v61 = Vec3{
		X: v65,
		Y: v70,
		Z: v75,
	}
	goto _62
_62:
	v78 = v61
	v79 = r
	v80 = Vec3{
		X: v78.X - v79.X,
		Y: v78.Y - v79.Y,
		Z: v78.Z - v79.Z,
	}
	goto _81
_81:
	v82 = center1
	v83 = center2
	v86 = v82.X
	v87 = v83.X
	if v86 > v87 {
		v90 = v86
	} else {
		v90 = v87
	}
	v88 = v90
	goto _89
_89:
	v91 = v82.Y
	v92 = v83.Y
	if v91 > v92 {
		v95 = v91
	} else {
		v95 = v92
	}
	v93 = v95
	goto _94
_94:
	v96 = v82.Z
	v97 = v83.Z
	if v96 > v97 {
		v100 = v96
	} else {
		v100 = v97
	}
	v98 = v100
	goto _99
_99:
	v84 = Vec3{
		X: v88,
		Y: v93,
		Z: v98,
	}
	goto _85
_85:
	v101 = v84
	v102 = r
	v103 = Vec3{
		X: v101.X + v102.X,
		Y: v101.Y + v102.Y,
		Z: v101.Z + v102.Z,
	}
	goto _104
_104:
	return AABB{
		LowerBound: v80,
		UpperBound: v103,
	}
}

// Temporary abbreviations for convenience.

// C documentation
//
//	// SIMD separating axis test based on an implementation developed by Cairn Overturf.
//	// See his article: https://cairnc.github.io/posts/improvements-to-the-separating-axis/
func b3ComputeSeparatingAxisW4(tls *_Stack, hullA uintptr, hullB uintptr, xfB Transform, earlyReturn uint8) (r b3AxisQuery) {
	bp := tls.Alloc(13200)
	defer tls.Free(13200)
	var A, B, C, D, DC, av0, b6, bv0, bv1, cA, cB, centerAinB, centerBinA, d, deltaCenter, direction, direction1, direction2, direction3, hA, hB, n, n0, n1, probeA, probeB, v16, v18, v19, v20, v23, v24, v27, v29, v30, v31, v34, v35, v38, v39, v41, v42, v43, v45, v46, v47, v49, v52, v53 Vec3
	var ADC, BDC, CBA, DBA, EPS, boundSlack, c, c1, centerDistance, floorA, floorB, inv, len2, linearSlop, maxCD, maxDotA, maxDotB, maxFaceSeparation, planeSeparation, planeSeparation1, planeSeparation2, planeSeparation3, probeBound, radius, radiusBound, seedSeparationA, seedSeparationB, separation, separation1, separation2, speculativeDistance, squaredTol, t, thresholdA, thresholdB, xw, xx, xy, xz, yw, yy, yz, zw, zz, v22, v33, v50, v54, v66, v68 float32
	var R, invR, out, v2, v37, v4, v5 Matrix3
	var cosinesA, cosinesB, edge, edge1, halfEdgesA, halfEdgesB, nxB, nyB, nzB, planesA, planesB, twin, twin1, vxA, vxB, vyA, vyB, vzA, vzB, v121, v13, v7, v8 uintptr
	var f0, f1, faceCountA, faceCountB, faceStrideB, halfEdgeCountA, halfEdgeCountB, i, i1, i11, i12, i2, i21, i22, i3, i4, i5, i6, j, k, na, nb, seedIndexA, seedIndexB, v0, v01, v11, v12, vertexA, vertexB, vertexStrideA, vertexStrideB, wideFaceCountB, wideVertexCountA, wideVertexCountB, v10 int32
	var plane, plane1, plane2, plane3 Plane
	var res b3AxisQuery
	var zero b3FloatW
	var zeroV b3Vec3W
	var v1 Quat
	var v15, v26 AABB
	var v231 b3Vec3W4
	var _ /* aDx at bp+9504 */ [132]float32
	var _ /* aDy at bp+10032 */ [132]float32
	var _ /* aDz at bp+10560 */ [132]float32
	var _ /* aN0x at bp+6336 */ [132]float32
	var _ /* aN0y at bp+6864 */ [132]float32
	var _ /* aN0z at bp+7392 */ [132]float32
	var _ /* aN1x at bp+7920 */ [132]float32
	var _ /* aN1y at bp+8448 */ [132]float32
	var _ /* aN1z at bp+8976 */ [132]float32
	var _ /* aTol at bp+12672 */ [132]float32
	var _ /* aV0x at bp+11088 */ [132]float32
	var _ /* aV0y at bp+11616 */ [132]float32
	var _ /* aV0z at bp+12144 */ [132]float32
	var _ /* bFNx at bp+3168 */ [132]float32
	var _ /* bFNy at bp+3696 */ [132]float32
	var _ /* bFNz at bp+4224 */ [132]float32
	var _ /* bWx at bp+4752 */ [132]float32
	var _ /* bWy at bp+5280 */ [132]float32
	var _ /* bWz at bp+5808 */ [132]float32
	var _ /* dotA at bp+0 */ [132]float32
	var _ /* dotB at bp+536 */ [132]float32
	var _ /* edgeIndicesA at bp+1072 */ [132]int32
	var _ /* edgeIndicesB at bp+1600 */ [128]int32
	var _ /* planeDotA at bp+2112 */ [132]float32
	var _ /* planeDotB at bp+2640 */ [132]float32
	var _ /* seedVertexA at bp+1064 */ int32
	var _ /* seedVertexB at bp+528 */ int32
	var _ /* vertexIndex at bp+1068 */ int32
	var _ /* vertexIndex at bp+532 */ int32
	_, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _, _ = A, ADC, B, BDC, C, CBA, D, DBA, DC, EPS, R, av0, b6, boundSlack, bv0, bv1, c, c1, cA, cB, centerAinB, centerBinA, centerDistance, cosinesA, cosinesB, d, deltaCenter, direction, direction1, direction2, direction3, edge, edge1, f0, f1, faceCountA, faceCountB, faceStrideB, floorA, floorB, hA, hB, halfEdgeCountA, halfEdgeCountB, halfEdgesA, halfEdgesB, i, i1, i11, i12, i2, i21, i22, i3, i4, i5, i6, inv, invR, j, k, len2, linearSlop, maxCD, maxDotA, maxDotB, maxFaceSeparation, n, n0, n1, na, nb, nxB, nyB, nzB, out, plane, plane1, plane2, plane3, planeSeparation, planeSeparation1, planeSeparation2, planeSeparation3, planesA, planesB, probeA, probeB, probeBound, radius, radiusBound, res, seedIndexA, seedIndexB, seedSeparationA, seedSeparationB, separation, separation1, separation2, speculativeDistance, squaredTol, t, thresholdA, thresholdB, twin, twin1, v0, v01, v11, v12, vertexA, vertexB, vertexStrideA, vertexStrideB, vxA, vxB, vyA, vyB, vzA, vzB, wideFaceCountB, wideVertexCountA, wideVertexCountB, xw, xx, xy, xz, yw, yy, yz, zero, zeroV, zw, zz, v1, v10, v121, v13, v15, v16, v18, v19, v2, v20, v22, v23, v231, v24, v26, v27, v29, v30, v31, v33, v34, v35, v37, v38, v39, v4, v41, v42, v43, v45, v46, v47, v49, v5, v50, v52, v53, v54, v66, v68, v7, v8
	v1 = xfB.Q
	xx = float32(v1.V.X * v1.V.X)
	yy = float32(v1.V.Y * v1.V.Y)
	zz = float32(v1.V.Z * v1.V.Z)
	xy = float32(v1.V.X * v1.V.Y)
	xz = float32(v1.V.X * v1.V.Z)
	xw = float32(v1.V.X * v1.S)
	yz = float32(v1.V.Y * v1.V.Z)
	yw = float32(v1.V.Y * v1.S)
	zw = float32(v1.V.Z * v1.S)
	v2 = Matrix3{
		Cx: Vec3{
			X: float32FromFloat32(1) - float32(float32FromFloat32(2)*(yy+zz)),
			Y: float32(float32FromFloat32(2) * (xy + zw)),
			Z: float32(float32FromFloat32(2) * (xz - yw)),
		},
		Cy: Vec3{
			X: float32(float32FromFloat32(2) * (xy - zw)),
			Y: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+zz)),
			Z: float32(float32FromFloat32(2) * (yz + xw)),
		},
		Cz: Vec3{
			X: float32(float32FromFloat32(2) * (xz + yw)),
			Y: float32(float32FromFloat32(2) * (yz - xw)),
			Z: float32FromFloat32(1) - float32(float32FromFloat32(2)*(xx+yy)),
		},
	}
	goto _3
_3:
	R = v2
	v4 = R
	out.Cx = Vec3{
		X: v4.Cx.X,
		Y: v4.Cy.X,
		Z: v4.Cz.X,
	}
	out.Cy = Vec3{
		X: v4.Cx.Y,
		Y: v4.Cy.Y,
		Z: v4.Cz.Y,
	}
	out.Cz = Vec3{
		X: v4.Cx.Z,
		Y: v4.Cy.Z,
		Z: v4.Cz.Z,
	}
	v5 = out
	goto _6
_6:
	invR = v5
	speculativeDistance = float32(float32FromFloat32(4) * float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)))
	res = b3AxisQuery{
		FaceA: b3SeparatingAxis{
			Normal:     b3Vec3_zero,
			Separation: -__builtin_inff(tls),
			IndexA:     -int32(1),
			IndexB:     -int32(1),
			Type1:      int32(b3_faceAxisA),
		},
		FaceB: b3SeparatingAxis{
			Normal:     b3Vec3_zero,
			Separation: -__builtin_inff(tls),
			IndexA:     -int32(1),
			IndexB:     -int32(1),
			Type1:      int32(b3_faceAxisB),
		},
		Edge: b3SeparatingAxis{
			Normal:     b3Vec3_zero,
			Separation: -__builtin_inff(tls),
			IndexA:     -int32(1),
			IndexB:     -int32(1),
			Type1:      int32(b3_edgePairAxis),
		},
	}
	faceCountA = (*HullData)(unsafe.Pointer(hullA)).FaceCount
	v7 = hullA
	if (*HullData)(unsafe.Pointer(v7)).PlaneOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _9
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).PlaneOffset))
	goto _9
_9:
	planesA = v8
	v10 = ((*HullData)(unsafe.Pointer(hullB)).VertexCount + int32(_B3_HULL_SOA_PADDING) - int32(1)) & ^(int32FromInt32(_B3_HULL_SOA_PADDING) - int32FromInt32(1))
	goto _11
_11:
	vertexStrideB = v10
	v121 = hullB
	if (*HullData)(unsafe.Pointer(v121)).SoaVertexOffset == int32FromInt32(0) {
		v13 = uintptrFromInt32(0)
		goto _14
	}
	v13 = uintptr(int64(v121) + int64((*HullData)(unsafe.Pointer(v121)).SoaVertexOffset))
	goto _14
_14:
	vxB = v13
	vyB = vxB + uintptr(vertexStrideB)*4
	vzB = vyB + uintptr(vertexStrideB)*4
	wideVertexCountB = b3RoundUpToSIMDWidth(tls, (*HullData)(unsafe.Pointer(hullB)).VertexCount)
	v15 = (*HullData)(unsafe.Pointer(hullB)).Aabb
	v18 = v15.UpperBound
	v19 = v15.LowerBound
	v20 = Vec3{
		X: v18.X + v19.X,
		Y: v18.Y + v19.Y,
		Z: v18.Z + v19.Z,
	}
	goto _21
_21:
	v22 = float32FromFloat32(0.5)
	v23 = v20
	v24 = Vec3{
		X: float32(v22 * v23.X),
		Y: float32(v22 * v23.Y),
		Z: float32(v22 * v23.Z),
	}
	goto _25
_25:
	v16 = v24
	goto _17
_17:
	cB = v16
	v26 = (*HullData)(unsafe.Pointer(hullB)).Aabb
	v29 = v26.UpperBound
	v30 = v26.LowerBound
	v31 = Vec3{
		X: v29.X - v30.X,
		Y: v29.Y - v30.Y,
		Z: v29.Z - v30.Z,
	}
	goto _32
_32:
	v33 = float32FromFloat32(0.5)
	v34 = v31
	v35 = Vec3{
		X: float32(v33 * v34.X),
		Y: float32(v33 * v34.Y),
		Z: float32(v33 * v34.Z),
	}
	goto _36
_36:
	v27 = v35
	goto _28
_28:
	hB = v27
	v37 = R
	v38 = (*HullData)(unsafe.Pointer(hullB)).Center
	b6 = Vec3{
		X: float32(v37.Cx.X*v38.X) + float32(v37.Cy.X*v38.Y) + float32(v37.Cz.X*v38.Z),
		Y: float32(v37.Cx.Y*v38.X) + float32(v37.Cy.Y*v38.Y) + float32(v37.Cz.Y*v38.Z),
		Z: float32(v37.Cx.Z*v38.X) + float32(v37.Cy.Z*v38.Y) + float32(v37.Cz.Z*v38.Z),
	}
	v39 = b6
	goto _40
_40:
	v41 = v39
	v42 = xfB.P
	v43 = Vec3{
		X: v41.X + v42.X,
		Y: v41.Y + v42.Y,
		Z: v41.Z + v42.Z,
	}
	goto _44
_44:
	v45 = v43
	v46 = (*HullData)(unsafe.Pointer(hullA)).Center
	v47 = Vec3{
		X: v45.X - v46.X,
		Y: v45.Y - v46.Y,
		Z: v45.Z - v46.Z,
	}
	goto _48
_48:
	// The hulls have a precomputed inner radius and centroid.
	// A given axis cannot achieve a separation larger than:
	// dot(axis, centerB - centerA) - innerRadiusA - innerRadiusB
	// So this is the upper bound for the separation of a candidate axis.
	// An axis can be skipped if it has an upper bound that is less than the current
	// best separation. This lets me skip many of the candidates without computing
	// support points.
	deltaCenter = v47
	v49 = deltaCenter
	v52 = v49
	v53 = v49
	v54 = float32(v52.X*v53.X) + float32(v52.Y*v53.Y) + float32(v52.Z*v53.Z)
	goto _55
_55:
	v50 = sqrtf(tls, v54)
	goto _51
_51:
	centerDistance = v50
	radius = (*HullData)(unsafe.Pointer(hullA)).InnerRadius + (*HullData)(unsafe.Pointer(hullB)).InnerRadius
	// Adjust the radius to ensure the best axis isn't skipped.
	radiusBound = radius - (float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls)) + float32(float32FromFloat32(0.001)*(centerDistance+radius)))
	b3GetFaceDots(tls, hullA, deltaCenter, bp)
	// Find the face of hullA that most aligns with centerDelta.
	seedIndexA = 0
	maxDotA = (*(*[132]float32)(unsafe.Add(nil, bp)))[0]
	i = int32(1)
	for {
		if !(i < faceCountA) {
			break
		}
		if (*(*[132]float32)(unsafe.Add(nil, bp)))[i] > maxDotA {
			maxDotA = (*(*[132]float32)(unsafe.Add(nil, bp)))[i]
			seedIndexA = i
		}
		goto _56
	_56:
		;
		i = i + 1
	}
	// Use the seed to get a lower bound on the separation for the faces of hullA.
	floorA = -__builtin_inff(tls)
	seedSeparationA = -__builtin_inff(tls)
	*(*int32)(unsafe.Add(nil, bp+528)) = 0
	if earlyReturn != 0 {
		plane = *(*Plane)(unsafe.Add(nil, planesA+uintptr(seedIndexA)*16))
		v2 = invR
		v16 = plane.Normal
		b6 = Vec3{
			X: float32(v2.Cx.X*v16.X) + float32(v2.Cy.X*v16.Y) + float32(v2.Cz.X*v16.Z),
			Y: float32(v2.Cx.Y*v16.X) + float32(v2.Cy.Y*v16.Y) + float32(v2.Cz.Y*v16.Z),
			Z: float32(v2.Cx.Z*v16.X) + float32(v2.Cy.Z*v16.Y) + float32(v2.Cz.Z*v16.Z),
		}
		v18 = b6
		goto _60
	_60:
		v19 = v18
		v20 = Vec3{
			X: -v19.X,
			Y: -v19.Y,
			Z: -v19.Z,
		}
		goto _63
	_63:
		direction = v20
		v23 = plane.Normal
		v24 = xfB.P
		v22 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y) + float32(v23.Z*v24.Z)
		goto _67
	_67:
		planeSeparation = v22 - plane.Offset
		seedSeparationA = b3GetFaceSeparation(tls, direction, planeSeparation, vxB, vyB, vzB, wideVertexCountB, cB, hB, bp+528)
		v22 = seedSeparationA
		v33 = speculativeDistance
		if v22 < v33 {
			v54 = v22
		} else {
			v54 = v33
		}
		v50 = v54
		goto _71
	_71:
		floorA = v50
	}
	// Test A's face planes against B's vertices.
	i1 = 0
	for {
		if !(i1 < faceCountA) {
			break
		}
		// The bound offset ensures the seed will be evaluated.
		v22 = floorA
		v33 = res.FaceA.Separation
		if v22 > v33 {
			v54 = v22
		} else {
			v54 = v33
		}
		v50 = v54
		goto _77
	_77:
		if (*(*[132]float32)(unsafe.Add(nil, bp)))[i1]-radiusBound < v50 {
			goto _73
		}
		plane1 = *(*Plane)(unsafe.Add(nil, planesA+uintptr(i1)*16))
		*(*int32)(unsafe.Add(nil, bp+532)) = *(*int32)(unsafe.Add(nil, bp+528))
		separation = seedSeparationA
		// Avoid recomputing the seed face separation.
		if int32FromUint8(earlyReturn) == false1 || i1 != seedIndexA {
			v2 = invR
			v16 = plane1.Normal
			b6 = Vec3{
				X: float32(v2.Cx.X*v16.X) + float32(v2.Cy.X*v16.Y) + float32(v2.Cz.X*v16.Z),
				Y: float32(v2.Cx.Y*v16.X) + float32(v2.Cy.Y*v16.Y) + float32(v2.Cz.Y*v16.Z),
				Z: float32(v2.Cx.Z*v16.X) + float32(v2.Cy.Z*v16.Y) + float32(v2.Cz.Z*v16.Z),
			}
			v18 = b6
			goto _82
		_82:
			v19 = v18
			v20 = Vec3{
				X: -v19.X,
				Y: -v19.Y,
				Z: -v19.Z,
			}
			goto _85
		_85:
			direction1 = v20
			v23 = plane1.Normal
			v24 = xfB.P
			v22 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y) + float32(v23.Z*v24.Z)
			goto _89
		_89:
			planeSeparation1 = v22 - plane1.Offset
			separation = b3GetFaceSeparation(tls, direction1, planeSeparation1, vxB, vyB, vzB, wideVertexCountB, cB, hB, bp+532)
		}
		if separation > res.FaceA.Separation {
			res.FaceA.Normal = plane1.Normal
			res.FaceA.Separation = separation
			res.FaceA.IndexA = i1
			res.FaceA.IndexB = *(*int32)(unsafe.Add(nil, bp+532))
			if separation > speculativeDistance && earlyReturn != 0 {
				res.SeparatedFeature = int32(b3_faceAxisA)
				return res
			}
		}
		goto _73
	_73:
		;
		i1 = i1 + 1
	}
	faceCountB = (*HullData)(unsafe.Pointer(hullB)).FaceCount
	v7 = hullB
	if (*HullData)(unsafe.Pointer(v7)).PlaneOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _92
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).PlaneOffset))
	goto _92
_92:
	planesB = v8
	v10 = ((*HullData)(unsafe.Pointer(hullA)).VertexCount + int32(_B3_HULL_SOA_PADDING) - int32(1)) & ^(int32FromInt32(_B3_HULL_SOA_PADDING) - int32FromInt32(1))
	goto _94
_94:
	vertexStrideA = v10
	v121 = hullA
	if (*HullData)(unsafe.Pointer(v121)).SoaVertexOffset == int32FromInt32(0) {
		v13 = uintptrFromInt32(0)
		goto _97
	}
	v13 = uintptr(int64(v121) + int64((*HullData)(unsafe.Pointer(v121)).SoaVertexOffset))
	goto _97
_97:
	vxA = v13
	vyA = vxA + uintptr(vertexStrideA)*4
	vzA = vyA + uintptr(vertexStrideA)*4
	wideVertexCountA = b3RoundUpToSIMDWidth(tls, (*HullData)(unsafe.Pointer(hullA)).VertexCount)
	v15 = (*HullData)(unsafe.Pointer(hullA)).Aabb
	v18 = v15.UpperBound
	v19 = v15.LowerBound
	v20 = Vec3{
		X: v18.X + v19.X,
		Y: v18.Y + v19.Y,
		Z: v18.Z + v19.Z,
	}
	goto _104
_104:
	v22 = float32FromFloat32(0.5)
	v23 = v20
	v24 = Vec3{
		X: float32(v22 * v23.X),
		Y: float32(v22 * v23.Y),
		Z: float32(v22 * v23.Z),
	}
	goto _108
_108:
	v16 = v24
	goto _100
_100:
	cA = v16
	v26 = (*HullData)(unsafe.Pointer(hullA)).Aabb
	v29 = v26.UpperBound
	v30 = v26.LowerBound
	v31 = Vec3{
		X: v29.X - v30.X,
		Y: v29.Y - v30.Y,
		Z: v29.Z - v30.Z,
	}
	goto _115
_115:
	v33 = float32FromFloat32(0.5)
	v34 = v31
	v35 = Vec3{
		X: float32(v33 * v34.X),
		Y: float32(v33 * v34.Y),
		Z: float32(v33 * v34.Z),
	}
	goto _119
_119:
	v27 = v35
	goto _111
_111:
	hA = v27
	v2 = invR
	v38 = deltaCenter
	b6 = Vec3{
		X: float32(v2.Cx.X*v38.X) + float32(v2.Cy.X*v38.Y) + float32(v2.Cz.X*v38.Z),
		Y: float32(v2.Cx.Y*v38.X) + float32(v2.Cy.Y*v38.Y) + float32(v2.Cz.Y*v38.Z),
		Z: float32(v2.Cx.Z*v38.X) + float32(v2.Cy.Z*v38.Y) + float32(v2.Cz.Z*v38.Z),
	}
	v39 = b6
	goto _123
_123:
	v41 = v39
	v42 = Vec3{
		X: -v41.X,
		Y: -v41.Y,
		Z: -v41.Z,
	}
	goto _126
_126:
	b3GetFaceDots(tls, hullB, v42, bp+536)
	seedIndexB = 0
	maxDotB = (*(*[132]float32)(unsafe.Add(nil, bp+536)))[0]
	i2 = int32(1)
	for {
		if !(i2 < faceCountB) {
			break
		}
		if (*(*[132]float32)(unsafe.Add(nil, bp+536)))[i2] > maxDotB {
			maxDotB = (*(*[132]float32)(unsafe.Add(nil, bp+536)))[i2]
			seedIndexB = i2
		}
		goto _127
	_127:
		;
		i2 = i2 + 1
	}
	// Get a lower bound on the separation for the faces of hullB.
	floorB = -__builtin_inff(tls)
	seedSeparationB = -__builtin_inff(tls)
	*(*int32)(unsafe.Add(nil, bp+1064)) = 0
	if earlyReturn != 0 {
		plane2 = *(*Plane)(unsafe.Add(nil, planesB+uintptr(seedIndexB)*16))
		v2 = R
		v16 = plane2.Normal
		b6 = Vec3{
			X: float32(v2.Cx.X*v16.X) + float32(v2.Cy.X*v16.Y) + float32(v2.Cz.X*v16.Z),
			Y: float32(v2.Cx.Y*v16.X) + float32(v2.Cy.Y*v16.Y) + float32(v2.Cz.Y*v16.Z),
			Z: float32(v2.Cx.Z*v16.X) + float32(v2.Cy.Z*v16.Y) + float32(v2.Cz.Z*v16.Z),
		}
		v18 = b6
		goto _131
	_131:
		v19 = v18
		v20 = Vec3{
			X: -v19.X,
			Y: -v19.Y,
			Z: -v19.Z,
		}
		goto _134
	_134:
		direction2 = v20
		v23 = direction2
		v24 = xfB.P
		v22 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y) + float32(v23.Z*v24.Z)
		goto _138
	_138:
		planeSeparation2 = v22 - plane2.Offset
		seedSeparationB = b3GetFaceSeparation(tls, direction2, planeSeparation2, vxA, vyA, vzA, wideVertexCountA, cA, hA, bp+1064)
		// Include the floor set by hull A faces.
		v22 = seedSeparationB
		v33 = res.FaceA.Separation
		if v22 > v33 {
			v54 = v22
		} else {
			v54 = v33
		}
		v50 = v54
		goto _142
	_142:
		floorB = v50
		v22 = floorB
		v33 = speculativeDistance
		if v22 < v33 {
			v54 = v22
		} else {
			v54 = v33
		}
		v50 = v54
		goto _147
	_147:
		floorB = v50
	}
	// Test B's face planes against A's vertices.
	i3 = 0
	for {
		if !(i3 < faceCountB) {
			break
		}
		v22 = floorB
		v33 = res.FaceB.Separation
		if v22 > v33 {
			v54 = v22
		} else {
			v54 = v33
		}
		v50 = v54
		goto _153
	_153:
		if (*(*[132]float32)(unsafe.Add(nil, bp+536)))[i3]-radiusBound < v50 {
			goto _149
		}
		plane3 = *(*Plane)(unsafe.Add(nil, planesB+uintptr(i3)*16))
		v2 = R
		v16 = plane3.Normal
		b6 = Vec3{
			X: float32(v2.Cx.X*v16.X) + float32(v2.Cy.X*v16.Y) + float32(v2.Cz.X*v16.Z),
			Y: float32(v2.Cx.Y*v16.X) + float32(v2.Cy.Y*v16.Y) + float32(v2.Cz.Y*v16.Z),
			Z: float32(v2.Cx.Z*v16.X) + float32(v2.Cy.Z*v16.Y) + float32(v2.Cz.Z*v16.Z),
		}
		v18 = b6
		goto _158
	_158:
		v19 = v18
		v20 = Vec3{
			X: -v19.X,
			Y: -v19.Y,
			Z: -v19.Z,
		}
		goto _161
	_161:
		direction3 = v20
		*(*int32)(unsafe.Add(nil, bp+1068)) = *(*int32)(unsafe.Add(nil, bp+1064))
		separation1 = seedSeparationB
		// Avoid recomputing the seed face separation.
		if int32FromUint8(earlyReturn) == false1 || i3 != seedIndexB {
			v23 = direction3
			v24 = xfB.P
			v22 = float32(v23.X*v24.X) + float32(v23.Y*v24.Y) + float32(v23.Z*v24.Z)
			goto _165
		_165:
			planeSeparation3 = v22 - plane3.Offset
			separation1 = b3GetFaceSeparation(tls, direction3, planeSeparation3, vxA, vyA, vzA, wideVertexCountA, cA, hA, bp+1068)
		}
		if separation1 > res.FaceB.Separation {
			res.FaceB.Normal = direction3
			res.FaceB.Separation = separation1
			res.FaceB.IndexA = *(*int32)(unsafe.Add(nil, bp+1068))
			res.FaceB.IndexB = i3
			if separation1 > speculativeDistance && earlyReturn != 0 {
				res.SeparatedFeature = int32(b3_faceAxisB)
				return res
			}
		}
		goto _149
	_149:
		;
		i3 = i3 + 1
	}
	// Gather edges of A that can feasibly create a separating axis that beats the maximum face separation.
	halfEdgeCountA = (*HullData)(unsafe.Pointer(hullA)).EdgeCount
	v7 = hullA
	if (*HullData)(unsafe.Pointer(v7)).EdgeOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _168
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).EdgeOffset))
	goto _168
_168:
	halfEdgesA = v8
	na = 0
	halfEdgeCountB = (*HullData)(unsafe.Pointer(hullB)).EdgeCount
	v121 = hullB
	if (*HullData)(unsafe.Pointer(v121)).EdgeOffset == int32FromInt32(0) {
		v13 = uintptrFromInt32(0)
		goto _171
	}
	v13 = uintptr(int64(v121) + int64((*HullData)(unsafe.Pointer(v121)).EdgeOffset))
	goto _171
_171:
	halfEdgesB = v13
	nb = 0
	v22 = res.FaceA.Separation
	v33 = res.FaceB.Separation
	if v22 > v33 {
		v54 = v22
	} else {
		v54 = v33
	}
	v50 = v54
	goto _175
_175:
	maxFaceSeparation = v50
	boundSlack = radius - radiusBound
	if earlyReturn != 0 {
		v66 = maxFaceSeparation + (*HullData)(unsafe.Pointer(hullB)).InnerRadius - boundSlack
	} else {
		v66 = -__builtin_inff(tls)
	}
	thresholdA = v66
	if earlyReturn != 0 {
		v68 = maxFaceSeparation + (*HullData)(unsafe.Pointer(hullA)).InnerRadius - boundSlack
	} else {
		v68 = -__builtin_inff(tls)
	}
	thresholdB = v68
	v2 = R
	v16 = (*HullData)(unsafe.Pointer(hullB)).Center
	b6 = Vec3{
		X: float32(v2.Cx.X*v16.X) + float32(v2.Cy.X*v16.Y) + float32(v2.Cz.X*v16.Z),
		Y: float32(v2.Cx.Y*v16.X) + float32(v2.Cy.Y*v16.Y) + float32(v2.Cz.Y*v16.Z),
		Z: float32(v2.Cx.Z*v16.X) + float32(v2.Cy.Z*v16.Y) + float32(v2.Cz.Z*v16.Z),
	}
	v18 = b6
	goto _182
_182:
	v19 = v18
	v20 = xfB.P
	v23 = Vec3{
		X: v19.X + v20.X,
		Y: v19.Y + v20.Y,
		Z: v19.Z + v20.Z,
	}
	goto _186
_186:
	centerBinA = v23
	v24 = (*HullData)(unsafe.Pointer(hullA)).Center
	v27 = xfB.P
	v29 = Vec3{
		X: v24.X - v27.X,
		Y: v24.Y - v27.Y,
		Z: v24.Z - v27.Z,
	}
	goto _190
_190:
	v4 = invR
	v30 = v29
	b6 = Vec3{
		X: float32(v4.Cx.X*v30.X) + float32(v4.Cy.X*v30.Y) + float32(v4.Cz.X*v30.Z),
		Y: float32(v4.Cx.Y*v30.X) + float32(v4.Cy.Y*v30.Y) + float32(v4.Cz.Y*v30.Z),
		Z: float32(v4.Cx.Z*v30.X) + float32(v4.Cy.Z*v30.Y) + float32(v4.Cz.Z*v30.Z),
	}
	v31 = b6
	goto _194
_194:
	centerAinB = v31
	b3GetFacePlaneSeparations(tls, hullA, centerBinA, bp+2112)
	b3GetFacePlaneSeparations(tls, hullB, centerAinB, bp+2640)
	v7 = hullA
	if (*HullData)(unsafe.Pointer(v7)).EdgeCosineOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _197
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).EdgeCosineOffset))
	goto _197
_197:
	cosinesA = v8
	v121 = hullB
	if (*HullData)(unsafe.Pointer(v121)).EdgeCosineOffset == int32FromInt32(0) {
		v13 = uintptrFromInt32(0)
		goto _200
	}
	v13 = uintptr(int64(v121) + int64((*HullData)(unsafe.Pointer(v121)).EdgeCosineOffset))
	goto _200
_200:
	cosinesB = v13
	i4 = 0
	for {
		if !(i4 < halfEdgeCountA) {
			break
		}
		i11 = int32FromUint8((*(*HullHalfEdge)(unsafe.Add(nil, halfEdgesA+uintptr(i4)*4))).Face)
		i21 = int32FromUint8((*(*HullHalfEdge)(unsafe.Add(nil, halfEdgesA+uintptr(i4+int32(1))*4))).Face)
		c = *(*float32)(unsafe.Add(nil, cosinesA+uintptr(i4>>int32(1))*4))
		(*(*[132]int32)(unsafe.Add(nil, bp+1072)))[na] = i4
		na = na + b3TestEdgeCandidateSorted(tls, (*(*[132]float32)(unsafe.Add(nil, bp+2112)))[i11], (*(*[132]float32)(unsafe.Add(nil, bp+2112)))[i21], c, thresholdA)
		goto _201
	_201:
		;
		i4 = i4 + int32(2)
	}
	i5 = 0
	for {
		if !(i5 < halfEdgeCountB) {
			break
		}
		i12 = int32FromUint8((*(*HullHalfEdge)(unsafe.Add(nil, halfEdgesB+uintptr(i5)*4))).Face)
		i22 = int32FromUint8((*(*HullHalfEdge)(unsafe.Add(nil, halfEdgesB+uintptr(i5+int32(1))*4))).Face)
		c1 = *(*float32)(unsafe.Add(nil, cosinesB+uintptr(i5>>int32(1))*4))
		(*(*[128]int32)(unsafe.Add(nil, bp+1600)))[nb] = i5
		nb = nb + b3TestEdgeCandidateSorted(tls, (*(*[132]float32)(unsafe.Add(nil, bp+2640)))[i12], (*(*[132]float32)(unsafe.Add(nil, bp+2640)))[i22], c1, thresholdB)
		goto _202
	_202:
		;
		i5 = i5 + int32(2)
	}
	// Apply additional culling use the support points for the best face axes.
	// The support vertices of the best faces are points on the other hull, so an edge pair axis cannot
	// have a larger separation than the plane separation of these points over the arc of the edge.
	// This can slow down boxes so skip this if there are not enough edge candidates.
	if earlyReturn != 0 && nb*((na+int32(3))>>int32(2)) >= int32(_B3_EDGE_PROBE_MIN_TESTS) {
		v22 = centerDistance + radius
		if v22 < float32FromInt32(0) {
			v50 = -v22
		} else {
			v50 = v22
		}
		v33 = v50
		goto _205
	_205:
		probeBound = maxFaceSeparation - (float32(float32FromFloat32(0.1)*float32(float32FromFloat32(0.005)*b3GetLengthUnitsPerMeter(tls))) + float32(float32FromFloat32(0.001)*v33))
		if res.FaceA.IndexB != -int32(1) {
			v10 = res.FaceA.IndexB
		} else {
			v10 = *(*int32)(unsafe.Add(nil, bp+528))
		}
		vertexB = v10
		probeB = Vec3{
			X: *(*float32)(unsafe.Add(nil, vxB+uintptr(vertexB)*4)),
			Y: *(*float32)(unsafe.Add(nil, vyB+uintptr(vertexB)*4)),
			Z: *(*float32)(unsafe.Add(nil, vzB+uintptr(vertexB)*4)),
		}
		v2 = R
		v16 = probeB
		b6 = Vec3{
			X: float32(v2.Cx.X*v16.X) + float32(v2.Cy.X*v16.Y) + float32(v2.Cz.X*v16.Z),
			Y: float32(v2.Cx.Y*v16.X) + float32(v2.Cy.Y*v16.Y) + float32(v2.Cz.Y*v16.Z),
			Z: float32(v2.Cx.Z*v16.X) + float32(v2.Cy.Z*v16.Y) + float32(v2.Cz.Z*v16.Z),
		}
		v18 = b6
		goto _211
	_211:
		v19 = v18
		v20 = xfB.P
		v23 = Vec3{
			X: v19.X + v20.X,
			Y: v19.Y + v20.Y,
			Z: v19.Z + v20.Z,
		}
		goto _215
	_215:
		b3GetFacePlaneSeparations(tls, hullA, v23, bp+2112)
		na = b3FilterEdgeCandidates(tls, hullA, bp+2112, probeBound, bp+1072, na)
		if res.FaceB.IndexA != -int32(1) {
			v10 = res.FaceB.IndexA
		} else {
			v10 = *(*int32)(unsafe.Add(nil, bp+1064))
		}
		vertexA = v10
		probeA = Vec3{
			X: *(*float32)(unsafe.Add(nil, vxA+uintptr(vertexA)*4)),
			Y: *(*float32)(unsafe.Add(nil, vyA+uintptr(vertexA)*4)),
			Z: *(*float32)(unsafe.Add(nil, vzA+uintptr(vertexA)*4)),
		}
		v16 = probeA
		v18 = xfB.P
		v19 = Vec3{
			X: v16.X - v18.X,
			Y: v16.Y - v18.Y,
			Z: v16.Z - v18.Z,
		}
		goto _220
	_220:
		v2 = invR
		v20 = v19
		b6 = Vec3{
			X: float32(v2.Cx.X*v20.X) + float32(v2.Cy.X*v20.Y) + float32(v2.Cz.X*v20.Z),
			Y: float32(v2.Cx.Y*v20.X) + float32(v2.Cy.Y*v20.Y) + float32(v2.Cz.Y*v20.Z),
			Z: float32(v2.Cx.Z*v20.X) + float32(v2.Cy.Z*v20.Y) + float32(v2.Cz.Z*v20.Z),
		}
		v23 = b6
		goto _224
	_224:
		b3GetFacePlaneSeparations(tls, hullB, v23, bp+2640)
		nb = b3FilterEdgeCandidates(tls, hullB, bp+2640, probeBound, bp+1600, nb)
	}
	if na == 0 || nb == 0 {
		// No edge candidates found.
		return res
	}
	v10 = (faceCountB + int32(_B3_HULL_SOA_PADDING) - int32(1)) & ^(int32FromInt32(_B3_HULL_SOA_PADDING) - int32FromInt32(1))
	goto _226
_226:
	faceStrideB = v10
	v7 = hullB
	if (*HullData)(unsafe.Pointer(v7)).SoaNormalOffset == int32FromInt32(0) {
		v8 = uintptrFromInt32(0)
		goto _229
	}
	v8 = uintptr(int64(v7) + int64((*HullData)(unsafe.Pointer(v7)).SoaNormalOffset))
	goto _229
_229:
	nxB = v8
	nyB = nxB + uintptr(faceStrideB)*4
	nzB = nyB + uintptr(faceStrideB)*4
	wideFaceCountB = b3RoundUpToSIMDWidth(tls, faceCountB)
	b3NegativeTransformFromSoA(tls, R, xfB.P, nxB, nyB, nzB, wideFaceCountB, bp+3168, bp+3696, bp+4224, boolUint8(false1 != 0))
	b3NegativeTransformFromSoA(tls, R, xfB.P, vxB, vyB, vzB, wideVertexCountB, bp+4752, bp+5280, bp+5808, boolUint8(true1 != 0))
	squaredTol = float32(float32FromFloat32(0.005) * float32FromFloat32(0.005))
	k = 0
	for {
		if !(k < na) {
			break
		}
		edge = halfEdgesA + uintptr((*(*[132]int32)(unsafe.Add(nil, bp+1072)))[k])*4
		twin = edge + uintptr(1)*4
		A = (*(*Plane)(unsafe.Add(nil, planesA+uintptr((*HullHalfEdge)(unsafe.Pointer(edge)).Face)*16))).Normal
		B = (*(*Plane)(unsafe.Add(nil, planesA+uintptr((*HullHalfEdge)(unsafe.Pointer(twin)).Face)*16))).Normal
		(*(*[132]float32)(unsafe.Add(nil, bp+6336)))[k] = A.X
		(*(*[132]float32)(unsafe.Add(nil, bp+6864)))[k] = A.Y
		(*(*[132]float32)(unsafe.Add(nil, bp+7392)))[k] = A.Z
		(*(*[132]float32)(unsafe.Add(nil, bp+7920)))[k] = B.X
		(*(*[132]float32)(unsafe.Add(nil, bp+8448)))[k] = B.Y
		(*(*[132]float32)(unsafe.Add(nil, bp+8976)))[k] = B.Z
		v0 = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge)).Origin)
		v11 = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(twin)).Origin)
		(*(*[132]float32)(unsafe.Add(nil, bp+9504)))[k] = *(*float32)(unsafe.Add(nil, vxA+uintptr(v11)*4)) - *(*float32)(unsafe.Add(nil, vxA+uintptr(v0)*4))
		(*(*[132]float32)(unsafe.Add(nil, bp+10032)))[k] = *(*float32)(unsafe.Add(nil, vyA+uintptr(v11)*4)) - *(*float32)(unsafe.Add(nil, vyA+uintptr(v0)*4))
		(*(*[132]float32)(unsafe.Add(nil, bp+10560)))[k] = *(*float32)(unsafe.Add(nil, vzA+uintptr(v11)*4)) - *(*float32)(unsafe.Add(nil, vzA+uintptr(v0)*4))
		(*(*[132]float32)(unsafe.Add(nil, bp+11088)))[k] = *(*float32)(unsafe.Add(nil, vxA+uintptr(v0)*4))
		(*(*[132]float32)(unsafe.Add(nil, bp+11616)))[k] = *(*float32)(unsafe.Add(nil, vyA+uintptr(v0)*4))
		(*(*[132]float32)(unsafe.Add(nil, bp+12144)))[k] = *(*float32)(unsafe.Add(nil, vzA+uintptr(v0)*4))
		(*(*[132]float32)(unsafe.Add(nil, bp+12672)))[k] = float32(squaredTol * (float32((*(*[132]float32)(unsafe.Add(nil, bp+9504)))[k]*(*(*[132]float32)(unsafe.Add(nil, bp+9504)))[k]) + float32((*(*[132]float32)(unsafe.Add(nil, bp+10032)))[k]*(*(*[132]float32)(unsafe.Add(nil, bp+10032)))[k]) + float32((*(*[132]float32)(unsafe.Add(nil, bp+10560)))[k]*(*(*[132]float32)(unsafe.Add(nil, bp+10560)))[k])))
		goto _230
	_230:
		;
		k = k + 1
	}
	// Zero the tail lanes.
	zero = b3ZeroW4(tls)
	zeroV = b3Vec3W{
		X: zero,
		Y: zero,
		Z: zero,
	}
	v231 = zeroV
	b3StoreW4(tls, bp+6336+uintptr(na)*4, v231.X)
	b3StoreW4(tls, bp+6864+uintptr(na)*4, v231.Y)
	b3StoreW4(tls, bp+7392+uintptr(na)*4, v231.Z)
	v231 = zeroV
	b3StoreW4(tls, bp+7920+uintptr(na)*4, v231.X)
	b3StoreW4(tls, bp+8448+uintptr(na)*4, v231.Y)
	b3StoreW4(tls, bp+8976+uintptr(na)*4, v231.Z)
	v231 = zeroV
	b3StoreW4(tls, bp+9504+uintptr(na)*4, v231.X)
	b3StoreW4(tls, bp+10032+uintptr(na)*4, v231.Y)
	b3StoreW4(tls, bp+10560+uintptr(na)*4, v231.Z)
	v231 = zeroV
	b3StoreW4(tls, bp+11088+uintptr(na)*4, v231.X)
	b3StoreW4(tls, bp+11616+uintptr(na)*4, v231.Y)
	b3StoreW4(tls, bp+12144+uintptr(na)*4, v231.Z)
	b3StoreW4(tls, bp+12672+uintptr(na)*4, zero)
	linearSlop = float32(float32FromFloat32(0.005) * b3GetLengthUnitsPerMeter(tls))
	// The SIMD emulated version of this code is very slow. This is a purely scalar version
	// for platforms that don't have SIMD capability. It is much faster than SIMD emulation.
	// WARNING: this math needs to match the SIMD version for cross platform determinism.
	EPS = float32(-linearSlop * linearSlop)
	j = 0
	for {
		if !(j < nb) {
			break
		}
		edge1 = halfEdgesB + uintptr((*(*[128]int32)(unsafe.Add(nil, bp+1600)))[j])*4
		twin1 = edge1 + uintptr(1)*4
		f0 = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge1)).Face)
		f1 = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(twin1)).Face)
		v01 = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(edge1)).Origin)
		v12 = int32FromUint8((*HullHalfEdge)(unsafe.Pointer(twin1)).Origin)
		C = Vec3{
			X: (*(*[132]float32)(unsafe.Add(nil, bp+3168)))[f0],
			Y: (*(*[132]float32)(unsafe.Add(nil, bp+3696)))[f0],
			Z: (*(*[132]float32)(unsafe.Add(nil, bp+4224)))[f0],
		}
		D = Vec3{
			X: (*(*[132]float32)(unsafe.Add(nil, bp+3168)))[f1],
			Y: (*(*[132]float32)(unsafe.Add(nil, bp+3696)))[f1],
			Z: (*(*[132]float32)(unsafe.Add(nil, bp+4224)))[f1],
		}
		bv0 = Vec3{
			X: (*(*[132]float32)(unsafe.Add(nil, bp+4752)))[v01],
			Y: (*(*[132]float32)(unsafe.Add(nil, bp+5280)))[v01],
			Z: (*(*[132]float32)(unsafe.Add(nil, bp+5808)))[v01],
		}
		bv1 = Vec3{
			X: (*(*[132]float32)(unsafe.Add(nil, bp+4752)))[v12],
			Y: (*(*[132]float32)(unsafe.Add(nil, bp+5280)))[v12],
			Z: (*(*[132]float32)(unsafe.Add(nil, bp+5808)))[v12],
		}
		v16 = bv1
		v18 = bv0
		v19 = Vec3{
			X: v16.X - v18.X,
			Y: v16.Y - v18.Y,
			Z: v16.Z - v18.Z,
		}
		goto _239
	_239:
		DC = v19
		i6 = 0
		for {
			if !(i6 < na) {
				break
			}
			d = Vec3{
				X: (*(*[132]float32)(unsafe.Add(nil, bp+9504)))[i6],
				Y: (*(*[132]float32)(unsafe.Add(nil, bp+10032)))[i6],
				Z: (*(*[132]float32)(unsafe.Add(nil, bp+10560)))[i6],
			}
			v20 = C
			v23 = d
			v22 = float32(v20.X*v23.X) + float32(v20.Y*v23.Y) + float32(v20.Z*v23.Z)
			goto _244
		_244:
			// CBA = C.dir, DBA = D.dir, where dir = B_x_A
			CBA = v22
			v24 = D
			v27 = d
			v33 = float32(v24.X*v27.X) + float32(v24.Y*v27.Y) + float32(v24.Z*v27.Z)
			goto _248
		_248:
			DBA = v33
			if float32(CBA*DBA) >= EPS {
				goto _240
			}
			n0 = Vec3{
				X: (*(*[132]float32)(unsafe.Add(nil, bp+6336)))[i6],
				Y: (*(*[132]float32)(unsafe.Add(nil, bp+6864)))[i6],
				Z: (*(*[132]float32)(unsafe.Add(nil, bp+7392)))[i6],
			}
			n1 = Vec3{
				X: (*(*[132]float32)(unsafe.Add(nil, bp+7920)))[i6],
				Y: (*(*[132]float32)(unsafe.Add(nil, bp+8448)))[i6],
				Z: (*(*[132]float32)(unsafe.Add(nil, bp+8976)))[i6],
			}
			v16 = n0
			v18 = DC
			v22 = float32(v16.X*v18.X) + float32(v16.Y*v18.Y) + float32(v16.Z*v18.Z)
			goto _252
		_252:
			// ADC = n0.DC, BDC = n1.DC, where DC = D_x_C
			ADC = v22
			v19 = n1
			v20 = DC
			v33 = float32(v19.X*v20.X) + float32(v19.Y*v20.Y) + float32(v19.Z*v20.Z)
			goto _256
		_256:
			BDC = v33
			if float32(ADC*BDC) >= EPS || float32(CBA*BDC) >= EPS {
				goto _240
			}
			v22 = float32(CBA * CBA)
			v33 = float32(DBA * DBA)
			if v22 > v33 {
				v54 = v22
			} else {
				v54 = v33
			}
			v50 = v54
			goto _260
		_260:
			// Reject near parallel edges
			maxCD = v50
			if maxCD <= (*(*[132]float32)(unsafe.Add(nil, bp+12672)))[i6] {
				goto _240
			}
			// t = -CBA / (DBA - CBA)
			t = -CBA / (DBA - CBA)
			v16 = D
			v18 = C
			v19 = Vec3{
				X: v16.X - v18.X,
				Y: v16.Y - v18.Y,
				Z: v16.Z - v18.Z,
			}
			goto _265
		_265:
			v20 = C
			v22 = t
			v23 = v19
			v24 = Vec3{
				X: v20.X + float32(v22*v23.X),
				Y: v20.Y + float32(v22*v23.Y),
				Z: v20.Z + float32(v22*v23.Z),
			}
			goto _270
		_270:
			// normal = lerp(t, C, D) = C + (D-C)*t
			n = v24
			v27 = n
			v29 = n
			v33 = float32(v27.X*v29.X) + float32(v27.Y*v29.Y) + float32(v27.Z*v29.Z)
			goto _274
		_274:
			len2 = v33
			inv = float32FromFloat32(1) / sqrtf(tls, len2)
			v50 = inv
			v30 = n
			v31 = Vec3{
				X: float32(v50 * v30.X),
				Y: float32(v50 * v30.Y),
				Z: float32(v50 * v30.Z),
			}
			goto _278
		_278:
			n = v31
			// separation = -dot(normal, av0 + bv0)
			av0 = Vec3{
				X: (*(*[132]float32)(unsafe.Add(nil, bp+11088)))[i6],
				Y: (*(*[132]float32)(unsafe.Add(nil, bp+11616)))[i6],
				Z: (*(*[132]float32)(unsafe.Add(nil, bp+12144)))[i6],
			}
			v16 = av0
			v18 = bv0
			v19 = Vec3{
				X: v16.X + v18.X,
				Y: v16.Y + v18.Y,
				Z: v16.Z + v18.Z,
			}
			goto _282
		_282:
			v20 = v19
			v23 = n
			v22 = float32(v20.X*v23.X) + float32(v20.Y*v23.Y) + float32(v20.Z*v23.Z)
			goto _286
		_286:
			separation2 = -v22
			if separation2 > res.Edge.Separation {
				res.Edge.Normal = n
				res.Edge.Separation = separation2
				// Half edge index
				res.Edge.IndexA = (*(*[132]int32)(unsafe.Add(nil, bp+1072)))[i6]
				res.Edge.IndexB = (*(*[128]int32)(unsafe.Add(nil, bp+1600)))[j]
				if separation2 > speculativeDistance && earlyReturn != 0 {
					res.SeparatedFeature = int32(b3_edgePairAxis)
					return res
				}
			}
			goto _240
		_240:
			;
			i6 = i6 + 1
		}
		goto _235
	_235:
		;
		j = j + 1
	}
	return res
}

/**@}*/ // debug_draw

// Snapshot image magic 'BNS3' and version

// C documentation
//
//	// Layout hash over all POD-copied structs + key constants.
//	// Changing a struct size updates this, catching ABI drift early.
func b3ComputeLayoutHash(tls *_Stack) (r uint32) {
	var h uint32_t
	_ = h
	h = uint32(2166136261)
	h = h ^ uint32(uint64FromInt64(96))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(252))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(64))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(208))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(216))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(268))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(72))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(444))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(64))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(4))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(12))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(12))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(56))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(8))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(88))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(120))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(112))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(32))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(24))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(152))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(1))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(4))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(1))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(16))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(96))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(32))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(12))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(96))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(96))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(32))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(20))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromUint64(0x6B1E39D4A87C25F3)&uint64FromUint32(0xFFFFFFFF))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromUint64(0x6B1E39D4A87C25F3)>>int32FromInt32(32))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromUint64(0xAAAB9A00F1A8AAF7)&uint64FromUint32(0xFFFFFFFF))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromUint64(0xAAAB9A00F1A8AAF7)>>int32FromInt32(32))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromUint64(0x8E41E5FB084848F8)&uint64FromUint32(0xFFFFFFFF))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromUint64(0x8E41E5FB084848F8)>>int32FromInt32(32))
	h = h * uint32(16777619)
	h = h ^ uint32((uint64FromUint64(0x9E4B17D3A25C68F1)^uint64FromUint64(0xAAAB9A00F1A8AAF7)^uint64FromUint64(0x6B1E39D4A87C25F3))&uint64FromUint32(0xFFFFFFFF))
	h = h * uint32(16777619)
	h = h ^ uint32((uint64FromUint64(0x9E4B17D3A25C68F1)^uint64FromUint64(0xAAAB9A00F1A8AAF7)^uint64FromUint64(0x6B1E39D4A87C25F3))>>int32FromInt32(32))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(16))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(24))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(40))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(12))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(20))
	h = h * uint32(16777619)
	h = h ^ uint32FromInt32(int32FromInt32(_B3_GRAPH_COLOR_COUNT))
	h = h * uint32(16777619)
	h = h ^ uint32FromInt32(int32(b3_bodyTypeCount))
	h = h * uint32(16777619)
	h = h ^ uint32(uint64FromInt64(8))
	h = h * uint32(16777619)
	return h
}
