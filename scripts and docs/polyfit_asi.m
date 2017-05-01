x=[0,15,30,40,50, 60, 80,130,160,190]
y=[0,15,30,60,90,120,180,270,315,360]
p=polyfit(x,y,6)
polyout(p,'x')

xx=0:1:190
f=polyval(p,xx)


plot(x,y,'o',xx,f,'-')
ylabel ("Degrees");
xlabel ("Speed (kts)");

