function jumpingOnClouds($c) {
    // Write your code here
    $jump = 0;
    for ($i = 0; $i < count($c); $i++){
        if (isset($c[$i + 2]) && $c[$i + 2] === 0){
            $jump++;
            $i++;
            continue;
        }
        
        if (isset($c[$i + 2]) && $c[$i + 2] === 1){
            $jump++;
            continue;
        }
        
        if(isset($c[$i + 1]) && $c[$i + 1] === 0){
            $jump++;
        }
    }
    
    return $jump;
}