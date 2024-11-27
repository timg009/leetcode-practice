class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        n1 = len(word1)
        n2 = len(word2)
        p1 = 0
        p2 = 0
        mword = ''

        while (p1 < n1 or p2 < n2):
            if p1 < n1:
                mword += word1[p1]
                p1 = p1 + 1
            if p2 < n2:
                mword += word2[p2]
                p2 = p2 + 1
            # print(mword)

        return mword