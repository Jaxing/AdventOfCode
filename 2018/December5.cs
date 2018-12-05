using System;
using System.Text.RegularExpressions;

public class December5 {
    public static string polymerReaction(string input) {
        bool noSplit = false;
        int splitIndex = 0;
        for (int i=0; i < input.Length - 1; i++) {
            if (Math.Abs(((int) input[i]) - ((int) input[i+1])) == 32) {
                splitIndex = i;
                noSplit = true;
                break;
            }
        }

        if (!noSplit) {
            return input;
        }

        input = input.Remove(splitIndex,2);
        return polymerReaction(input);
    } 

    public static int unitsLeft(string input) {
        return polymerReaction(input).Length;
    }

    public static int findBestPolymer(string input) {
        int minLenght = int.MaxValue;

        for (char c='a'; c<='z'; c++) {
            var re = new Regex(c.ToString() ,RegexOptions.IgnoreCase);
            var res = re.Split(input);
            var poly = polymerReaction(string.Join("", res));

            if (poly.Length < minLenght) {
                minLenght = poly.Length;
            } 
        }

        return minLenght;
    }
}