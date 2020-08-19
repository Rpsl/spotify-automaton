-- select artists by genre
SELECT *
FROM artists a
         INNER JOIN artists_to_genres ag ON ag.artist_id = a.id
         INNER JOIN genres g ON g.id = ag.genre_id
         INNER JOIN artists_to_tracks att ON att.artist_id = a.id
         INNER JOIN tracks t ON t.id = att.track_id
WHERE g.name like '%rap%'
   OR g.name like '%hip hop%'
ORDER BY t.added DESC;

-- select artists and their genres
SELECT *
FROM artists a
         INNER JOIN artists_to_genres ag ON ag.artist_id = a.id
         INNER JOIN genres g ON g.id = ag.genre_id
         INNER JOIN artists_to_tracks att ON att.artist_id = a.id;


