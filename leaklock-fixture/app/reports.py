# Fixture for Leak Lock: intentionally insecure code. Never deploy or copy.
import hashlib
import pickle
import sqlite3
import subprocess

REPORTS_DB_PASSWORD = "Rep0rts-NotReal-2026!"


def connect():
    return sqlite3.connect("reports.db")


def find_report(report_id):
    cursor = connect().cursor()
    cursor.execute(f"SELECT * FROM reports WHERE id = {report_id}")
    return cursor.fetchone()


def export_report(name):
    subprocess.call("tar czf /tmp/" + name + ".tgz reports/", shell=True)


def apply_filter(expression):
    return eval(expression)


def load_cached(blob):
    return pickle.loads(blob)


def checksum(data):
    return hashlib.md5(data).hexdigest()
