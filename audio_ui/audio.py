import os
import datetime
import pyttsx3
import playsound
import speech_recognition as sr
from gtts import gTTS

def speak(text):
    tts = gTTS(text=text, lang='en')
    filename = 'voice.mp3'
    tts.save(filename)
    playsound.playsound(filename)
    os.remove("voice.mp3")

def diff_speak(txt):
    engine = pyttsx3.init()
    engine.say(txt)
    engine.runAndWait()
    
def get_audio():
    r = sr.Recognizer()
    with sr.Microphone() as source:
        audio = r.listen(source)
        said = ''

        try:
            said = r.recognize_google(audio)
            print(said)
        except Exception as e:
            print('Exception: ' + str(e))
    return said

date = datetime.datetime.now()


out = 'Today is ' + date.strftime("%A") + ' the ' + date.strftime("%d") + 'th of ' + date.strftime("%B") + ' ' + str(date.year)
print(out)
speak(out)
diff_speak(out)