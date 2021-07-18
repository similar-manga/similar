import os
import csv
import sqlite3
from collections import defaultdict

# parameters
basedir = "../../similar_data/mapping/"
sqlitefile = basedir+"mappings.db"
csvfiles = {
    "al" : basedir+"anilist2mdex.csv",
    "ap" : basedir+"animeplanet2mdex.csv",
    "bw" : basedir+"bookwalker2mdex.csv",
    "mu" : basedir+"mangaupdates2mdex.csv",
    "nu" : basedir+"novelupdates2mdex.csv",
    "kt" : basedir+"kitsu2mdex.csv",
    "mal" : basedir+"myanimelist2mdex.csv",
}

# remove db file if exists so we are fresh
if os.path.exists(sqlitefile):
  os.remove(sqlitefile)

# open the connection to file
con = sqlite3.connect(sqlitefile)
cur = con.cursor()

# create table and read in data for each csv file
# this seems to make the database be much larger then needed!
# for table in csvfiles:
#     cur.execute("CREATE TABLE IF NOT EXISTS "+table+" (idMdex, idExt);")
#     with open(csvfiles[table],'r') as fin:
#         dr = csv.DictReader(fin, fieldnames=['idExt','idMdex'])
#         to_db = [(i['idMdex'], i['idExt']) for i in dr]
#     cur.executemany("INSERT INTO "+table+" (idMdex, idExt) VALUES (?, ?);", to_db)
#     print("wrote "+str(len(to_db))+" to db for "+table+"...")
#     con.commit()


# collect all manga into a large "mapping"
cur.execute("CREATE TABLE IF NOT EXISTS mappings (mdex, al, ap, bw, mu, nu, kt, mal);")
mangas = defaultdict(lambda: None)
for table in csvfiles:
    with open(csvfiles[table], 'r', encoding="utf8", errors='replace') as fin:
        dr = csv.DictReader(fin, fieldnames=['idExt','idMdex'])
        for i in dr:
            if i['idMdex'] not in mangas:
                mangas[i['idMdex']] = defaultdict(lambda: None)
            mangas[i['idMdex']][table] = i['idExt']
# finally write to database
to_db = [(i, mangas[i]["al"], mangas[i]["ap"], mangas[i]["bw"], mangas[i]["mu"], mangas[i]["nu"], mangas[i]["kt"], mangas[i]["mal"]) for i in mangas]
cur.executemany("INSERT INTO mappings (mdex, al, ap, bw, mu, nu, kt, mal) VALUES (?, ?, ?, ?, ?, ?, ?, ?);", to_db)
print("wrote "+str(len(to_db))+" to db...")
con.commit()


# finally close the file
con.close()



