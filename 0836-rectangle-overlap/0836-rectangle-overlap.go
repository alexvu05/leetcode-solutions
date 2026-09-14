func isRectangleOverlap(rec1 []int, rec2 []int) bool {
    noOverlap := rec1[2] <= rec2[0] || // rec1 left of rec2
                 rec1[0] >= rec2[2] || // rec1 right of rec2
                 rec1[3] <= rec2[1] || // rec1 below rec2
                 rec1[1] >= rec2[3]    // rec1 above rec2
    return !noOverlap
}