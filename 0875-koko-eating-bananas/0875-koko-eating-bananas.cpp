class Solution {
public:
    int minEatingSpeed(vector<int>& piles, int h) {
        int left  = 1;
        int right = *max_element(piles.begin(), piles.end());
        // Binary search for the smallest k satisfying hoursNeeded(k) <= h
        while (left < right) {
            int mid = left + (right - left) / 2;
            if (hoursNeeded(piles, mid) <= h) {
                right = mid;       // k = mid works; try smaller k
            } else {
                left = mid + 1;    // k = mid too slow; need larger k
            }
        }
        return left;
    }
private:
    // Total hours needed to eat all piles at speed k
    long long hoursNeeded(vector<int>& piles, int k) {
        long long hours = 0;
        for (int p : piles) {
            // Ceiling division: hours to finish pile p at speed k
            hours += (p + k - 1) / k;
        }
        return hours;
    }
};