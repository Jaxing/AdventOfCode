using System;
using static System.Math;
using System.Collections.Generic;
using System.Text.RegularExpressions;

public class OverlapPosition {
    public int x,y;

    public OverlapPosition(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public override bool Equals(object obj)
    {
        //
        // See the full list of guidelines at
        //   http://go.microsoft.com/fwlink/?LinkID=85237
        // and also the guidance for operator== at
        //   http://go.microsoft.com/fwlink/?LinkId=85238
        //
        
        if (obj == null || GetType() != obj.GetType())
        {
            return false;
        }

        

        return this.x == ((OverlapPosition) obj).x && this.y == ((OverlapPosition) obj).y;
    }
    
    // override object.GetHashCode
    public override int GetHashCode()
    {
        return this.x*43 + this.y*59;
    }
}
public class Claim {
    public int id,distanceFromLeft,distanceFromTop,height,width;

    public Claim(int id, int x, int y, int height, int width) {
        this.id = id;
        this.distanceFromLeft = x;
        this.distanceFromTop = y;
        this.height = height;
        this.width = width;
    }

    public List<OverlapPosition> overlap(Claim otherClaim) {
        int upper, lower, leftMost, rightMost;

        upper = Max(this.distanceFromTop, otherClaim.distanceFromTop);
        lower = Min(this.distanceFromTop + this.height, otherClaim.distanceFromTop + otherClaim.height);
        leftMost = Max(this.distanceFromLeft, otherClaim.distanceFromLeft);
        rightMost = Min(this.distanceFromLeft + this.width, otherClaim.distanceFromLeft + otherClaim.width);

        List<OverlapPosition> overlapPositions = new List<OverlapPosition>();

        for (int x = leftMost; x < rightMost; x++) {
            for (int y = upper; y < lower; y++) {
                overlapPositions.Add(new OverlapPosition(x,y));
            }
        }

        return overlapPositions;
    }
}
public class December3 {

    private static Claim parseClaim(string claim) {
        var parts = Regex.Split(claim, " @ ");
        var id = Int32.Parse(parts[0].Substring(1,parts[0].Length-1));
        parts = Regex.Split(parts[1], ": ");
        var startParts = Regex.Split(parts[0], ",");
        var dimParts = Regex.Split(parts[1], "x");

        var left = Int32.Parse(startParts[0]);
        var top = Int32.Parse(startParts[1]);
        var width = Int32.Parse(dimParts[0]);
        var height = Int32.Parse(dimParts[1]);

        return new Claim(id, left, top, height, width);
    }

    private static Claim[] parseClaims(string[] stringClaims) {
        Claim[] claims = new Claim[stringClaims.Length];

        for (int i=0; i < stringClaims.Length; i++) {
            claims[i] = parseClaim(stringClaims[i]);
        }
        return claims;
    }

    public static int firstClaimNotOverlapping(string[] stringClaims) {
        Claim[] claims = parseClaims(stringClaims);
        bool overlap = true;

        for (int i=0; i < claims.Length; i++) {
            overlap = false;
            for (int j=0; j < claims.Length; j++) {
                if (j == i) {
                    continue;
                }
                var overlaps = claims[i].overlap(claims[j]);
                if (overlaps.Count != 0) {
                    overlap = true;
                    break;
                }
            }
            if (!overlap) {
                return claims[i].id;
            }
        }
        return 0;
    }

    public static Dictionary<OverlapPosition, bool> overlap(string[] stringClaims) {
        Claim[] claims = parseClaims(stringClaims);
        Dictionary<OverlapPosition, bool> overlapPositions = new Dictionary<OverlapPosition, bool>();


        for (int i=0; i < claims.Length; i++) {
            for (int j= i+1; j < claims.Length; j++) {
                var overlaps = claims[i].overlap(claims[j]);
                //Console.WriteLine("Overlaps between {0} and {1}", claims[i].id, claims[j].id);
                foreach (var overlap in overlaps) {
                    //Console.WriteLine("X: {0} Y:{1}", overlap.x, overlap.y);
                    overlapPositions[overlap] = true;
                }
            } 
        }

        return overlapPositions;
    }

    public static void drawOverlap(Dictionary<OverlapPosition, bool> overlaps, int sheetSize) {
        for (int y=0; y < sheetSize; y++) {
            for (int x=0; x < sheetSize; x++) {
                if (overlaps.ContainsKey(new OverlapPosition(x, y))) {
                    Console.Write("x ");
                } else {
                    Console.Write(". ");
                }
            }
            Console.WriteLine();
        }
    }
}