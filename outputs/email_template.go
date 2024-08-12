package outputs

var TestTemplate = `<!DOCTYPE html>
<html style="height: 100%; width: 100%; background-color: #eaeaea">
	<head>
		<title>Email Template</title>
	</head>

	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
	<link href="https://fonts.googleapis.com/css?family=Inter" rel="stylesheet" />
	<meta charset="UTF-8" />
	<meta http-equiv="X-UA-Compatible" content="IE=edge" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />

	<body style="margin: 0; font-family: 'Inter'">
		<div
			class="main-body"
			style="
				position: absolute;
				left: 50%;
				top: 50%;
				width: 700px;
				-webkit-transform: translate(-50%, -50%);
				transform: translate(-50%, -50%);
				background-color: #eaeaea;
			"
		>
			<div
				class="title-image"
				style="text-align: center; padding: 20px 0px 0px"
			>
				<img
					src= {{.HeaderLogo}}
					alt="logo"
					style="width: 180px"
				/>
			</div>
			<div
				class="main-content"
				style="
					background-color: #f3f3f3;
					margin: 10px 20px;
					padding: 0px 0px;
					border-radius: 5px;
					border: 2px solid #fff;
					border-left-color: #fd0000;
				"
			>
				<ul>
				Test Email: If you see this email it means that email is configured as channel integration successfully.  
				</ul>
			</div>
			<div class="out-box">
				<div
					class="temp-button"
					style="display: grid; justify-content: center; align-items: center"
				>
					<button
						style="
							background-color: #05147d;
							padding: 15px 10px;
							color: #fff;
							margin: 20px 30px;
							border: 0px solid;
							cursor: pointer;
							font-size: 15px;
							font-weight: 700;
						"
					>
					<a
							style="color: #fff; text-decoration: none"
							href={{.Link}}
							>Login</a
						>
					</button>
					<p
						style="
							color: #5e5e5e;
							font-size: 15px;
							margin: 20px 30px;
							text-align: center;
						"
					>
						This email was sent from tenant <b> {{.TenantName}} </b>
					</p>
				</div>
			</div>
			<div class="footer" style="margin: 20px 0px; text-align: center">
				<p style="font-size: 14px; color: #969595">
					You're receiving this email because you have configured Email as an integration from AccuKnox . 
				</p>
				<p style="font-size: 14px; color: #969595; padding-bottom: 15px">
					AccuKnox . 7772 Orogrande Place . Cupertino, CA 95014
				</p>
			</div>
		</div>
	</body>
</html>`

var AlertTemplate = `<!DOCTYPE html>
<html style="height: 100%; width: 100%; background-color: #eaeaea">
	<head>
		<title>Email Template</title>
	</head>

	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
	<link href="https://fonts.googleapis.com/css?family=Inter" rel="stylesheet" />
	<meta charset="UTF-8" />
	<meta http-equiv="X-UA-Compatible" content="IE=edge" />
	<meta name="viewport" content="width=device-width, initial-scale=1.0" />

	<body style="margin: 0; font-family: 'Inter'">
		<div
			class="main-body"
			style="
				position: absolute;
				left: 50%;
				top: 50%;
				width: 700px;
				-webkit-transform: translate(-50%, -50%);
				transform: translate(-50%, -50%);
				background-color: #eaeaea;
			"
		>
			<div
				class="title-image"
				style="text-align: center; padding: 20px 0px 0px"
			>
				<img
					src= {{.HeaderLogo}}
					alt="logo"
					style="width: 180px"
				/>
			</div>
			<div class="title-text">
				<h3 style="text-align: center; font-weight: 900">
					Monitor Alert : {{.TriggerName}}
				</h3>
			</div>
			<div
				class="main-content"
				style="
					background-color: #f3f3f3;
					margin: 10px 20px;
					padding: 0px 0px;
					border-radius: 5px;
					border: 2px solid #fff;
					border-left-color: #fd0000;
				"
			>
				<ul style="list-style-type: none">
					<li style="padding: 5px 0px; font-size: 15px">
						<b>Sev :</b>{{.Severity}}
					</li>
					<li style="padding: 5px 0px; font-size: 15px">
						<b>Policy-name :</b> {{.PolicyName}}
					</li>
					<li style="padding: 5px 0px; font-size: 15px">
						<b>Message :</b> {{.Message}}
					</li>
					<li style="padding: 5px 0px; font-size: 15px">
						<b>Cluster :</b> {{.Cluster}}
					</li>
					<li style="padding: 5px 0px; font-size: 15px">
						<b>Action :</b> {{.Action}}
					</li>
					<li style="padding: 5px 0px; font-size: 15px">
						<b>Result :</b> {{.Result}}
					</li>
				</ul>
			</div>
			<div class="out-box">
				<div
					class="temp-button"
					style="display: grid; justify-content: center; align-items: center"
				>
					<button
						style="
							background-color: #05147d;
							padding: 15px 10px;
							color: #fff;
							margin: 20px 30px;
							border: 0px solid;
							cursor: pointer;
							font-size: 15px;
							font-weight: 700;
						"
					>
					<a
							style="color: #fff; text-decoration: none"
							href={{.Link}}
							>View Alerts on Accuknox</a
						>
					</button>
					<p
						style="
							color: #5e5e5e;
							font-size: 15px;
							margin: 20px 30px;
							text-align: center;
						"
					>
						This alert was raised by tenant <b> {{.TenantName}} </b>
					</p>
				</div>
			</div>
			<div class="footer" style="margin: 20px 0px; text-align: center">
				<p style="font-size: 14px; color: #969595">
					You're receiving this email because you have configured Email as an integration from AccuKnox . 
				</p>
				<p style="font-size: 14px; color: #969595; padding-bottom: 15px">
					AccuKnox . 7772 Orogrande Place . Cupertino, CA 95014
				</p>
			</div>
		</div>
	</body>
</html>`
