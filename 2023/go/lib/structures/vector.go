package structures

type Number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Vector[T Number] []T

func (vec Vector[T]) Times(multiplier T) Vector[T] {
	res := make([]T, len(vec))
	for i, val := range vec{
		res[i] = val * multiplier
	}
	return res
}
func (vec Vector[T]) Divide(divisor T) Vector[T] {
	res := make([]T, len(vec))
	for i, val := range vec{
		res[i] = val / divisor
	}
	return res
}
func (vec Vector[T]) Plus(other Vector[T]) Vector[T] {
	res := make([]T, len(vec))
	for i, val := range vec{
		res[i] = val + other[i]
	}
	return res
}
func (vec Vector[T]) Minus(other Vector[T]) Vector[T] {
	res := make([]T, len(vec))
	for i, val := range vec{
		res[i] = val - other[i]
	}
	return res
}

func (vec Vector[T]) SimpleCross(other Vector[T]) T {
	if len(vec) != 2 || len(other) != 2 {
		panic("Can't Simple Cross These")
	}
	return (vec[0]*other[1]) - (vec[1]*other[0])
}

/**
Part 2. Ah, part 2. What a beautiful, largely redundant, system of nonlinear equations. The obvious dirty solution was to feed all this to a constraint solver and come back to the result. But that wouldn't be satisfying, nor very much in the spirit of the advent, at least not as I approach it. For two days, I toyed with rings, looking for a way to bring in the chinese remainder theorem which has been strangely absent this year. After all, for each hailstone with coordinates (a, b, c), if our initial position is (x, y, z) and the paths cross at time t, we have x = a mod t, y = b mod t and z = c mod t. But as we have no idea what t is, what can we do with that? Probably loads, but I didn't find any.

I then tried to focus on pairs of paths which had one velocity in common, trying to see if I could use that fact to remove some of the unknowns from the system. Couldn't find a way either.

Finally, I realised that, if I concentrated on the plane, I could completely write out the time of the system of equations. If (dx, dy) represents our velocity and (da, db) that of the path we consider in the plane (ignoring the 3rd dimension), we have

x + t*dx = a + t*da

y + t* dy = b + t* db

Which we can rewrite as

t = (a - x) / (dx - da)

t = (a -y) / (dy - da)

By combining both equations, we get rid of the t and get

(a - x) / (dx - da) = (b - y) / (dy - db)

a*dxy- a*db - x*dy + x * db = b * dx - b * da - y * dx - y * da

That's still a non-linear equation, but we can rearrange it as

y*dx - x * dy = b*dx - b * da + y * da - a * dy + a * db - x * db

Let's then bring in a second path, (d, e, whatever) (dd, de, dwhatever). We also have

y * dx - x * dy = e * dx - e * dd + y * dd - d * dy + d * de - x * de

Bringing those two togethers, we get the beautiful equation

b * dx - b * da +y * da -a * dy + a * db - x * db = e * dx - e * dd + y * dd - d * dy + d * de - x * de

Which we can rearrange as

(db - de) * x + ( e - b ) * dx + (dd - da) * y + (a - d) * dy = db * a - da * b + dd * e - de * d

Finally, a linear equation, with four unknowns.

By using five paths total, we can generate four such equations, and get a nice linear system, which we solve by using Gaussian elimination (my current write out of the solving part is not the most elegant, tbh, I'm sure it could be made more general).

We now have values for x, dx, y and dy. Using the first two, we can get the times t0 and t1 at which our rock crosses the first and second paths. From there, we have to linear equations for z and dz, and we can get them (although we don't need the latter, nor do we need dy). Just add x y and z, and you have your solution. If you have remembered that the order you were getting your solutions was (x dx y dy) and not (x y dx dy), the first bug I encountered. And if you have remembered to check whether the numbers you were toying with were within the Haskell bounds for Ints, which they are not, the second bug I encountered.

Anyhow, we consume a whopping 15 equations when we have only eleven unknowns (the three coordinates, the three velocity values and one time per path). Solving the nonlinear system would probably require only nine equations (three paths), but that's beyond my scope for now.

*/