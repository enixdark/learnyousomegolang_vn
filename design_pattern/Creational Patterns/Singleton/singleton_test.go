package singleton

import "testing"

func TestGetCoach(t *testing.T) {
        c1 := GetCoach()
        if c1 == nil {
                t.Error("expected pointer to Singleton after calling GetCoach(), not nil")
        }

        c2 := GetCoach()

        if c1 != c2 {
                t.Error("expected between c1 and c2 are'diff")
        }

        name := c1.GetName()

        if name != "A" {
                t.Error("expected between name of coach instance 's not A")
        }

        age := c1.GetAge()

        c1.IncreaseAge()

        ageNew := c1.GetAge()

        if age + 1 != ageNew {
                t.Error("expected the age of coach doesn't increase")
        }
}