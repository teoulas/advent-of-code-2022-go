package main

import "fmt"

const input = `cvtvbvfbvfbfqfrrtnnmqqmcqqpfplfpphchpcpcfppmwwhhscccjwjnnctcjtjgjbgjjdzjjwzwfffnhhgppjmpmgppppzbzbrzznttfqqhtqqbssrjjffgssnqnsqnsqnnmlmmztmtqmqccqscsshrsrtstrrjwrrswwpnnftfwwwzznsstbbbtzzncnbcnbccwtwctwwjgjcjtccblclgcgjgngwgbggpbbbbbgcczzfsfbbwlwwnqwnwbnwnmmbjjnfnqfqfvvgvqqvcvggpnpssptpnpjjvvpdpqpcqpcqqnhhnqncqclcfcfjcjhjdjqddhndnmmvmfmsmdsdwsdswsnsgsmsrmrjjpzzzzcppczcscsgcczqcqnnpzpbbdqqqqsqlslzljzjdjbjdjvvlhvvnbbjtbbqbdqdggzwgghhdssmttjtttmwtmtjjcpjpjptjpjmpptnntgttjpttljjwtjwwrgrvgrvrmvvrvdrdcdmmvvbpbddfbdbppdccjdjvjpjjgzjzdzfddqcqlccbvbqvqvjqjggqrqdrqrhhwggmvgmgvgrvvzwzddlzlbbdpbdbtbhthnthttjwtwffwvvrccnpcnclcgcrrgdrrgzgmzggzhhftfppnwppqhhmddzhdzzzjqjhqqcgqqmjjwppfrprlrjlrjjlhjlhjjnbnlblbzlbbfjjwfjjdcdndvvbllfmlmfmvmnnlcnncsnslsvlsvlsltslslqsqzssslqlssqrrpdpwdppsfpsprpffchhcvcpvvjpjvjmvvrjrqrzqqqlvvzzggcgrcgcmmvnmmchmmswsvwsvwswnwswdwmwzmmqfqrqtrqqwqlqlmqmsmggpzpjjlvllzplzlbzzbhhglhhwwhfwwpwgwrwlwplppfvflftltffccmcrcgrcgghgmmbrbppwvvghvhphhztzpzzwbbtpppcnpccnvcvbccvfvmffmtmbttdqdrrfnrnffsgstttbhhgmmtffhthllsbllsqsllptpwpvpdpsddsgsrrngnpggwbwwppfqfmqfmfvfzzzldltdtvdvtvtbvtgpvtnvcsdbwjrqwrfzdhwfwnmtrltdhsfzwwfpscrlsffchnfszqdddsztwzcbwdwfnljpfdsqgpnjffqfgsprggfwgvfldvzmnrpnzvfscswnzbljtvshzmcztqqwtzdstsqggqqvvlmspjjcmllcmsndwpllrvlrglphcqjgwzntwqffthrfmfgpcvjvclvdrjrpnwtrlshlnwgbbznpflsfmgvbwmnwjpnqhmdnsmwnmblqzjfzfcslhntnjhbrvvlfbffsbhrtvjltqrqszsrpmdghmhsrqlmrjdsgtsgwggdlqptqrvgmwrpcjrnmmmcdpgjzgghqmshhmddfvfrqwnwvwbvzzpmwzmvvgsqlsgvmtczwnbzrtnzjnpjjghdzsmcgvqpdhltvltbfvgsgzhvqjjnpspngqpvqwvtqrczqzwrclqbnplmvlnvcwdbzlrnwbppvdfzmblnqfpfqlngttrrcjzfjnrnmssdrrvbvhtrfzblzsjbqwbttsshmcmfhrfbfqspvcfmqbwjszmvwqnpszcfhtnqtftvtpvcjrcnlmfnvlpjzspcnmvvshpsdfhtbhgjcmjmtpmtcbwnlgtnvlwlggbnzfrrqljmjrwjtzfmgbllmmjtfhrszfrfpftmpzhzdhfwdbdnwqspjbrlglgdwtfzcrdhwgqltvthbnrqfngjgnddqlcpjflqzvbhmlzsvpvlzdvhspmchrjlzfgnlpwdnszqdqpnrchdrswfwnflslcpfmbfrplsmchbgjhhhlzjlvztpdhslrpmmcthhjwfczjtmbmsvzvsqrvzrcnflqldfnhjzlcsfhgrtqtscnbzqjglgjnqbtprrprrdtqqtlqwdqnmgbglvwtrpnpjqdsmvczbnldvmgmtpmmqbwgjbfsqbjvnjhhlrppptlrjsptpvghbgtvwvsrnggznjsbjpnwpdsqmmjjqjdwshqwzvqgpnlhdldwzmvmdnhrnhpscrbnbdtmfqwlmrlszljpjdqgjpgcdwtqzsbqsndjdslrstjwbbfnjvrjlbrphqspltjzpvsgccfrmqhfhjmlslncpgpflhjlrzqsffrtcgsvzpjmlqjznscgjrjhzfjbbchvgjgvfsrpnzdrstdtsltqldfldvfdjqdwsltcbnnvcgvnjpgdjzsbzgfvgvtmghwfmrmblscfdqlsdgcdbgsphfvvvnffcqwlrvnsrjmvvdrjdljlppsfsrmwwvmsqcgtgtfjhqbqvlsnwdwzhncgprbspjzqnbllwcjnqpfzlhscdtvlldffngwttntsbtlctdfbzjgclgbhhjtlffvtgbdsswczlgmwntrrnnbdtflplfbtqcfrhtgjnrltvzqlcvtdthldqrvjdmlntdjzcncrplgmdbqhlrrfdszrnbrllhgpzzgjnqzrjrtjjntqvdphgvjtbtqfjwvlpprrcvmpjcdmntvtqjvjrhngbtmdtflpqmlmmvjjjpdlcvntlvlgstbblwfwjrbcqjgzvjjzflfhdvgnhpnqprcvmhsrpmgfcgjpdbqpmpvlnmzdngzqcqtdsbbqbdhtgpqmqwtcctsvhwmlhpgzgpwwsqngsrwzghfhzbqcpspfhqbllbvpwgwdnghqnvjtfqtwgvnbfbpcffbwwbfqgqnftpqdtwhmrprpgdlnwrcmdndllwbvlrrjvmwlbgfwgbhtmbqhfvhlhlthbrljwnwcgsdgsbbzgnwzmjfgnfjcqhspjwctntznjptsnbpwtfzmrqwjzrmdthfwdvttrcghbjndfmtswhsmdwwwffmgsvshmlhdbwdscjnvrvqnwttlsmfbssnsscftcrpwfbbplbpmvzmttvbgpnjtqqzcfdfwqdphggwbnhfvssmdjvbvtlrcghltthqrqdbvjnbqpzjcggrbnrswhvdcrmgncsmhsmstmcbwgbfwvzqtdrpjbgljdgcncgcdslhftptnvfhgpsjqqmnsjmqhztnfzqqqhccbnqjlhqcwwbfzgpmbvjszswjlhcqshtqdmwmhpgqgzbwbtchrdwgnvfqcgcmngtddpzfgjrrssdmcnbtpbhfvfbdgzmcbjqtdqrnrdchtrcfpmtnvfczdlnjplwccfmzjwwvdsrcjznrcbczdqwshjgvhjpgdgjvtvmdzncrqgffzhcbwstvltqbrzfbnrrvlntjmzctvnqhdcgvvmmbfplzpwgsdzdtvnwnshnwbwqplbctntslczwhblsnzdcsgltvlfvbtvthvlsmzmfgzhpvrvczdhmcpmgqzbdhbcdbstrrtzgrbhhwwghzlgcdtlcjtjdcwdtqpvpwzrwmhczfsdmwjlgtmzvdbwzdnsztvqzhjpfrqzllgvrsbmhvrlzwtfmpzchftrzlmndpwqgcmhdzmjrjvpbnrcmjvzbzlqngmrzcbdwfcwmlnlqsdgbvmvqcsjtwgvtpmqdtdqfrshmlgnjqdwsrpjcfnfpjzlhslsvppmgprqlpcfztcngmqgrvqvpcszzgtbpcrghmgnvgplctgdljhwjgbzmjrdsccdtrwjrrbttzmpvmzlmqdfwdccfzqztdjmcgrcdwghgmmntdwrclshlgnsmhmltwcczncndtglnqvcvbwlvcgqtzpnqnlllchwjwrlhcdlbjlzvhvtjsttqhvhfqbsqjpfcmpsfshgrrwgprjtfnhsjqqntrjfgjjmjbqwdjlcjjdljpppgvqvgrrwvclbmncmqglhbtjcfrbjmhqgmlpzgljntqvrcggmrgtbrftvhrnjbvzstlnzrqcgncnjhdrnqcnwcrstvdsgnwzjzmqsdtwzzjmfgzpslhqclhtjnrlwwfrqgmhsttfdprhgmphzfdfqntwztltnnfqwgdnnbrlgzmtwqsvhbjjqtcgmghnbchdfgfrdtjnvlgtqdgrfwmgnmpslqmqzdgqztnppbc`

func main() {
	findUnique(input, 4)
	findUnique(input, 14)
}

func findUnique(s string, l int) {
	size := len(input)
	for i := l - 1; i < size; i++ {
		substr := input[i-(l-1) : i+1]
		if allDiff(substr) {
			fmt.Printf("%s at position %d\n", substr, i+1)
			break
		}
	}
}

func allDiff(s string) bool {
	seen := make(map[rune]bool)
	for _, r := range s {
		seen[r] = true
	}
	return len(seen) == len(s)
}
