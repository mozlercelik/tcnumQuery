package tcnumquery

import (
	"strconv"
	"strings"
)

func Querytc(tc string) bool {
	var tektop, cifttop, toplam, kontrol int = 0, 0, 0, 0

	tcnum := strings.Split(tc, "") // tcnum = array

	if firstnumber, _ := strconv.Atoi(tcnum[0]); (len(tc) != 11) || (firstnumber == 0) {
		return false
	}

	for i := 0; i < 11; i++ {
		intdeger, err := strconv.Atoi(tcnum[i])

		if err != nil {
			return false
		} else if i < 9 {
			if (i % 2) == 0 {
				tektop += intdeger // 1,3,5,7,9. elemanı toplama
			} else {
				cifttop += intdeger //2,4,6,8. elemanı toplama
			}
		}
		
		if i < 10 {
			toplam += intdeger // 1. elemandan 10. elemana(dahil) kadar toplama
		}
	}

	kontrol = ((tektop * 7) - cifttop) % 10 //10. eleman kontrol koşulu
	onuncu, _ := strconv.Atoi(tcnum[9])     //10. eleman
	onbirinci, _ := strconv.Atoi(tcnum[10]) //11. eleman

	if (kontrol != onuncu) || (toplam%10 != onbirinci) { //10, 11. eleman kontrol sorgusu
		return false
	}

	return true
}
