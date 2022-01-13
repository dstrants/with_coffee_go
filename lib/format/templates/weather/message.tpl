{{ $forecast := index .Forecast.Forecastday 0}}

*{{ .Location.Name }}* - _{{$forecast.Day.Condition.Text}}_
:sunny: {{ $forecast.Astro.Sunrise }} | :moon: {{ $forecast.Astro.Sunset }}
:thermometer:  Max: {{ $forecast.Day.MaxtempC }} | :thermometer: Min: {{ $forecast.Day.MintempC }}
:dash: {{$forecast.Day.MaxwindKph}} kph | :rain_cloud: {{ $forecast.Day.TotalprecipMm }} mm
:eyes: {{ $forecast.Day.AvgvisKm}} km | :droplet: {{$forecast.Day.Avghumidity}} %
:sunglasses: {{$forecast.Day.Uv}} | {{ $forecast.Astro.MoonPhase }}