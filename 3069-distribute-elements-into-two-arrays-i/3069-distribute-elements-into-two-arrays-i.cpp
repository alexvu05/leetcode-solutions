class Solution {
public:
    vector<int> resultArray(vector<int>& nums) {
        vector<int> arr1 = {nums[0]};
        vector<int> arr2 = {nums[1]};
        for (int i = 2; i < (int)nums.size(); i++) {
            // Append to arr1 if its last element is greater, else append to arr2
            if (arr1.back() > arr2.back()) {
                arr1.push_back(nums[i]);
            } else {
                arr2.push_back(nums[i]);
            }
        }
        // Concatenate arr1 and arr2
        for (int x : arr2) arr1.push_back(x);
        return arr1;
    }
};