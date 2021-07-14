// Copyright 2021 Mohammad Shafiee and The SwEphGo Authors
//
// Licensed under the GNU General Public License, Version 3.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.gnu.org/licenses/gpl-3.0.html
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package swephgo

const (
	// SeAunitToKm as defined in src/swephexp.h:88
	SeAunitToKm = 1.495978707e+08
	// SeAunitToLightyear as defined in src/swephexp.h:89
	SeAunitToLightyear = 1.5812507409819728e-05
	// SeAunitToParsec as defined in src/swephexp.h:90
	SeAunitToParsec = 4.848136811095274e-06
	// SeJulCal as defined in src/swephexp.h:93
	SeJulCal = 0
	// SeGregCal as defined in src/swephexp.h:94
	SeGregCal = 1
	// SeEclNut as defined in src/swephexp.h:99
	SeEclNut = -1
	// SeSun as defined in src/swephexp.h:101
	SeSun = 0
	// SeMoon as defined in src/swephexp.h:102
	SeMoon = 1
	// SeMercury as defined in src/swephexp.h:103
	SeMercury = 2
	// SeVenus as defined in src/swephexp.h:104
	SeVenus = 3
	// SeMars as defined in src/swephexp.h:105
	SeMars = 4
	// SeJupiter as defined in src/swephexp.h:106
	SeJupiter = 5
	// SeSaturn as defined in src/swephexp.h:107
	SeSaturn = 6
	// SeUranus as defined in src/swephexp.h:108
	SeUranus = 7
	// SeNeptune as defined in src/swephexp.h:109
	SeNeptune = 8
	// SePluto as defined in src/swephexp.h:110
	SePluto = 9
	// SeMeanNode as defined in src/swephexp.h:111
	SeMeanNode = 10
	// SeTrueNode as defined in src/swephexp.h:112
	SeTrueNode = 11
	// SeMeanApog as defined in src/swephexp.h:113
	SeMeanApog = 12
	// SeOscuApog as defined in src/swephexp.h:114
	SeOscuApog = 13
	// SeEarth as defined in src/swephexp.h:115
	SeEarth = 14
	// SeChiron as defined in src/swephexp.h:116
	SeChiron = 15
	// SePholus as defined in src/swephexp.h:117
	SePholus = 16
	// SeCeres as defined in src/swephexp.h:118
	SeCeres = 17
	// SePallas as defined in src/swephexp.h:119
	SePallas = 18
	// SeJuno as defined in src/swephexp.h:120
	SeJuno = 19
	// SeVesta as defined in src/swephexp.h:121
	SeVesta = 20
	// SeIntpApog as defined in src/swephexp.h:122
	SeIntpApog = 21
	// SeIntpPerg as defined in src/swephexp.h:123
	SeIntpPerg = 22
	// SeNplanets as defined in src/swephexp.h:125
	SeNplanets = 23
	// SePlmoonOffset as defined in src/swephexp.h:127
	SePlmoonOffset = 9000
	// SeAstOffset as defined in src/swephexp.h:128
	SeAstOffset = 10000
	// SeVaruna as defined in src/swephexp.h:129
	SeVaruna = 30000
	// SeFictOffset as defined in src/swephexp.h:131
	SeFictOffset = 40
	// SeFictOffset1 as defined in src/swephexp.h:132
	SeFictOffset1 = 39
	// SeFictMax as defined in src/swephexp.h:133
	SeFictMax = 999
	// SeNfictElem as defined in src/swephexp.h:134
	SeNfictElem = 15
	// SeCometOffset as defined in src/swephexp.h:136
	SeCometOffset = 1000
	// SeNallNatPoints as defined in src/swephexp.h:138
	SeNallNatPoints = 38
	// SeCupido as defined in src/swephexp.h:141
	SeCupido = 40
	// SeHades as defined in src/swephexp.h:142
	SeHades = 41
	// SeZeus as defined in src/swephexp.h:143
	SeZeus = 42
	// SeKronos as defined in src/swephexp.h:144
	SeKronos = 43
	// SeApollon as defined in src/swephexp.h:145
	SeApollon = 44
	// SeAdmetos as defined in src/swephexp.h:146
	SeAdmetos = 45
	// SeVulkanus as defined in src/swephexp.h:147
	SeVulkanus = 46
	// SePoseidon as defined in src/swephexp.h:148
	SePoseidon = 47
	// SeIsis as defined in src/swephexp.h:150
	SeIsis = 48
	// SeNibiru as defined in src/swephexp.h:151
	SeNibiru = 49
	// SeHarrington as defined in src/swephexp.h:152
	SeHarrington = 50
	// SeNeptuneLeverrier as defined in src/swephexp.h:153
	SeNeptuneLeverrier = 51
	// SeNeptuneAdams as defined in src/swephexp.h:154
	SeNeptuneAdams = 52
	// SePlutoLowell as defined in src/swephexp.h:155
	SePlutoLowell = 53
	// SePlutoPickering as defined in src/swephexp.h:156
	SePlutoPickering = 54
	// SeVulcan as defined in src/swephexp.h:157
	SeVulcan = 55
	// SeWhiteMoon as defined in src/swephexp.h:158
	SeWhiteMoon = 56
	// SeProserpina as defined in src/swephexp.h:159
	SeProserpina = 57
	// SeWaldemath as defined in src/swephexp.h:160
	SeWaldemath = 58
	// SeFixstar as defined in src/swephexp.h:162
	SeFixstar = -10
	// SeAsc as defined in src/swephexp.h:164
	SeAsc = 0
	// SeMc as defined in src/swephexp.h:165
	SeMc = 1
	// SeArmc as defined in src/swephexp.h:166
	SeArmc = 2
	// SeVertex as defined in src/swephexp.h:167
	SeVertex = 3
	// SeEquasc as defined in src/swephexp.h:168
	SeEquasc = 4
	// SeCoasc1 as defined in src/swephexp.h:169
	SeCoasc1 = 5
	// SeCoasc2 as defined in src/swephexp.h:170
	SeCoasc2 = 6
	// SePolasc as defined in src/swephexp.h:171
	SePolasc = 7
	// SeNascmc as defined in src/swephexp.h:172
	SeNascmc = 8
	// SeflgJpleph as defined in src/swephexp.h:186
	SeflgJpleph = 1
	// SeflgSwieph as defined in src/swephexp.h:187
	SeflgSwieph = 2
	// SeflgMoseph as defined in src/swephexp.h:188
	SeflgMoseph = 4
	// SeflgHelctr as defined in src/swephexp.h:190
	SeflgHelctr = 8
	// SeflgTruepos as defined in src/swephexp.h:191
	SeflgTruepos = 16
	// SeflgJ2000 as defined in src/swephexp.h:192
	SeflgJ2000 = 32
	// SeflgNonut as defined in src/swephexp.h:193
	SeflgNonut = 64
	// SeflgSpeed3 as defined in src/swephexp.h:194
	SeflgSpeed3 = 128
	// SeflgSpeed as defined in src/swephexp.h:196
	SeflgSpeed = 256
	// SeflgNogdefl as defined in src/swephexp.h:197
	SeflgNogdefl = 512
	// SeflgNoaberr as defined in src/swephexp.h:198
	SeflgNoaberr = 1024
	// SeflgAstrometric as defined in src/swephexp.h:199
	SeflgAstrometric = 1536
	// SeflgEquatorial as defined in src/swephexp.h:202
	SeflgEquatorial = 2048
	// SeflgXyz as defined in src/swephexp.h:203
	SeflgXyz = 4096
	// SeflgRadians as defined in src/swephexp.h:204
	SeflgRadians = 8192
	// SeflgBaryctr as defined in src/swephexp.h:205
	SeflgBaryctr = 16384
	// SeflgTopoctr as defined in src/swephexp.h:206
	SeflgTopoctr = 32768
	// SeflgOrbelAa as defined in src/swephexp.h:207
	SeflgOrbelAa = 32768
	// SeflgTropical as defined in src/swephexp.h:209
	SeflgTropical = 0
	// SeflgSidereal as defined in src/swephexp.h:210
	SeflgSidereal = 65536
	// SeflgIcrs as defined in src/swephexp.h:211
	SeflgIcrs = 131072
	// SeflgDpsideps1980 as defined in src/swephexp.h:212
	SeflgDpsideps1980 = 262144
	// SeflgJplhor as defined in src/swephexp.h:214
	SeflgJplhor = 262144
	// SeflgJplhorApprox as defined in src/swephexp.h:215
	SeflgJplhorApprox = 524288
	// SeflgCenterBody as defined in src/swephexp.h:216
	SeflgCenterBody = 1048576
	// SeflgTestPlmoon as defined in src/swephexp.h:218
	SeflgTestPlmoon = 2228280
	// SeSidbits as defined in src/swephexp.h:221
	SeSidbits = 256
	// SeSidbitEclT0 as defined in src/swephexp.h:223
	SeSidbitEclT0 = 256
	// SeSidbitSsyPlane as defined in src/swephexp.h:225
	SeSidbitSsyPlane = 512
	// SeSidbitUserUt as defined in src/swephexp.h:227
	SeSidbitUserUt = 1024
	// SeSidbitEclDate as defined in src/swephexp.h:230
	SeSidbitEclDate = 2048
	// SeSidbitNoPrecOffset as defined in src/swephexp.h:233
	SeSidbitNoPrecOffset = 4096
	// SeSidbitPrecOrig as defined in src/swephexp.h:235
	SeSidbitPrecOrig = 8192
	// SeSidmFaganBradley as defined in src/swephexp.h:238
	SeSidmFaganBradley = 0
	// SeSidmLahiri as defined in src/swephexp.h:239
	SeSidmLahiri = 1
	// SeSidmDeluce as defined in src/swephexp.h:240
	SeSidmDeluce = 2
	// SeSidmRaman as defined in src/swephexp.h:241
	SeSidmRaman = 3
	// SeSidmUshashashi as defined in src/swephexp.h:242
	SeSidmUshashashi = 4
	// SeSidmKrishnamurti as defined in src/swephexp.h:243
	SeSidmKrishnamurti = 5
	// SeSidmDjwhalKhul as defined in src/swephexp.h:244
	SeSidmDjwhalKhul = 6
	// SeSidmYukteshwar as defined in src/swephexp.h:245
	SeSidmYukteshwar = 7
	// SeSidmJnBhasin as defined in src/swephexp.h:246
	SeSidmJnBhasin = 8
	// SeSidmBabylKugler1 as defined in src/swephexp.h:247
	SeSidmBabylKugler1 = 9
	// SeSidmBabylKugler2 as defined in src/swephexp.h:248
	SeSidmBabylKugler2 = 10
	// SeSidmBabylKugler3 as defined in src/swephexp.h:249
	SeSidmBabylKugler3 = 11
	// SeSidmBabylHuber as defined in src/swephexp.h:250
	SeSidmBabylHuber = 12
	// SeSidmBabylEtpsc as defined in src/swephexp.h:251
	SeSidmBabylEtpsc = 13
	// SeSidmAldebaran15tau as defined in src/swephexp.h:252
	SeSidmAldebaran15tau = 14
	// SeSidmHipparchos as defined in src/swephexp.h:253
	SeSidmHipparchos = 15
	// SeSidmSassanian as defined in src/swephexp.h:254
	SeSidmSassanian = 16
	// SeSidmGalcent0sag as defined in src/swephexp.h:255
	SeSidmGalcent0sag = 17
	// SeSidmJ2000 as defined in src/swephexp.h:256
	SeSidmJ2000 = 18
	// SeSidmJ1900 as defined in src/swephexp.h:257
	SeSidmJ1900 = 19
	// SeSidmB1950 as defined in src/swephexp.h:258
	SeSidmB1950 = 20
	// SeSidmSuryasiddhanta as defined in src/swephexp.h:259
	SeSidmSuryasiddhanta = 21
	// SeSidmSuryasiddhantaMsun as defined in src/swephexp.h:260
	SeSidmSuryasiddhantaMsun = 22
	// SeSidmAryabhata as defined in src/swephexp.h:261
	SeSidmAryabhata = 23
	// SeSidmAryabhataMsun as defined in src/swephexp.h:262
	SeSidmAryabhataMsun = 24
	// SeSidmSsRevati as defined in src/swephexp.h:263
	SeSidmSsRevati = 25
	// SeSidmSsCitra as defined in src/swephexp.h:264
	SeSidmSsCitra = 26
	// SeSidmTrueCitra as defined in src/swephexp.h:265
	SeSidmTrueCitra = 27
	// SeSidmTrueRevati as defined in src/swephexp.h:266
	SeSidmTrueRevati = 28
	// SeSidmTruePushya as defined in src/swephexp.h:267
	SeSidmTruePushya = 29
	// SeSidmGalcentRgilbrand as defined in src/swephexp.h:268
	SeSidmGalcentRgilbrand = 30
	// SeSidmGalequIau1958 as defined in src/swephexp.h:269
	SeSidmGalequIau1958 = 31
	// SeSidmGalequTrue as defined in src/swephexp.h:270
	SeSidmGalequTrue = 32
	// SeSidmGalequMula as defined in src/swephexp.h:271
	SeSidmGalequMula = 33
	// SeSidmGalalignMardyks as defined in src/swephexp.h:272
	SeSidmGalalignMardyks = 34
	// SeSidmTrueMula as defined in src/swephexp.h:273
	SeSidmTrueMula = 35
	// SeSidmGalcentMulaWilhelm as defined in src/swephexp.h:274
	SeSidmGalcentMulaWilhelm = 36
	// SeSidmAryabhata522 as defined in src/swephexp.h:275
	SeSidmAryabhata522 = 37
	// SeSidmBabylBritton as defined in src/swephexp.h:276
	SeSidmBabylBritton = 38
	// SeSidmTrueSheoran as defined in src/swephexp.h:277
	SeSidmTrueSheoran = 39
	// SeSidmGalcentCochrane as defined in src/swephexp.h:278
	SeSidmGalcentCochrane = 40
	// SeSidmGalequFiorenza as defined in src/swephexp.h:279
	SeSidmGalequFiorenza = 41
	// SeSidmValensMoon as defined in src/swephexp.h:280
	SeSidmValensMoon = 42
	// SeSidmLahiri1940 as defined in src/swephexp.h:281
	SeSidmLahiri1940 = 43
	// SeSidmLahiriVp285 as defined in src/swephexp.h:282
	SeSidmLahiriVp285 = 44
	// SeSidmKrishnamurtiVp291 as defined in src/swephexp.h:283
	SeSidmKrishnamurtiVp291 = 45
	// SeSidmLahiriIcrc as defined in src/swephexp.h:284
	SeSidmLahiriIcrc = 46
	// SeSidmUser as defined in src/swephexp.h:286
	SeSidmUser = 255
	// SeNsidmPredef as defined in src/swephexp.h:288
	SeNsidmPredef = 47
	// SeNodbitMean as defined in src/swephexp.h:291
	SeNodbitMean = 1
	// SeNodbitOscu as defined in src/swephexp.h:292
	SeNodbitOscu = 2
	// SeNodbitOscuBar as defined in src/swephexp.h:293
	SeNodbitOscuBar = 4
	// SeNodbitFopoint as defined in src/swephexp.h:294
	SeNodbitFopoint = 256
	// SeflgDefaulteph as defined in src/swephexp.h:297
	SeflgDefaulteph = 2
	// SeMaxStname as defined in src/swephexp.h:299
	SeMaxStname = 256
	// SeEclCentral as defined in src/swephexp.h:307
	SeEclCentral = 1
	// SeEclNoncentral as defined in src/swephexp.h:308
	SeEclNoncentral = 2
	// SeEclTotal as defined in src/swephexp.h:309
	SeEclTotal = 4
	// SeEclAnnular as defined in src/swephexp.h:310
	SeEclAnnular = 8
	// SeEclPartial as defined in src/swephexp.h:311
	SeEclPartial = 16
	// SeEclAnnularTotal as defined in src/swephexp.h:312
	SeEclAnnularTotal = 32
	// SeEclHybrid as defined in src/swephexp.h:313
	SeEclHybrid = 32
	// SeEclPenumbral as defined in src/swephexp.h:314
	SeEclPenumbral = 64
	// SeEclAlltypesSolar as defined in src/swephexp.h:315
	SeEclAlltypesSolar = 63
	// SeEclAlltypesLunar as defined in src/swephexp.h:316
	SeEclAlltypesLunar = 84
	// SeEclVisible as defined in src/swephexp.h:317
	SeEclVisible = 128
	// SeEclMaxVisible as defined in src/swephexp.h:318
	SeEclMaxVisible = 256
	// SeEcl1stVisible as defined in src/swephexp.h:319
	SeEcl1stVisible = 512
	// SeEclPartbegVisible as defined in src/swephexp.h:320
	SeEclPartbegVisible = 512
	// SeEcl2ndVisible as defined in src/swephexp.h:321
	SeEcl2ndVisible = 1024
	// SeEclTotbegVisible as defined in src/swephexp.h:322
	SeEclTotbegVisible = 1024
	// SeEcl3rdVisible as defined in src/swephexp.h:323
	SeEcl3rdVisible = 2048
	// SeEclTotendVisible as defined in src/swephexp.h:324
	SeEclTotendVisible = 2048
	// SeEcl4thVisible as defined in src/swephexp.h:325
	SeEcl4thVisible = 4096
	// SeEclPartendVisible as defined in src/swephexp.h:326
	SeEclPartendVisible = 4096
	// SeEclPenumbbegVisible as defined in src/swephexp.h:327
	SeEclPenumbbegVisible = 8192
	// SeEclPenumbendVisible as defined in src/swephexp.h:328
	SeEclPenumbendVisible = 16384
	// SeEclOccBegDaylight as defined in src/swephexp.h:329
	SeEclOccBegDaylight = 8192
	// SeEclOccEndDaylight as defined in src/swephexp.h:330
	SeEclOccEndDaylight = 16384
	// SeEclOneTry as defined in src/swephexp.h:331
	SeEclOneTry = 32768
	// SeCalcRise as defined in src/swephexp.h:336
	SeCalcRise = 1
	// SeCalcSet as defined in src/swephexp.h:337
	SeCalcSet = 2
	// SeCalcMtransit as defined in src/swephexp.h:338
	SeCalcMtransit = 4
	// SeCalcItransit as defined in src/swephexp.h:339
	SeCalcItransit = 8
	// SeBitDiscCenter as defined in src/swephexp.h:340
	SeBitDiscCenter = 256
	// SeBitDiscBottom as defined in src/swephexp.h:343
	SeBitDiscBottom = 8192
	// SeBitGeoctrNoEclLat as defined in src/swephexp.h:346
	SeBitGeoctrNoEclLat = 128
	// SeBitNoRefraction as defined in src/swephexp.h:349
	SeBitNoRefraction = 512
	// SeBitCivilTwilight as defined in src/swephexp.h:351
	SeBitCivilTwilight = 1024
	// SeBitNauticTwilight as defined in src/swephexp.h:352
	SeBitNauticTwilight = 2048
	// SeBitAstroTwilight as defined in src/swephexp.h:353
	SeBitAstroTwilight = 4096
	// SeBitFixedDiscSize as defined in src/swephexp.h:354
	SeBitFixedDiscSize = 16384
	// SeBitForceSlowMethod as defined in src/swephexp.h:357
	SeBitForceSlowMethod = 32768
	// SeBitHinduRising as defined in src/swephexp.h:361
	SeBitHinduRising = 896
	// SeEcl2hor as defined in src/swephexp.h:364
	SeEcl2hor = 0
	// SeEqu2hor as defined in src/swephexp.h:365
	SeEqu2hor = 1
	// SeHor2ecl as defined in src/swephexp.h:366
	SeHor2ecl = 0
	// SeHor2equ as defined in src/swephexp.h:367
	SeHor2equ = 1
	// SeTrueToApp as defined in src/swephexp.h:370
	SeTrueToApp = 0
	// SeAppToTrue as defined in src/swephexp.h:371
	SeAppToTrue = 1
	// SeDeNumber as defined in src/swephexp.h:377
	SeDeNumber = 431
	// SeFnameDe200 as defined in src/swephexp.h:378
	SeFnameDe200 = "de200.eph"
	// SeFnameDe403 as defined in src/swephexp.h:379
	SeFnameDe403 = "de403.eph"
	// SeFnameDe404 as defined in src/swephexp.h:380
	SeFnameDe404 = "de404.eph"
	// SeFnameDe405 as defined in src/swephexp.h:381
	SeFnameDe405 = "de405.eph"
	// SeFnameDe406 as defined in src/swephexp.h:382
	SeFnameDe406 = "de406.eph"
	// SeFnameDe431 as defined in src/swephexp.h:383
	SeFnameDe431 = "de431.eph"
	// SeFnameDft as defined in src/swephexp.h:384
	SeFnameDft = "de431.eph"
	// SeFnameDft2 as defined in src/swephexp.h:385
	SeFnameDft2 = "de406.eph"
	// SeStarfileOld as defined in src/swephexp.h:386
	SeStarfileOld = "fixstars.cat"
	// SeStarfile as defined in src/swephexp.h:387
	SeStarfile = "sefstars.txt"
	// SeAstnamfile as defined in src/swephexp.h:388
	SeAstnamfile = "seasnam.txt"
	// SeFictfile as defined in src/swephexp.h:389
	SeFictfile = "seorbel.txt"
	// SeEphePath as defined in src/swephexp.h:406
	SeEphePath = ".:/users/ephe2/:/users/ephe/"
	// SeSplitDegRoundSec as defined in src/swephexp.h:416
	SeSplitDegRoundSec = 1
	// SeSplitDegRoundMin as defined in src/swephexp.h:417
	SeSplitDegRoundMin = 2
	// SeSplitDegRoundDeg as defined in src/swephexp.h:418
	SeSplitDegRoundDeg = 4
	// SeSplitDegZodiacal as defined in src/swephexp.h:419
	SeSplitDegZodiacal = 8
	// SeSplitDegNakshatra as defined in src/swephexp.h:420
	SeSplitDegNakshatra = 1024
	// SeSplitDegKeepSign as defined in src/swephexp.h:421
	SeSplitDegKeepSign = 16
	// SeSplitDegKeepDeg as defined in src/swephexp.h:424
	SeSplitDegKeepDeg = 32
	// SeHeliacalRising as defined in src/swephexp.h:429
	SeHeliacalRising = 1
	// SeHeliacalSetting as defined in src/swephexp.h:430
	SeHeliacalSetting = 2
	// SeMorningFirst as defined in src/swephexp.h:431
	SeMorningFirst = 1
	// SeEveningLast as defined in src/swephexp.h:432
	SeEveningLast = 2
	// SeEveningFirst as defined in src/swephexp.h:433
	SeEveningFirst = 3
	// SeMorningLast as defined in src/swephexp.h:434
	SeMorningLast = 4
	// SeAcronychalRising as defined in src/swephexp.h:435
	SeAcronychalRising = 5
	// SeAcronychalSetting as defined in src/swephexp.h:436
	SeAcronychalSetting = 6
	// SeCosmicalSetting as defined in src/swephexp.h:437
	SeCosmicalSetting = 6
	// SeHelflagLongSearch as defined in src/swephexp.h:439
	SeHelflagLongSearch = 128
	// SeHelflagHighPrecision as defined in src/swephexp.h:440
	SeHelflagHighPrecision = 256
	// SeHelflagOpticalParams as defined in src/swephexp.h:441
	SeHelflagOpticalParams = 512
	// SeHelflagNoDetails as defined in src/swephexp.h:442
	SeHelflagNoDetails = 1024
	// SeHelflagSearch1Period as defined in src/swephexp.h:443
	SeHelflagSearch1Period = 2048
	// SeHelflagVislimDark as defined in src/swephexp.h:444
	SeHelflagVislimDark = 4096
	// SeHelflagVislimNomoon as defined in src/swephexp.h:445
	SeHelflagVislimNomoon = 8192
	// SeHelflagVislimPhotopic as defined in src/swephexp.h:447
	SeHelflagVislimPhotopic = 16384
	// SeHelflagVislimScotopic as defined in src/swephexp.h:448
	SeHelflagVislimScotopic = 32768
	// SeHelflagAv as defined in src/swephexp.h:449
	SeHelflagAv = 65536
	// SeHelflagAvkindVr as defined in src/swephexp.h:450
	SeHelflagAvkindVr = 65536
	// SeHelflagAvkindPto as defined in src/swephexp.h:451
	SeHelflagAvkindPto = 131072
	// SeHelflagAvkindMin7 as defined in src/swephexp.h:452
	SeHelflagAvkindMin7 = 262144
	// SeHelflagAvkindMin9 as defined in src/swephexp.h:453
	SeHelflagAvkindMin9 = 524288
	// SeHelflagAvkind as defined in src/swephexp.h:454
	SeHelflagAvkind = 983040
	// TjdInvalid as defined in src/swephexp.h:455
	TjdInvalid = 9.9999999e+07
	// SimulateVictorvb as defined in src/swephexp.h:456
	SimulateVictorvb = 1
	// SePhotopicFlag as defined in src/swephexp.h:474
	SePhotopicFlag = 0
	// SeScotopicFlag as defined in src/swephexp.h:475
	SeScotopicFlag = 1
	// SeMixedopicFlag as defined in src/swephexp.h:476
	SeMixedopicFlag = 2
	// SeTidalDe200 as defined in src/swephexp.h:483
	SeTidalDe200 = -23.8946
	// SeTidalDe403 as defined in src/swephexp.h:484
	SeTidalDe403 = -25.58
	// SeTidalDe404 as defined in src/swephexp.h:485
	SeTidalDe404 = -25.58
	// SeTidalDe405 as defined in src/swephexp.h:486
	SeTidalDe405 = -25.826
	// SeTidalDe406 as defined in src/swephexp.h:487
	SeTidalDe406 = -25.826
	// SeTidalDe421 as defined in src/swephexp.h:488
	SeTidalDe421 = -25.85
	// SeTidalDe422 as defined in src/swephexp.h:489
	SeTidalDe422 = -25.85
	// SeTidalDe430 as defined in src/swephexp.h:490
	SeTidalDe430 = -25.82
	// SeTidalDe431 as defined in src/swephexp.h:491
	SeTidalDe431 = -25.8
	// SeTidalDe441 as defined in src/swephexp.h:492
	SeTidalDe441 = -25.936
	// SeTidal26 as defined in src/swephexp.h:493
	SeTidal26 = -26
	// SeTidalStephenson2016 as defined in src/swephexp.h:494
	SeTidalStephenson2016 = -25.85
	// SeTidalDefault as defined in src/swephexp.h:495
	SeTidalDefault = -25.8
	// SeTidalAutomatic as defined in src/swephexp.h:496
	SeTidalAutomatic = 999999
	// SeTidalMoseph as defined in src/swephexp.h:497
	SeTidalMoseph = -25.58
	// SeTidalSwieph as defined in src/swephexp.h:498
	SeTidalSwieph = -25.8
	// SeTidalJpleph as defined in src/swephexp.h:499
	SeTidalJpleph = -25.8
	// SeDeltatAutomatic as defined in src/swephexp.h:502
	SeDeltatAutomatic = -1e-10
	// SeModelDeltat as defined in src/swephexp.h:504
	SeModelDeltat = 0
	// SeModelPrecLongterm as defined in src/swephexp.h:505
	SeModelPrecLongterm = 1
	// SeModelPrecShortterm as defined in src/swephexp.h:506
	SeModelPrecShortterm = 2
	// SeModelNut as defined in src/swephexp.h:507
	SeModelNut = 3
	// SeModelBias as defined in src/swephexp.h:508
	SeModelBias = 4
	// SeModelJplhorMode as defined in src/swephexp.h:509
	SeModelJplhorMode = 5
	// SeModelJplhoraMode as defined in src/swephexp.h:510
	SeModelJplhoraMode = 6
	// SeModelSidt as defined in src/swephexp.h:511
	SeModelSidt = 7
	// NseModels as defined in src/swephexp.h:512
	NseModels = 8
	// SemodNprec as defined in src/swephexp.h:515
	SemodNprec = 11
	// SemodPrecIau1976 as defined in src/swephexp.h:516
	SemodPrecIau1976 = 1
	// SemodPrecLaskar1986 as defined in src/swephexp.h:517
	SemodPrecLaskar1986 = 2
	// SemodPrecWillEpsLask as defined in src/swephexp.h:518
	SemodPrecWillEpsLask = 3
	// SemodPrecWilliams1994 as defined in src/swephexp.h:519
	SemodPrecWilliams1994 = 4
	// SemodPrecSimon1994 as defined in src/swephexp.h:520
	SemodPrecSimon1994 = 5
	// SemodPrecIau2000 as defined in src/swephexp.h:521
	SemodPrecIau2000 = 6
	// SemodPrecBretagnon2003 as defined in src/swephexp.h:522
	SemodPrecBretagnon2003 = 7
	// SemodPrecIau2006 as defined in src/swephexp.h:523
	SemodPrecIau2006 = 8
	// SemodPrecVondrak2011 as defined in src/swephexp.h:524
	SemodPrecVondrak2011 = 9
	// SemodPrecOwen1990 as defined in src/swephexp.h:525
	SemodPrecOwen1990 = 10
	// SemodPrecNewcomb as defined in src/swephexp.h:526
	SemodPrecNewcomb = 11
	// SemodPrecDefault as defined in src/swephexp.h:527
	SemodPrecDefault = 9
	// SemodPrecDefaultShort as defined in src/swephexp.h:532
	SemodPrecDefaultShort = 9
	// SemodNnut as defined in src/swephexp.h:535
	SemodNnut = 5
	// SemodNutIau1980 as defined in src/swephexp.h:536
	SemodNutIau1980 = 1
	// SemodNutIauCorr1987 as defined in src/swephexp.h:537
	SemodNutIauCorr1987 = 2
	// SemodNutIau2000a as defined in src/swephexp.h:539
	SemodNutIau2000a = 3
	// SemodNutIau2000b as defined in src/swephexp.h:540
	SemodNutIau2000b = 4
	// SemodNutWoolard as defined in src/swephexp.h:541
	SemodNutWoolard = 5
	// SemodNutDefault as defined in src/swephexp.h:542
	SemodNutDefault = 4
	// SemodNsidt as defined in src/swephexp.h:545
	SemodNsidt = 4
	// SemodSidtIau1976 as defined in src/swephexp.h:546
	SemodSidtIau1976 = 1
	// SemodSidtIau2006 as defined in src/swephexp.h:547
	SemodSidtIau2006 = 2
	// SemodSidtIersConv2010 as defined in src/swephexp.h:548
	SemodSidtIersConv2010 = 3
	// SemodSidtLongterm as defined in src/swephexp.h:549
	SemodSidtLongterm = 4
	// SemodSidtDefault as defined in src/swephexp.h:550
	SemodSidtDefault = 4
	// SemodNbias as defined in src/swephexp.h:554
	SemodNbias = 3
	// SemodBiasNone as defined in src/swephexp.h:555
	SemodBiasNone = 1
	// SemodBiasIau2000 as defined in src/swephexp.h:556
	SemodBiasIau2000 = 2
	// SemodBiasIau2006 as defined in src/swephexp.h:557
	SemodBiasIau2006 = 3
	// SemodBiasDefault as defined in src/swephexp.h:558
	SemodBiasDefault = 3
	// SemodNjplhor as defined in src/swephexp.h:562
	SemodNjplhor = 2
	// SemodJplhorLongAgreement as defined in src/swephexp.h:563
	SemodJplhorLongAgreement = 1
	// SemodJplhorDefault as defined in src/swephexp.h:567
	SemodJplhorDefault = 1
	// SemodNjplhora as defined in src/swephexp.h:582
	SemodNjplhora = 3
	// SemodJplhora1 as defined in src/swephexp.h:583
	SemodJplhora1 = 1
	// SemodJplhora2 as defined in src/swephexp.h:584
	SemodJplhora2 = 2
	// SemodJplhora3 as defined in src/swephexp.h:585
	SemodJplhora3 = 3
	// SemodJplhoraDefault as defined in src/swephexp.h:586
	SemodJplhoraDefault = 3
	// SemodNdeltat as defined in src/swephexp.h:604
	SemodNdeltat = 5
	// SemodDeltatStephensonMorrison1984 as defined in src/swephexp.h:605
	SemodDeltatStephensonMorrison1984 = 1
	// SemodDeltatStephenson1997 as defined in src/swephexp.h:606
	SemodDeltatStephenson1997 = 2
	// SemodDeltatStephensonMorrison2004 as defined in src/swephexp.h:607
	SemodDeltatStephensonMorrison2004 = 3
	// SemodDeltatEspenakMeeus2006 as defined in src/swephexp.h:608
	SemodDeltatEspenakMeeus2006 = 4
	// SemodDeltatStephensonEtc2016 as defined in src/swephexp.h:609
	SemodDeltatStephensonEtc2016 = 5
	// SemodDeltatDefault as defined in src/swephexp.h:611
	SemodDeltatDefault = 5
	// Malloc as defined in src/swephexp.h:633
	Malloc = 0
	// Calloc as defined in src/swephexp.h:634
	Calloc = 0
	// Free as defined in src/swephexp.h:635
	Free = 0
	// CallConv as defined in src/swephexp.h:667

	// Exp32 as defined in src/swephexp.h:668

	// MyTrue as defined in src/sweodef.h:74
	MyTrue = 1
	// MyFalse as defined in src/sweodef.h:75
	MyFalse = 0
	// Tls as defined in src/sweodef.h:84
	Tls = 0
	// Msdos as defined in src/sweodef.h:157
	Msdos = 0
	// Hpunix as defined in src/sweodef.h:158
	Hpunix = 1
	// UnixFs as defined in src/sweodef.h:162
	UnixFs = 1
	// Abs4 as defined in src/sweodef.h:220
	Abs4 = 0
	// True as defined in src/sweodef.h:233
	True = 1
	// False as defined in src/sweodef.h:234
	False = 0
	// Ok as defined in src/sweodef.h:238
	Ok = 0
	// Err as defined in src/sweodef.h:239
	Err = -1
	// Ucp as defined in src/sweodef.h:250
	//Ucp = ( UCHAR * )( )
	// OdegreeString as defined in src/sweodef.h:253
	OdegreeString = "°"
	// AsMaxch as defined in src/sweodef.h:266
	AsMaxch = 256
	// Radtodeg as defined in src/sweodef.h:272
	Radtodeg = 57.29577951308232
	// Degtorad as defined in src/sweodef.h:273
	Degtorad = 0.017453292519943295
	// Cs as defined in src/sweodef.h:276
	Cs = 0
	// Csec as defined in src/sweodef.h:277
	Csec = 0
	// Deg as defined in src/sweodef.h:279
	Deg = 360000
	// Deg730 as defined in src/sweodef.h:280
	Deg730 = 2700000
	// Deg15 as defined in src/sweodef.h:281
	Deg15 = 5400000
	// Deg24 as defined in src/sweodef.h:282
	Deg24 = 8640000
	// Deg30 as defined in src/sweodef.h:283
	Deg30 = 10800000
	// Deg60 as defined in src/sweodef.h:284
	Deg60 = 21600000
	// Deg90 as defined in src/sweodef.h:285
	Deg90 = 32400000
	// Deg120 as defined in src/sweodef.h:286
	Deg120 = 43200000
	// Deg150 as defined in src/sweodef.h:287
	Deg150 = 54000000
	// Deg180 as defined in src/sweodef.h:288
	Deg180 = 64800000
	// Deg270 as defined in src/sweodef.h:289
	Deg270 = 97200000
	// Deg360 as defined in src/sweodef.h:290
	Deg360 = 129600000
	// Cstorad as defined in src/sweodef.h:294
	Cstorad = 4.84813681109536e-08
	// Radtocs as defined in src/sweodef.h:295
	Radtocs = 2.0626480624709636e+07
	// Cs2deg as defined in src/sweodef.h:297
	Cs2deg = 2.777777777777778e-06
	// BfileRAccess as defined in src/sweodef.h:301
	BfileRAccess = "r"
	// BfileRwAccess as defined in src/sweodef.h:302
	BfileRwAccess = "r+"
	// BfileWCreate as defined in src/sweodef.h:303
	BfileWCreate = "w"
	// BfileAAccess as defined in src/sweodef.h:304
	BfileAAccess = "a+"
	// FileRAccess as defined in src/sweodef.h:305
	FileRAccess = "r"
	// FileRwAccess as defined in src/sweodef.h:306
	FileRwAccess = "r+"
	// FileWCreate as defined in src/sweodef.h:307
	FileWCreate = "w"
	// FileAAccess as defined in src/sweodef.h:308
	FileAAccess = "a+"
	// OBinary as defined in src/sweodef.h:309
	OBinary = 0
	// OpenMode as defined in src/sweodef.h:310
	OpenMode = 438
	// DirGlue as defined in src/sweodef.h:311
	DirGlue = "/"
	// PathSeparator as defined in src/sweodef.h:312
	PathSeparator = ";:"

	// SeVersion as defined in src/sweph.h:65
	SeVersion = "2.10.01"
	// J2000 as defined in src/sweph.h:67
	J2000 = 2.451545e+06
	// B1950 as defined in src/sweph.h:68
	B1950 = 2.43328242345905e+06
	// J1900 as defined in src/sweph.h:69
	J1900 = 2.41502e+06
	// B1850 as defined in src/sweph.h:70
	B1850 = 2.396758203581e+06
	// MpcCeres as defined in src/sweph.h:72
	MpcCeres = 1
	// MpcPallas as defined in src/sweph.h:73
	MpcPallas = 2
	// MpcJuno as defined in src/sweph.h:74
	MpcJuno = 3
	// MpcVesta as defined in src/sweph.h:75
	MpcVesta = 4
	// MpcChiron as defined in src/sweph.h:76
	MpcChiron = 2060
	// MpcPholus as defined in src/sweph.h:77
	MpcPholus = 5145
	// SeNameSun as defined in src/sweph.h:79
	SeNameSun = "Sun"
	// SeNameMoon as defined in src/sweph.h:80
	SeNameMoon = "Moon"
	// SeNameMercury as defined in src/sweph.h:81
	SeNameMercury = "Mercury"
	// SeNameVenus as defined in src/sweph.h:82
	SeNameVenus = "Venus"
	// SeNameMars as defined in src/sweph.h:83
	SeNameMars = "Mars"
	// SeNameJupiter as defined in src/sweph.h:84
	SeNameJupiter = "Jupiter"
	// SeNameSaturn as defined in src/sweph.h:85
	SeNameSaturn = "Saturn"
	// SeNameUranus as defined in src/sweph.h:86
	SeNameUranus = "Uranus"
	// SeNameNeptune as defined in src/sweph.h:87
	SeNameNeptune = "Neptune"
	// SeNamePluto as defined in src/sweph.h:88
	SeNamePluto = "Pluto"
	// SeNameMeanNode as defined in src/sweph.h:89
	SeNameMeanNode = "mean Node"
	// SeNameTrueNode as defined in src/sweph.h:90
	SeNameTrueNode = "true Node"
	// SeNameMeanApog as defined in src/sweph.h:91
	SeNameMeanApog = "mean Apogee"
	// SeNameOscuApog as defined in src/sweph.h:92
	SeNameOscuApog = "osc. Apogee"
	// SeNameIntpApog as defined in src/sweph.h:93
	SeNameIntpApog = "intp. Apogee"
	// SeNameIntpPerg as defined in src/sweph.h:94
	SeNameIntpPerg = "intp. Perigee"
	// SeNameEarth as defined in src/sweph.h:95
	SeNameEarth = "Earth"
	// SeNameCeres as defined in src/sweph.h:96
	SeNameCeres = "Ceres"
	// SeNamePallas as defined in src/sweph.h:97
	SeNamePallas = "Pallas"
	// SeNameJuno as defined in src/sweph.h:98
	SeNameJuno = "Juno"
	// SeNameVesta as defined in src/sweph.h:99
	SeNameVesta = "Vesta"
	// SeNameChiron as defined in src/sweph.h:100
	SeNameChiron = "Chiron"
	// SeNamePholus as defined in src/sweph.h:101
	SeNamePholus = "Pholus"
	// SeNameCupido as defined in src/sweph.h:104
	SeNameCupido = "Cupido"
	// SeNameHades as defined in src/sweph.h:105
	SeNameHades = "Hades"
	// SeNameZeus as defined in src/sweph.h:106
	SeNameZeus = "Zeus"
	// SeNameKronos as defined in src/sweph.h:107
	SeNameKronos = "Kronos"
	// SeNameApollon as defined in src/sweph.h:108
	SeNameApollon = "Apollon"
	// SeNameAdmetos as defined in src/sweph.h:109
	SeNameAdmetos = "Admetos"
	// SeNameVulkanus as defined in src/sweph.h:110
	SeNameVulkanus = "Vulkanus"
	// SeNamePoseidon as defined in src/sweph.h:111
	SeNamePoseidon = "Poseidon"
	// SeNameIsis as defined in src/sweph.h:112
	SeNameIsis = "Isis"
	// SeNameNibiru as defined in src/sweph.h:113
	SeNameNibiru = "Nibiru"
	// SeNameHarrington as defined in src/sweph.h:114
	SeNameHarrington = "Harrington"
	// SeNameNeptuneLeverrier as defined in src/sweph.h:115
	SeNameNeptuneLeverrier = "Leverrier"
	// SeNameNeptuneAdams as defined in src/sweph.h:116
	SeNameNeptuneAdams = "Adams"
	// SeNamePlutoLowell as defined in src/sweph.h:117
	SeNamePlutoLowell = "Lowell"
	// SeNamePlutoPickering as defined in src/sweph.h:118
	SeNamePlutoPickering = "Pickering"
	// SeNameVulcan as defined in src/sweph.h:119
	SeNameVulcan = "Vulcan"
	// SeNameWhiteMoon as defined in src/sweph.h:120
	SeNameWhiteMoon = "White Moon"
	// Pi as defined in src/sweph.h:126
	Pi = 3.141592653589793
	// Twopi as defined in src/sweph.h:127
	Twopi = 6.283185307179586
	// Endmark as defined in src/sweph.h:129
	Endmark = -99
	// SeiEpsilon as defined in src/sweph.h:131
	SeiEpsilon = -2
	// SeiNutation as defined in src/sweph.h:132
	SeiNutation = -1
	// SeiEmb as defined in src/sweph.h:133
	SeiEmb = 0
	// SeiEarth as defined in src/sweph.h:134
	SeiEarth = 0
	// SeiSun as defined in src/sweph.h:135
	SeiSun = 0
	// SeiMoon as defined in src/sweph.h:136
	SeiMoon = 1
	// SeiMercury as defined in src/sweph.h:137
	SeiMercury = 2
	// SeiVenus as defined in src/sweph.h:138
	SeiVenus = 3
	// SeiMars as defined in src/sweph.h:139
	SeiMars = 4
	// SeiJupiter as defined in src/sweph.h:140
	SeiJupiter = 5
	// SeiSaturn as defined in src/sweph.h:141
	SeiSaturn = 6
	// SeiUranus as defined in src/sweph.h:142
	SeiUranus = 7
	// SeiNeptune as defined in src/sweph.h:143
	SeiNeptune = 8
	// SeiPluto as defined in src/sweph.h:144
	SeiPluto = 9
	// SeiSunbary as defined in src/sweph.h:145
	SeiSunbary = 10
	// SeiAnybody as defined in src/sweph.h:146
	SeiAnybody = 11
	// SeiChiron as defined in src/sweph.h:147
	SeiChiron = 12
	// SeiPholus as defined in src/sweph.h:148
	SeiPholus = 13
	// SeiCeres as defined in src/sweph.h:149
	SeiCeres = 14
	// SeiPallas as defined in src/sweph.h:150
	SeiPallas = 15
	// SeiJuno as defined in src/sweph.h:151
	SeiJuno = 16
	// SeiVesta as defined in src/sweph.h:152
	SeiVesta = 17
	// SeiNplanets as defined in src/sweph.h:154
	SeiNplanets = 18
	// SeiMeanNode as defined in src/sweph.h:156
	SeiMeanNode = 0
	// SeiTrueNode as defined in src/sweph.h:157
	SeiTrueNode = 1
	// SeiMeanApog as defined in src/sweph.h:158
	SeiMeanApog = 2
	// SeiOscuApog as defined in src/sweph.h:159
	SeiOscuApog = 3
	// SeiIntpApog as defined in src/sweph.h:160
	SeiIntpApog = 4
	// SeiIntpPerg as defined in src/sweph.h:161
	SeiIntpPerg = 5
	// SeiNnodeEtc as defined in src/sweph.h:163
	SeiNnodeEtc = 6
	// SeiFlgHelio as defined in src/sweph.h:165
	SeiFlgHelio = 1
	// SeiFlgRotate as defined in src/sweph.h:166
	SeiFlgRotate = 2
	// SeiFlgEllipse as defined in src/sweph.h:167
	SeiFlgEllipse = 4
	// SeiFlgEmbhel as defined in src/sweph.h:168
	SeiFlgEmbhel = 8
	// SeiFilePlanet as defined in src/sweph.h:173
	SeiFilePlanet = 0
	// SeiFileMoon as defined in src/sweph.h:174
	SeiFileMoon = 1
	// SeiFileMainAst as defined in src/sweph.h:175
	SeiFileMainAst = 2
	// SeiFileAnyAst as defined in src/sweph.h:176
	SeiFileAnyAst = 3
	// SeiFileFixstar as defined in src/sweph.h:177
	SeiFileFixstar = 4
	// SeiFilePlmoon as defined in src/sweph.h:178
	SeiFilePlmoon = 5
	// SeiFileTestEndian as defined in src/sweph.h:183
	SeiFileTestEndian = 6382179
	// SeiFileBigendian as defined in src/sweph.h:184
	SeiFileBigendian = 0
	// SeiFileNoreord as defined in src/sweph.h:185
	SeiFileNoreord = 0
	// SeiFileLitendian as defined in src/sweph.h:186
	SeiFileLitendian = 1
	// SeiFileReord as defined in src/sweph.h:187
	SeiFileReord = 2
	// SeiFileNmaxplan as defined in src/sweph.h:189
	SeiFileNmaxplan = 50
	// SeiFileEfposbegin as defined in src/sweph.h:190
	SeiFileEfposbegin = 500
	// SeFileSuffix as defined in src/sweph.h:192
	SeFileSuffix = "se1"
	// SeiNephfiles as defined in src/sweph.h:194
	SeiNephfiles = 7
	// SeiCurrFpos as defined in src/sweph.h:195
	SeiCurrFpos = -1
	// SeiNmodels as defined in src/sweph.h:196
	SeiNmodels = 8
	// SeiEclGeoaltMax as defined in src/sweph.h:198
	SeiEclGeoaltMax = 25000
	// SeiEclGeoaltMin as defined in src/sweph.h:199
	SeiEclGeoaltMin = -500
	// ChironStart as defined in src/sweph.h:207
	ChironStart = 1.9676015e+06
	// ChironEnd as defined in src/sweph.h:208
	ChironEnd = 3.4194375e+06
	// PholusStart as defined in src/sweph.h:216
	PholusStart = 640648.5
	// PholusEnd as defined in src/sweph.h:217
	PholusEnd = 4.3906175e+06
	// MoshplephStart as defined in src/sweph.h:219
	MoshplephStart = 625000.5
	// MoshplephEnd as defined in src/sweph.h:220
	MoshplephEnd = 2.8180005e+06
	// MoshluephStart as defined in src/sweph.h:221
	MoshluephStart = 625000.5
	// MoshluephEnd as defined in src/sweph.h:222
	MoshluephEnd = 2.8180005e+06
	// MoshndephStart as defined in src/sweph.h:225
	MoshndephStart = -3.1000155e+06
	// MoshndephEnd as defined in src/sweph.h:226
	MoshndephEnd = 8.0000165e+06
	// JplDe431Start as defined in src/sweph.h:233
	JplDe431Start = -3.0272155e+06
	// JplDe431End as defined in src/sweph.h:234
	JplDe431End = 7.9301925e+06
	// Maxord as defined in src/sweph.h:247
	Maxord = 40
	// Ncties as defined in src/sweph.h:249
	Ncties = 6
	// NotAvailable as defined in src/sweph.h:253
	NotAvailable = -2
	// BeyondEphLimits as defined in src/sweph.h:254
	BeyondEphLimits = -3
	// JToJ2000 as defined in src/sweph.h:256
	JToJ2000 = 1
	// J2000ToJ as defined in src/sweph.h:257
	J2000ToJ = -1
	// MoonMeanDist as defined in src/sweph.h:260
	MoonMeanDist = 3.844e+08
	// MoonMeanIncl as defined in src/sweph.h:261
	MoonMeanIncl = 5.1453964
	// MoonMeanEcc as defined in src/sweph.h:262
	MoonMeanEcc = 0.054900489
	// SunEarthMrat as defined in src/sweph.h:264
	SunEarthMrat = 332946.050895
	// EarthMoonMrat as defined in src/sweph.h:265
	EarthMoonMrat = 81.30055985272827
	// Aunit as defined in src/sweph.h:273
	Aunit = 1.495978707e+11
	// Clight as defined in src/sweph.h:274
	Clight = 2.99792458e+08
	// Helgravconst as defined in src/sweph.h:278
	Helgravconst = 1.32712440017987e+20
	// Geogconst as defined in src/sweph.h:279
	Geogconst = 3.98600448e+14
	// Kgauss as defined in src/sweph.h:280
	Kgauss = 0.01720209895
	// SunRadius as defined in src/sweph.h:281
	SunRadius = 0.00465241752803144
	// EarthRadius as defined in src/sweph.h:282
	EarthRadius = 6.3781366e+06
	// EarthOblateness as defined in src/sweph.h:284
	EarthOblateness = 0.003352819697896193
	// EarthRotSpeed as defined in src/sweph.h:285
	EarthRotSpeed = 6.300387486748799
	// LighttimeAunit as defined in src/sweph.h:287
	LighttimeAunit = 0.005775518331437499
	// ParsecToAunit as defined in src/sweph.h:288
	ParsecToAunit = 206264.8062471
	// SsyPlaneNodeE2000 as defined in src/sweph.h:291
	SsyPlaneNodeE2000 = 1.8776700468039835
	// SsyPlaneNode as defined in src/sweph.h:293
	SsyPlaneNode = 1.8777793895872257
	// SsyPlaneIncl as defined in src/sweph.h:295
	SsyPlaneIncl = 0.027553530354526998
	// KmSToAuCty as defined in src/sweph.h:297
	KmSToAuCty = 21.095
	// MoonSpeedIntv as defined in src/sweph.h:298
	MoonSpeedIntv = 5e-05
	// PlanSpeedIntv as defined in src/sweph.h:299
	PlanSpeedIntv = 0.0001
	// MeanNodeSpeedIntv as defined in src/sweph.h:300
	MeanNodeSpeedIntv = 0.001
	// NodeCalcIntv as defined in src/sweph.h:301
	NodeCalcIntv = 0.0001
	// NodeCalcIntvMosh as defined in src/sweph.h:302
	NodeCalcIntvMosh = 0.1
	// NutSpeedIntv as defined in src/sweph.h:303
	NutSpeedIntv = 0.0001
	// DeflSpeedIntv as defined in src/sweph.h:304
	DeflSpeedIntv = 5e-07
	// SeLapseRate as defined in src/sweph.h:306
	SeLapseRate = 0.0065
	// Ndiam as defined in src/sweph.h:314
	Ndiam = 21
	// Str as defined in src/sweph.h:666
	Str = 4.84813681109536e-06
	// SwiStarLength as defined in src/sweph.h:775
	SwiStarLength = 40
	// SweDataDpsiDeps as defined in src/sweph.h:785
	SweDataDpsiDeps = 36525
)
