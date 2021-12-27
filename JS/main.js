
var courses = [
    {
        id: 1, 
        name: 'JS', 
        coin: 0
    },
    {
        id: 2, 
        name: 'A', 
        coin: 0
    },
    {
        id: 3, 
        name: 'B', 
        coin: 10
    },
    {
        id: 4, 
        name: 'C', 
        coin: 20
    }
]

courses.forEach(function(course, index) {
    console.log(course, index);
})