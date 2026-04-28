conn_str = (
    "DRIVER={ODBC Driver 18 for SQL Server};"
    "SERVER=VWIBSSQLPRD101;"
    "DATABASE=GHRBI;"
    "UID=your_user;"
    "PWD=your_password;"
    "Encrypt=no;"
    "TrustServerCertificate=yes;"
)
conn = pyodbc.connect(conn_str)