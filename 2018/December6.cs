using System;
using System.Collections.Generic;
using System.Text.RegularExpressions;

public class December6
{
    private static (int,int) topLeftbound((int,int)[] positions) {
        int mostLeft = int.MaxValue;
        int top = 0;

        foreach (var pos in positions) {
            if (pos.Item1 < mostLeft) {
                mostLeft = pos.Item1;
            }
            if (pos.Item2 > top) {
                top = pos.Item2;
            }
        }
        return (mostLeft, top);
    }

    private static (int,int) topRightbound((int,int)[] positions) {
        int mostRight = 0;
        int top = 0;

        foreach (var pos in positions) {
            if (pos.Item1 > mostRight) {
                mostRight = pos.Item1;
            }
            if (pos.Item2 > top) {
                top = pos.Item2;
            }
        }
        return (mostRight, top);
    }

    private static (int,int) bottomLeftbound((int,int)[] positions) {
        int mostLeft = 0;
        int bottom = 0;

        foreach (var pos in positions) {
            if (pos.Item1 < mostLeft) {
                mostLeft = pos.Item1;
            }
            if (pos.Item2 < bottom) {
                bottom = pos.Item2;
            }
        }
        return (mostLeft, bottom);
    }

    private static (int,int) bottomRightbound((int,int)[] positions) {
        int mostRight = 0;
        int bottom = 0;

        foreach (var pos in positions) {
            if (pos.Item1 > mostRight) {
                mostRight = pos.Item1;
            }
            if (pos.Item2 < bottom) {
                bottom = pos.Item2;
            }
        }
        return (mostRight, bottom);
    }
    private static int findingBottomRight((int, int)[] positions){
        (int,int) corner = bottomRightbound(positions);
        return indexOfClosest(positions, corner);
    }

    private static int findingBottomLeft((int, int)[] positions){
        (int,int) corner = bottomLeftbound(positions);
        return indexOfClosest(positions, corner);
    }

    private static int findingTopRight((int, int)[] positions){
        (int,int) corner = topRightbound(positions);
        return indexOfClosest(positions, corner);
    }

    private static int findingTopLeft((int, int)[] positions){
        (int,int) corner = topLeftbound(positions);
        return indexOfClosest(positions, corner);
    }

    private static int indexOfClosest((int,int)[] positions, (int,int) target) {
        int index = 0;
        int closestDistToCorner = int.MaxValue;

        for (int i=0; i < positions.Length; i++) {
            (int,int) pos = positions[i];
            int distToCorner = distance(pos.Item1, pos.Item2, target.Item1, target.Item2);
            if (distToCorner < closestDistToCorner) {
                closestDistToCorner = distToCorner;
                index = i;
            }
        }
        return index;
    }

    private static int distance(int aX, int aY, int bX, int bY) {
        return Math.Abs(aY - bY) + Math.Abs(aX - bX);
    }

    private static (int, int) parsePos(string pos) {
        string[] parts = Regex.Split(pos, ", ");
        return (Int32.Parse(parts[1]), Int32.Parse(parts[0]));
    }



    public static int largestArea(string[] input) {
        (int, int)[] intInput = new (int, int)[input.Length];
        Dictionary<int, int> areaOfIndex = new Dictionary<int, int>(); 

        for (int i=0; i < input.Length; i++) {
            intInput[i] = parsePos(input[i]);
            Console.WriteLine(intInput[i]);
        }

        //unbounded
        var iBottomRight = findingBottomRight(intInput);
        var iBottomLeft = findingBottomLeft(intInput);
        var iTopRight = findingTopRight(intInput);
        var iTopLeft = findingTopLeft(intInput);

        int rowStart = Math.Min(intInput[iBottomLeft].Item2, intInput[iBottomRight].Item2);
        int rowEnd = Math.Max(intInput[iTopLeft].Item2, intInput[iTopRight].Item2);
        int colStart = Math.Min(intInput[iBottomLeft].Item1, intInput[iTopLeft].Item1);
        int colEnd = Math.Max(intInput[iBottomRight].Item1, intInput[iTopRight].Item1);
        Console.WriteLine("{0},{1},{2},{3}", intInput[iBottomLeft], intInput[iBottomRight], intInput[iTopLeft], intInput[iTopRight]);
        Console.WriteLine("{0},{1},{2},{3}",rowStart, rowEnd, colStart, colEnd);
        Console.WriteLine("{0},{1},{2},{3}",iBottomLeft, iBottomRight, iTopLeft, iTopRight);
        for (int row=rowStart; row <= rowEnd; row++) {
            for (int col=colStart; col <=colEnd; col++) {
                int closestIndex = -1;
                int closestDistance = int.MaxValue;
                bool closestSame = false;

                for (int i=0; i < intInput.Length; i++) {
                    int dist = distance(col, row, intInput[i].Item1, intInput[i].Item2);
                    if (dist < closestDistance) {
                        //Console.WriteLine(i);
                        closestSame = false;
                        closestIndex = i;
                        closestDistance = dist;
                    }else if (dist == closestDistance) {
                        //Console.WriteLine("{0}, {1}", i, closestIndex);
                        closestSame = true;
                    }
                }
                if (!closestSame && closestIndex >= 0) {
                    areaOfIndex[closestIndex] = areaOfIndex.GetValueOrDefault(closestIndex) + 1;
                }
            }
        }

        int largestArea = 0;

        foreach (var key in areaOfIndex.Keys) {
            Console.WriteLine("Key: {0} Value: {1}", key, areaOfIndex[key]);
            if (key == iBottomLeft || key == iBottomRight || key == iTopLeft || key == iTopRight) {
                continue;
            }
            if (areaOfIndex[key] > largestArea) {
                largestArea = areaOfIndex[key];
            }
        }

        return largestArea;
    }
}