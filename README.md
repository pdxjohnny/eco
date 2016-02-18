Eco
---

The purpose of this tool is to provide a way to map an ecosystem.

Current Functionality
---

* Place markers at points of interest. Markers have several attributes.
 * Name
 * Description
 * longitude and latitude
* Find markers within a distance of a location

TODO
---

* Add user accounts
* Client to store data on users devices so I don't have to store their data sets
* Per user maps -> pdxjohnny/frogs
* Public vs. Private maps, Share maps with some users give some write access
* Custom parameters
* Find markers which match a parameter
* Sensor data - locations could be sensors that report and update data
  * This data would probably be relevant over time which posses the question of how
    it should be stored. Not sure yet how but a thought is to have a separate area for
    logs. This would keep track of all edits made to anything. That way you can view the
    changes over time. This would enable the tracking of wildlife so as to map their
    locations or the plotting of sensor data over time.

Building
---

```bash
go build -tags netgo . && \
docker-compose build && \
docker-compose up
```
