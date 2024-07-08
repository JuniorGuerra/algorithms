/**
 * @param {string} s
 * @return {boolean}
 */
const isPalindrome = function(s) {
    s = s.toLowerCase()
    s = s.replace(/\s+/g, "")
    s = s.replace(/[^a-zA-Z0-9]/g, "")

    for (let i = 0; i < s.length/2; i++) {
       if (s[i] !== s[(s.length)-i-1]) {
           return false
       }
    }
    return true
};

console.log(isPalindrome("A man, a plan, a canal: Panama") === true)
console.log(isPalindrome("race a car") === false)
console.log(isPalindrome(" ") === true)