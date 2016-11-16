package kmeans

var assignmentData = [][]float64{
	[]float64{-112.0707922, 33.4516246},
	[]float64{-112.0655423, 33.4492979},
	[]float64{-112.0739312, 33.4564905},
	[]float64{-112.0748658, 33.4701155},
	[]float64{-80.5256905, 43.4770992},
	[]float64{-80.5266413, 43.4858748},
	[]float64{-80.846495, 35.225825},
	[]float64{-112.0744277, 33.4484911},
	[]float64{-112.0731006, 33.4299067},
	[]float64{-80.8396357, 35.2270542},
	[]float64{-112.0738147, 33.4487514},
	[]float64{-80.84395, 35.225394},
	[]float64{-80.5361832, 43.4722801},
	[]float64{-112.0742828, 33.4787931},
	[]float64{-80.8404318988, 35.2274789984},
	[]float64{-80.8419989013, 35.2283137726},
	[]float64{-80.846935, 35.22586},
	[]float64{-112.074115, 33.475583},
	[]float64{-80.968306, 35.283424},
	[]float64{-80.5408376, 43.484427},
	[]float64{-80.527637205, 43.5026116046},
	[]float64{-80.5231354, 43.4679471},
	[]float64{-80.4937878, 43.4894705},
	[]float64{-80.8427865, 35.2207294},
	[]float64{-112.0686202, 33.4521542},
	[]float64{-80.5256905, 43.4770992},
	[]float64{-112.0696144, 33.4523351},
	[]float64{-80.5211216, 43.4793516},
	[]float64{-112.067518, 33.4660304},
	[]float64{-112.074326, 33.46084},
	[]float64{-80.5293258, 43.49749},
	[]float64{-80.843155, 35.218214},
	[]float64{-112.0686202, 33.4521542},
	[]float64{-112.0686777, 33.4579128},
	[]float64{-112.0657878, 33.4768722},
	[]float64{-80.8440182, 35.2275943},
	[]float64{-80.5203846, 43.4787316},
	[]float64{-112.0686202, 33.4521542},
	[]float64{-80.5256905, 43.4770992},
	[]float64{-80.8530782034, 35.2260468097},
	[]float64{-80.840783, 35.23361},
	[]float64{-80.8502645514, 35.2336634617},
	[]float64{-80.849647, 35.223622},
	[]float64{-80.8343082, 35.2262527},
	[]float64{-80.8411478, 35.227822},
	[]float64{-80.5377496406, 43.4723280992},
	[]float64{-112.0742162, 33.4796505},
	[]float64{-80.8500343, 35.2341598},
	[]float64{-112.0743321, 33.4580837},
	[]float64{-112.0721778, 33.4663026},
	[]float64{-80.5231407, 43.4768135},
	[]float64{-80.844616, 35.223805},
	[]float64{-112.065825, 33.442472},
	[]float64{-80.8670549, 35.089035},
	[]float64{-80.5238819, 43.4648096},
	[]float64{-80.84468, 35.226433},
	[]float64{-112.0655975, 33.4653345},
	[]float64{-80.7791381, 35.0952753},
	[]float64{-80.5392510355, 43.4721726021},
	[]float64{-80.843388, 35.2279912},
	[]float64{-112.072117722, 33.4448019849},
	[]float64{-80.5243892, 43.4752375},
	[]float64{-80.8414916, 35.2278131},
	[]float64{-112.0657286, 33.4657117},
	[]float64{-80.5230849, 43.4677521},
	[]float64{-80.846592, 35.228193},
	[]float64{-80.8456832, 35.2223511},
	[]float64{-112.0691069, 33.4458023},
	[]float64{-80.5227058, 43.4670204},
	[]float64{-80.843388, 35.2279912},
	[]float64{-80.5379276723, 43.4722519425},
	[]float64{-112.0717019, 33.4483992},
	[]float64{-80.8547192, 35.2283983},
	[]float64{-80.537082, 43.471891},
	[]float64{-112.0722009, 33.4724737},
	[]float64{-112.0750849, 33.4685474},
	[]float64{-112.073929, 33.429424},
	[]float64{-80.8464246, 35.2246161},
	[]float64{-80.5353761, 43.4723904},
	[]float64{-112.0657967, 33.4365057},
	[]float64{-112.0742794, 33.4722371},
	[]float64{-80.8419492, 35.2262406},
	[]float64{-112.0613638, 33.4660928},
	[]float64{-80.5232652, 43.4692451},
	[]float64{-112.0740375, 33.4501458},
	[]float64{-112.0255034, 33.5303271},
	[]float64{-80.5223045, 43.4646939},
	[]float64{-112.071886, 33.449858},
	[]float64{-80.5219541118, 43.4628223628},
	[]float64{-112.0744277, 33.4484911},
	[]float64{-80.536591, 43.472091},
	[]float64{-80.5377811566, 43.4727816313},
	[]float64{-112.0699472, 33.4506843},
	[]float64{-80.5244182, 43.486349},
	[]float64{-80.5246119, 43.4646469},
	[]float64{-80.5409781, 43.5071296},
	[]float64{-80.8473278, 35.2239303},
	[]float64{-80.8394745, 35.2278411},
	[]float64{-112.0996404, 33.5176328},
	[]float64{-112.0705878, 33.4500609},
	[]float64{-112.0706266, 33.4450359},
	[]float64{-80.519187, 43.4772196},
	[]float64{-112.0678078, 33.4585108},
	[]float64{-112.0740511, 33.4499524},
	[]float64{-80.847333, 35.234098},
	[]float64{-80.847, 35.217205},
	[]float64{-112.070122, 33.443909},
	[]float64{-80.5398649, 43.4691285},
	[]float64{-112.0772, 33.449188},
	[]float64{-112.070655, 33.4555253},
	[]float64{-80.5273068, 43.463366},
	[]float64{-80.8426239, 35.2284099},
	[]float64{-112.0740962, 33.4794824},
	[]float64{-80.5371146277, 43.4725006073},
	[]float64{-112.072706876, 33.4483756752},
	[]float64{-112.0686337, 33.4516878},
	[]float64{-80.5182138, 43.483281},
	[]float64{-112.0656044, 33.4372015},
	[]float64{-80.5194703, 43.476998},
	[]float64{-80.528095, 43.474541},
	[]float64{-80.8421891375, 35.2247714304},
	[]float64{-80.8414916, 35.2278131},
	[]float64{-112.0736968, 33.476644},
	[]float64{-80.5137443, 43.471126},
	[]float64{-80.8480944, 35.056435},
	[]float64{-80.5233044, 43.4679985},
	[]float64{-80.84581, 35.214505},
	[]float64{-80.5381164327, 43.4721891678},
	[]float64{-80.8393449, 35.2216474},
	[]float64{-80.5235285, 43.4636098},
	[]float64{-80.838659, 35.228234},
	[]float64{-80.5545309, 43.5093982},
	[]float64{-80.851868, 35.233064},
	[]float64{-112.0737923, 33.4568607},
	[]float64{-112.0741097, 33.4505535},
	[]float64{-80.5183755, 43.4782831},
	[]float64{-80.52628, 43.484483},
	[]float64{-112.0686202, 33.4521542},
	[]float64{-112.0737003, 33.4757646},
	[]float64{-80.5178879, 43.4758781},
	[]float64{-80.84395, 35.225394},
	[]float64{-80.8449462, 35.225705},
	[]float64{-112.0709775, 33.4470733},
	[]float64{-80.5221597, 43.4642893},
	[]float64{-80.8464728, 35.2240131},
	[]float64{-80.84104, 35.2244979},
	[]float64{-80.8411478, 35.227822},
	[]float64{-112.073569, 33.4640982988},
	[]float64{-80.841786, 35.2340477},
	[]float64{-80.8174063386, 35.0578530619},
	[]float64{-80.5193889, 43.4596341},
	[]float64{-80.8439988, 35.2329575},
	[]float64{-80.809002, 35.220697},
	[]float64{-80.523446, 43.4667473},
	[]float64{-80.5365122, 43.4720116},
	[]float64{-80.5238087, 43.466708},
	[]float64{-80.5250886, 43.477593},
	[]float64{-80.5228895, 43.467046},
	[]float64{-80.857545, 35.117183},
	[]float64{-80.783798, 35.096526},
	[]float64{-80.527828, 43.489061},
	[]float64{-112.0696051, 33.464388},
	[]float64{-80.5351869, 43.5030738},
	[]float64{-112.0709775, 33.4470733},
	[]float64{-80.843388, 35.2279912},
	[]float64{-80.53118, 43.502488},
	[]float64{-80.5213927, 43.4659522},
	[]float64{-80.521901, 43.4647024},
	[]float64{-80.5249727, 43.4845752},
	[]float64{-80.5243842, 43.5048639},
	[]float64{-80.5373674259, 43.4730259144},
	[]float64{-112.070585498, 33.4507074659},
	[]float64{-112.074248, 33.47626},
	[]float64{-80.922732, 35.237471},
	[]float64{-80.849538, 35.227078},
	[]float64{-112.074802052, 33.448424817},
	[]float64{-80.5364492, 43.4720225},
	[]float64{-80.8428142, 35.2265794},
	[]float64{-80.845804, 35.2263469},
	[]float64{-80.5223193, 43.4648949},
	[]float64{-112.06909, 33.46544},
	[]float64{-80.555717, 43.452243},
	[]float64{-80.8319153, 35.1879068},
	[]float64{-80.8530782034, 35.2260468097},
	[]float64{-80.5391085148, 43.4719665363},
	[]float64{-80.8401304, 35.2273178},
	[]float64{-80.843388, 35.2279912},
	[]float64{-80.5243892, 43.4752375},
	[]float64{-112.073577, 33.44455},
	[]float64{-80.5260902, 43.483456},
	[]float64{-80.8461559, 35.22533},
	[]float64{-80.844191, 35.226995},
	[]float64{-112.0736854, 33.4793911},
	[]float64{-112.071322, 33.449909},
	[]float64{-112.073417, 33.477939},
	[]float64{-80.845006058, 35.2286502258},
	[]float64{-112.06649, 33.4667869},
	[]float64{-112.0741767, 33.4558849},
	[]float64{-80.8411478, 35.227822},
	[]float64{-112.081565, 33.465499},
	[]float64{-80.5252447306, 43.4763190134},
	[]float64{-112.0739284, 33.4739691},
	[]float64{-80.851336, 35.238088},
	[]float64{-112.074153, 33.46425},
	[]float64{-80.538629, 43.472371},
	[]float64{-80.9405309, 35.239674},
	[]float64{-112.072914, 33.477594},
	[]float64{-80.531231, 43.5028859},
	[]float64{-80.84119, 35.228915},
	[]float64{-80.728384, 35.2771265},
	[]float64{-80.537061, 43.472566},
	[]float64{-80.839671, 35.228985},
	[]float64{-80.8473, 35.230158},
	[]float64{-112.0686202, 33.4521542},
	[]float64{-80.5221353, 43.4657569},
	[]float64{-80.5211055, 43.4793595},
	[]float64{-112.0736809, 33.4794762},
	[]float64{-80.5388363109, 43.4717869083},
	[]float64{-112.0741467, 33.4633319},
	[]float64{-80.5210309, 43.462251},
	[]float64{-112.072092, 33.47066},
	[]float64{-80.5217836, 43.4629195},
	[]float64{-80.5222866, 43.4610385},
	[]float64{-80.536591, 43.471998},
	[]float64{-80.8394131, 35.2302318},
	[]float64{-80.844683, 35.223866},
	[]float64{-80.8416248, 35.2286559},
	[]float64{-112.0741767, 33.4558849},
	[]float64{-80.828252, 35.1467079},
	[]float64{-112.073106, 33.473847},
	[]float64{-112.0686202, 33.4521542},
	[]float64{-80.842303, 35.226746},
	[]float64{-80.84547, 35.224456},
	[]float64{-80.536612, 43.472091},
	[]float64{-80.5318027, 43.4991122},
	[]float64{-80.843784, 35.2275289},
	[]float64{-112.0712574, 33.4484654},
	[]float64{-80.8461559, 35.22533},
	[]float64{-112.071075886, 33.4468850573},
	[]float64{-80.8412724, 35.2288472},
	[]float64{-112.0664327, 33.476323},
	[]float64{-112.0744277, 33.4484911},
	[]float64{-80.8381685, 35.2310911},
	[]float64{-80.8464246, 35.2246161},
	[]float64{-80.8424708, 35.2292799},
	[]float64{-80.8393335, 35.2250606},
	[]float64{-80.840985, 35.228281},
	[]float64{-80.8396379, 35.231684},
	[]float64{-80.834638, 35.228506},
	[]float64{-112.0740373, 33.4483771},
	[]float64{-80.8464728, 35.2240131},
	[]float64{-80.8388937, 35.2274621},
	[]float64{-80.5238819, 43.4648096},
	[]float64{-80.53827703, 43.4725940388},
	[]float64{-80.5365704, 43.471988},
	[]float64{-80.5250037834, 43.4774177316},
	[]float64{-80.8489878, 35.2335751},
	[]float64{-112.0674357, 33.4503971},
	[]float64{-80.8407079, 35.2172053},
	[]float64{-112.0744197, 33.4771144},
	[]float64{-80.5353761, 43.4723904},
	[]float64{-112.0865171, 33.4547647},
	[]float64{-80.8406527, 35.2274197},
	[]float64{-80.5274971, 43.4980824},
	[]float64{-80.5408262, 43.4844189},
	[]float64{-112.0732082, 33.5087398},
	[]float64{-80.5184485, 43.4657108},
	[]float64{-112.07194, 33.4493093},
	[]float64{-80.5228895, 43.467046},
	[]float64{-112.074558593, 33.4492572677},
	[]float64{-80.529250574, 43.5037624514},
	[]float64{-80.8361303, 35.2327235},
	[]float64{-112.0743321, 33.4580837},
	[]float64{-80.8452366, 35.2254586},
	[]float64{-80.5253004, 43.4763761},
	[]float64{-80.5224956, 43.4697612},
	[]float64{-112.073481, 33.439268},
	[]float64{-112.066333, 33.450875},
	[]float64{-80.5150582, 43.4703532},
	[]float64{-80.5229421, 43.4667083},
	[]float64{-80.84104, 35.2244979},
	[]float64{-80.84972, 35.21873},
	[]float64{-80.8414487, 35.227987},
	[]float64{-80.851151, 35.220527},
	[]float64{-80.5223735, 43.4639601},
	[]float64{-80.5247756, 43.4770643},
	[]float64{-80.8473977, 35.2215512},
	[]float64{-80.5551852, 43.5064923},
	[]float64{-80.5361832, 43.4722801},
	[]float64{-112.0736806, 33.4794817},
	[]float64{-80.5393988, 43.5011946},
	[]float64{-112.0789482, 33.4800317},
	[]float64{-80.864092, 35.093322},
	[]float64{-112.0705438, 33.4657432},
	[]float64{-80.5178155, 43.4755885},
	[]float64{-112.0732961, 33.479418},
	[]float64{-112.0669197, 33.4670273},
	[]float64{-112.0733849, 33.4691037},
	[]float64{-80.5340476, 43.4728924},
	[]float64{-112.0693184, 33.4657177},
}