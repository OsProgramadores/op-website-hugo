#AUTHOR: Andre Rosa
#DATE: 22OCT2019
#OBJECTIVE: Search recursively for broken links in files. 

from pathlib import Path

# THIS PROGRAM:
# 1 - Search recursively files with .html and .yaml extensions and save themn in a array;
# 2 - Search each file from the array for links using REGEX to find links in the file text;
# 3 - Saves each link found in an array of siteLink class;  
# 4 - Tests each link for broken links
# 5 - Generates a .csv report of broken links with file name.  

#------------------------------------------------------
# CLASS SITELINK
#------------------------------------------------------
class siteLink:
    ''' A class to store a link URL, file and boolean if the link works or is broken. '''
    def __init__(self, pURL, pFile):
        self.URL = pURL  # Link URL.
        self.file = pFile # file where the URL is written.
        self.broken = False # TRUE if link is broken, FALSE is link is fine. 

#------------------------------------------------------

#------------------------------------------------------
def searchForFiles(pExtensions):
    ''' Search for files, with .html and .yaml extensions, recursively.
        Param: an array with strings holding the extensions to be searched 
        Return: an array with file names '''

    lfiles = [] # creates an empty array to store file names

    for ext in pExtensions:
        # get files recursively
        for filename in Path('').glob('**/*' + ext):
            lfiles.append(filename)

    return lfiles # returns all files
#-------------------------------------------------------

#-------------------------------------------------------
def printFileList (pFiles):
    '''' Print list of files just for testing '''

    for filename in pFiles:
        print(filename)
#-------------------------------------------------------

#-------------------------------------------------------
def printLinkList (pLinks):
    '''' Print the list of links just for testing '''

    for m in pLinks:
        print (m.URL) 
#-------------------------------------------------------

#-------------------------------------------------------
def findLinksInFile(pFileName):
    ''' Searches for links in each lines of the file defined in pFileName parameter. ''' 
    import re
    lSiteLinkArray = []

    f = open(pFileName, "r")   # open the file
    
    for line in f:         # loops the lines of the opened file
        
        # regex taken from http://urlregex.com/
        lLink = re.findall("(http[s]?://(?:[a-zA-Z]|[0-9]|[$-_@.&+]|[!*\(\),]|(?:%[0-9a-fA-F][0-9a-fA-F]))+)", line)   #finds all urls in the line
        
        for any in lLink:
            any = any.split(',')[0] 
            l = siteLink(any, pFileName)   # creates a siteLink object with URL and File Name
            lSiteLinkArray.append(l)

    return lSiteLinkArray                 
#-------------------------------------------------------

#-------------------------------------------------------
def searchForLinks(pFiles):
    ''' Loops the list of files to search for URL links '''

    lSiteLinkArray = [] # an empty list of siteLink objects

    for filename in pFiles: # loops the files
        lSiteLinkArray += findLinksInFile(filename)[:]

    return lSiteLinkArray
#-------------------------------------------------------

#-------------------------------------------------------
def testForBrokenLinks(pLinks):
    ''' Tests if links in the pLinks are not broken '''
    import requests

    count = 0
    for link in pLinks:
        count +=1
        if (count % 10 == 0):
            print ('Testing links', count)
        try: 
            request = requests.get(link.URL)
            if request.status_code != 200:
                print ('Found broken link in ' + link.URL)
                link.broken = True
        except:
            print ('could not test ', link.URL)
    return pLinks

#-------------------------------------------------------

#-------------------------------------------------------
def generateReportOnBrokenLinks(pLinks):
    ''' Saves the broken links to a csv file. '''
    import datetime 
    today = datetime.datetime.now()
    today = today.strftime("%x")
    reportStr = 'DATE, FILE, BROKEN URL\n'
    reportFile = 'brokenLinks.csv'
    count = 0

    for link in pLinks:
        if (link.broken == True):
            reportStr = reportStr + today + ',' + str(link.file) + ',' + link.URL + '\n'
            count += 1
    
    print ('found ' + str(count) + 'broken links, save report in brokenLinks.csv file')
    f = open(reportFile, 'w+')  
    f.write(reportStr)  
    f.close() 
#-------------------------------------------------------

#-------------------------------------------------------
def main():

    #file extensions to search for links
    lExtensions = ['.html', '.yaml']

    # 1. gets recursevely all files with extensions defined in lExtensions
    print ('Searching for files')
    mFiles = searchForFiles(lExtensions) 
    print ('Found ' + str(len(mFiles)) + ' files')

    # 2. gets all links found in the files
    print ('Searching for URL links in files')
    mAllLinks = searchForLinks(mFiles) 
    print ('Found ' + str(len(mAllLinks)) + ' URLs in files')

    # 3. test the links
    print ('Testing links list')
    mTestedLinks = testForBrokenLinks (mAllLinks[:])
    print ('Tested ' + str(len(mTestedLinks)) + ' URLs in files')

    # 4. print the .csv report with the broken links
    print ('Saving report')
    generateReportOnBrokenLinks (mTestedLinks[:])

    #print files found for testing
    #printFileList(mFiles)
    #print links found for testing
    #printLinkList(mAllLinks)

#-------------------------------------------------------

if __name__== "__main__":
  main()