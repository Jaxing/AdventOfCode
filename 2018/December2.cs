using System;
using System.IO;
using System.Text;
using System.Collections.Generic;
public class December2 {
    public static int checksum(string[] ids) {
        int dups = 0;
        int trips = 0;

        foreach (var id in ids) {
            Dictionary<char, int> occurencies = new Dictionary<char, int>();

            foreach (char c in id) {
                occurencies[c] = occurencies.GetValueOrDefault(c) + 1;
            }
            if (occurencies.ContainsValue(2)) {
                dups++;
            }
            if (occurencies.ContainsValue(3)) {
                trips++;
            }
        }

        return dups * trips;
    }

    public static string findPackages(string[] ids) {
        List<string> idList = new List<string>(ids);
        idList.Sort(string.Compare);

        for (int i=0; i < ids.Length - 1; i++) {
            string common = commonChars(idList[i], idList[i+1]);

            if (common.Length == idList[i].Length - 1) {
                return common;
            }
        }

        return "";
    }

    public static string commonChars(string a, string b) {
        StringBuilder common = new StringBuilder();

        for (int i=0; i < a.Length && i < b.Length; i++) {
            if (a[i] == b[i]) {
                common.Append(a[i]);
            }
        }

        return common.ToString();
    }
}