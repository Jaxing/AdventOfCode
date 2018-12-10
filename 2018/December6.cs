using System;
using System.Collections.Generic;
using System.Text.RegularExpressions;

public class December6
{

    private static int distance(int aX, int aY, int bX, int bY) {
        return Math.Abs(aY - bY) + Math.Abs(aX - bX);
    }

    private static (int, int) parsePos(string pos) {
        string[] parts = Regex.Split(pos, ", ");
        return (Int32.Parse(parts[0]), Int32.Parse(parts[1]));
    }

    private static Dictionary<int,int> computeAreas(int rowStart, int rowEnd, int colStart, int colEnd, (int,int)[] intInput) {
        Dictionary<int, int> areaOfIndex = new Dictionary<int, int>(); 

        for (int row=rowStart; row <= rowEnd; row++) {
            for (int col=colStart; col <=colEnd; col++) {
                int closestIndex = -1;
                int closestDistance = int.MaxValue;
                bool closestSame = false;

                for (int i=0; i < intInput.Length; i++) {
                    int dist = distance(col, row, intInput[i].Item1, intInput[i].Item2);
                    if (dist < closestDistance) {
                        closestSame = false;
                        closestIndex = i;
                        closestDistance = dist;
                    }else if (dist == closestDistance) {
                        closestSame = true;
                    }
                }
                if (!closestSame && closestIndex >= 0) {
                    areaOfIndex[closestIndex] = areaOfIndex.GetValueOrDefault(closestIndex) + 1;
                }
            }
        }

        return areaOfIndex;
    }

    private static Dictionary<int,int> removeInfiniteAreas(int rowStart, int rowEnd, int colStart, int colEnd, Dictionary<int,int> areaOfIndex, (int,int)[] positions) {
        for (int i=0; i < positions.Length; i++) {
            int leftBlock = 0;
            int rightBlock = 0;
            int topBlock = 0;
            int bottomBlock = 0;

            foreach (var pos in positions) {
                if (pos.Item1 < positions[i].Item1) {
                    leftBlock++;
                }
                if (pos.Item1 > positions[i].Item1) {
                    rightBlock++;
                }
                if (pos.Item2 < positions[i].Item2) {
                    bottomBlock++;
                }
                if (pos.Item2 > positions[i].Item2) {
                    topBlock++;
                }
            }

            if (leftBlock < 2 || rightBlock < 2 || topBlock < 2 || bottomBlock < 2 ) {
                areaOfIndex[i] = -1;
            }
        }

        return areaOfIndex;
    }

    public static int largestArea(string[] input) {
        (int, int)[] intInput = new (int, int)[input.Length];
        Dictionary<int, int> areaOfIndex = new Dictionary<int, int>(); 

        int rowStart = int.MaxValue;
        int rowEnd = 0;
        int colStart = int.MaxValue;
        int colEnd = 0;

        for (int i=0; i < input.Length; i++) {
            intInput[i] = parsePos(input[i]);
            rowStart = Math.Min(rowStart, intInput[i].Item2);
            rowEnd = Math.Max(rowEnd, intInput[i].Item2);
            colStart = Math.Min(colStart, intInput[i].Item1);
            colEnd = Math.Max(colEnd, intInput[i].Item1);
        }

        areaOfIndex = computeAreas(rowStart,rowEnd,colStart,colEnd, intInput);
        areaOfIndex = removeInfiniteAreas(rowStart,rowEnd,colStart,colEnd, areaOfIndex, intInput);
        int largestArea = 0;

        foreach (var key in areaOfIndex.Keys) {
            largestArea = Math.Max(areaOfIndex[key], largestArea);
        }

        return largestArea;
    }
}