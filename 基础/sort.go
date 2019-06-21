package main
import (
	"fmt"
	"sort"
)

type MyStringList []string

func (m MyStringList) Len() int {
	return len(m)
}

func (m MyStringList) Less(i,j int) bool {
	return m[i] < m[j]
}

func (m MyStringList) Swap(i,j int) {
	m[i],m[j] = m[j],m[i]
}

/*func main() {
	// names := sort.StringSlice{
 //        "3. Triple Kill",
 //        "5. Penta Kill",
 //        "2. Double Kill",
 //        "4. Quadra Kill",
 //        "1. First Blood",
	// }

	// sort.Sort(names)


	names := [] string{
        "3. Triple Kill",
        "5. Penta Kill",
        "2. Double Kill",
        "4. Quadra Kill",
        "1. First Blood",
	}

	sort.Strings(names)
	for _,v := range names{
		fmt.Printf("%s\n",v)
	}
}*/

type HeroKind int

const (
	None HeroKind = iota
	Tank
	Assassin
	Mage
)

type Hero struct {
	Name string
	Kind HeroKind
}

type Heros []*Hero

func (s Heros) Len() int {
	return len(s)
}

func (s Heros) Less(i,j int)bool{
	if s[i].Kind != s[j].Kind {
		return s[i].Kind < s[j].Kind
	}

	return s[i].Name < s[j].Name
}

func (s Heros) Swap(i,j int){
	s[i],s[j] = s[j],s[i]
}

func main() {
	// heros := Heros{
 //        &Hero{"吕布", Tank},
 //        &Hero{"李白", Assassin},
 //        &Hero{"妲己", Mage},
 //        &Hero{"貂蝉", Assassin},
 //        &Hero{"关羽", Tank},
 //        &Hero{"诸葛亮", Mage},
	// }

	// sort.Sort(heros)


    heros := []*Hero{
        {"吕布", Tank},
        {"李白", Assassin},
        {"妲己", Mage},
        {"貂蝉", Assassin},
        {"关羽", Tank},
        {"诸葛亮", Mage},
    }

    sort.Slice(heros,func(i,j int) bool {
    	if heros[i].Kind != heros[j].Kind{
    		return heros[i].Kind < heros[j].Kind
    	}

    	return heros[i].Name < heros[j].Name
	})

	for _,v := range heros {
		fmt.Printf("%+v\n",v)
	}
}