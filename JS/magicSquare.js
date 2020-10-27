/* We define a magic square to be an  matrix of distinct positive integers from 
 to  where the sum of any row, column, or diagonal of length  is always equal to the 
 same number: the magic constant.

You will be given a  matrix  of integers in the inclusive range . We can convert any digit
  to any other digit  in the range  at cost of . Given , convert it into a magic square at
   minimal cost. Print this cost on a new line.

Note: The resulting magic square must contain distinct integers in the inclusive range .

For example, we start with the following matrix :

5 3 4
1 5 8
6 4 2
We can convert it to the following magic square:

8 3 4
1 5 9
6 7 2
This took three replacements at a cost of .

Function Description

Complete the formingMagicSquare function in the editor below. It should return an integer that
 represents the minimal total cost of converting the input square to a magic square.

formingMagicSquare has the following parameter(s):

s: a  array of integers
Input Format

Each of the lines contains three space-separated integers of row .

Constraints

Output Format

Print an integer denoting the minimum cost of turning matrix  into a magic square.

Sample Input 0

4 9 2
3 5 7
8 1 5
Sample Output 0

1
Explanation 0

If we change the bottom right value, , from  to  at a cost of ,  becomes a magic square at the minimum possible cost.

Sample Input 1

4 8 2
4 5 7
6 1 6
Sample Output 1

4

COST = Sum of differences to arrive magic square

*/

function formMagicSquare(matrix){
    const map = new Map();
    const rows = new Map();
    const columns = new Map();
    const diagonals = new Map();
    const repeated = new Map();
    const range = matrix.length;
    if(range != 3){
        throw new Error('Invalid matrix range: ', range);
    }
    const magicConstant = (((range**2)+1)/2)*range
    const center = 5//getCenter(range, magicConstant)
    const minNumber = 1;
    const maxNumber = range**2;
    const interval = [minNumber, maxNumber]
    //diagonals.set('diagonal',0)
    //diagonals.set('antidiagonal',0)
    //for(let rangeIndex = 0;rangeIndex < range; rangeIndex++){
    //    rows.set(rangeIndex, 0)
    //    columns.set(rangeIndex, 0)
    //}
    //matrix.map( (row, rowIndex) => {
    //    row.map( (value, columnIndex) => {
    //        if(value < minNumber || value > maxNumber){
    //            throw new Error('Invalid matrix value: ', value);
    //        }
    //        if(map.get(value) === undefined){
    //            map.set(value, [[rowIndex, columnIndex]])
    //        }else{
    //            map.set(value, map.get(value).push([rowIndex, columnIndex]))
    //            repeated.set([rowIndex, columnIndex], value)
    //        }
    //        map.set([rowIndex, columnIndex], value)
    //        rows.set(rowIndex, rows.get(rowIndex) + value)
    //        columns.set(columnIndex, columns.get(columnIndex) + value)
    //        if(rowIndex == columnIndex){
    //            diagonals.set('diagonal',
    //            ((diagonals.get('diagonal') === undefined)? 0 : diagonals.get('diagonal')) + value)
    //        }
    //        if(rowIndex + 1 + columnIndex == range){
    //            diagonals.set('antidiagonal',
    //            ((diagonals.get('antidiagonal') === undefined)? 0 : diagonals.get('antidiagonal')) + value)
    //        }
    //    })
    //})
    let cost = 0;
    const usedNumbers = []
    const oddNumbers = [1,3,7,9]
    const evenNumbers = [2,4,6,8]

    let dif = getDif(matrix[0,1], oddNumbers, []);
    matrix[0,1] -= dif
    oddNumbers.splice(oddNumbers.indexOf(matrix[0,1]), 1)
    usedNumbers.push(matrix[0,1])
    cost += Math.abs(dif)

    let difOpposite01 = matrix[3,1] - (magicConstant - matrix[0,1])
    matrix[3,1] -= difOpposite01
    oddNumbers.splice(oddNumbers.indexOf(matrix[3,1]), 1)
    cost += Math.abs(difOpposite01);

    let difOpposite10 = getDif(matrix[1,0], oddNumbers, usedNumbers);
    matrix[1,0] += difOpposite10
    oddNumbers.splice(oddNumbers.indexOf(matrix[1,0]), 1)
    cost += Math.abs(difOpposite10);
    
    let difOpposite12 = matrix[1,2] - (magicConstant - matrix[1,0])
    matrix[1,2] -= difOpposite12
    oddNumbers.splice(oddNumbers.indexOf(matrix[1,2]), 1)
    cost += Math.abs(difOpposite12);


    let difOpposite00 = getDif(matrix[0,0], evenNumbers, usedNumbers);
    matrix[0,0] += difOpposite00
    evenNumbers.splice(evenNumbers.indexOf(matrix[0,0]), 1)
    cost += Math.abs(difOpposite00);
    
    let value02 = magicConstant - matrix[0,0] - matrix[0,1]
    let value20 = magicConstant - matrix[0,0] - matrix[1,0]
    let value22 = matrix[2,2] - evenNumbers[0]

    let dif02 = matrix[0,2] - value02;
    let dif20 = matrix[2,0] - value20;
    let dif22 = matrix[2,2] - value22;
}
