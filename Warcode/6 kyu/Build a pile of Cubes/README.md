Your task is to construct a building which will be a pile of n cubes. The cube at the bottom will have a volume of n^3 , the cube above will have volume of (n-1)^3 and so on until the top which will have a volume of 1^3.

You are given the total volume m of the building. Being given m can you find the number of cubes n required for the building?

The parameter of the function will be an integer m and you have to return an integer n such that:

n^3 + (n-1)^3 + ... + 1^3 = m if such a n exists or -1 if there is no such n.

Examples:
find_nb(1071225) --> 45
find_nb(91716553919377) --> -1
