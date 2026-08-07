static const int CONTRIB[10][4] = {
    {0,0,0,0}, // 0 (không dùng, số không được chứa 0)
    {0,0,0,0}, // 1
    {1,0,0,0}, // 2
    {0,1,0,0}, // 3
    {2,0,0,0}, // 4
    {0,0,1,0}, // 5
    {1,1,0,0}, // 6
    {0,0,0,1}, // 7
    {3,0,0,0}, // 8
    {0,2,0,0}, // 9
};
class Solution {
public:    
    unordered_map<long long,int> memo[10]; // memo[v]: state -> số chữ số tối thiểu dùng digit >= v
 
    static long long key(int a,int b,int c,int d){
        return (((long long)a*100 + b)*100 + c)*100 + d;
    }

    int getMin(int v,int a,int b,int c,int d){
        if(a<=0 && b<=0 && c<=0 && d<=0) return 0;
        long long k = key(a,b,c,d);
        auto it = memo[v].find(k);
        if(it != memo[v].end()) return it->second;
        int best = INT_MAX/2;
        for(int w=v; w<=9; w++){
            int na = max(0, a - CONTRIB[w][0]);
            int nb = max(0, b - CONTRIB[w][1]);
            int nc = max(0, c - CONTRIB[w][2]);
            int nd = max(0, d - CONTRIB[w][3]);
            if(na==a && nb==b && nc==c && nd==d) continue; // chữ số này không giúp ích gì, bỏ qua
            int r = getMin(v, na, nb, nc, nd);
            if(r < INT_MAX/2) best = min(best, 1+r);
        }
        return memo[v][k] = best;
    }
 
    // Dựng chuỗi S ký tự nhỏ nhất (không số 0) phủ đủ (a,b,c,d). Giả sử S >= getMin(2,a,b,c,d).
    string fillMinimal(int a,int b,int c,int d,int S){
        int M = getMin(2,a,b,c,d);
        string ones(max(0,S-M), '1');
        string suffix; suffix.reserve(M);
        int ca=a, cb=b, cc=c, cd=d, remaining=M, floorV=2;
        for(int k=0;k<M;k++){
            for(int v=floorV; v<=9; v++){
                int na = max(0, ca - CONTRIB[v][0]);
                int nb = max(0, cb - CONTRIB[v][1]);
                int nc = max(0, cc - CONTRIB[v][2]);
                int nd = max(0, cd - CONTRIB[v][3]);
                if(getMin(v, na, nb, nc, nd) <= remaining-1){
                    suffix.push_back(char('0'+v));
                    ca=na; cb=nb; cc=nc; cd=nd;
                    remaining--; floorV=v;
                    break;
                }
            }
        }
        return ones + suffix;
    }
 
    string smallestNumber(string num, long long t) {
        long long tt = t;
        int A=0,B=0,C=0,D=0;
        while(tt%2==0){tt/=2;A++;}
        while(tt%3==0){tt/=3;B++;}
        while(tt%5==0){tt/=5;C++;}
        while(tt%7==0){tt/=7;D++;}
        if(tt != 1) return "-1";
 
        int n = (int)num.size();
        int M = getMin(2,A,B,C,D);
        if(M > n) return fillMinimal(A,B,C,D,M);
 
        size_t zpos = num.find('0');
        bool hasZero = (zpos != string::npos);
        if(!hasZero){
            long long a0=0,b0=0,c0=0,d0=0;
            for(char ch: num){
                int dg = ch-'0';
                a0+=CONTRIB[dg][0]; b0+=CONTRIB[dg][1]; c0+=CONTRIB[dg][2]; d0+=CONTRIB[dg][3];
            }
            if(a0>=A && b0>=B && c0>=C && d0>=D) return num;
        }
 
        vector<int> prefA(n+1,0), prefB(n+1,0), prefC(n+1,0), prefD(n+1,0);
        for(int i=0;i<n;i++){
            int dg = num[i]-'0';
            prefA[i+1]=prefA[i]+CONTRIB[dg][0];
            prefB[i+1]=prefB[i]+CONTRIB[dg][1];
            prefC[i+1]=prefC[i]+CONTRIB[dg][2];
            prefD[i+1]=prefD[i]+CONTRIB[dg][3];
        }
 
        int firstZero = hasZero ? (int)zpos : n;
        int startI = hasZero ? firstZero : n-1;
 
        for(int i=startI; i>=0; i--){
            int dorig = num[i]-'0';
            int L = n-1-i;
            for(int dprime = dorig+1; dprime<=9; dprime++){
                int needA = max(0, A - prefA[i] - CONTRIB[dprime][0]);
                int needB = max(0, B - prefB[i] - CONTRIB[dprime][1]);
                int needC = max(0, C - prefC[i] - CONTRIB[dprime][2]);
                int needD = max(0, D - prefD[i] - CONTRIB[dprime][3]);
                if(getMin(2, needA, needB, needC, needD) <= L){
                    string ans = num.substr(0,i);
                    ans.push_back(char('0'+dprime));
                    ans += fillMinimal(needA,needB,needC,needD,L);
                    return ans;
                }
            }
        }
        return fillMinimal(A,B,C,D,n+1); // không tìm được cùng độ dài -> tăng thêm 1 chữ số
    }
};