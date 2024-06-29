/**
 * @param {string} s
 * @return {boolean}
 */
var isPalindrome = function(s) {
    s = s.toLowerCase()
    s = s.replace(/\s+/g, "")
    s = s.replace(/[^a-zA-Z0-9]/g, "")

    for (let i = 0; i < s.length; i++) {
       if (s[i] !== s[(s.length)-i-1]) {
           return false
       }
    }
    return true
};

isPalindrome("A man, a plan, a canal: Panama")