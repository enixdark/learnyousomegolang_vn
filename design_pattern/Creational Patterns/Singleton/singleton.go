package singleton

type Coach struct{
        name string
        age int
}

func (c *Coach) IncreaseAge() int {
        c.age = c.age + 1
        return c.age
}

func (c* Coach) GetAge() int {
        return c.age
}

func (c* Coach) GetName() string {
        return c.name
}

func GetCoach() *Coach {
        if (coach == nil ) {
            coach = &Coach{"A", 50}
        }

        return coach
}

var coach *Coach

