package main

import (
	"fmt"
	"strconv"
	"strings"
)

var input = "8417\n8501\n5429\n2112\n6482\n7971\n9636\n4003\n\n4697\n2941\n3275\n6060\n4879\n7158\n5066\n3196\n5780\n3143\n2510\n7073\n\n3820\n6710\n4781\n6241\n3975\n6308\n5289\n6575\n2662\n4804\n5352\n4610\n\n18980\n3731\n16643\n\n6500\n2810\n13252\n5288\n2110\n12393\n\n5354\n5521\n6015\n2501\n6067\n2465\n2201\n3864\n2661\n4078\n2399\n3212\n2831\n3268\n2726\n\n4701\n4530\n1699\n3015\n3682\n2537\n4554\n2826\n1895\n2652\n4074\n\n3650\n2113\n4750\n6639\n5177\n3823\n1529\n5671\n3679\n7061\n\n1984\n3240\n7139\n1931\n6740\n5695\n6572\n3264\n2904\n7199\n\n4982\n3483\n3358\n2762\n3983\n1341\n7420\n5830\n5709\n6140\n5363\n2636\n\n8943\n9452\n14560\n11685\n9792\n\n51549\n\n4032\n3303\n6291\n2010\n1099\n4304\n3308\n3226\n7038\n5911\n\n3627\n4414\n3217\n3388\n5294\n1914\n4938\n1120\n3825\n4163\n5549\n1712\n2589\n1550\n\n3683\n5829\n3937\n6935\n8588\n2765\n9142\n6662\n\n2440\n1029\n5550\n4674\n3176\n1791\n3867\n2779\n5259\n4676\n5868\n2371\n4557\n1058\n\n5431\n7001\n1789\n7292\n2550\n6087\n6538\n3913\n6977\n2262\n5710\n2555\n\n10586\n1560\n9521\n6677\n7222\n7567\n9516\n\n2974\n4438\n5184\n1797\n3472\n3685\n4252\n3512\n6003\n3274\n6482\n5545\n6207\n3442\n\n2542\n2128\n3556\n3399\n2355\n2355\n2887\n5069\n6026\n2508\n1716\n1556\n2295\n3955\n\n25725\n15128\n7375\n\n14056\n1829\n18860\n\n55324\n\n6072\n6692\n4264\n6261\n1588\n1143\n7373\n6603\n5488\n3320\n2545\n5954\n\n18062\n29180\n\n3677\n6522\n4883\n7097\n5284\n2549\n\n48546\n\n3552\n4460\n1059\n2662\n4123\n4032\n4430\n3578\n3441\n3895\n3487\n1717\n5564\n2195\n5687\n\n3760\n3157\n2112\n5937\n5018\n4986\n1568\n7508\n2288\n4252\n4226\n\n8943\n35829\n\n3615\n4087\n16628\n9725\n\n18133\n3552\n19658\n8840\n\n5401\n2506\n2988\n4760\n5149\n2556\n1375\n2458\n1349\n6012\n3101\n1719\n4705\n\n1915\n4351\n4834\n6064\n1864\n6366\n1081\n4066\n4200\n6982\n1148\n6685\n\n6497\n19940\n5418\n19755\n\n6763\n4697\n4844\n6105\n7080\n7110\n3403\n1635\n2373\n4737\n1538\n4517\n\n55242\n\n5390\n4892\n9004\n6322\n2309\n\n14406\n2408\n3189\n13705\n15881\n\n8001\n4336\n7789\n3403\n7771\n5339\n3603\n2882\n2789\n5457\n\n5291\n1004\n4796\n5162\n3295\n5628\n6048\n8055\n2805\n2979\n6347\n\n4342\n2131\n3900\n9516\n9941\n10085\n9855\n4946\n\n3766\n12591\n6027\n7521\n2532\n2965\n\n6640\n8130\n2235\n3616\n1513\n4982\n5757\n6450\n7833\n2535\n\n11114\n9350\n5189\n7416\n2510\n12480\n\n2161\n1329\n6055\n5304\n1213\n5987\n1021\n1484\n2736\n1594\n3897\n1298\n3701\n3500\n5264\n\n1393\n8359\n7732\n1479\n1367\n9665\n4956\n8989\n2423\n\n14947\n26722\n\n4382\n1257\n4970\n3681\n4299\n2421\n3996\n4978\n3178\n5039\n3989\n5297\n1322\n1941\n\n2901\n5088\n5876\n5770\n4000\n5408\n5932\n3787\n4428\n3649\n3000\n4528\n3172\n5852\n4803\n\n3841\n6352\n7104\n3285\n8693\n7082\n8890\n10598\n\n32634\n\n5064\n4469\n2675\n5662\n1483\n5918\n3117\n4976\n3964\n2474\n5210\n1632\n4618\n2106\n3214\n\n1940\n5559\n2138\n6436\n1860\n5369\n5147\n2794\n1714\n3409\n1892\n5806\n3899\n2305\n\n4500\n2147\n4251\n4609\n3619\n2567\n5201\n3390\n1319\n4424\n3005\n1397\n1693\n4714\n\n2415\n4424\n3809\n3615\n3269\n1986\n2018\n4848\n5122\n1968\n3075\n2470\n5922\n1937\n4125\n\n6317\n6637\n4533\n5743\n4806\n3129\n1306\n3957\n4663\n7623\n\n14487\n\n34032\n11976\n\n5839\n5692\n6716\n1473\n1572\n6622\n7735\n2335\n5670\n6461\n2705\n\n5643\n1721\n3150\n6460\n5134\n1124\n3925\n1463\n1951\n6910\n6680\n3038\n2685\n\n3205\n2335\n4418\n2358\n5683\n2334\n1910\n4780\n1315\n1613\n1523\n2882\n\n3266\n12924\n8477\n7172\n13876\n9290\n\n18737\n7661\n5950\n\n61407\n\n24968\n23181\n\n4341\n5288\n9306\n10062\n9350\n4045\n\n4607\n2565\n2507\n4183\n3101\n5356\n3384\n4707\n2040\n5622\n2513\n2429\n2496\n1069\n\n3590\n5290\n5371\n6878\n6657\n8246\n4479\n4906\n3830\n\n7894\n7018\n5222\n7556\n7754\n3433\n4862\n9010\n6891\n\n6337\n21986\n2791\n\n1075\n5278\n3373\n3820\n1277\n3217\n1349\n1758\n2934\n1298\n5885\n1544\n4560\n4384\n4804\n\n6345\n1952\n1082\n5949\n3946\n5991\n6474\n6288\n4526\n5059\n6081\n1155\n1388\n5735\n\n1166\n6023\n6465\n2215\n7056\n7369\n6983\n1051\n6578\n6131\n6334\n6491\n\n4486\n5666\n5937\n5460\n2587\n1803\n3938\n1771\n4157\n5144\n2624\n4098\n5403\n2063\n2746\n\n6953\n11646\n2940\n6280\n4667\n1413\n\n5913\n2811\n1258\n5507\n6685\n1464\n3137\n6716\n4217\n3656\n6306\n4198\n4557\n\n5846\n2427\n3635\n6365\n5987\n5964\n4458\n4119\n4522\n3972\n6216\n5551\n1705\n\n2125\n5680\n1450\n1701\n5256\n4673\n5045\n5515\n1432\n3565\n3246\n6050\n5755\n3390\n5847\n\n1655\n5405\n6132\n3714\n1627\n7477\n10344\n6058\n\n4914\n4975\n4158\n2466\n5685\n2484\n3594\n1037\n3369\n1022\n3917\n3185\n5233\n5571\n5222\n\n6996\n4804\n1923\n3688\n4808\n1448\n2836\n5346\n5266\n4603\n2173\n1082\n\n5730\n9574\n5903\n6668\n5482\n5088\n8905\n5333\n7135\n\n5377\n2192\n6250\n1581\n5354\n3088\n1125\n2676\n3780\n3977\n4634\n3810\n4168\n\n3173\n4961\n1384\n5868\n4666\n2902\n7158\n8275\n5856\n6056\n\n7098\n12466\n4255\n10181\n13726\n4866\n\n1681\n1910\n5175\n4773\n7702\n5183\n5759\n2788\n2596\n5676\n5361\n\n6560\n5352\n5182\n13627\n7612\n10355\n\n4780\n1702\n5111\n1101\n6959\n4120\n5441\n4818\n8091\n6480\n\n8332\n5999\n6744\n9027\n7948\n3470\n9144\n6675\n2996\n\n6961\n19157\n1485\n19127\n\n22926\n\n17211\n14166\n1463\n\n5074\n9836\n4387\n10138\n10072\n5791\n3683\n9806\n\n5623\n3668\n4324\n9990\n4865\n1793\n10752\n\n9739\n18699\n14526\n4157\n\n2785\n1804\n4527\n4430\n3373\n10323\n6855\n\n2075\n4828\n2026\n3909\n6410\n4035\n3709\n5177\n3259\n6033\n5534\n6097\n2578\n\n2246\n6146\n5899\n1864\n2881\n4628\n3176\n5663\n3088\n3834\n2935\n\n16262\n11795\n14098\n15140\n13228\n\n4323\n8674\n5497\n8305\n6261\n7846\n2593\n10156\n\n5309\n3112\n4563\n1434\n1107\n7130\n5727\n4950\n7466\n6218\n3641\n\n6379\n1200\n2050\n6235\n4257\n2532\n3132\n6179\n4883\n6282\n6576\n\n11097\n12482\n7593\n3300\n3342\n2184\n\n1950\n2753\n4277\n8130\n5451\n7991\n4993\n7978\n8004\n\n20042\n19430\n5442\n\n7059\n6666\n5832\n6485\n2405\n2744\n4710\n1726\n1121\n\n2277\n1842\n1575\n1955\n1906\n4176\n4114\n5040\n5552\n4857\n2775\n5033\n3488\n3971\n3550\n\n11844\n9703\n6549\n7779\n9117\n1536\n\n8861\n25738\n\n3649\n10095\n4132\n1888\n6453\n7269\n5898\n9769\n\n5465\n8855\n12706\n6698\n3456\n4949\n\n3093\n8425\n8580\n8693\n1401\n6281\n8069\n7607\n6368\n7467\n\n6885\n1469\n1315\n2335\n6669\n3226\n5261\n2477\n1931\n6691\n2575\n1377\n5029\n\n5220\n4639\n5985\n1669\n1162\n4269\n6532\n4630\n4169\n3164\n6896\n6432\n\n12705\n4248\n5011\n7544\n2040\n7767\n\n60119\n\n58181\n\n69195\n\n10643\n5605\n3433\n3275\n4741\n2858\n8151\n10582\n\n4299\n10093\n12455\n2583\n2604\n11520\n\n4609\n8245\n1989\n17043\n\n4998\n8893\n3978\n4961\n2847\n8109\n9266\n1034\n4808\n\n6593\n3208\n3966\n1649\n5952\n1961\n4517\n3911\n4239\n3255\n4144\n\n3721\n2249\n3195\n1827\n3942\n2713\n4481\n2475\n3566\n2701\n2832\n2372\n5281\n3315\n6029\n\n5817\n6495\n8621\n7101\n3936\n1675\n3309\n1873\n8737\n4558\n\n6533\n4655\n4205\n6051\n2902\n2566\n2173\n1535\n5827\n2807\n6391\n7194\n\n5844\n2233\n4210\n2293\n2529\n3081\n1524\n5844\n5931\n6456\n6427\n4873\n3195\n\n3101\n3627\n9503\n1451\n9148\n1990\n2748\n5949\n\n1357\n\n4305\n4183\n8665\n9225\n5169\n8711\n7360\n4352\n8762\n\n19297\n6700\n3922\n13161\n\n7177\n5207\n12649\n10712\n3766\n\n3197\n11717\n4128\n14139\n\n5844\n4035\n9789\n5874\n6510\n6863\n9172\n3398\n\n3289\n3003\n7463\n5109\n6382\n5082\n2587\n5094\n7545\n2487\n2585\n\n15763\n18524\n4143\n8602\n\n4828\n2314\n3592\n3742\n4802\n3369\n3590\n5778\n5439\n4261\n4063\n3117\n3004\n2803\n4503\n\n3356\n1306\n6227\n5953\n1962\n1596\n1031\n6259\n1314\n5303\n4053\n6606\n6141\n\n23367\n13468\n\n7561\n3664\n10113\n11060\n7173\n6770\n\n14496\n3114\n2958\n10451\n12975\n\n2779\n8809\n7880\n9646\n5169\n11946\n10951\n\n1808\n8038\n4003\n2639\n4307\n2437\n3374\n6114\n3179\n2414\n6346\n\n3683\n4271\n5347\n4651\n4209\n3216\n4860\n5151\n3676\n3964\n5855\n5537\n2334\n2847\n\n3036\n4491\n2992\n5222\n1007\n3398\n4412\n5978\n4514\n5111\n2173\n4115\n3297\n3917\n4087\n\n1120\n9199\n13194\n8285\n12725\n2925\n\n4647\n8917\n5082\n2858\n7782\n3193\n7282\n2432\n6432\n\n8143\n1337\n7526\n4612\n3111\n2914\n2283\n7143\n\n13219\n11724\n11762\n13753\n11229\n9784\n\n2649\n8191\n6390\n2745\n7430\n6038\n7916\n7448\n2535\n\n16115\n13575\n13440\n12313\n7022\n\n3839\n5604\n3500\n1413\n4189\n2483\n4394\n4244\n5173\n4677\n2620\n3916\n2756\n3345\n5480\n\n3136\n4946\n2793\n3039\n4646\n6070\n4631\n1187\n1349\n5338\n3411\n2576\n3330\n\n1836\n11177\n12819\n12092\n11081\n5211\n\n3380\n3522\n6853\n4943\n3560\n6687\n2925\n1849\n7189\n2430\n4195\n4027\n\n17162\n17251\n\n12148\n6285\n8465\n9842\n2740\n9321\n\n5968\n1099\n5670\n12712\n4090\n10669\n\n5599\n11408\n2321\n8643\n7657\n3127\n\n5527\n5415\n3154\n7523\n7113\n7944\n1210\n4328\n2435\n7446\n\n9572\n2600\n4927\n7038\n2863\n8806\n6209\n5458\n6332\n\n14202\n11728\n4854\n\n2445\n4571\n5676\n4766\n6356\n3013\n7160\n4432\n4641\n5337\n2949\n7099\n\n5446\n5189\n2921\n4698\n4187\n2921\n3446\n4539\n3540\n5492\n2071\n1449\n1923\n4901\n5579\n\n13333\n5770\n5180\n9001\n7376\n8291\n\n3865\n4909\n3073\n1720\n1261\n5913\n1136\n2052\n5530\n5411\n1702\n1550\n2913\n5178\n2584\n\n10717\n2353\n1830\n5310\n8508\n10132\n10255\n\n9848\n11390\n1490\n8553\n4819\n2115\n\n2023\n1214\n1606\n9515\n5918\n1391\n\n32192\n31035\n\n6595\n5614\n7624\n8271\n4408\n5528\n1851\n1879\n7449\n6260\n\n6197\n7995\n9114\n\n1713\n5615\n1103\n6443\n4641\n4952\n6313\n1337\n6825\n3202\n5031\n4031\n4673\n\n17667\n14248\n18603\n7973\n\n4832\n5010\n1017\n2983\n1808\n6429\n5847\n5974\n2629\n2190\n5018\n5594\n3316\n2445\n\n8857\n1902\n6649\n5100\n2108\n2331\n7523\n9370\n2149\n\n4877\n2080\n3225\n1529\n6630\n3562\n2262\n2460\n3366\n5855\n1527\n2721\n5946\n\n32414\n14299\n\n6015\n2017\n6133\n9382\n3482\n7670\n1090\n4911\n4220\n\n4475\n5519\n2909\n1638\n5199\n3699\n3179\n1962\n4396\n1750\n2375\n5751\n3907\n2693\n3037\n\n2510\n6869\n3477\n4511\n1178\n4193\n6442\n5839\n3640\n4681\n4547\n4390\n4146\n\n4965\n3818\n1807\n1184\n2238\n5347\n3971\n5439\n5111\n6372\n2632\n2137\n1616\n\n5784\n6903\n3123\n5558\n6147\n3532\n4366\n2303\n6913\n1180\n6775\n2658\n5388\n\n3315\n4958\n3185\n2507\n4022\n3424\n1123\n3518\n5329\n5409\n5895\n4483\n5391\n3281\n1485\n\n5753\n2212\n6587\n4668\n1712\n6688\n1846\n3009\n3829\n5361\n6819\n1247\n\n3500\n23879\n\n3471\n4339\n3831\n4004\n5783\n5540\n5242\n5588\n4065\n5878\n1615\n1441\n4345\n3932\n3752\n\n4206\n5411\n8355\n5696\n1314\n1549\n8447\n10241\n\n4882\n4973\n3042\n7802\n5574\n4666\n2738\n7655\n7766\n\n2498\n2714\n2695\n4724\n4075\n5588\n6234\n5988\n5455\n4229\n4890\n3590\n1722\n5413\n\n1034\n2896\n4742\n4024\n5667\n6076\n5041\n8378\n3271\n8737\n\n1411\n23272\n6602\n\n3363\n12044\n11777\n7395\n4856\n2881\n12001\n\n2064\n4239\n2992\n6806\n6814\n3037\n3173\n3336\n5565\n\n2053\n2518\n5120\n4318\n1240\n3674\n7302\n1160\n3537\n4288\n2983\n\n5826\n5926\n7843\n2996\n6079\n7832\n2716\n8027\n1470\n8005\n\n10435\n24340\n11044\n\n7277\n4053\n1588\n2648\n5715\n3133\n5720\n6702\n7375\n2244\n7196\n2918\n\n20699\n12028\n\n11053\n11400\n3604\n7989\n7666\n7379\n10696\n\n5916\n13989\n\n6510\n4784\n2679\n11144\n\n5017\n3999\n6453\n1013\n3002\n4568\n6255\n6327\n5108\n2654\n1390\n3828\n5913\n2765\n\n1242\n8316\n3256\n1006\n6026\n10319\n3552\n6538\n\n2266\n2011\n7466\n2302\n5729\n1468\n1527\n6297\n7443\n3490\n6348\n\n2767\n5816\n5452\n2946\n3461\n1647\n5520\n5237\n3193\n6009\n5393\n5583\n2746\n4326\n3446\n\n2147\n1913\n3254\n5992\n7458\n4791\n7670\n3815\n2950\n6429\n1796\n\n17982\n7519\n\n33743\n\n2744\n8866\n15440\n10433\n15680\n\n1710\n3609\n1849\n4343\n5780\n4180\n3584\n2376\n4885\n4627\n1885\n3348\n2099\n1006\n2550\n\n22016\n29227\n\n22422\n12606\n11998\n\n5357\n2121\n2744\n4713\n6531\n5164\n5774\n3553\n4974\n2279\n3932\n2702\n\n4518\n3189\n3337\n1722\n2096\n5583\n2532\n1486\n2548\n2643\n5666\n6102\n4889\n5184\n\n4730\n6382\n3208\n6277\n2724\n5944\n2098\n5082\n1984\n5999\n\n6747\n7644\n8933\n2013\n7112\n3608\n7377\n5398\n1166\n\n9919\n6479\n7831\n6544\n5801\n10986\n10245\n\n9314\n10790\n5728\n3835\n9597\n7245\n\n12744\n1502\n11126\n8626\n13645\n8998\n\n5883\n2823\n2589\n5695\n6506\n2769\n4070\n4731\n5193\n2890\n2243\n\n5437\n6241\n1398\n3651\n7865\n1854\n4169\n2921\n5701\n2839\n2277\n\n4279\n4987\n5933\n1476\n2572\n1320\n2304\n1148\n4617\n4571\n1202\n3799\n3033\n2079\n5217\n\n4247\n3918\n7174\n7333\n2215\n8351\n1339\n7229\n5266\n\n1357\n3298\n3348\n6580\n1364\n6996\n6527\n3579\n4631\n1355\n7165\n4424\n\n19948\n18739\n15360\n\n2666\n3271\n1581\n6748\n3957\n6221\n1543\n2058\n5185\n1264\n2211\n1826\n\n9309\n5622\n9303\n2516\n12117\n4126\n2700\n\n9949\n16352\n18404\n\n3351\n10708\n10017\n2594\n4711\n1749\n9432\n\n8877\n7738\n10940\n6554\n6358\n\n4961\n7610\n1594\n5133\n5937\n2581\n5131\n4044\n9096\n\n2157\n6653\n5601\n1394\n5724\n6487\n4100\n3519\n2273\n5199\n6924\n6953\n4903\n\n26443\n31228\n\n2953\n10062\n5321\n4145\n2668\n2525\n4886\n2965\n"

func main() {
	chunks := strings.Split(input, "\n")

	max1 := 0
	max2 := 0
	max3 := 0

	current := 0

	for _, line := range chunks {
		if line == "" {

			if current > max1 {
				max3 = max2
				max2 = max1
				max1 = current
			} else if current > max2 {
				max3 = max2
				max2 = current
			} else if current > max3 {
				max3 = current
			}

			current = 0
			continue
		}

		i, _ := strconv.ParseInt(line, 10, 32)

		current = current + int(i)
	}

	fmt.Println(max1 + max2 + max3)
}