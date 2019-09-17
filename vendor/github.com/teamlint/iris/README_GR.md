# Iris <a href="README.md"> <img width="20px" src="https://iris-go.com/images/flag-unitedkingdom.svg?v=10" /></a> <a href="README_ZH.md"><img width="20px" src="https://iris-go.com/images/flag-china.svg?v=10" /></a> <a href="README_ES.md"><img width="20px" src="https://iris-go.com/images/flag-spain.png" /></a> <a href="README_KO.md"><img width="20px" src="https://iris-go.com/images/flag-south-korea.svg" />

[![build status](https://img.shields.io/travis/kataras/iris/master.svg?style=for-the-badge)](https://travis-ci.org/kataras/iris) [![report card](https://img.shields.io/badge/report%20card-a%2B-ff3333.svg?style=for-the-badge)](https://goreportcard.com/report/github.com/kataras/iris)<!--[![godocs](https://img.shields.io/badge/go-%20docs-488AC7.svg?style=for-the-badge)](https://godoc.org/github.com/kataras/iris)--> [![view examples](https://img.shields.io/badge/learn%20by-examples-0077b3.svg?style=for-the-badge)](https://github.com/kataras/iris/tree/master/_examples) [![chat](https://img.shields.io/gitter/room/iris_go/community.svg?color=blue&logo=gitter&style=for-the-badge)](https://gitter.im/iris_go/community) [![release](https://img.shields.io/badge/release%20-v11.2-0077b3.svg?style=for-the-badge)](https://github.com/kataras/iris/releases)

Το Iris είναι ένα γρήγορο, απλό αλλά και πλήρως λειτουργικό και πολύ αποδοτικό web framework για τη Go γλώσσα προγραμματισμού. Παρέχει ένα εκφραστικό και εύχρηστο υπόβαθρο για την επόμενη ιστοσελίδα σας.

Μάθετε τι [λένε οι άλλοι για το Iris](https://iris-go.com/testimonials/) και δώστε ένα **αστεράκι** στο GitHub.

## Μαθαίνοντας το Iris

<<<<<<< HEAD
Επιτέλους, ένα πραγματικά ισάξιο (και με το παραπάνω) expressjs web framework για τη γλώσσα προγραμματισμού Go.

Μάθετε τι [λένε οι άλλοι για το Iris](#%CE%A5%CF%80%CE%BF%CF%83%CF%84%CE%AE%CF%81%CE%B9%CE%BE%CE%B7) και [δώστε ένα αστέρι](https://github.com/teamlint/iris/stargazers) στο github repository για να μένετε [πάντα ενημερωμένοι](https://facebook.com/iris.framework).

## Yποστηρικτές

Eυχαριστούμε όλους τους υποστηρικτές μας! 🙏 [Γίνετε ένας από αυτούς](https://iris-go.com/donate)

<a href="https://iris-go.com/donate" target="_blank"><img src="https://iris-go.com/backers.svg?v=2"/></a>
=======
<details>
<summary>Γρήγορο ξεκίνημα</summary>
>>>>>>> upstream/master

```sh
# υποθέτοντας ότι ο παρακάτω κώδικας
# βρίσκεται στο example.go αρχείο
#
$ cat example.go
```

```go
package main

import "github.com/teamlint/iris"

func main() {
    app := iris.Default()
    app.Get("/ping", func(ctx iris.Context) {
        ctx.JSON(iris.Map{
            "message": "pong",
        })
    })

    app.Run(iris.Addr(":8080"))
}
```

```sh
# τρέξτε το example.go και
# επισκεφτείτε την σελίδα http://localhost:8080/ping
# στο πρόγραμμα περιήγησης σας
#
$ go run example.go
```

<<<<<<< HEAD
## Εγκατάσταση

Η μόνη απαίτηση είναι η [Go Γλώσσα Προγραμματισμού](https://golang.org/dl/)

```sh
$ go get -u github.com/teamlint/iris
```

Το Iris εκμεταλλεύεται τη λεγόμενη λειτουργία [vendor directory](https://docs.google.com/document/d/1Bz5-UB7g2uPBdOx-rw5t9MxJwkfpx90cqG9AFL0JAYo). Παίρνετε πλήρως αναπαραγωγίσιμα builds, καθώς αυτή η μέθοδος προστατεύει από τις upstream μετονομασίες και διαγραφές.

[![Iris vs .NET Core(C#) vs Node.js (Express)](https://iris-go.com/images/benchmark-new-gray.png)](_benchmarks/README_UNIX.md)

_Η τελευταία ενημέρωση έγινε την [Τρίτη, 21 Νοεμβρίου του 2017](_benchmarks/README_UNIX.md)_

<details>
<summary>Στοιχεία αναφοράς από τρίτες πηγές σε σχέση με τα υπόλοιπα web frameworks</summary>

![Comparison with other frameworks](https://raw.githubusercontent.com/smallnest/go-web-framework-benchmark/4db507a22c964c9bc9774c5b31afdc199a0fe8b7/benchmark.png)

</details>

## Υποστήριξη

- To [HISTORY](HISTORY_GR.md#fr-11-january-2019--v1111) αρχείο είναι ο καλύτερος σας φίλος, περιέχει πληροφορίες σχετικά με τις τελευταίες λειτουργίες(features) και αλλαγές
- Μήπως τυχαίνει να βρήκατε κάποιο bug; Δημοσιεύστε το στα [github issues](https://github.com/kataras/iris/issues)
- Έχετε οποιεσδήποτε ερωτήσεις ή πρέπει να μιλήσετε με κάποιον έμπειρο για την επίλυση ενός προβλήματος σε πραγματικό χρόνο; Ελάτε μαζί μας στην [συνομιλία κοινότητας](https://chat.iris-go.com)
- Συμπληρώστε την αναφορά εμπειρίας χρήστη κάνοντας κλικ [εδώ](https://docs.google.com/forms/d/e/1FAIpQLSdCxZXPANg_xHWil4kVAdhmh7EBBHQZ_4_xSZVDL-oCC_z5pA/viewform?usp=sf_link)
- Σας αρέσει το Iris; Τιτιβίστε κάτι για αυτό! Άνθρωποι από ολόκληρο τον πλανήτη έχουνε μιλήσει για αυτό ακριβώς:

<a href="https://twitter.com/gelnior/status/769100480706379776"> 
    <img src="https://comments.iris-go.com/comment27_mini.png" width="350px">
</a>

<a href="https://twitter.com/MeAlex07/status/822799954188075008"> 
    <img src="https://comments.iris-go.com/comment28_mini.png" width="350px">
</a>

<a href="https://twitter.com/_mgale/status/818591490305761280"> 
    <img src="https://comments.iris-go.com/comment29_mini.png" width="350px">
</a>
<a href="https://twitter.com/VeayoX/status/813273328550973440"> 
    <img src="https://comments.iris-go.com/comment30_mini.png" width="350px">
</a>

<a href="https://twitter.com/pvsukale/status/745328224876408832"> 
    <img src="https://comments.iris-go.com/comment31_mini.png" width="350px">
</a>

<a href="https://twitter.com/blainsmith/status/745338092211560453"> 
    <img src="https://comments.iris-go.com/comment32_mini.png" width="350px">
</a>

<a href="https://twitter.com/tjbyte/status/758287014210867200"> 
    <img src="https://comments.iris-go.com/comment33_mini.png" width="350px">
</a>

<a href="https://twitter.com/tangzero/status/751050577220698112"> 
    <img src="https://comments.iris-go.com/comment34_mini.png" width="350px">
</a>

<a href="https://twitter.com/tjbyte/status/758287244947972096"> 
    <img src="https://comments.iris-go.com/comment33_2_mini.png" width="350px">
</a>

<a href="https://twitter.com/ferarias/status/902468752364773376"> 
    <img src="https://comments.iris-go.com/comment41.png" width="350px">
</a>

<br/><br/>

Για περισσότερες πληροφορίες σχετικά με τη συμβολή στο Iris, διαβάστε το [CONTRIBUTING.md](CONTRIBUTING.md) αρχείο.

[Κατάλογος όλων των Συνεργατών](https://github.com/teamlint/iris/graphs/contributors)

## Μάθηση

Πρώτα απ 'όλα, ο πιο σωστός τρόπος για να ξεκινήσετε με ένα web framework είναι να μάθετε τα βασικά της γλώσσας προγραμματισμού και των τυπικών της δυνατοτήτων `http`, αν η εφαρμογή σας είναι ένα πολύ απλό προσωπικό έργο χωρίς απαιτήσεις επιδόσεων και συντηρησιμότητας, ίσως να θέλετε να προχωρήσετε μόνο με τα τυπικά πακέτα, εαν οχι τότε ακολουθήστε τις παρακάτω οδηγίες:

- Πλοηγηθείτε μέσω των **100+1** **[παραδειγμάτων](_examples)** και μερικές [απλές εφαρμογές για αρχάριους](#iris-starter-kits) που δημιουργήσαμε για εσάς
- Διαβάστε τα [godocs](https://godoc.org/github.com/teamlint/iris) για οποιαδήποτε λεπτομέρεια
- Ετοιμάστε ένα φλιτζάνι καφέ ή τσάι, ό,τι σας ευχαριστεί περισσότερο και διαβάστε κάποια [άρθρα](#articles) που βρήκαμε για εσάς

### Iris starter kits

<!-- table form 
| Description | Link |
| -----------|-------------|
| Hasura hub starter project with a ready to deploy golang helloworld webapp with IRIS! | https://hasura.io/hub/project/hasura/hello-golang-iris |
| A basic web app built in Iris for Go |https://github.com/gauravtiwari/go_iris_app |
| A mini social-network created with the awesome Iris💖💖 | https://github.com/iris-contrib/Iris-Mini-Social-Network |
| Iris isomorphic react/hot reloadable/redux/css-modules starter kit | https://github.com/iris-contrib/iris-starter-kit |
| Demo project with react using typescript and Iris | https://github.com/ionutvilie/react-ts |
| Self-hosted Localization Management Platform built with Iris and Angular | https://github.com/iris-contrib/parrot |
| Iris + Docker and Kubernetes | https://github.com/iris-contrib/cloud-native-go |
| Quickstart for Iris with Nanobox | https://guides.nanobox.io/golang/iris/from-scratch |
-->

1. [snowlyg/IrisApiProject: Iris + gorm + jwt + sqlite3](https://github.com/snowlyg/IrisApiProject) **NEW-Chinese**
2. [yz124/superstar: Iris + xorm to implement the star library](https://github.com/yz124/superstar) **NEW-Chinese**
3. [jebzmos4/Iris-golang: A basic CRUD API in golang with Iris](https://github.com/jebzmos4/Iris-golang)
4. [gauravtiwari/go_iris_app: A basic web app built in Iris for Go](https://github.com/gauravtiwari/go_iris_app)
5. [A mini social-network created with the awesome Iris💖💖](https://github.com/iris-contrib/Iris-Mini-Social-Network)
6. [Iris isomorphic react/hot reloadable/redux/css-modules starter kit](https://github.com/iris-contrib/iris-starter-kit)
7. [ionutvilie/react-ts: Demo project with react using typescript and Iris](https://github.com/ionutvilie/react-ts)
8. [Self-hosted Localization Management Platform built with Iris and Angular](https://github.com/iris-contrib/parrot)
9. [Iris + Docker and Kubernetes](https://github.com/iris-contrib/cloud-native-go)
10. [nanobox.io: Quickstart for Iris with Nanobox](https://guides.nanobox.io/golang/iris/from-scratch)
11. [hasura.io: A Hasura starter project with a ready to deploy Golang hello-world web app with IRIS](https://hasura.io/hub/project/hasura/hello-golang-iris)

> Έχετε χτίσει κάτι παρόμοιο; [Ενημέρωσέ μας](https://github.com/teamlint/iris/pulls)!
=======
> Η δρομολόγηση τροφοδοτείται από το [muxie](https://github.com/kataras/muxie), το πιο ισχυρό και ταχύτερο λογισμικό βασισμένο σε trie αλγόριθμο που γράφτηκε σε Go.

</details>

Το Iris περιέχει εκτενείς και λεπτομερείς **[wiki](https://github.com/kataras/iris/wiki)** καθιστώντας το εύκολο στην εκμάθηση.
>>>>>>> upstream/master

Για λεπτομερέστερη τεχνική τεκμηρίωση μπορείτε να κατευθυνθείτε προς τα [godocs](https://godoc.org/github.com/kataras/iris) μας. Και για εκτελέσιμο κώδικα μπορείτε πάντα να επισκέπτεστε τα [παραδείγματα](_examples/).

### Σας αρέσει να διαβάζετε ενώ ταξιδεύετε;

Μπορείτε να [ζητήσετε](https://bit.ly/iris-req-book) σήμερα την PDF έκδοση και την online πρόσβαση στο Ηλεκτρονικό μας **Βιβλίο(E-Book)** και να συμμετάσχετε στην ανάπτυξη του Iris.

[![https://iris-go.com/images/iris-book-overview.png](https://iris-go.com/images/iris-book-overview.png)](https://bit.ly/iris-req-book)

## Συνεισφορά

Θα θέλαμε να δούμε τη συμβολή σας στο Iris Web Framework! Για περισσότερες πληροφορίες σχετικά με το πως μπορείτε να συμβάλετε, δείτε το [CONTRIBUTING.md](CONTRIBUTING.md) αρχείο.

[Κατάλογος όλων των συνεισφορών](https://github.com/kataras/iris/graphs/contributors).

## Αδυναμίες Ασφάλειας

Εάν εντοπίσετε κάποια αδυναμία ασφαλείας του Iris, στείλτε ένα μήνυμα ηλεκτρονικού ταχυδρομείου στο [iris-go@outlook.com](mailto:iris-go@outlook.com). Όλες οι τυχών αδυναμίες ασφαλείας θα αντιμετωπιστούν άμεσα.

## Άδεια Χρήσης

Το όνομα "Iris" εμπνεύστηκε από την ελληνική μυθολογία, από την θεά Ίριδα.

Το Iris Web Framework είναι δωρεάν λογισμικό ανοιχτού λογισμικού με άδεια χρήσης [3-Clause BSD](LICENSE).
