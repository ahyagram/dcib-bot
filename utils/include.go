package utils


func Include(v string, a []string) bool {
    for i := 0; i < len(a); i++ {
        if a[i] == v {
            return true
        }
    }
    
    return false
}