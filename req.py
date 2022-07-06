#insert spotify api code
import requests
import spotify_secrets


client_id = spotify_secrets.id

#http://paltos.github.com/callback/
#encoded:
uri = 'http%3A%2F%2Fpaltos.github.com%2Fcallback%2F'

x = requests.get(f'https://accounts.spotify.com/authorize?client_id={client_id}&response_type=code&redirect_uri={uri}')




print(x.status_code)




