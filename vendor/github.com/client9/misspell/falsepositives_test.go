package misspell

import (
	"testing"
)

func TestFalsePositives(t *testing.T) {
	cases := []string{
		"importEnd",
		"drinkeries",
		"subscripting",
		"unprojected",
		"updaters",
		"templatize",
		"requesters",
		"requestors",
		"replicaset",
		"parallelise",
		"parallelize",
		"perceptron",  // http://foldoc.org/perceptron
		"perceptrons", // ^^
		"convertors",  // alt spelling
		"adventurers",
		" s.svc.GetObject ",
		"infinitie.net",
		"foo summaries\n",
		"thru",
		"publically",
		"6YUO5",   // base64
		"cleaner", // triggered by "cleane->cleanser" and partial word FP
		"  http.Redirect(w, req, req.URL.Path, http.StatusFound) ",
		"url is http://zeebra.com ",
		"path is /zeebra?zeebra=zeebra ",
		"Malcom_McLean",
		"implementor", // alt spelling, see https://github.com/client9/misspell/issues/46
		"searchtypes",
		" witness",
		"returndata",
		"UNDERSTOOD",
		"textinterface",
		" committed ",
		"committed",
		"Bengali",
		"Portuguese",
		"scientists",
		"causally",
		"embarrassing",
		"setuptools", // python package
		"committing",
		"guises",
		"disguise",
		"begging",
		"cmo",
		"cmos",
		"borked",
		"hadn't",
		"Iceweasel",
		"summarised",
		"autorenew",
		"travelling",
		"republished",
		"fallthru",
		"pruning",
		"deb.VersionDontCare",
		"authtag",
		"intrepid",
		"usefully",
		"there",
		"definite",
		"earliest",
		"Japanese",
		"international",
		"excellent",
		"gracefully",
		"carefully",
		"class",
		"include",
		"process",
		"address",
		"attempt",
		"large",
		"although",
		"specific",
		"taste",
		"against",
		"successfully",
		"unsuccessfully",
		"occurred",
		"agree",
		"controlled",
		"publisher",
		"strategy",
		"geoposition",
		"paginated",
		"happened",
		"relative",
		"computing",
		"language",
		"manual",
		"token",
		"into",
		"nothing",
		"datatool",
		"propose",
		"learnt",
		"tolerant",
		"whitehat",
		"monotonic",
		"comprised",
		"indemnity",
		"flattened",
		"interrupted",
		"inotify",
		"occasional",
		"forging",
		"ampersand",
		"decomposition",
		"commit",
		"programmer", // "grammer"
		//		"requestsinserted",
		"seeked",      // technical word
		"bodyreader",  // variable name
		"cantPrepare", // variable name
		"dontPrepare", // variable name
		"\\nto",       // https://github.com/client9/misspell/issues/93
		"4f8b42c22dd3729b519ba6f68d2da7cc5b2d606d05daed5ad5128cc03e6c6358", // https://github.com/client9/misspell/issues/97
	}
	r := New()
	r.Debug = true
	for casenum, tt := range cases {
		got, _ := r.Replace(tt)
		if got != tt {
			t.Errorf("%d: %q got converted to %q", casenum, tt, got)
		}
	}
}
