func checkOverlap(radius int, xCenter int, yCenter int, x1 int, y1 int, x2 int, y2 int) bool {
    // Find closest point in rectangle to circle center
    closestX := clamp(xCenter, x1, x2)
    closestY := clamp(yCenter, y1, y2)

    // Check if distance from center to closest point <= radius
    dx := closestX - xCenter
    dy := closestY - yCenter
    return dx*dx + dy*dy <= radius*radius
}

func clamp(v, lo, hi int) int {
    if v < lo { return lo }
    if v > hi { return hi }
    return v
}