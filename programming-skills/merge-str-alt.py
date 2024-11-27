class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        shortlength = 0
        extendword = word1
        if len(word1) >= len(word2):
            shortlength = len(word2)
            if len(word1) == len(word2):
                extendword = None
        else:
            shortlength = len(word1)
            extendword = word2

        print(range(shortlength))
        concatword = ''
        for i in range(shortlength):
            concatword += word1[i] + word2[i]
        print(concatword)
        if extendword is not None:
            concatword += extendword[shortlength:len(extendword)]
        print(concatword)
        return concatword
